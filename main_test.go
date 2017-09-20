package nettable_test

import (
	"os"
	"testing"

	nt "github.com/nleiva/nettable"
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
	var container nt.Parser

	for _, tc := range tt {
		switch tc.name {
		case "isis-int":
			container = new(nt.Iints)
		case "isis-nbr":
			container = new(nt.Inbrs)
		case "int-count":
			container = new(nt.Icounts)
		case "int-rate":
			container = new(nt.Irates)
		default:
		}
		file, err := os.Open(tc.file)
		if err != nil {
			t.Fatalf("could not open the file: %s; %v", tc.file, err)
		}
		err = nt.ParseTelemetry(container, file, os.Stdout)
		if err != nil {
			t.Fatalf("could not parse: %s, file: %s", tc.name, tc.file)
		}
	}

}
