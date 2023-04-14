package tracto

const scribaToken = "https://scribaproject.space/api/v1.0/schedule"

type Day struct {
	Id        int
	DayNumber int
	Weekday   string
}

type LessonTime struct {
	Id           int
	LessonNumber int
	HourStart    int
	MinuteStart  int
	HourEnd      int
	MinuteEnd    int
}

type Department struct {
	Id        int
	FullName  string
	ShortName string
	Url       string
}

type Teacher struct {
	Id         int
	Surname    string
	Name       string
	Patronymic string
}

type StudentGroup struct {
	Id             int
	GroupNumber    string
	GroupNumberRus string
	Department     Department
	EducationForm  string
	GroupType      string
}

type Lesson struct {
	Id               int
	Name             string
	Place            string
	Department       Department
	StudentGroup     StudentGroup
	Subgroup         string
	Day              Day
	LessonTime       LessonTime
	Teacher          Teacher
	WeekType         string
	LessonType       string
	UpdatedTimestamp int
}

type Schedule struct {
	Lessons      []Lesson
	StudentGroup StudentGroup
}
