package converter

import (
	"calar-go/parser"
	"calar-go/tracto"
	"fmt"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
)

func MakerRule(event *ics.VEvent, curTime time.Time, lesson tracto.Lesson,
	semesterEnd time.Month, semesterDayEnd int) {
	lastLessonDate := time.Date(curTime.Year(), semesterEnd,
		semesterDayEnd, 23, 59, 59, 0, curTime.Location())
	rRuleEnd := lastLessonDate.UTC().Format("20060102T150405Z")

	var interval int
	switch lesson.WeekType {
	case "FULL":
		interval = 1
	default:
		interval = 2
	}

	rRule := fmt.Sprintf("FREQ=WEEKLY;INTERVAL=%d;UNTIL=%s",
		interval, rRuleEnd)
	event.AddRrule(rRule)

}

func makeLessonTime(event *ics.VEvent, curTime time.Time, lesson tracto.Lesson,
	semesterBegin, semesterEnd time.Month, semesterDayBegin,
	semesterDayEnd int) {
	firstLessonDate := time.Date(curTime.Year(), semesterBegin,
		semesterDayBegin, 0, 0, 0, 0, curTime.Location())

	firstLessonDate =
		firstLessonDate.AddDate(0, 0,
			(lesson.Day.Id - int(firstLessonDate.Weekday())))

	if "DENOM" == lesson.WeekType {
		firstLessonDate = firstLessonDate.Add(time.Hour * 24 * 7)
	}

	lessonTimeBegin :=
		firstLessonDate.Add(time.Hour*
			time.Duration(lesson.LessonTime.HourStart) +
			time.Minute*time.Duration(lesson.LessonTime.MinuteStart))

	lessonTimeEnd :=
		firstLessonDate.Add(time.Hour*
			time.Duration(lesson.LessonTime.HourEnd) +
			time.Minute*time.Duration(lesson.LessonTime.MinuteEnd))

	event.SetStartAt(lessonTimeBegin)
	event.SetEndAt(lessonTimeEnd)

	MakerRule(event, curTime, lesson, semesterEnd, semesterDayEnd)
}
func MakeCalendar(request parser.Request, schedule tracto.Schedule,
	cal *ics.Calendar) string {
	loc, _ := time.LoadLocation(timeZone)
	curTime := time.Now().In(loc)
	for _, lesson := range schedule.Lessons {
		if (len(request.Subgroups) == 0 ||
			Contains(request.Subgroups, lesson.Subgroup) ||
			"" == lesson.Subgroup) &&
			(!strings.Contains(lesson.Name, translatorSubstr) ||
				request.Translator) {
			event := cal.AddEvent(fmt.Sprintf("%d", lesson.Id))
			summary := fmt.Sprintf("%s: %s", lesson.Name, lesson.LessonType)

			event.SetSummary(summary)

			teacher :=
				fmt.Sprintf("%s %s %s", lesson.Teacher.Surname,
					lesson.Teacher.Name, lesson.Teacher.Patronymic)
			event.SetDescription(teacher)

			event.SetLocation(lesson.Place)

			if (curTime.Month() >= firstSemesterMonthBegin) ||
				(curTime.Month() <= firstSemesterMonthEnd) {
				makeLessonTime(event, curTime, lesson, firstSemesterMonthBegin,
					firstSemesterMonthEnd, semesterDayBegin, semesterDayEnd)
			} else {
				makeLessonTime(event, curTime, lesson, secondSemesterMonthBegin,
					secondSemesterMonthEnd, semesterDayBegin, semesterDayEnd)
			}
		}
	}
	calSerialaze := cal.Serialize()
	return calSerialaze
}
