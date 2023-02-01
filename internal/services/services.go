package services

import (
	"strconv"
	"time"
)

func Calculator(year string) int {
	intYear, _ := strconv.Atoi(year)

	now := time.Now()

	then := time.Date(
		intYear, 1, 1, 0, 0, 0, 0, time.UTC)

	diff := now.Sub(then)

	result := diff.Hours() / 24

	return int(result)
}
