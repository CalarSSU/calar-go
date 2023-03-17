package main

import (
	"calar-go/converter"
	"calar-go/tracto"
	"fmt"
	"log"
	"os"

	ics "github.com/arran4/golang-ical"
)

func main() {
	var schedule tracto.Schedule
	education := "full"
	department := "knt"
	group := "351"
	tracto.ParseJson(&schedule, education, department, group)
	iCalString := converter.MakeCalendar(schedule, &ics.Calendar{})

	file, err :=
		os.OpenFile(fmt.Sprintf("%s_%s.ics", department, group),
			os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file.WriteString(iCalString)
	file.Close()

}
