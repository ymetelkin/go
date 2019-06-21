package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/ymetelkin/go/json"

	_ "github.com/denisenkom/go-mssqldb"
)

type alerts struct {
	Config  sso
	FeedURL string
}

type ss struct {
	ID          int64
	Query       string
	TransportID int
}

func (ss *ss) GetTransports() []string {
	ts := make([]string, 0)
	if ss.TransportID&1 != 0 {
		ts = append(ts, "Email")
	}
	if ss.TransportID&2 != 0 {
		ts = append(ts, "Webfeeds")
	}
	if ss.TransportID&4 != 0 {
		ts = append(ts, "SMS")
	}
	if ss.TransportID&8 != 0 {
		ts = append(ts, "IM")
	}
	if ss.TransportID&16 != 0 {
		ts = append(ts, "Pager")
	}
	if ss.TransportID&32 != 0 {
		ts = append(ts, "Mobile")
	}
	return ts
}

func (a *alerts) Fix() {
	fmt.Println("Connecting to SQL...")
	db, err := getDb(a.Config.Server, a.Config.UserID, a.Config.Password)
	if err != nil {
		fmt.Printf("Cannot connect to SQL: %s\n", err.Error())
		return
	}

	fmt.Println("Reading saved searches from SQL...")
	sss, err := getAlerts(db)
	if err != nil {
		fmt.Printf("Cannot read SQL: %s\n", err.Error())
		return
	}

	fmt.Println("Fixing queries...")
	sss = fixAlerts(sss)
	err = updateAlerts(sss, db, a.FeedURL)
	if err != nil {
		fmt.Printf("Cannot fix queries: %s\n", err.Error())
		return
	}
}

func newAlertsFix() alerts {
	var (
		s, p string
		e    env
	)

	for {
		if p == "" {
			p = "Select your environment (qa/prod): "
		}
		fmt.Print(p)
		fmt.Scanln(&s)
		s = strings.ToLower(s)
		if s == "qa" {
			p = ""
			e = qa()
			if verifySSO(e) {
				return alerts{
					Config:  e.SSO,
					FeedURL: e.FeedURL,
				}
			}
		} else if s == "prod" {
			p = ""
			e = prod()
			if verifySSO(e) {
				return alerts{
					Config:  e.SSO,
					FeedURL: e.FeedURL,
				}
			}
		} else {
			p = fmt.Sprintf("[%s] is invalid imput. Try again: ", s)
		}
	}
}

func verifySSO(e env) bool {
	fmt.Printf("Updating [%s] and [%s]. Is this correct? (y/n)", e.SSO.Server, e.FeedURL)
	var s string
	fmt.Scanln(&s)
	s = strings.ToLower(s)
	return s == "y"
}

func getDb(host string, uid string, pwd string) (*sql.DB, error) {
	query := url.Values{}
	query.Add("app name", "Alerts")

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(uid, pwd),
		Host:     fmt.Sprintf("%s:%d", host, 1433),
		RawQuery: query.Encode(),
	}
	return sql.Open("sqlserver", u.String())
}

func getAlerts(db *sql.DB) ([]ss, error) {
	s := `
SELECT 
	ss.Search_ID,
	ss.APQLFielded,
	mus.TransportType_ID
FROM tbl_SavedSearches ss (NOLOCK)
	inner join map_User_SavedSearch mus (NOLOCK) ON (ss.Search_Id = mus.Search_ID AND ss.Owner_ID = mus.User_ID)
	inner join map_User_Application mua (NOLOCK) ON (mus.User_ID = mua.User_ID AND mus.Application_ID = mua.Application_ID AND mua.Status = 1)
	inner join tbl_Organizations o (NOLOCK) ON o.Organization_ID = mua.Organization_ID
    inner join map_Organization_Application moa (NOLOCK) on moa.Organization_ID = mua.Organization_ID and moa.Application_ID = mua.Application_ID and moa.Status = 1 
WHERE
    ss.Status = 1
AND ss.Application_ID in (1, 4096)
AND (mus.TransportType_ID & 1 = 1 OR mus.TransportType_ID & 2 = 2 OR mus.TransportType_ID & 32 = 32)
ORDER BY 
	ss.Application_ID, 
	o.Organization_ID, 
	ss.Owner_ID,
	ss.Search_ID
	`
	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sss := make([]ss, 0)

	for rows.Next() {
		s := ss{}
		err = rows.Scan(&s.ID, &s.Query, &s.TransportID)
		if err != nil {
			fmt.Printf("Cannot read SQL: %s\n", err.Error())
			return nil, nil
		}
		sss = append(sss, s)
	}

	return sss, nil
}

func updateAlerts(sss []ss, db *sql.DB, feed string) error {
	client := &http.Client{}

	for i, ss := range sss {
		s := "UPDATE tbl_SavedSearches SET APQLFielded = @QUERY WHERE Search_ID = @ID"

		_, err := db.Query(s, sql.Named("ID", ss.ID), sql.Named("QUERY", ss.Query))
		if err != nil {
			return err
		}

		url := fmt.Sprintf("%s/api/alerts/%d", feed, ss.ID)

		req, err := http.NewRequest(http.MethodPut, url, nil)
		if err != nil {
			return err
		}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode > 300 {
			return fmt.Errorf("Failed to re-feed: [%s]", resp.Status)
		}

		fmt.Printf("%d\tUpdated %d\n", i+1, ss.ID)
	}

	return nil
}

func fixAlerts(sss []ss) []ss {
	var i int

	fixed := make([]ss, 0)

	for _, s := range sss {
		jo, err := json.ParseJSONObject(s.Query)
		if err != nil {
			continue
		}

		query, err := jo.GetObject("query")
		if err != nil {
			continue
		}
		b, err := query.GetObject("bool")
		if err != nil {
			continue
		}

		must, err := b.GetObject("must")
		if err != nil {
			continue
		}

		qs, err := must.GetObject("query_string")
		if err != nil {
			continue
		}

		q, err := qs.GetString("query")
		if err != nil {
			continue
		}

		fx, ok := fixQuery(q)
		if ok {
			qs.SetString("query", fx)
			fixed = append(fixed, ss{ID: s.ID, Query: jo.InlineString()})
			i++
			fmt.Printf("%d\tSearch ID: %d\tTransports: %d => %v\n%s\n%s\n", i, s.ID, s.TransportID, s.GetTransports(), q, fx)
		}
	}

	return fixed
}
