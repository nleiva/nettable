package nettable

import (
	"fmt"
	"io"
	"net"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

// Iints contains a list of IS-IS interfaces
type Iints struct {
	list []*iint
}

type iint struct {
	hostname  string
	name      string
	config    string
	status    string
	fwAddress net.IP
	prefixes  []prefix
}

type neighbor struct {
	remoteID string
	metric   uint32
	mtid     string
}

type prefix struct {
	ip     string
	mask   uint8
	metric uint32
}

func (d *Iints) Read(r io.Reader) error {
	data := new(ISISInt)
	err := decodeTelemetry(data, r)
	if err != nil {
		return errors.Wrap(err, "error decoding JSON file")
	}
	for _, b := range data.Rows {
		i := new(iint)
		// Hostname
		i.hostname = data.Telemetry.NodeIDStr
		// Interface Name
		i.name = b.Keys.InterfaceName
		// Config Status
		i.config = b.Content.InterfaceStatusAndData.Enabled.AdjacencyFormStatus.Status
		af := b.Content.InterfaceStatusAndData.Enabled.PerAddressFamilyData
		if len(af) > 0 {
			// Protocol Status
			afd := af[0].AfStatus.AfData
			i.status = afd.ProtocolStatus.Status
			address := afd.ForwardingAddressStatus.ForwardingAddressData.ForwardingAddress
			if len(address) > 0 {
				// Forwarding Address
				i.fwAddress = address[0].Ipv6.Value
			}
			pfxs := afd.PrefixStatus.PrefixData.Prefix
			if len(pfxs) > 0 {
				// Prefix
				for _, pfx := range pfxs {
					i.prefixes = append(i.prefixes, prefix{
						ip:   pfx.Ipv6.Prefix.String(),
						mask: pfx.Ipv6.PrefixLength,
					})
				}
			}
		}
		d.list = append(d.list, i)
	}
	return nil
}

// DisplayTable pretty prints the info populated.
func (d *Iints) DisplayTable(w io.Writer) {
	var data [][]string
	for _, s := range d.list {
		for _, p := range s.prefixes {
			pf := p.ip + "/" + fmt.Sprint(p.mask)
			data = append(data, []string{s.hostname, s.name, s.config,
				s.status, s.fwAddress.String(), pf})
		}
	}
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"hostname", "interface", "config", "status", "FW address", "Prefix"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
