package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

func createHTTPClient(proxy string) *http.Client {
	httpClient := &http.Client{}

	if strings.TrimSpace(proxy) != "" {
		proxyURL, err := url.Parse(proxy)
		if err == nil {
			httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		}
	}

	return httpClient
}
