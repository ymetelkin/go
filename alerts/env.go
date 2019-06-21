package main

type sso struct {
	Server   string
	UserID   string
	Password string
}

type amazon struct {
	DynamoTable string
}

type env struct {
	SSO     sso
	AWS     amazon
	FeedURL string
}

func qa() env {
	return env{
		SSO: sso{
			Server:   "RQXSSOVIR3.RNDEXT.LOCAL",
			UserID:   "sso_user",
			Password: "sso_user",
		},
		AWS: amazon{
			DynamoTable: "apnews-dev-apcapdevelopment-us-east-1-BusinessObjects",
		},
		FeedURL: "http://aspen-feeder-internal.aptechlab.com",
	}
}

func prod() env {
	return env{
		SSO: sso{
			Server:   "NYCEAPXSSOVIR1.DMZEXT.LOCAL",
			UserID:   "SSOUser",
			Password: "^SSOUs3r^",
		},
		AWS: amazon{
			DynamoTable: "apnews-dev-apcapdevelopment-us-east-1-BusinessObjects",
		},
		FeedURL: "http://aspen-feeder-internal.associatedpress.com",
	}
}
