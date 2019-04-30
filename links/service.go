package links

import (
	"fmt"
	"sync"

	"github.com/ymetelkin/go/es"
	"github.com/ymetelkin/go/json"
)

//Statuses and Codes
const (
	StatusSuccess             int = 200
	StatusFailure             int = 400
	CodeDynamoError           int = 40001
	CodeLinkAddError          int = 40002
	CodeLinkAddReverseError   int = 40003
	CodeLinkRemoveError       int = 40004
	CodeLinRemoveReverseError int = 40005
	CodeNoCollectionError     int = 40006
	CodeNoLinkError           int = 40007
	CodeElasticsearchError    int = 40008
)

//LinkRequest link CRUD request
type LinkRequest struct {
	CollectionID string `json:"doc_id"`
	LinkID       string `json:"link_id"`
	UserID       string `json:"user_id"`
	Seq          int    `json:"seq"`
}

//LinkResponse link crud response
type LinkResponse struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Result string `json:"result"`
}

//ResetRequest to reset links for col
type ResetRequest struct {
	CollectionID string   `json:"doc_id"`
	LinkIDs      []string `json:"link_ids"`
	UserID       string   `json:"user_id"`
}

//GetCollectionRequest get collection from db and es request
type GetCollectionRequest struct {
	CollectionID string   `json:"doc_id"`
	Fields       []string `json:"fields"`
	UserID       string   `json:"user_id"`
}

//GetCollectionResponse get collection from db and es response
type GetCollectionResponse struct {
	Status     int
	Code       int
	Result     string
	Collection Collection
	Documents  json.Array
}

//Service is links facade
type Service struct {
	db   db
	rd   db
	appl es.ApplService
}

type remove func(string, string) (int, error)

//New is shortcut constructor for Service
func New(elasticseachClusterURL string) (Service, error) {
	appl, err := es.NewApplService(elasticseachClusterURL)
	if err != nil {
		return Service{}, err
	}
	return Service{
		db:   newDb("LinkCollections"),
		rd:   newDb("LinkReversedCollections"),
		appl: appl,
	}, nil
}

//AddLink adds link to collection
func (svc *Service) AddLink(req LinkRequest) LinkResponse {
	var (
		wg  sync.WaitGroup
		res LinkResponse
	)

	wg.Add(2)

	seq := make(chan int)

	go addLink(req.LinkID, req.CollectionID, req.UserID, seq, 0, svc.db, false, &res, &wg)
	go addLink(req.CollectionID, req.LinkID, req.UserID, seq, 0, svc.rd, true, &res, &wg)

	wg.Wait()

	if res.Status > 0 {
		return res
	}

	return LinkResponse{
		Status: StatusSuccess,
		Code:   StatusSuccess,
		Result: "Added",
	}
}

//ResetLinks reset links in collection
func (svc *Service) ResetLinks(req ResetRequest) LinkResponse {
	col, err := svc.db.GetCollection(req.CollectionID)
	if err != nil {
		return LinkResponse{
			Status: StatusFailure,
			Code:   CodeDynamoError,
			Result: err.Error(),
		}
	}

	rms := []Link{}
	if col.Links != nil && len(col.Links) > 0 {
		for _, id := range req.LinkIDs {
			for _, lnk := range col.Links {
				if lnk.ID == id {
					rms = append(rms, lnk)
				}
			}
		}
	}

	c1 := len(rms)
	c2 := len(req.LinkIDs)

	upd := NewUpdateHistory(req.UserID)
	links := make([]Link, c2)
	for i, id := range req.LinkIDs {
		links[i] = Link{
			ID:      id,
			Seq:     i,
			Updated: upd,
		}
	}
	col = Collection{
		ID:      req.CollectionID,
		Links:   links,
		Updated: upd,
	}

	err = svc.db.SaveCollection(col)
	if err != nil {
		return LinkResponse{
			Status: StatusFailure,
			Code:   CodeDynamoError,
			Result: err.Error(),
		}
	}

	if c1 > 0 || c2 > 0 {
		var (
			wg  sync.WaitGroup
			res LinkResponse
		)

		wg.Add(c1 + c2)

		if c1 > 0 {
			for _, lnk := range rms {
				go removeLink(lnk.ID, req.CollectionID, req.UserID, svc.rd, true, &res, &wg)
			}
		}

		if c2 > 0 {
			for _, lnk := range links {
				go addLink(lnk.ID, req.CollectionID, req.UserID, nil, lnk.Seq, svc.rd, true, &res, &wg)
			}
		}

		wg.Wait()

		if res.Status > 0 {
			return res
		}
	}

	return LinkResponse{
		Status: StatusSuccess,
		Code:   StatusSuccess,
		Result: "Reseted",
	}
}

//MoveLink moves link in collection
func (svc *Service) MoveLink(req LinkRequest) LinkResponse {
	col, err := svc.db.GetCollection(req.CollectionID)
	if err != nil {
		return LinkResponse{
			Status: StatusFailure,
			Code:   CodeDynamoError,
			Result: err.Error(),
		}
	}

	if col.ID == "" {
		col.ID = req.CollectionID
	}

	mv, err := col.Move(req.LinkID, req.Seq, req.UserID)
	if err != nil {
		return LinkResponse{
			Status: StatusFailure,
			Code:   CodeLinkAddError,
			Result: err.Error(),
		}
	}

	err = svc.db.SaveCollection(col)
	if err != nil {
		return LinkResponse{
			Status: StatusFailure,
			Code:   CodeDynamoError,
			Result: err.Error(),
		}
	}

	if mv != nil && len(mv) > 0 {
		var (
			wg  sync.WaitGroup
			res LinkResponse
		)

		wg.Add(len(mv))

		for _, lnk := range mv {
			go addLink(req.CollectionID, req.LinkID, req.UserID, nil, lnk.Seq, svc.rd, true, &res, &wg)
		}

		if res.Status > 0 {
			return res
		}
	}

	return LinkResponse{
		Status: StatusSuccess,
		Code:   StatusSuccess,
		Result: "Moved",
	}
}

//RemoveLink removes link from collection using goroutines
func (svc *Service) RemoveLink(req LinkRequest) LinkResponse {
	var (
		wg  sync.WaitGroup
		res LinkResponse
	)

	wg.Add(2)

	go removeLink(req.LinkID, req.CollectionID, req.UserID, svc.db, false, &res, &wg)
	go removeLink(req.CollectionID, req.LinkID, req.UserID, svc.rd, true, &res, &wg)

	wg.Wait()

	if res.Status > 0 {
		return res
	}

	return LinkResponse{
		Status: StatusSuccess,
		Code:   StatusSuccess,
		Result: "Removed",
	}
}

func addLink(id string, cid string, uid string, sqch chan int, seq int, db db, rev bool, res *LinkResponse, wg *sync.WaitGroup) {
	var code int

	defer wg.Done()

	col, err := db.GetCollection(cid)
	if err != nil {
		res = &LinkResponse{
			Status: StatusFailure,
			Code:   CodeDynamoError,
			Result: err.Error(),
		}
		return
	}

	if col.ID == "" {
		col.ID = cid
	}

	if rev {
		var i int
		if sqch == nil {
			i = seq
		} else {
			i = <-sqch
		}
		if i < 0 {
			res = &LinkResponse{
				Status: StatusFailure,
				Code:   CodeLinkAddReverseError,
				Result: fmt.Sprintf("Invalid sequence: [%d]", i),
			}
			return
		}
		_, err = col.AddReversed(id, i, uid)
		if err != nil {
			code = CodeLinkAddReverseError
		}
	} else {
		i, err := col.Append(id, uid)
		if sqch != nil {
			sqch <- i
		}
		if err != nil {
			code = CodeLinkAddError
		}
	}

	if err != nil {
		res = &LinkResponse{
			Status: StatusFailure,
			Code:   code,
			Result: err.Error(),
		}
		return
	}

	err = db.SaveCollection(col)
	if err != nil {
		res = &LinkResponse{
			Status: StatusFailure,
			Code:   CodeDynamoError,
			Result: err.Error(),
		}
	}
}

func removeLink(id string, cid string, uid string, db db, rev bool, res *LinkResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	col, err := db.GetCollection(cid)
	if err != nil {
		res = &LinkResponse{
			Status: StatusFailure,
			Code:   CodeDynamoError,
			Result: err.Error(),
		}
		return
	}

	if col.ID == "" {
		res = &LinkResponse{
			Status: StatusFailure,
			Code:   CodeNoCollectionError,
			Result: fmt.Sprintf("Collection [%s] does not exist", cid),
		}
		return
	}

	if rev {
		_, err = col.RemoveReversed(id, uid)
	} else {
		_, err = col.Remove(id, uid)
	}

	if err != nil {
		res = &LinkResponse{
			Status: StatusFailure,
			Code:   CodeLinkRemoveError,
			Result: err.Error(),
		}
		return
	}

	err = db.SaveCollection(col)
	if err != nil {
		res = &LinkResponse{
			Status: StatusFailure,
			Code:   CodeDynamoError,
			Result: err.Error(),
		}
	}
}

//GetCollection gets collection by its ID
func (svc *Service) GetCollection(req GetCollectionRequest) GetCollectionResponse {
	return getCollection(req, svc.db, svc.appl)
}

//GetReversedCollection gets reverese collection by its ID
func (svc *Service) GetReversedCollection(req GetCollectionRequest) GetCollectionResponse {
	return getCollection(req, svc.rd, svc.appl)
}

func getCollection(req GetCollectionRequest, db db, appl es.ApplService) GetCollectionResponse {
	var (
		size int
		ids  []string
		docs json.Array
	)

	col, err := db.GetCollection(req.CollectionID)
	if err != nil {
		return getCollectionResponse(StatusFailure, CodeDynamoError, err.Error(), col, docs)
	}

	if col.Links != nil {
		size = len(col.Links)
	}

	if size == 0 {
		return getCollectionResponse(StatusSuccess, StatusSuccess, "Empty collection", col, docs)
	}

	ids = make([]string, size)
	for i, link := range col.Links {
		ids[i] = link.ID
	}

	docs, err = appl.GetDocuments(ids, req.Fields)
	if err != nil {
		return getCollectionResponse(StatusFailure, CodeElasticsearchError, err.Error(), col, docs)
	}
	return getCollectionResponse(StatusSuccess, StatusSuccess, "Collection", col, docs)
}

func getCollectionResponse(status int, code int, result string, col Collection, docs json.Array) GetCollectionResponse {
	return GetCollectionResponse{
		Status:     StatusSuccess,
		Code:       StatusSuccess,
		Result:     result,
		Collection: col,
		Documents:  docs,
	}
}

//ToString JSON serializes GetCollectionResponse
func (res *GetCollectionResponse) ToString() string {
	jo := json.Object{}
	jo.AddInt("status", res.Status)
	jo.AddInt("code", res.Code)
	jo.AddString("result", res.Result)

	ja := json.Array{}
	if res.Collection.Links != nil {
		for _, link := range res.Collection.Links {
			l := json.Object{}
			l.AddInt("seq", link.Seq)
			l.AddString("id", link.ID)

			upd := json.Object{}
			upd.AddString("by", link.Updated.ID)
			upd.AddString("ts", link.Updated.DateTime())
			l.AddObject("linked", upd)

			ja.AddObject(l)
		}
	}
	jo.AddArray("links", ja)

	jo.AddArray("docs", res.Documents)

	return jo.ToString()
}
