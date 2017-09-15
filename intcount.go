package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

// Icounts contains a list of interface counters
type Icounts struct {
	list []*icount
}

type icount struct {
	hostname   string
	name       string
	pktsRecv   uint64
	pktsSent   uint64
	bytsRecv   uint64
	bytsSent   uint64
	muPktsRecv uint64
	muPktsSent uint64
	brPktsRecv uint64
	brPktsSent uint64
	inDrops    uint32
	inErrs     uint32
	outDrops   uint32
	outErrs    uint32
	crcErrs    uint32
	carrTrans  uint32
}

func (d *Icounts) Read(file *string) error {
	data := new(IntCount)
	err := decodeTelemetry(data, *file)
	if err != nil {
		return errors.Wrap(err, "error decoding JSON file")
	}
	for _, b := range data.Rows {
		i := new(icount)
		// Hostname
		i.hostname = data.Telemetry.NodeIDStr
		// Interface Name
		i.name = b.Keys.InterfaceName
		// Counters
		intf := b.Content
		i.pktsRecv = intf.PacketsReceived
		i.pktsSent = intf.PacketsSent
		i.bytsRecv = intf.BytesReceived
		i.bytsSent = intf.BytesSent
		i.muPktsRecv = intf.MulticastPacketsReceived
		i.muPktsSent = intf.MulticastPacketsSent
		i.carrTrans = intf.CarrierTransitions
		i.inErrs = intf.InputErrors
		i.outErrs = intf.OutputErrors
		d.list = append(d.list, i)
	}
	return nil
}

// DisplayTable pretty prints the info populated.
func (d *Icounts) DisplayTable() {
	var data [][]string
	for _, s := range d.list {
		data = append(data, []string{s.hostname, s.name, fmt.Sprint(s.pktsRecv),
			fmt.Sprint(s.pktsSent), fmt.Sprint(s.carrTrans), fmt.Sprint(s.inErrs), fmt.Sprint(s.outErrs)})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"hostname", "interface", "pkts sent",
		"pkts recv", "transitions", "in errs", "out errs"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
