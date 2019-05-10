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
		CollectionID:   id,
		CollectionHref: "/doc",
		LinkID:         "0",
		LinkHref:       "/test",
		UserID:         "YM",
	}
	res := svc.AddLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	req = LinkRequest{
		CollectionID:   id,
		CollectionHref: "/doc",
		LinkID:         "1",
		LinkHref:       "/test",
		UserID:         "YM",
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
		CollectionID:   id,
		CollectionHref: "/doc",
		LinkID:         "0",
		LinkHref:       "/test",
		UserID:         "YM",
	}
	res = svc.AddLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	req = LinkRequest{
		CollectionID:   id,
		CollectionHref: "/doc",
		LinkID:         "1",
		LinkHref:       "/link",
		UserID:         "YM",
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
	req = LinkRequest{CollectionID: "A", LinkID: "0", UserID: "YM"}
	res = svc.MoveLink(req)
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
	req = LinkRequest{CollectionID: id, LinkID: "0", UserID: "YM"}
	res = svc.RemoveLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	req = LinkRequest{CollectionID: id, LinkID: "1", UserID: "YM"}
	res = svc.RemoveLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}

	id = "B"
	req = LinkRequest{CollectionID: id, LinkID: "0", UserID: "YM"}
	res = svc.RemoveLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	req = LinkRequest{CollectionID: id, LinkID: "1", UserID: "YM"}
	res = svc.RemoveLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
}
