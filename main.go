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

// Parser is a Telemetry message that can be read and transformed to a
// summary table
type Parser interface {
	Read(file *string) error
	DisplayTable()
}

func parseTelemetry(p Parser, f *string) error {
	err := p.Read(f)
	if err != nil {
		return errors.Wrap(err, "error reading file")
	}
	p.DisplayTable()
	return nil
}

func main() {
	var container Parser
	file := flag.String("f", "input/isis-int.json", "Input file")
	info := flag.String("i", "isis-int", "Type of information")
	flag.Parse()

	switch *info {
	case "isis-int":
		container = new(Iints)
	case "isis-nbr":
		container = new(Inbrs)
	case "int-count":
		container = new(Icounts)
	case "int-rate":
		container = new(Irates)
	default:
		log.Fatalf("Type of info unknown: %s", *info)
	}

	err := parseTelemetry(container, file)
	check(err, "Error Parsing Telemetry data")
}
