package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl := template.Must(template.ParseFiles("tmpl/dayn.go"))
	out := bytes.Buffer{}
	dayFlag := flag.Int("d", 0, "Day to add")
	flag.Parse()
	day := *dayFlag
	if day == 0 {
		log.Fatal("Day cannot be zero")
		os.Exit(1)
	}
	var inp struct {
		Day int
	}
	inp.Day = day
	err := tmpl.Execute(&out, inp)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	err = os.Mkdir(fmt.Sprintf("days/%d", day), 0o777)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok {
			log.Fatal(err.Error())
			os.Exit(1)
		}
		log.Println("Directory already exists, continuing")
	}
	err = os.WriteFile(fmt.Sprintf("days/%d/%d.go", day, day), out.Bytes(), 0o644)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
