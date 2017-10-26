package client

const (
	baseURL = "https://www.myfxbook.com/"

	// calendar
	userLoginURL = baseURL + "login.html?loginEmail=%s&loginPassword=%s&remember=false&z=0.4554484698084875&locale=ru"
	calendarURL  = baseURL + "calendar_statement.xml?"

	// api
	apiURL           = baseURL + "api/"
	loginURL         = "login.json?email=%s&password=%s"
	logoutURL        = "logout.json?session=%s"
	getMyAccountsURL = "get-my-accounts.json?session=%s"

	// ErrorInvalidResponseType error text
	ErrorInvalidResponseType = "invalid response type"
)
