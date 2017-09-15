// nettable creates a "summary" table from a Telemetry message.
package main

import (
	"flag"
	"log"

	"github.com/pkg/errors"
)

func check(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

type parser interface {
	Read(file *string) error
	displayTable()
}

func parseTelemetry(p parser, f *string) error {
	err := p.Read(f)
	if err != nil {
		return errors.Wrap(err, "error reading file")
	}
	p.displayTable()
	return nil
}

func main() {
	var container parser
	file := flag.String("f", "input/isis-int.json", "Input file")
	info := flag.String("i", "isis-int", "Type of information")
	flag.Parse()

	switch *info {
	case "isis-int":
		container = new(intfs)
	case "isis-nbr":
		container = new(nbrs)
	default:
		log.Fatalf("Type of info unknown: %s", *info)
	}

	err := parseTelemetry(container, file)
	check(err, "Error Parsing Telemetry data")
}
