package client

const (
	baseURL = "https://www.myfxbook.com/"

	// calendar
	userLoginURL        = baseURL + "login.json"
	calendarNextPageURL = baseURL + "get-more-calendar-events.json?"

	// api
	apiURL           = baseURL + "api/"
	loginURL         = "login.json?email=%s&password=%s"
	logoutURL        = "logout.json?session=%s"
	getMyAccountsURL = "get-my-accounts.json?session=%s"
)
