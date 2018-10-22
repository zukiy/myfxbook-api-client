package client

import (
	"encoding/xml"
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

func getSystemLoginURL(user, password string) string {
	return fmt.Sprintf(userLoginURL, user, password)
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

// UnmarshalXML unmarshal string to time
func (cd *CalendarDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var v string
	layout := "2006, January 02, 15:04"

	err = d.DecodeElement(&v, &start)
	if err != nil {
		return
	}

	t, err := time.Parse(layout, v)
	if err != nil {
		return
	}

	*cd = CalendarDate{t}
	return
}
