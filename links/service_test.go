package links

import (
	"fmt"
	"testing"
)

func TestService(t *testing.T) {
	svc, err := New()
	if err != nil {
		t.Error(err.Error())
	}

	id := "A"
	req := LinkRequest{
		Collection: Link{
			ID:   id,
			Href: "/doc",
		},
		Link: Link{
			ID:   "0",
			Href: "/test",
		},
		UserID: "YM",
	}
	res := svc.AddLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	req = LinkRequest{
		Collection: Link{
			ID:   id,
			Href: "/doc",
		},
		Link: Link{
			ID:   "1",
			Href: "/test",
		},
		UserID: "YM",
	}
	res = svc.AddLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col := svc.GetCollection(GetCollectionRequest{CollectionID: id})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	fmt.Printf("%s should have 2 links\n", id)
	fmt.Println(col.String())

	id = "B"
	req = LinkRequest{
		Collection: Link{
			ID:   id,
			Href: "/doc",
		},
		Link: Link{
			ID:   "0",
			Href: "/test",
		},
		UserID: "YM",
	}
	res = svc.AddLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	req = LinkRequest{
		Collection: Link{
			ID:   id,
			Href: "/doc",
		},
		Link: Link{
			ID:   "1",
			Href: "/test",
		},
		UserID: "YM",
	}
	res = svc.AddLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col = svc.GetCollection(GetCollectionRequest{CollectionID: id})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	fmt.Printf("%s should have 2 links\n", id)
	fmt.Println(col.String())

	id = "0"
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	fmt.Printf("\n%s should be linked to 2 docs\n", id)
	fmt.Println(col.String())
	mreq := MoveRequest{CollectionID: "A", LinkID: "0", UserID: "YM"}
	res = svc.MoveLink(mreq)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	fmt.Printf("\n%s should be moved and linked to 2 docs\n", id)
	fmt.Println(col.String())

	id = "1"
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	fmt.Printf("\n%s should be linked to 2 docs\n", id)
	fmt.Println(col.String())

	id = "A"
	mreq = MoveRequest{CollectionID: id, LinkID: "0", UserID: "YM"}
	res = svc.RemoveLink(mreq)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	mreq = MoveRequest{CollectionID: id, LinkID: "1", UserID: "YM"}
	res = svc.RemoveLink(mreq)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}

	id = "B"
	mreq = MoveRequest{CollectionID: id, LinkID: "0", UserID: "YM"}
	res = svc.RemoveLink(mreq)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	mreq = MoveRequest{CollectionID: id, LinkID: "1", UserID: "YM"}
	res = svc.RemoveLink(mreq)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
}
