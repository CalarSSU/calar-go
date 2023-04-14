package converter

import "time"

const firstSemesterMonthBegin = time.September
const semesterDayBegin = 1
const semesterDayEnd = 31
const firstSemesterMonthEnd = time.January
const secondSemesterMonthBegin = time.February
const secondSemesterMonthEnd = time.May
const timeZone = "Europe/Saratov"
const translatorSubstr = "перевод."

func Contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
