package links

import (
	"fmt"
)

//Service is links facade
type Service struct {
	db db
	rd db
}

//New is shortcut constructor for Service
func New() Service {
	return Service{db: newDb("LinkCollections"), rd: newDb("LinkReversedCollections")}
}

//AddLink adds link to collection
func (svc *Service) AddLink(collectionID string, linkID string, byID string) error {
	ex := make(chan error)
	seq := make(chan int, 1)

	go func() {
		col, err := svc.db.GetCollection(collectionID)
		if err != nil {
			ex <- err
			return
		}

		if col.ID == "" {
			col.ID = collectionID
		}

		i, err := col.Append(linkID, byID)
		if err == nil {
			seq <- i
		} else {
			ex <- err
			return
		}

		err = svc.db.SaveCollection(col)
		if err != nil {
			ex <- err
			return
		}

		ex <- nil
	}()

	go func() {
		col, err := svc.rd.GetCollection(linkID)
		if err != nil {
			ex <- err
			return
		}

		if col.ID == "" {
			col.ID = linkID
		}

		_, err = col.AddReversed(collectionID, <-seq, byID)
		if err != nil {
			ex <- err
			return
		}

		err = svc.rd.SaveCollection(col)
		if err != nil {
			ex <- err
			return
		}

		ex <- nil
	}()

	return <-ex
}

//RemoveLink removes link from collection using goroutines
func (svc *Service) RemoveLink(collectionID string, linkID string, byID string) error {
	ex := make(chan error)

	go func() {
		col, err := svc.db.GetCollection(collectionID)
		if err != nil {
			ex <- err
			return
		}
		if col.ID == "" {
			ex <- fmt.Errorf("Collection [%s] does not exist", collectionID)
			return
		}

		_, err = col.Remove(linkID, byID)
		if err != nil {
			ex <- err
			return
		}

		err = svc.db.SaveCollection(col)
		if err != nil {
			ex <- err
			return
		}

		ex <- nil
	}()

	go func() {
		col, err := svc.rd.GetCollection(linkID)
		if err != nil {
			ex <- err
			return
		}
		if col.ID == "" {
			ex <- fmt.Errorf("Link [%s] does not exist", linkID)
		}

		_, err = col.RemoveReversed(collectionID, byID)
		if err != nil {
			ex <- err
		}

		err = svc.rd.SaveCollection(col)
		if err != nil {
			ex <- err
		}

		ex <- nil
	}()

	return <-ex
}

//GetCollection gets collection by its ID
func (svc *Service) GetCollection(collectionID string) (Collection, error) {
	col, err := svc.db.GetCollection(collectionID)
	return col, err
}

//GetReversedCollection gets reverese collection by its ID
func (svc *Service) GetReversedCollection(linkID string) (Collection, error) {
	col, err := svc.rd.GetCollection(linkID)
	return col, err
}
