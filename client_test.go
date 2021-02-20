package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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
	client := NewCalendar("", "", time.UTC)
	items, err := client.FetchCalendarItems(timeFromString("2021-02-26 00:00"), timeFromString("2021-02-26 00:00"))
	require.NoError(t, err)
	fmt.Printf("%d, %+v", len(items), items)
}
