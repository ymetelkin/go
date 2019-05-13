package links

import (
	"fmt"
	"sync"

	"github.com/ymetelkin/go/json"
)

//Statuses and Codes
const (
	StatusSuccess              int = 200
	StatusFailure              int = 400
	CodeDynamoError            int = 40001
	CodeLinkAddError           int = 40002
	CodeLinkAddReverseError    int = 40003
	CodeLinkMoveError          int = 40004
	CodeLinkRemoveError        int = 40005
	CodeLinkRemoveReverseError int = 40006
	CodeNoCollectionError      int = 40007
	CodeNoLinkError            int = 40008
	CodeElasticsearchError     int = 40009
)

//LinkRequest link request
type LinkRequest struct {
	Collection doc    `json:"doc"`
	Link       doc    `json:"link"`
	UserID     string `json:"user_id"`
}

//LinkResponse link crud response
type LinkResponse struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Result string `json:"result"`
}

//MoveRequest move request
type MoveRequest struct {
	CollectionID string `json:"doc_id"`
	LinkID       string `json:"link_id"`
	Seq          int    `json:"seq"`
	UserID       string `json:"user_id"`
}

//ResetRequest to reset links for col
type ResetRequest struct {
	Collection doc    `json:"doc"`
	Links      []doc  `json:"links"`
	UserID     string `json:"user_id"`
}

//CollectionRequest get collection from db and es request
type CollectionRequest struct {
	CollectionID string `json:"doc_id"`
	UserID       string `json:"user_id"`
}

//CollectionResponse get collection from db and es response
type CollectionResponse struct {
	Status     int
	Code       int
	Result     string
	Collection collection
}

//Service is links facade
type Service struct {
	db db
	rd db
}

//New is shortcut constructor for Service
func New() (Service, error) {
	return Service{
		db: newDb("LinkCollections"),
		rd: newDb("LinkReversedCollections"),
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

	go addLink(req.Link.ID, req.Collection.ID, req.Link.Href, req.Collection.Href, req.UserID, seq, 0, svc.db, false, &res, &wg)
	go addLink(req.Collection.ID, req.Link.ID, req.Collection.Href, req.Link.Href, req.UserID, seq, 0, svc.rd, true, &res, &wg)

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
	col, err := svc.db.GetCollection(req.Collection.ID)
	if err != nil {
		return LinkResponse{
			Status: StatusFailure,
			Code:   CodeDynamoError,
			Result: err.Error(),
		}
	}

	rms := []link{}
	if col.Links != nil && len(col.Links) > 0 {
		for _, lnk := range col.Links {
			var exists bool
			for _, l := range req.Links {
				if lnk.ID == l.ID {
					exists = true
					break
				}
			}
			if !exists {
				rms = append(rms, lnk)
			}
		}
	}

	c1 := len(rms)
	c2 := len(req.Links)

	upd := newAudit(req.UserID)
	links := make([]link, c2)
	for i, l := range req.Links {
		links[i] = newLink(l.ID, i, l.Href, upd)
	}
	col = collection{
		ID:      req.Collection.ID,
		Href:    req.Collection.Href,
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
				go removeLink(req.Collection.ID, lnk.ID, req.UserID, svc.rd, true, &res, &wg)
			}
		}

		if c2 > 0 {
			for _, lnk := range links {
				go addLink(req.Collection.ID, lnk.ID, req.Collection.Href, lnk.Href, req.UserID, nil, lnk.Seq, svc.rd, true, &res, &wg)
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

//MoveLink moves link in Collection
func (svc *Service) MoveLink(req MoveRequest) LinkResponse {
	col, err := svc.db.GetCollection(req.CollectionID)
	if err != nil {
		return LinkResponse{
			Status: StatusFailure,
			Code:   CodeDynamoError,
			Result: err.Error(),
		}
	}

	if col.ID == "" {
		return LinkResponse{
			Status: StatusFailure,
			Code:   CodeLinkMoveError,
			Result: fmt.Sprintf("Collection [%s] not found", req.CollectionID),
		}
	}

	mv, err := col.move(req.LinkID, req.Seq, req.UserID)
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
			go addLink(col.ID, lnk.ID, col.Href, lnk.Href, req.UserID, nil, lnk.Seq, svc.rd, true, &res, &wg)
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

//RemoveLink removes link from Collection using goroutines
func (svc *Service) RemoveLink(req MoveRequest) LinkResponse {
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

func addLink(id string, cid string, href string, chref string, uid string, sqch chan int, seq int, db db, rev bool, res *LinkResponse, wg *sync.WaitGroup) {
	var (
		code, pos int
		result    string
		ch        bool
	)

	defer wg.Done()

	col, err := db.GetCollection(cid)
	if err != nil {
		code = CodeDynamoError
		result = err.Error()
	} else {
		if col.ID == "" {
			col.ID = cid
			col.Href = chref
		}

		if rev {
			if sqch == nil {
				pos = seq
			} else {
				pos = <-sqch
				ch = true
			}
			if pos < 0 {
				code = CodeLinkAddReverseError
				result = fmt.Sprintf("Invalid sequence: [%d]", pos)
			} else {
				_, err = col.addReversed(id, pos, href, uid)
				if err != nil {
					code = CodeLinkAddReverseError
					result = err.Error()
				}
			}
		} else {
			i, err := col.append(id, href, uid)
			if sqch != nil {
				sqch <- i
				ch = true
			}
			if err != nil {
				code = CodeLinkAddError
				result = err.Error()
			}
		}

		if code == 0 {
			err = db.SaveCollection(col)
			if err != nil {
				code = CodeDynamoError
				result = err.Error()
			}
		}
	}

	if code > 0 {
		res = &LinkResponse{
			Status: StatusFailure,
			Code:   code,
			Result: result,
		}
	}

	if sqch != nil && !ch {
		if rev {
			<-sqch
		} else {
			sqch <- -1
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
		_, err = col.removeReversed(id, uid)
	} else {
		_, err = col.remove(id, uid)
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

//GetCollection gets Collection by its ID
func (svc *Service) GetCollection(req CollectionRequest) CollectionResponse {
	return getCollection(req, svc.db)
}

//GetReversedCollection gets reverese Collection by its ID
func (svc *Service) GetReversedCollection(req CollectionRequest) CollectionResponse {
	return getCollection(req, svc.rd)
}

func getCollection(req CollectionRequest, db db) CollectionResponse {
	var (
		size int
		ids  []string
	)

	col, err := db.GetCollection(req.CollectionID)
	if err != nil {
		return getFailureResponse(CodeDynamoError, err.Error(), col)
	}

	if col.Links != nil {
		size = len(col.Links)
	}

	if size == 0 {
		return getSuccessResponse("Empty Collection", col)
	}

	ids = make([]string, size)
	for i, link := range col.Links {
		ids[i] = link.ID
	}

	return getSuccessResponse("Collection", col)
}

func getSuccessResponse(result string, col collection) CollectionResponse {
	return CollectionResponse{
		Status:     StatusSuccess,
		Code:       StatusSuccess,
		Result:     result,
		Collection: col,
	}
}

func getFailureResponse(code int, result string, col collection) CollectionResponse {
	return CollectionResponse{
		Status:     StatusFailure,
		Code:       code,
		Result:     result,
		Collection: col,
	}
}

//String JSON serializes CollectionResponse
func (res *CollectionResponse) String() string {
	jo := json.Object{}
	jo.AddInt("status", res.Status)
	jo.AddInt("code", res.Code)
	jo.AddString("result", res.Result)
	jo.AddString("href", res.Collection.Href)

	ja := json.Array{}
	if res.Collection.Links != nil {
		for _, link := range res.Collection.Links {
			l := json.Object{}
			l.AddInt("seq", link.Seq)
			l.AddString("id", link.ID)
			l.AddString("href", link.Href)

			upd := json.Object{}
			upd.AddString("by", link.Updated.ID)
			upd.AddString("ts", link.Updated.DateTime)
			l.AddObject("linked", upd)

			ja.AddObject(l)
		}
	}
	jo.AddArray("links", ja)

	return jo.String()
}
