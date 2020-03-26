package client

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Calendar api
type Calendar struct {
	email    string
	password string
	client   *http.Client
	cookie   []*http.Cookie
}

// NewCalendar return Calendar instance
func NewCalendar(email, password, proxy string) *Calendar {
	return &Calendar{
		email:    email,
		password: password,
		client:   createHTTPClient(proxy),
	}
}

func (c *Calendar) reLogin() (err error) {
	req, err := http.NewRequest(http.MethodGet, getSystemLoginURL(c.email, c.password), nil)
	if err != nil {
		return
	}

	req.Header.Set("Origin", "https://www.myfxbook.com")
	req.Header.Set("Referer", "https://www.myfxbook.com")

	rawResponse, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer rawResponse.Body.Close()

	response, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return
	}

	// Login responses
	type XML struct {
		Error   bool   `xml:"error,attr"`
		Code    string `xml:"code,attr"`
		Message string `xml:"message,attr"`
	}

	x := &XML{}
	err = xml.Unmarshal(response, x)
	if err != nil {
		return
	}

	if x.Error {
		err = errors.New(x.Message)
		return
	}

	c.cookie = append(c.cookie, rawResponse.Cookies()...)

	return
}

func (c *Calendar) fetchCalendar(start, end time.Time) (calendarItems []*EconomicCalendarItem, err error) {
	calendarItems = make([]*EconomicCalendarItem, 0)

	data, err := c.loadData(start, end)
	if err != nil {
		return
	}

	calendarItems, err = c.parseXML(data)
	if err != nil {
		return
	}

	return
}

func (c *Calendar) loadData(start, end time.Time) (data []byte, err error) {
	err = c.reLogin()
	if err != nil {
		return
	}

	p := map[string]string{
		"start":     timeToString(start),
		"end":       timeToString(end),
		"calPeriod": "-1",
		"filter":    "0-1-2-3_ANG-ARS-AUD-BRL-CAD-CHF-CLP-CNY-COP-CZK-DKK-EEK-EUR-GBP-HKD-HUF-IDR-INR-ISK-JPY-KPW-KRW-MXN-NOK-NZD-PEI-PLN-QAR-ROL-RUB-SEK-SGD-TRY-USD-ZAR",
	}

	m := make([]string, 0)
	for key, value := range p {
		m = append(m, fmt.Sprintf("%s=%s", key, value))
	}

	uri := strings.Join(m, "&")
	t := &url.URL{Path: uri}

	requestURL := fmt.Sprintf("%s%s", calendarURL, t.String()[2:])

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return
	}

	for _, cookie := range c.cookie {
		req.AddCookie(cookie)
	}

	response, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	return
}

func (c *Calendar) parseXML(data []byte) (content []*EconomicCalendarItem, err error) {
	type Events struct {
		Event []*EconomicCalendarItem `xml:"event"`
	}

	type Response struct {
		Error   bool   `xml:"error,attr"`
		Message string `xml:"message,attr"`
		Events  Events `xml:"events"`
	}

	x := Response{}
	err = xml.Unmarshal(data, &x)
	content = x.Events.Event

	return
}
