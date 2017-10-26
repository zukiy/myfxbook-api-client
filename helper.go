package client

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

func loadData(url string, responseContainer interface{}) error {
	body, err := getRequest(url)
	if err != nil {
		return err
	}

	switch c := responseContainer.(type) {
	case *LoginResponse, *Response, *GetMyAccountsResponse:
		err = json.Unmarshal([]byte(body), c)
	default:
		return errors.New(ErrorInvalidResponseType)
	}

	return err
}

func getRequest(url string) ([]byte, error) {
	var err error

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, err
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
