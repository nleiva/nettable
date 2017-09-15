package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

type irates struct {
	list []*irate
}

type irate struct {
	hostname    string
	name        string
	bw          uint32
	inDataRate  uint64
	outDataRate uint64
	inPktRate   uint64
	outPktRate  uint64
}

func (d *irates) Read(file *string) error {
	data := new(IntRate)
	err := decodeTelemetry(data, *file)
	if err != nil {
		return errors.Wrap(err, "error decoding JSON file")
	}
	for _, b := range data.Rows {
		i := new(irate)
		// Hostname
		i.hostname = data.Telemetry.NodeIDStr
		// Interface Name
		i.name = b.Keys.InterfaceName
		// Data Rate
		intf := b.Content
		i.bw = intf.Bandwidth
		i.inDataRate = intf.InputDataRate
		i.outDataRate = intf.OutputDataRate
		i.inPktRate = intf.InputPacketRate
		i.outPktRate = intf.OutputPacketRate
		d.list = append(d.list, i)
	}
	return nil
}

func (d *irates) DisplayTable() {
	var data [][]string
	for _, s := range d.list {
		data = append(data, []string{s.hostname, s.name, fmt.Sprint(s.inDataRate),
			fmt.Sprint(s.outDataRate), fmt.Sprint(s.bw)})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"hostname", "interface", "in data rate",
		"out data rate", "bw"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
