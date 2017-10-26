package client

// MyFxBookClient struct
type MyFxBookClient struct {
	Email    string
	Password string
	Session  string
}

// NewClient create new myfxbook client
func NewClient(email, password string) *MyFxBookClient {
	return &MyFxBookClient{
		Email:    email,
		Password: password,
	}
}

// Login to account
func (c *MyFxBookClient) Login() error {
	var err error

	loginResponseContainer := new(LoginResponse)
	err = callGetRequest(getLoginURL(c.Email, c.Password), loginResponseContainer)
	if err != nil {
		return err
	}

	c.Session = loginResponseContainer.Session

	return err
}

// Logout from account
func (c *MyFxBookClient) Logout() error {
	var err error

	logoutResponseContainer := new(Response)
	err = callGetRequest(getLogoutURL(c.Session), logoutResponseContainer)
	if err != nil {
		return err
	}

	return err
}
