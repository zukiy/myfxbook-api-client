package client

import (
	"fmt"
	"time"
)

func getLoginURL(email, password string) string {
	return apiURL + fmt.Sprintf(loginURL, email, password)
}

func getLogoutURL(session string) string {
	return apiURL + fmt.Sprintf(logoutURL, session)
}

func getGetMyAccountsURL(session string) string {
	return apiURL + fmt.Sprintf(getMyAccountsURL, session)
}

func timeToString(t time.Time) string {
	return t.Format("2006-01-02 00:00")
}

func timeFromString(s string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 04:05", s, time.Local)
	if err != nil {
		t = time.Now()
	}

	return t
}
