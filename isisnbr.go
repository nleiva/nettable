package nettable

import (
	"io"
	"net"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

// Inbrs contains a list of IS-IS neighbors
type Inbrs struct {
	list []*inbr
}

type inbr struct {
	remoteID  string
	fwAddress net.IP
	local
}

type local struct {
	hostname string
	intName  string
	area     string
}

func (d *Inbrs) Read(r io.Reader) error {
	data := new(ISISNbr)
	err := decodeTelemetry(data, r)
	if err != nil {
		return errors.Wrap(err, "error decoding JSON file")
	}
	for _, b := range data.Rows {
		i := new(inbr)
		// Hostname
		i.hostname = data.Telemetry.NodeIDStr
		// Interface Name
		i.intName = b.Content.LocalInterface
		// Area
		i.area = b.Content.NeighborActiveAreaAddresses[0].Value
		// Remote SystemID
		i.remoteID = b.Content.NeighborSystemID
		af := b.Content.NeighborPerAddressFamilyData
		if len(af) > 0 {
			// FW Address neighbor
			i.fwAddress = af[0].Ipv6.NextHop
		}
		d.list = append(d.list, i)
	}
	return nil
}

// DisplayTable pretty prints the info populated.
func (d *Inbrs) DisplayTable(w io.Writer) {
	var data [][]string
	for _, s := range d.list {
		data = append(data, []string{s.hostname, s.intName, s.area,
			s.remoteID, s.fwAddress.String()})
	}
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"hostname", "interface", "area", "remote id", "FW address"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
