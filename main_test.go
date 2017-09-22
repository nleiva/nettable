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
		ty   string
		err  string
	}{
		{name: "ISIS Interface 1", ty: "isis-int", file: "input/isis-int.json"},
		{name: "ISIS Neighbor 1", ty: "isis-nbr", file: "input/isis-nbr.json"},
		{name: "ISIS LSP 1", ty: "isis-lsp", file: "input/isis-lsp.json"},
		{name: "BGP Neighbor 1", ty: "bgp-nbr", file: "input/bgp-nbr.json"},
		{name: "Interface Counters 1", ty: "int-count", file: "input/int-count.json"},
		{name: "Interface Data Rates 1", ty: "int-rate", file: "input/int-rate.json"},
		{name: "IPv6 RIB Table", ty: "rib-ipv6", file: "input/rib-ipv6.json"},
		{name: "ISIS Interface 2", ty: "isis-int", file: "input/isis-int2.json"},
		{name: "ISIS Neighbor 2", ty: "isis-nbr", file: "input/isis-nbr2.json"},
		{name: "ISIS LSP 2", ty: "isis-lsp", file: "input/isis-lsp2.json"},
		{name: "BGP Neighbor 2", ty: "bgp-nbr", file: "input/bgp-nbr2.json"},
		{name: "Interface Counters 2", ty: "int-count", file: "input/int-count2.json"},
		{name: "Interface Data Rates 2", ty: "int-rate", file: "input/int-rate2.json"},
		{name: "IPv6 RIB Table 2", ty: "rib-ipv6", file: "input/rib-ipv62.json"},
	}
	var container nt.Parser

	for _, tc := range tt {
		switch tc.ty {
		case "isis-int":
			container = new(nt.Iints)
		case "isis-nbr":
			container = new(nt.Inbrs)
		case "isis-lsp":
			container = new(nt.ILSPs)
		case "bgp-nbr":
			container = new(nt.Bnbrs)
		case "int-count":
			container = new(nt.Icounts)
		case "int-rate":
			container = new(nt.Irates)
		case "rib-ipv6":
			container = new(nt.RIPv6s)
		default:
		}
		file, err := os.Open(tc.file)
		if err != nil {
			t.Fatalf("could not open the file: %s; %v", tc.file, err)
		}
		err = nt.ParseTelemetry(container, file, os.Stdout)
		if err != nil {
			t.Fatalf("could not parse: %s, file: %s: %s", tc.name, tc.file, err)
		}
	}

}
