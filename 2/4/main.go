package main

import (
	"time"
	"strconv"
)

//CalendarStructure
type CalendarStruct struct {
	month   string
	quarter int
}

//Returns Calendar's current quater
func (cs CalendarStruct) CurrentQuarter() int {
	return cs.quarter
}

//Creates new CalendarStruct
func NewCalendar(time time.Time) CalendarStruct {
	month := int64(time.Month())
	monthString := ""
	if month < 10 {
		monthString = "0" + strconv.FormatInt(month, 10)
	} else {
		monthString = strconv.FormatInt(month, 10)
	}
	quoter := 1
	switch {
	case month > 9:
		quoter = 4
	case month > 6:
		quoter = 3
	case month > 3:
		quoter = 2
	}

	return CalendarStruct{month: monthString, quarter: quoter}
}

func main() {

}
