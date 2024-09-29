package spotify

import (
	"strconv"
	"strings"
	"time"
)

func (c *Client) ReleaseDateTime(ReleaseDatePrecision, ReleaseDate string) *time.Time {
	if ReleaseDatePrecision == "day" {
		result, _ := time.Parse("2006-01-02", ReleaseDate)
		return &result
	}
	if ReleaseDatePrecision == "month" {
		ym := strings.Split(ReleaseDate, "-")
		year, _ := strconv.Atoi(ym[0])
		month, _ := strconv.Atoi(ym[1])

		date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

		return &date
	}
	year, _ := strconv.Atoi(ReleaseDate)

	date := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	return &date
}
