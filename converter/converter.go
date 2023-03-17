package converter

import (
	"calar-go/tracto"
	"fmt"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
)

func MakeCalendar(schedule tracto.Schedule, cal *ics.Calendar) string {
	loc, _ := time.LoadLocation(timeZone)
	curTime := time.Now().In(loc)

	for _, lesson := range schedule.Lessons {
		event := cal.AddEvent(fmt.Sprintf("%d", lesson.Id))
		summary := fmt.Sprintf("%s: %s", lesson.Name, lesson.LessonType)
		if lesson.Subgroup != "" {
			summary =
				fmt.Sprintf("(%s) %s",
					strings.ReplaceAll(strings.Trim(lesson.Subgroup, " ."),
						" ", "_"),
					summary)
		}
		event.SetSummary(summary)

		teacher :=
			fmt.Sprintf("%s %s %s", lesson.Teacher.Surname,
				lesson.Teacher.Name, lesson.Teacher.Patronymic)
		event.SetDescription(teacher)

		event.SetLocation(lesson.Place)

		var lessonTimeBegin, lessonTimeEnd, semesterBegin time.Time
		if (curTime.Month() >= firstSemesterMonthBegin) ||
			(curTime.Month() <= firstSemesterMonthEnd) {
			semesterBegin =
				semesterBegin.AddDate(0, 0,
					(lesson.Day.Id - int(semesterBegin.Weekday())))

			lessonTimeBegin =
				semesterBegin.Add(time.Hour*
					time.Duration(lesson.LessonTime.HourStart) +
					time.Minute*time.Duration(lesson.LessonTime.MinuteStart))

			lessonTimeEnd =
				semesterBegin.Add(time.Hour*
					time.Duration(lesson.LessonTime.HourEnd) +
					time.Minute*time.Duration(lesson.LessonTime.MinuteEnd))
		} else {
			semesterBegin =
				time.Date(curTime.Year(), secondSemesterMonthBegin,
					semesterDayBegin, 0, 0, 0, 0, curTime.Location())
			semesterBegin =
				semesterBegin.AddDate(0, 0,
					(lesson.Day.Id - int(semesterBegin.Weekday())))

			lessonTimeBegin =
				semesterBegin.Add(time.Hour*
					time.Duration(lesson.LessonTime.HourStart) +
					time.Minute*time.Duration(lesson.LessonTime.MinuteStart))

			lessonTimeEnd =
				semesterBegin.Add(time.Hour*
					time.Duration(lesson.LessonTime.HourEnd) +
					time.Minute*time.Duration(lesson.LessonTime.MinuteEnd))
		}
		event.SetStartAt(lessonTimeBegin)
		event.SetEndAt(lessonTimeEnd)
	}
	return cal.Serialize()
}
