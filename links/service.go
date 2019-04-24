package links

import (
	"errors"
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

//AddLink2 adds link to collection
func (svc *Service) AddLink2(collectionID string, linkID string, byID string) error {
	col, err := svc.db.GetCollection(collectionID)
	if err != nil {
		return err
	}

	rvs, err := svc.rd.GetCollection(linkID)
	if err != nil {
		return err
	}

	if col.ID == "" {
		col.ID = collectionID
	}

	if rvs.ID == "" {
		rvs.ID = linkID
	}

	seq, err := col.Append(linkID, byID)
	if err != nil {
		return err
	}

	_, err = rvs.AddReversed(collectionID, seq, byID)
	if err != nil {
		return err
	}

	err = svc.db.SaveCollection(col)
	if err != nil {
		return err
	}

	err = svc.rd.SaveCollection(rvs)
	if err != nil {
		return err
	}

	return nil
}

//AddLink adds link to collection
func (svc *Service) AddLink(collectionID string, linkID string, byID string) error {
	ex := make(chan error)
	seq := make(chan int)

	go svc.addForward(collectionID, linkID, byID, ex, seq)
	go svc.addReverse(collectionID, linkID, byID, ex, seq)

	err := <-ex
	return err
}

func (svc *Service) addForward(collectionID string, linkID string, byID string, ex chan<- error, seq chan<- int) {
	col, err := svc.db.GetCollection(collectionID)
	if err != nil {
		ex <- err
	}

	if col.ID == "" {
		col.ID = collectionID
	}

	i, err := col.Append(linkID, byID)
	if err == nil {
		seq <- i
	} else {
		ex <- err
	}

	err = svc.db.SaveCollection(col)
	if err != nil {
		ex <- err
	}

	ex <- nil
}

func (svc *Service) addReverse(collectionID string, linkID string, byID string, ex chan<- error, seq <-chan int) {
	i := <-seq

	col, err := svc.rd.GetCollection(linkID)
	if err != nil {
		ex <- err
	}

	if col.ID == "" {
		col.ID = linkID
	}

	_, err = col.AddReversed(collectionID, i, byID)
	if err != nil {
		ex <- err
	}

	err = svc.rd.SaveCollection(col)
	if err != nil {
		ex <- err
	}

	ex <- nil
}

//RemoveLink removes link from collection
func (svc *Service) RemoveLink(collectionID string, linkID string, byID string) error {
	col, err := svc.db.GetCollection(collectionID)
	if err != nil {
		return err
	}
	if col.ID == "" {
		s := fmt.Sprintf("Collection [%s] does not exist", collectionID)
		return errors.New(s)
	}

	rvs, err := svc.rd.GetCollection(linkID)
	if err != nil {
		return err
	}

	_, err = col.Remove(linkID, byID)
	if err != nil {
		return err
	}

	rvs.RemoveReversed(collectionID, byID)

	err = svc.db.SaveCollection(col)
	if err != nil {
		return err
	}

	err = svc.rd.SaveCollection(rvs)
	if err != nil {
		return err
	}

	return nil
}

//RemoveLink2 removes link from collection using goroutines
func (svc *Service) RemoveLink2(collectionID string, linkID string, byID string) error {
	ex := make(chan error)

	go svc.removeForward(collectionID, linkID, byID, ex)
	go svc.removeReverse(collectionID, linkID, byID, ex)

	err := <-ex
	return err
}

func (svc *Service) removeForward(collectionID string, linkID string, byID string, ex chan<- error) {
	col, err := svc.db.GetCollection(collectionID)
	if err != nil {
		ex <- err
	}
	if col.ID == "" {
		s := fmt.Sprintf("Collection [%s] does not exist", collectionID)
		ex <- errors.New(s)
	}

	_, err = col.Remove(linkID, byID)
	if err != nil {
		ex <- err
	}

	err = svc.db.SaveCollection(col)
	if err != nil {
		ex <- err
	}

	ex <- nil
}

func (svc *Service) removeReverse(collectionID string, linkID string, byID string, ex chan<- error) {
	col, err := svc.rd.GetCollection(linkID)
	if err != nil {
		ex <- err
	}
	if col.ID == "" {
		s := fmt.Sprintf("Link [%s] does not exist", linkID)
		ex <- errors.New(s)
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
