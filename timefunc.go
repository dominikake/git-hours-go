package main

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"time"
)

// DTF is Default Time format
const DTF string = "2006-01-02"

var timeFormat = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var iso8601 = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} [+-]\d{4}$`)
var findISO8601 = regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} [+-]\d{4}`)
var rfc3339 = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[+-]\d{2}:\d{2}$`)

func timeZoneOffset() string {
	_, offset := time.Now().Zone()
	if offset > 0 {
		return fmt.Sprintf("+%04d", offset/60/60*100)
	}
	return fmt.Sprintf("-%04d", (-1*offset)/60/60*100)
}

// ISO8601ToRFC3339 converts time format to RFC3339 format.
func ISO8601ToRFC3339(t string) (string, error) {
	if !iso8601.MatchString(t) {
		return t, errors.New("time string is not ISO8601 format")
	}
	return fmt.Sprintf("%sT%s%s:%s", t[0:10], t[11:19], t[20:23], t[23:25]), nil
}

func beforeMonth() (string, string) {
	y, m, _ := time.Now().Date()
	if m == 1 {
		y--
		m = 12
	} else {
		m--
	}
	since := time.Date(y, m, 1, 0, 0, 0, 0, time.Now().Location())
	return fmt.Sprintf(since.Format(DTF)), fmt.Sprintf(since.AddDate(0, 1, 0).Add(-time.Nanosecond).Format(DTF))
}

func thisMonth() (string, string) {
	y, m, _ := time.Now().Date()
	since := time.Date(y, m, 1, 0, 0, 0, 0, time.Now().Location())
	return fmt.Sprintf(since.Format(DTF)), fmt.Sprintf(since.AddDate(0, 1, 0).Add(-time.Nanosecond).Format(DTF))
}

// getFirstAndLatestCommitDates retrieves the first and latest commit dates from the git log.
// It returns the dates in YYYY-MM-DD format.
func getFirstAndLatestCommitDates() (string, string, error) {
	// Get the first commit date
	cmdFirst := exec.Command("git", "log", "--reverse", "--format=%cd", "--date=iso-local")
	outputFirst, err := cmdFirst.Output()
	if err != nil {
		return "", "", fmt.Errorf("failed to get first commit date: %w", err)
	}
	firstCommitFull := string(outputFirst)
	firstCommitDate := ""
	if len(firstCommitFull) >= 10 {
		firstCommitDate = firstCommitFull[0:10]
	} else {
		return "", "", errors.New("could not parse first commit date")
	}

	// Get the latest commit date
	cmdLatest := exec.Command("git", "log", "-1", "--format=%cd", "--date=iso-local")
	outputLatest, err := cmdLatest.Output()
	if err != nil {
		return "", "", fmt.Errorf("failed to get latest commit date: %w", err)
	}
	latestCommitFull := string(outputLatest)
	latestCommitDate := ""
	if len(latestCommitFull) >= 10 {
		latestCommitDate = latestCommitFull[0:10]
	} else {
		return "", "", errors.New("could not parse latest commit date")
	}

	return firstCommitDate, latestCommitDate, nil
}
