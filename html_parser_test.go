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

	expect := []EconomicCalendarItem{
		{
			Date:     time.Date(2021, 2, 21, 23, 50, 0, 0, loc),
			Name:     "Corporate Service Price Index (YoY)",
			Impact:   "Low",
			Currency: "JPY",
		},
		{
			Date:     time.Date(2021, 2, 22, 1, 30, 0, 0, loc),
			Name:     "PBoC Interest Rate Decision",
			Impact:   "High",
			Currency: "CNY",
		},
	}

	assert.Equal(t, expect, events)
}
