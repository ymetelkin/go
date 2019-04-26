package links

import (
	"fmt"

	"github.com/ymetelkin/go/es"
)

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
func (svc *Service) AddLink(collectionID string, linkID string, byID string) error {
	ex1 := make(chan error, 1)
	ex2 := make(chan error, 1)
	seq := make(chan int, 1)

	go func() {
		col, err := svc.db.GetCollection(collectionID)
		if err != nil {
			ex1 <- err
			return
		}

		if col.ID == "" {
			col.ID = collectionID
		}

		i, err := col.Append(linkID, byID)
		if err == nil {
			seq <- i
		} else {
			ex1 <- err
			return
		}

		err = svc.db.SaveCollection(col)
		if err != nil {
			ex1 <- err
			return
		}

		ex1 <- nil
	}()

	go func() {
		col, err := svc.rd.GetCollection(linkID)
		if err != nil {
			ex2 <- err
			return
		}

		if col.ID == "" {
			col.ID = linkID
		}

		_, err = col.AddReversed(collectionID, <-seq, byID)
		if err != nil {
			ex2 <- err
			return
		}

		err = svc.rd.SaveCollection(col)
		if err != nil {
			ex2 <- err
			return
		}

		ex2 <- nil
	}()

	err := <-ex1
	if err == nil {
		return <-ex2
	}

	return err
}

//RemoveLink removes link from collection using goroutines
func (svc *Service) RemoveLink(collectionID string, linkID string, byID string) error {
	ex1 := make(chan error, 1)
	ex2 := make(chan error, 1)

	go func() {
		col, err := svc.db.GetCollection(collectionID)
		if err != nil {
			ex1 <- err
			return
		}
		if col.ID == "" {
			ex1 <- fmt.Errorf("Collection [%s] does not exist", collectionID)
			return
		}

		_, err = col.Remove(linkID, byID)
		if err != nil {
			ex1 <- err
			return
		}

		err = svc.db.SaveCollection(col)
		if err != nil {
			ex1 <- err
			return
		}

		ex1 <- nil
	}()

	go func() {
		col, err := svc.rd.GetCollection(linkID)
		if err != nil {
			ex2 <- err
			return
		}
		if col.ID == "" {
			ex2 <- fmt.Errorf("Link [%s] does not exist", linkID)
		}

		_, err = col.RemoveReversed(collectionID, byID)
		if err != nil {
			ex2 <- err
		}

		err = svc.rd.SaveCollection(col)
		if err != nil {
			ex2 <- err
		}

		ex2 <- nil
	}()

	err := <-ex1
	if err == nil {
		return <-ex2
	}

	return err
}

//GetCollection gets collection by its ID
func (svc *Service) GetCollection(collectionID string) (Collection, []es.ApplDocument, error) {
	return getCollection(collectionID, svc.db, svc.appl)
}

//GetReversedCollection gets reverese collection by its ID
func (svc *Service) GetReversedCollection(linkID string) (Collection, []es.ApplDocument, error) {
	return getCollection(linkID, svc.rd, svc.appl)
}

func getCollection(id string, db db, appl es.ApplService) (Collection, []es.ApplDocument, error) {
	col, err := db.GetCollection(id)
	if err != nil {
		return col, nil, err
	}

	var (
		size int
		ids  []string
	)

	if col.Links != nil {
		size = len(col.Links)
	}

	if size == 0 {
		return col, nil, nil
	}

	ids = make([]string, size)
	for i, link := range col.Links {
		ids[i] = link.ID
	}

	docs, err := appl.GetDocuments(ids)
	return col, docs, err
}
