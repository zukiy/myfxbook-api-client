package client

import (
	"time"
)

// Response type
type Response struct {
	Error   bool
	Message string
}

// LoginResponse type
type LoginResponse struct {
	Response
	Session string
}

// GetMyAccountsResponse type
type GetMyAccountsResponse struct {
	Response
	Accounts []Account
}

// Account type
type Account struct {
	ID             uint64
	Name           string
	Description    string
	AccountID      uint64
	Gain           float64
	AbsGain        float64
	Daily          float64
	Monthly        float64
	Withdrawals    int
	Deposits       float64
	Interest       float64
	Profit         float64
	Balance        float64
	Drawdown       float64
	Equity         float64
	EquityPercent  int
	Demo           bool
	LastUpdateDate string
	CreationDate   string
	FirstTradeDate string
	Tracking       int
	Views          int
	Commission     float64
	Currency       string
	ProfitFactor   float64
	Pips           float64
	InvitationURL  string
	Server         Server
}

// Server for Account struct
type Server struct {
	Name string
}

// EconomicCalendarItem model
type EconomicCalendarItem struct {
	Date     time.Time
	Name     string
	Impact   string
	Previous string
	Actual   string
	Currency string
}
