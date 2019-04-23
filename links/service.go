package links

import (
	"errors"
	"fmt"
)

//Service is links facade
type Service struct {
}

//New is shortcut constructor for Service
func New() Service {
	return Service{}
}

//AddLink adds link to collection
func (svc *Service) AddLink(collectionID string, linkID string) error {
	db := newDb()
	col, err := db.GetCollection(collectionID)
	if err != nil {
		return err
	}

	if col.ID == "" {
		col.ID = collectionID
	}

	_, err = col.Append(linkID)
	if err != nil {
		return err
	}

	err = db.SaveCollection(col)
	if err != nil {
		return err
	}

	return nil
}

//RemoveLink removes link from collection
func (svc *Service) RemoveLink(collectionID string, linkID string) error {
	db := newDb()
	col, err := db.GetCollection(collectionID)
	if err != nil {
		return err
	}
	if col.ID == "" {
		s := fmt.Sprintf("Collection [%s] does not exist", collectionID)
		return errors.New(s)
	}

	_, err = col.Remove(linkID)
	if err != nil {
		return err
	}

	err = db.SaveCollection(col)
	if err != nil {
		return err
	}

	return nil
}

//GetCollection gets collection by its ID
func (svc *Service) GetCollection(collectionID string) (Collection, error) {
	db := newDb()
	col, err := db.GetCollection(collectionID)
	return col, err
}
