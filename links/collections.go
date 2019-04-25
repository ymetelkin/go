package links

import (
	"fmt"
)

//Collection struct
type Collection struct {
	ID      string
	Links   []Link
	Updated UpdateHistory
}

//Append new link to collection
func (col *Collection) Append(id string, by string) (int, error) {
	if col.Links == nil {
		link := NewLink(id, 0, by)
		col.Links = []Link{link}
		return 0, nil
	}

	size := len(col.Links)

	if size > 0 {
		for i, l := range col.Links {
			if l.ID == id {
				return i, fmt.Errorf("Link [%s] already exists in [%s] collection", id, col.ID)
			}
		}
	}

	link := NewLink(id, size, by)
	col.Updated = link.Updated
	col.Links = append(col.Links, link)

	return size, nil
}

//Insert new link into the collection at specified position
func (col *Collection) Insert(id string, pos int, by string) error {
	if col.Links == nil {
		if pos == 0 {
			link := NewLink(id, 0, by)
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
				links[i] = NewLink(id, pos, by)
			}
			link.Seq++
			links[i+1] = link
		}
	}

	col.Updated = NewUpdateHistory(by)
	col.Links = links

	return nil
}

//Move existing link into new position
func (col *Collection) Move(id string, pos int, by string) error {
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
			link.Updated = NewUpdateHistory(by)
			links[i] = lnk
		}
	}

	col.Updated = NewUpdateHistory(by)
	col.Links = links

	return nil
}

//Remove link from collection
func (col *Collection) Remove(id string, by string) (int, error) {
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

	col.Updated = NewUpdateHistory(by)
	col.Links = links
	return cur, nil
}

//AddReversed adds new entry to reverse collection
func (col *Collection) AddReversed(id string, seq int, by string) (int, error) {
	var (
		size, i     int
		add, resize bool
	)

	new := NewLink(id, seq, by)

	if col.Links != nil {
		size = len(col.Links)
	}

	if size == 0 {
		col.Links = []Link{new}
		return 0, nil
	}

	links := make([]Link, size+1)

	for _, link := range col.Links {
		if link.ID == id {
			resize = true
		} else {
			if !add && (link.Seq > seq || (link.Seq == seq && link.Updated.Unix <= new.Updated.Unix)) {
				links[i] = new
				add = true
				i++
			}
			links[i] = link
			i++
		}
	}

	if !add {
		if resize {
			links[size-1] = new
		} else {
			links[size] = new
		}
	}
	if resize {
		links = links[0:size]
	}

	col.Updated = new.Updated
	col.Links = links

	return len(links), nil
}

//RemoveReversed removes entry from reverse collection
func (col *Collection) RemoveReversed(id string, by string) (int, error) {
	if col.Links == nil {
		return 0, invalidLinkError(id, col)
	}

	size := len(col.Links)
	if size == 0 {
		return 0, invalidLinkError(id, col)
	}

	cur := -1
	for i, link := range col.Links {
		if link.ID == id {
			cur = i
			break
		}
	}
	if cur == -1 {
		return 0, invalidLinkError(id, col)
	}

	col.Links = append(col.Links[:cur], col.Links[cur+1:]...)
	col.Updated = NewUpdateHistory(by)

	return len(col.Links), nil
}

func invalidLinkError(id string, col *Collection) error {
	return fmt.Errorf("Link [%s] does not exists in [%s] collection", id, col.ID)
}

func invalidPositionError(pos int, col *Collection) error {
	return fmt.Errorf("Invalid position [%d] in [%s] collection", pos, col.ID)
}