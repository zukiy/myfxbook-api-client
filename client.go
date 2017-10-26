package client

import (
	"time"
)

// MyFxBookClient struct
type MyFxBookClient struct {
	email    string
	password string
	session  string
}

// NewClient create new myfxbook client
func NewClient(email, password string) *MyFxBookClient {
	return &MyFxBookClient{
		email:    email,
		password: password,
	}
}

// Login to account
func (c *MyFxBookClient) Login() (*LoginResponse, error) {
	var err error

	loginResponseContainer := new(LoginResponse)
	err = loadData(getLoginURL(c.email, c.password), loginResponseContainer)
	if err != nil {
		return nil, err
	}

	c.session = loginResponseContainer.Session
	return loginResponseContainer, err
}

// Logout from account
func (c *MyFxBookClient) Logout() (*Response, error) {
	var err error

	logoutResponseContainer := new(Response)
	err = loadData(getLogoutURL(c.session), logoutResponseContainer)
	if err != nil {
		return nil, err
	}

	return logoutResponseContainer, err
}

// GetMyAccounts get a list of my accounts and their data
func (c *MyFxBookClient) GetMyAccounts() (*GetMyAccountsResponse, error) {
	var err error

	getMyAccountsResponse := new(GetMyAccountsResponse)
	err = loadData(getGetMyAccountsURL(c.session), getMyAccountsResponse)
	if err != nil {
		return nil, err
	}

	return getMyAccountsResponse, err
}

// FetchEconomicCalendar load Calendar
func (c *MyFxBookClient) FetchEconomicCalendar(start, end time.Time) (calendarItems []*EconomicCalendarItem, err error) {
	calendar := NewCalendar(c.email, c.password)

	calendarItems, err = calendar.fetchCalendar(start, end)

	if err != nil {
		return
	}

	return
}
