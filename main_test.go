package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	tt := []struct {
		name string
		file string
		err  string
	}{
		{name: "isis-int", file: "input/isis-int.json"},
		{name: "isis-nbr", file: "input/isis-nbr.json"},
		{name: "int-count", file: "input/int-count.json"},
		{name: "int-rate", file: "input/int-rate.json"},
	}
	var container Parser

	for _, tc := range tt {
		switch tc.name {
		case "isis-int":
			container = new(Iints)
		case "isis-nbr":
			container = new(Inbrs)
		case "int-count":
			container = new(Icounts)
		case "int-rate":
			container = new(Irates)
		default:
		}
		err := parseTelemetry(container, &tc.file)
		if err != nil {
			t.Fatalf("could not parse: %s, file: %s", tc.name, tc.file)
		}
	}

}
