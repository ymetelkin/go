package links

import (
	"fmt"

	"github.com/ymetelkin/go/es"
	"github.com/ymetelkin/go/json"
)

//Statuses and Codes
const (
	Success               int = 200
	Failure               int = 400
	DynamoError           int = 40001
	LinkAddError          int = 40002
	LinkAddReverseError   int = 40003
	LinkRemoveError       int = 40004
	LinRemoveReverseError int = 40005
	NoCollectionError     int = 40006
	NoLinkError           int = 40007
	ElasticsearchError    int = 40008
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
	res1 := make(chan LinkResponse, 1)
	res2 := make(chan LinkResponse, 1)
	seq := make(chan int, 1)

	go func() {
		col, err := svc.db.GetCollection(req.CollectionID)
		if err != nil {
			res1 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}

		if col.ID == "" {
			col.ID = req.CollectionID
		}

		i, err := col.Append(req.LinkID, req.UserID)
		if err == nil {
			seq <- i
		} else {
			res1 <- LinkResponse{Status: Failure, Code: LinkAddError, Result: err.Error()}
			return
		}

		err = svc.db.SaveCollection(col)
		if err != nil {
			res1 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}

		res1 <- LinkResponse{Status: Success, Code: Success, Result: "Added"}
	}()

	go func() {
		col, err := svc.rd.GetCollection(req.LinkID)
		if err != nil {
			res2 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}

		if col.ID == "" {
			col.ID = req.LinkID
		}

		_, err = col.AddReversed(req.CollectionID, <-seq, req.UserID)
		if err != nil {
			res2 <- LinkResponse{Status: Failure, Code: LinkAddReverseError, Result: err.Error()}
			return
		}

		err = svc.rd.SaveCollection(col)
		if err != nil {
			res2 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}

		res2 <- LinkResponse{Status: Success, Code: Success, Result: "Added"}
	}()

	res := <-res1
	if res.Status != Success {
		return res
	}

	return <-res2
}

//MoveLink moves link in collection
func (svc *Service) MoveLink(req LinkRequest) LinkResponse {
	res1 := make(chan LinkResponse, 1)
	res2 := make(chan LinkResponse, 1)

	go func() {
		col, err := svc.db.GetCollection(req.CollectionID)
		if err != nil {
			res1 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}

		if col.ID == "" {
			col.ID = req.CollectionID
		}

		err = col.Move(req.LinkID, req.Seq, req.UserID)
		if err != nil {
			res1 <- LinkResponse{Status: Failure, Code: LinkAddError, Result: err.Error()}
			return
		}

		err = svc.db.SaveCollection(col)
		if err != nil {
			res1 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}

		res1 <- LinkResponse{Status: Success, Code: Success, Result: "Moved"}
	}()

	go func() {
		col, err := svc.rd.GetCollection(req.LinkID)
		if err != nil {
			res2 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}

		if col.ID == "" {
			col.ID = req.LinkID
		}

		_, err = col.AddReversed(req.CollectionID, req.Seq, req.UserID)
		if err != nil {
			res2 <- LinkResponse{Status: Failure, Code: LinkAddReverseError, Result: err.Error()}
			return
		}

		err = svc.rd.SaveCollection(col)
		if err != nil {
			res2 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}

		res2 <- LinkResponse{Status: Success, Code: Success, Result: "Added"}
	}()

	res := <-res1
	if res.Status != Success {
		return res
	}

	return <-res2
}

//RemoveLink removes link from collection using goroutines
func (svc *Service) RemoveLink(req LinkRequest) LinkResponse {
	res1 := make(chan LinkResponse, 1)
	res2 := make(chan LinkResponse, 1)

	go func() {
		col, err := svc.db.GetCollection(req.CollectionID)
		if err != nil {
			res1 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}
		if col.ID == "" {
			res1 <- LinkResponse{Status: Failure, Code: NoCollectionError, Result: fmt.Sprintf("Collection [%s] does not exist", req.CollectionID)}
			return
		}

		_, err = col.Remove(req.LinkID, req.UserID)
		if err != nil {
			res1 <- LinkResponse{Status: Failure, Code: LinkRemoveError, Result: err.Error()}
			return
		}

		err = svc.db.SaveCollection(col)
		if err != nil {
			res1 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}

		res1 <- LinkResponse{Status: Success, Code: Success, Result: "Removed"}
	}()

	go func() {
		col, err := svc.rd.GetCollection(req.LinkID)
		if err != nil {
			res2 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
			return
		}
		if col.ID == "" {
			res2 <- LinkResponse{Status: Failure, Code: NoLinkError, Result: fmt.Sprintf("Link [%s] does not exist", req.LinkID)}
		}

		_, err = col.RemoveReversed(req.CollectionID, req.UserID)
		if err != nil {
			res2 <- LinkResponse{Status: Failure, Code: LinRemoveReverseError, Result: err.Error()}
		}

		err = svc.rd.SaveCollection(col)
		if err != nil {
			res2 <- LinkResponse{Status: Failure, Code: DynamoError, Result: err.Error()}
		}

		res2 <- LinkResponse{Status: Success, Code: Success, Result: "Removed"}
	}()

	res := <-res1
	if res.Status != Success {
		return res
	}

	return <-res2
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
		return getCollectionResponse(Failure, DynamoError, err.Error(), col, docs)
	}

	if col.Links != nil {
		size = len(col.Links)
	}

	if size == 0 {
		return getCollectionResponse(Success, Success, "Empty collection", col, docs)
	}

	ids = make([]string, size)
	for i, link := range col.Links {
		ids[i] = link.ID
	}

	docs, err = appl.GetDocuments(ids, req.Fields)
	if err != nil {
		return getCollectionResponse(Failure, ElasticsearchError, err.Error(), col, docs)
	}
	return getCollectionResponse(Success, Success, "Collection", col, docs)
}

func getCollectionResponse(status int, code int, result string, col Collection, docs json.Array) GetCollectionResponse {
	return GetCollectionResponse{
		Status:     Success,
		Code:       Success,
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
