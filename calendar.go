package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type (
	// XMLAuthResponse model
	XMLAuthResponse struct {
		Error   bool   `xml:"error,attr"`
		Code    string `xml:"code,attr"`
		Message string `xml:"message,attr"`
	}

	// NextPageResponse model
	NextPageResponse struct {
		Error   bool    `json:"error"`
		Content Content `json:"content"`
	}

	// Content
	Content struct {
		Rows             []string `json:"rows"`
		PageNumber       int      `json:"pageNumber"`
		HasMore          bool     `json:"hasMore"`
		NewCnt           int      `json:"newCnt"`
		LastDay          string   `json:"lastDay"`
		FutureEventFound int      `json:"futureEventFound"`
	}

	// Calendar engine
	Calendar struct {
		loc      *time.Location
		email    string
		password string
		client   *http.Client
		cookie   []*http.Cookie
	}
)

// NewCalendar creates a economic calendar client, https://www.myfxbook.com/forex-economic-calendar
func NewCalendar(email, password string, loc *time.Location) *Calendar {
	return &Calendar{
		email:    email,
		password: password,
		client:   &http.Client{},
		loc:      loc,
	}
}

// NewCalendarWithHTTPClient creates a economic calendar client by HTTP client
func NewCalendarWithHTTPClient(email, password string, loc *time.Location, client *http.Client) *Calendar {
	return &Calendar{
		email:    email,
		password: password,
		client:   client,
		loc:      loc,
	}
}

func (c *Calendar) login() error {
	var err error

	form := map[string]string{
		"loginEmail":    c.email,
		"loginPassword": c.password,
		"remember":      "false",
		"z":             "0.5537966003240857",
	}

	headers := map[string]string{
		"Origin":       "https://www.myfxbook.com",
		"Referer":      "https://www.myfxbook.com",
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
	}

	response := XMLAuthResponse{}

	err, cookie := c.postRequest(userLoginURL, form, headers, &response)
	if err != nil {
		return err
	}

	if response.Error {
		return errors.Wrap(err, "login error")
	}

	// replacing the cookies slice
	c.cookie = append(c.cookie[:0], cookie...)

	return err
}

// FetchCalendarItems economic calendar's items
func (c *Calendar) FetchCalendarItems(start, end time.Time) (calendarItems []EconomicCalendarItem, err error) {
	err = c.login()
	if err != nil {
		return
	}

	data, err := c.loadData(start, end)
	if err != nil {
		return
	}

	calendarItems, err = parseHTML(data, c.loc)
	return
}

func (c *Calendar) loadData(start, end time.Time) (data []byte, err error) {
	var (
		page int
		buf  bytes.Buffer
	)

	buf.WriteString(`<table id="economicCalendarTable">`)

	for {
		payload := map[string]string{
			"pageNumber":       strconv.Itoa(page),
			"filter":           "0-1-2-3_ANG-ARS-AUD-BRL-CAD-CHF-CLP-CNY-COP-CZK-DKK-EEK-EUR-GBP-HKD-HUF-IDR-INR-ISK-JPY-KPW-KRW-MXN-NOK-NZD-PEI-PLN-QAR-ROL-RUB-SEK-SGD-TRY-USD-ZAR",
			"start":            timeToString(start),
			"end":              timeToString(end),
			"cnt":              "40",
			"futureEventFound": "1",
			"isMobile":         "false",
			"z":                "0.8090088713380079",
			"lastDay":          end.Format("Monday, Jan 02, 2006"),
		}

		headers := map[string]string{
			"referer": "https://www.myfxbook.com/",
		}

		response := NextPageResponse{}

		if err = c.getRequest(calendarNextPageURL, payload, headers, &response); err != nil {
			return nil, err
		}

		buf.Write([]byte(strings.Join(response.Content.Rows, ",")))

		if !response.Content.HasMore {
			break
		}

		page++
	}

	buf.WriteString(`</table>`)

	return buf.Bytes(), nil
}

func (c *Calendar) postRequest(u string, payload, headers map[string]string, response interface{}) (error, []*http.Cookie) {
	data := url.Values{}
	for k, v := range payload {
		data.Set(k, v)
	}

	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader(data.Encode()))
	if err != nil {
		return err, nil
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	for _, cookie := range c.cookie {
		req.AddCookie(cookie)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err, nil
	}

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	err = xml.Unmarshal(raw, response)
	if err != nil {
		return err, nil
	}

	return err, resp.Cookies()
}

func (c *Calendar) getRequest(u string, payload, headers map[string]string, response interface{}) error {
	params := url.Values{}
	for k, v := range payload {
		params.Set(k, v)
	}

	req, err := http.NewRequest(http.MethodGet, u+params.Encode(), nil)
	if err != nil {
		return err
	}

	for _, cookie := range c.cookie {
		req.AddCookie(cookie)
	}

	for k, v := range headers {
		req.Header[k] = []string{v}
	}

	for _, cookie := range c.cookie {
		req.AddCookie(cookie)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if response == nil {
		return nil
	}

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer func() {
		err = resp.Body.Close()
	}()

	err = json.Unmarshal(raw, response)
	return err
}
