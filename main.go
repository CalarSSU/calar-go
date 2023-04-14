package main

import (
	"calar-go/converter"
	"calar-go/parser"
	"calar-go/tracto"
	"fmt"
	"log"
	"os"
	"strings"

	ics "github.com/arran4/golang-ical"
)

func main() {
	var schedule tracto.Schedule
	var request parser.Request
	err := parser.ParseArguments(&request)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	tracto.ParseJson(&schedule, request)
	iCalString := converter.MakeCalendar(&request, &schedule, &ics.Calendar{})

	var isTranslator string
	if request.Translator {
		isTranslator = "translator"
	}

	file, err :=
		os.OpenFile(
			fmt.Sprintf("%s_%s_%s_%s.ics", request.Department, request.Group,
				strings.Join(request.Subgroups, "_"), isTranslator),
			os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file.WriteString(iCalString)
	file.Close()

}
