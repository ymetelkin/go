package links

import (
	"errors"
	"fmt"
)

//Link struct
type Link struct {
	ID  string
	Seq int
}

//Collection struct
type Collection struct {
	ID    string
	Links []Link
}

//Append new link to collection
func (col *Collection) Append(id string) (int, error) {
	if col.Links == nil {
		link := Link{ID: id}
		col.Links = []Link{link}
		return 0, nil
	}

	size := len(col.Links)

	if size > 0 {
		for i, l := range col.Links {
			if l.ID == id {
				s := fmt.Sprintf("Link [%s] already exists in [%s] collection", id, col.ID)
				return i, errors.New(s)
			}
		}
	}

	link := Link{ID: id, Seq: size}
	col.Links = append(col.Links, link)

	return size, nil
}

//Insert new link into the collection at specified position
func (col *Collection) Insert(id string, pos int) error {
	if col.Links == nil {
		if pos == 0 {
			link := Link{ID: id}
			col.Links = []Link{link}
			return nil
		}
		return invalidPositionError(pos, col)
	}

	size := len(col.Links)
	if pos >= size {
		return invalidPositionError(pos, col)
	}

	size++
	links := make([]Link, size)

	for i, link := range col.Links {
		if i < pos {
			links[i] = link
		} else {
			if i == pos {
				links[i] = Link{ID: id, Seq: pos}
			}
			link.Seq++
			links[i+1] = link
		}
	}

	col.Links = links

	return nil
}

//Move existing link into new position
func (col *Collection) Move(id string, pos int) error {
	if col.Links == nil {
		return invalidLinkError(id, col)
	}

	size := len(col.Links)
	if size == 0 {

	}
	if pos >= size {
		return invalidPositionError(pos, col)
	}

	cur := -1
	if size > 0 {
		for i, l := range col.Links {
			if l.ID == id {
				cur = i
				break
			}
		}
	}
	if cur == -1 {
		return invalidLinkError(id, col)
	}
	if cur == pos {
		return nil
	}

	links := make([]Link, size)

	for i, link := range col.Links {
		if i == pos {
			links[i] = Link{ID: id, Seq: pos}
		} else {
			var lnk Link
			if pos < cur { //moving left
				if i < pos || i > cur {
					links[i] = link
					continue
				} else {
					lnk = col.Links[i-1]
				}
			} else { //moving right
				if i < cur || i > pos {
					links[i] = link
					continue
				} else {
					lnk = col.Links[i+1]
				}
			}
			lnk.Seq = i
			links[i] = lnk
		}
	}

	col.Links = links

	return nil
}

//Remove link from collection
func (col *Collection) Remove(id string) (int, error) {
	if col.Links == nil {
		return 0, invalidLinkError(id, col)
	}

	size := len(col.Links)
	if size == 0 {
		return 0, invalidLinkError(id, col)
	}

	cur := -1
	for i, l := range col.Links {
		if l.ID == id {
			cur = i
			break
		}
	}
	if cur == -1 {
		return 0, invalidLinkError(id, col)
	}

	links := make([]Link, size-1)

	for i, link := range col.Links {
		if i != cur {
			idx := i
			if i > cur {
				link.Seq--
				idx--
			}
			links[idx] = link
		}
	}

	col.Links = links
	return cur, nil
}

func invalidLinkError(id string, col *Collection) error {
	s := fmt.Sprintf("Link [%s] does not exists in [%s] collection", id, col.ID)
	return errors.New(s)
}

func invalidPositionError(pos int, col *Collection) error {
	s := fmt.Sprintf("Invalid position [%d] in [%s] collection", pos, col.ID)
	return errors.New(s)
}
