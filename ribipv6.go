package nettable

import (
	"fmt"
	"io"
	"net"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

// RIPv6s contains a list of RIB entries
type RIPv6s struct {
	list []*ripv6
}

type ripv6 struct {
	hostname      string
	prefix        net.IP
	lenght        uint32
	protocol      string
	adminDistance uint32
	metric        uint32
	paths         uint32
	nxHops        []nxHop
}

type nxHop struct {
	address net.IP
	source  net.IP
}

func (d *RIPv6s) Read(r io.Reader) error {
	data := new(RIBIPv6)
	err := decodeTelemetry(data, r)
	if err != nil {
		return errors.Wrap(err, "error decoding JSON file")
	}
	for _, b := range data.Rows {
		i := new(ripv6)
		i.hostname = data.Telemetry.NodeIDStr
		c := b.Content
		i.prefix = c.Prefix
		i.lenght = c.PrefixLength
		i.protocol = c.ProtocolName
		i.adminDistance = c.Distance
		i.metric = c.Metric
		i.paths = c.PathsCount
		nh := c.RoutePath.Ipv6RibEdmPath
		if len(nh) > 0 {
			for _, n := range nh {
				i.nxHops = append(i.nxHops, nxHop{
					address: n.Address,
					source:  n.InformationSource,
				})
			}
		}
		d.list = append(d.list, i)
	}
	return nil
}

// DisplayTable pretty prints the info populated.
func (d *RIPv6s) DisplayTable(w io.Writer) {
	var data [][]string
	for _, s := range d.list {
		for _, p := range s.nxHops {
			pf := s.prefix.String() + "/" + fmt.Sprint(s.lenght)
			data = append(data, []string{s.hostname, pf, s.protocol,
				p.address.String(), p.address.String(), fmt.Sprint(s.metric)})
		}
	}
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"hostname", "prefix", "protocol", "next Hop", "source", "metric"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
