package main

import (
	"calar-go/converter"
	"calar-go/parser"
	"calar-go/tracto"
	"fmt"
	"log"
	"os"

	ics "github.com/arran4/golang-ical"
)

func main() {
	var schedule tracto.Schedule
	var cfg parser.Config
	parser.ParseArguments(&cfg)
	tracto.ParseJson(&schedule, cfg)
	fmt.Printf("%d\n", len(cfg.Subgroups))
	for _, s := range cfg.Subgroups {
		fmt.Printf("%s\n", s)
	}
	iCalString := converter.MakeCalendar(schedule, &ics.Calendar{})

	file, err :=
		os.OpenFile(fmt.Sprintf("%s_%s.ics", cfg.Department, cfg.Group),
			os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file.WriteString(iCalString)
	file.Close()

}
