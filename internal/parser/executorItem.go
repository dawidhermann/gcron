package parser

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	Minute int = iota
	Hour
	Dom
	Month
	Dow
)

// Validation
var (
	minuteBounds = Bounds{0, 59}
	hourBounds   = Bounds{0, 23}
	domBounds    = Bounds{0, 31}
	monthBounds  = Bounds{1, 12}
	dowBounds    = Bounds{0, 6}
)

type CronTask struct {
	//Schedule              time.Time
	PreviousExecutionTime time.Time
	NextExecutionTime     time.Time
	Task                  string
}

type ExecutorItem struct {
	ExecutionTime time.Time
	Task          interface{}
}

type Bounds struct {
	Min, Max uint
}

func ParseEntry(entry string) (CronTask, error) {
	r, _ := regexp.Compile("(.\\S*)")
	matchResult := r.FindAllString(entry, -1)
	if matchResult == nil {
		return CronTask{}, errors.New("no matches")
	}
	minute, err := strconv.Atoi(matchResult[Minute])
	if err != nil {
		return CronTask{}, errors.New("could not parse minutes")
	}
	hour, err := strconv.Atoi(matchResult[Hour])
	if err != nil {
		return CronTask{}, errors.New("could not parse hours")
	}
	dayMonth, err := strconv.Atoi(matchResult[Dom])
	if err != nil {
		return CronTask{}, errors.New("could not parse day of month")
	}
	month, err := strconv.Atoi(matchResult[Month])
	if err != nil {
		return CronTask{}, errors.New("could not parse month")
	}
	//dayWeek, err := strconv.Atoi(matchResult[Dow])
	//if err != nil {
	//	return ExecutorItem{}, errors.New("could not parse day of week")
	//}
	task := strings.Join(matchResult[5:6], " ")
	now := time.Now()
	executionTime := time.Date(now.Year(), time.Month(month), dayMonth, hour, minute, 0, 0, time.UTC)
	return CronTask{Task: task, NextExecutionTime: executionTime, PreviousExecutionTime: now}, nil
}
