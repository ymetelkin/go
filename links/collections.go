package links

import (
	"fmt"
)

type collection struct {
	ID      string
	Href    string
	Links   []link
	Updated audit
}

func (col *collection) append(id string, href string, by string) (int, error) {
	var lnk link

	audit := newAudit(by)

	if col.Links == nil {
		lnk = newLink(id, 0, href, audit)
		col.Links = []link{lnk}
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

	lnk = newLink(id, size, href, audit)
	col.Updated = lnk.Updated
	col.Links = append(col.Links, lnk)

	return size, nil
}

func (col *collection) insert(id string, pos int, href string, by string) error {
	var lnk link

	audit := newAudit(by)

	if col.Links == nil {
		if pos == 0 {
			lnk = newLink(id, 0, href, audit)
			col.Links = []link{lnk}
			return nil
		}
		return invalidPositionError(pos, col)
	}

	size := len(col.Links)
	if pos >= size {
		return invalidPositionError(pos, col)
	}

	size++
	links := make([]link, size)

	for i, lnk := range col.Links {
		if i < pos {
			links[i] = lnk
		} else {
			if i == pos {
				links[i] = newLink(id, pos, href, audit)
			}
			lnk.Seq++
			links[i+1] = lnk
		}
	}

	col.Updated = newAudit(by)
	col.Links = links

	return nil
}

func (col *collection) move(id string, pos int, by string) ([]link, error) {
	if col.Links == nil {
		return nil, invalidLinkError(id, col)
	}

	size := len(col.Links)
	if size == 0 {
		return nil, nil
	}
	if pos >= size {
		return nil, invalidPositionError(pos, col)
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
		return nil, invalidLinkError(id, col)
	}
	if cur == pos {
		return nil, nil
	}

	links := make([]link, size)
	moved := []link{}

	var lnk, mv link

	for i, l := range col.Links {
		if l.ID == id {
			mv = newLink(id, pos, l.Href, newAudit(by))
		}

		if i == pos {
<<<<<<< HEAD
			lnk = Link{
				ID:      id,
				Seq:     pos,
				Updated: NewUpdateHistory(by),
			}
		} else {
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
=======
			continue
		}

		if pos < cur { //moving left
			if i < pos || i > cur {
				links[i] = l
				continue
			} else {
				lnk = col.Links[i-1]
			}
		} else if pos > cur { //moving right
			if i < cur || i > pos {
				links[i] = l
				continue
			} else {
				lnk = col.Links[i+1]
>>>>>>> 5f47947789048c5e033d95409fb25ea7dbbfa033
			}
		}
		lnk.Seq = i
		lnk.Updated = newAudit(by)

		links[i] = lnk
		moved = append(moved, lnk)
	}

	links[pos] = mv
	moved = append(moved, mv)

	col.Updated = newAudit(by)
	col.Links = links

	return moved, nil
}

func (col *collection) remove(id string, by string) (int, error) {
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

	links := make([]link, size-1)

	for i, l := range col.Links {
		if i != cur {
			idx := i
			if i > cur {
				l.Seq--
				idx--
			}
			links[idx] = l
		}
	}

	col.Updated = newAudit(by)
	col.Links = links
	return cur, nil
}

func (col *collection) addReversed(id string, seq int, href string, by string) (int, error) {
	var (
		size, i     int
		add, resize bool
	)

	new := newLink(id, seq, href, newAudit(by))

	if col.Links != nil {
		size = len(col.Links)
	}

	if size == 0 {
		col.Links = []link{new}
		return 0, nil
	}

	links := make([]link, size+1)

	for _, l := range col.Links {
		if l.ID == id {
			resize = true
		} else {
			if !add && (l.Seq > seq || (l.Seq == seq && l.Updated.Timestamp <= new.Updated.Timestamp)) {
				links[i] = new
				add = true
				i++
			}
			links[i] = l
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

func (col *collection) removeReversed(id string, by string) (int, error) {
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

	col.Links = append(col.Links[:cur], col.Links[cur+1:]...)
	col.Updated = newAudit(by)

	return len(col.Links), nil
}

func invalidLinkError(id string, col *collection) error {
	return fmt.Errorf("Link [%s] does not exists in [%s] collection", id, col.ID)
}

func invalidPositionError(pos int, col *collection) error {
	return fmt.Errorf("Invalid position [%d] in [%s] collection", pos, col.ID)
}
