package client

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_extractNewsFromHTML(t *testing.T) {
	rawHtml, err := ioutil.ReadFile("./response.html")
	require.NoError(t, err)

	loc := time.UTC
	events, err := parseHTML(rawHtml, loc)
	require.NoError(t, err)

	var expect = []EconomicCalendarItem{
		{
			Date:      time.Date(2021, 2, 21, 23, 50, 0, 0, loc),
			TimeLeft:  "2 days",
			Name:      "Corporate Service Price Index (YoY)",
			Impact:    "Low",
			Currency:  "JPY",
			Previous:  "-0.4%",
			Consensus: "-1.1%",
		},
		{
			Date:      time.Date(2021, 2, 22, 1, 30, 0, 0, loc),
			TimeLeft:  "2 weeks",
			Name:      "PBoC Interest Rate Decision",
			Impact:    "High",
			Currency:  "CNY",
			Previous:  "3.85%",
			Consensus: "",
		},
	}
	assert.Equal(t, expect, events)
}
