package links

import (
	"fmt"
	"testing"
)

func TestService(t *testing.T) {
	svc, err := New("http://proteus-qa-uno-esdata.aptechlab.com:9200")
	if err != nil {
		t.Error(err.Error())
	}

	id := "1a087fa501d8445ab3d319fcbc72b709"
	req := LinkRequest{CollectionID: id, LinkID: "664259da4f1f429bab16307eea9a582f", UserID: "YM"}
	res := svc.AddLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	req = LinkRequest{CollectionID: id, LinkID: "b416041bc1de48799ff18894836e14c6", UserID: "YM"}
	res = svc.AddLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col := svc.GetCollection(GetCollectionRequest{CollectionID: id, Fields: []string{"itemid", "headline", "type"}})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	fmt.Printf("%s should have 2 links\n", id)
	fmt.Println(col.ToString())

	id = "bbade2c1b43b4184bf0bee9eebdf9dce"
	req = LinkRequest{CollectionID: id, LinkID: "b416041bc1de48799ff18894836e14c6", UserID: "YM"}
	res = svc.AddLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	req = LinkRequest{CollectionID: id, LinkID: "664259da4f1f429bab16307eea9a582f", UserID: "YM"}
	res = svc.AddLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col = svc.GetCollection(GetCollectionRequest{CollectionID: id, Fields: []string{"itemid", "headline", "type"}})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	fmt.Printf("%s should have 2 links\n", id)
	fmt.Println(col.ToString())

	id = "664259da4f1f429bab16307eea9a582f"
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id, Fields: []string{"headline", "type"}})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	fmt.Printf("\n%s should be linked to 2 docs\n", id)
	fmt.Println(col.ToString())
	req = LinkRequest{CollectionID: "1a087fa501d8445ab3d319fcbc72b709", LinkID: "b416041bc1de48799ff18894836e14c6", UserID: "YM"}
	res = svc.MoveLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id, Fields: []string{"headline", "type"}})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	fmt.Printf("\n%s should be moved and linked to 2 docs\n", id)
	fmt.Println(col.ToString())

	id = "b416041bc1de48799ff18894836e14c6"
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id, Fields: []string{"headline", "type"}})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	fmt.Printf("\n%s should be linked to 2 docs\n", id)
	fmt.Println(col.ToString())

	id = "1a087fa501d8445ab3d319fcbc72b709"
	req = LinkRequest{CollectionID: id, LinkID: "664259da4f1f429bab16307eea9a582f", UserID: "YM"}
	res = svc.RemoveLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	req = LinkRequest{CollectionID: id, LinkID: "b416041bc1de48799ff18894836e14c6", UserID: "YM"}
	res = svc.RemoveLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id, Fields: []string{"headline", "type"}})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}

	id = "bbade2c1b43b4184bf0bee9eebdf9dce"
	req = LinkRequest{CollectionID: id, LinkID: "664259da4f1f429bab16307eea9a582f", UserID: "YM"}
	res = svc.RemoveLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	req = LinkRequest{CollectionID: id, LinkID: "b416041bc1de48799ff18894836e14c6", UserID: "YM"}
	res = svc.RemoveLink(req)
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
	col = svc.GetReversedCollection(GetCollectionRequest{CollectionID: id, Fields: []string{"headline", "type"}})
	if res.Status != StatusSuccess {
		t.Error(res.Result)
	}
}
