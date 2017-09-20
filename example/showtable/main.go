// nettable creates a "summary" table from a Telemetry message.
package main

import (
	"flag"
	"log"
	"os"

	nt "github.com/nleiva/nettable"
)

func check(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {
	var container nt.Parser
	file := flag.String("f", "../../input/isis-int.json", "Input file")
	info := flag.String("i", "isis-int", "Type of information")
	flag.Parse()

	switch *info {
	case "isis-int":
		container = new(nt.Iints)
	case "isis-nbr":
		container = new(nt.Inbrs)
	case "isis-lsp":
		container = new(nt.ILSPs)
	case "int-count":
		container = new(nt.Icounts)
	case "int-rate":
		container = new(nt.Irates)
	default:
		log.Fatalf("Type of info unknown: %s", *info)
	}
	f, err := os.Open(*file)
	check(err, "Could not open the file")

	err = nt.ParseTelemetry(container, f, os.Stdout)
	check(err, "Error Parsing Telemetry data")
}
