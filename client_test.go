package client

import (
	"fmt"
	"testing"
)

func TestGetLoginURL(t *testing.T) {
	var testCases = []struct {
		login    string
		password string
		expect   string
	}{
		{
			"test@test.ru",
			"pass",
			"https://www.myfxbook.com/api/login.json?email=test@test.ru&password=pass",
		},
		{
			"test@test.ru",
			"",
			"https://www.myfxbook.com/api/login.json?email=test@test.ru&password=",
		},
	}

	for _, tc := range testCases {
		result := getLoginURL(tc.login, tc.password)

		if result != tc.expect {
			t.Errorf("getLoginURL(%s, %s) expect %s got %s", tc.login, tc.password, tc.expect, result)
		}
	}
}

func TestGetLogoutURL(t *testing.T) {
	var testCases = []struct {
		session string
		expect  string
	}{
		{
			"",
			"https://www.myfxbook.com/api/logout.json?session=",
		},
		{
			"something-session-string",
			"https://www.myfxbook.com/api/logout.json?session=something-session-string",
		},
	}

	for _, tc := range testCases {
		result := getLogoutURL(tc.session)

		if result != tc.expect {
			t.Errorf("getLogoutURL(%s) expect %s got %s", tc.session, tc.expect, result)
		}
	}
}

func TestGetGetMyAccountsURL(t *testing.T) {
	var testCases = []struct {
		session string
		expect  string
	}{
		{
			"",
			"https://www.myfxbook.com/api/get-my-accounts.json?session=",
		},
		{
			"something-session-string",
			"https://www.myfxbook.com/api/get-my-accounts.json?session=something-session-string",
		},
	}

	for _, tc := range testCases {
		result := getGetMyAccountsURL(tc.session)
		if result != tc.expect {
			t.Errorf("getGetMyAccountsURL(%s) expect %s got %s", tc.session, tc.expect, result)
		}
	}
}

func TestLoadEconomicCalendar(t *testing.T) {
	client := NewClient("e.mitroshin@space307.com", "asdQWE123")
	items, _ := client.FetchEconomicCalendar(timeFromString("2018-01-08 00:00"), timeFromString("2018-01-14 00:00"))
	fmt.Printf("%+v", items)
}
