package client

import (
	"bytes"
	"errors"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	tdDate int = iota
	tdTimeLeft
	_ // country's flag
	tdCurrency
	tdEvent
	tdImpact
	tdPrevious
	tdConsensus
)

var (
	// ErrEventsDateNoFound ...
	ErrEventsDateNoFound = errors.New("event's time not found ")

	// ErrEventsDateParseError ...
	ErrEventsDateParseError = errors.New("event's time parse error")

	// ErrEventsDateIsZero ...
	ErrEventsDateIsZero = errors.New("event's time is zero")
)

func parseHTML(htmlText []byte, loc *time.Location) ([]EconomicCalendarItem, error) {
	var err error

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlText))
	if err != nil {
		return nil, err
	}

	selection := doc.Find("table#economicCalendarTable tr.economicCalendarRow")
	events := make([]EconomicCalendarItem, 0, len(selection.Nodes))

	selection.Each(func(i int, s *goquery.Selection) {
		e, err := parseEvent(s, loc)
		if err != nil {
			return
		}

		events = append(events, e)
	})

	return events, err
}

func parseEvent(s *goquery.Selection, loc *time.Location) (EconomicCalendarItem, error) {
	var (
		err error
		m   EconomicCalendarItem
	)

	s.Find("td").Each(func(i int, selection *goquery.Selection) {
		switch i {
		case tdDate:
			if m.Date, err = extractEventDate(selection, loc); err != nil {
				return
			}

		case tdTimeLeft:
			m.TimeLeft = extractEventLeftTime(selection)

		case tdCurrency:
			m.Currency = extractText(selection)

		case tdEvent:
			m.Name = extractText(selection)

		case tdImpact:
			m.Impact = extractImpact(selection)

		case tdPrevious:
			m.Previous = extractPrevious(selection)

		case tdConsensus:
			m.Consensus = extractConsensus(selection)
		}
	})

	return m, err
}

func extractEventDate(s *goquery.Selection, loc *time.Location) (time.Time, error) {
	ds, ok := s.Find("div.calendarDateTd").Attr("data-calendardatetd")
	if !ok {
		return time.Time{}, ErrEventsDateNoFound
	}

	dt, err := time.ParseInLocation("2006-01-02 15:04:05.9", ds, time.UTC)
	if err != nil {
		return time.Time{}, ErrEventsDateParseError
	}

	if dt.IsZero() {
		return dt, ErrEventsDateIsZero
	}

	return dt.In(loc), nil
}

func extractEventLeftTime(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("span.calendarLeft").Text())
}

func extractText(s *goquery.Selection) string {
	return strings.TrimSpace(s.Text())
}

func extractImpact(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("div").Text())
}

func extractPrevious(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("span").Text())
}

func extractConsensus(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("div").Text())
}
