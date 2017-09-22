package nettable

import (
	"fmt"
	"io"
	"net"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

// Bnbrs contains a list of BGP neighbors
type Bnbrs struct {
	list []*bnbr
}

type bnbr struct {
	hostname    string
	af          string
	remoteAddr  net.IP
	remoteAS    uint32
	remotePort  uint32
	state       string
	description string
	pfxsRecv    uint32
	maxPrf      uint32
	blocal
}

type blocal struct {
	hostname  string
	localAS   uint32
	localAddr net.IP
	localPort uint32
	pfxsAdv   uint32
}

func (d *Bnbrs) Read(r io.Reader) error {
	data := new(BGPNbr)
	err := decodeTelemetry(data, r)
	if err != nil {
		return errors.Wrap(err, "error decoding JSON file")
	}
	for _, b := range data.Rows {
		i := new(bnbr)
		i.hostname = data.Telemetry.NodeIDStr
		i.af = b.Keys.AfName
		c := b.Content
		// ASN
		i.remoteAS = c.RemoteAs
		i.localAS = c.LocalAs
		// IPv6 Address
		i.remoteAddr = c.ConnectionRemoteAddress.Ipv6Address.Value
		i.localAddr = c.ConnectionLocalAddress.Ipv6Address.Value
		// TCP Connection port
		i.remotePort = c.ConnectionRemotePort
		i.localPort = c.ConnectionLocalPort
		i.state = c.ConnectionState
		i.description = c.Description
		// Prefixes Advertised/Received
		af := c.AfData
		if len(af) > 0 {
			for _, n := range af {
				if n.Value.AfName != "" {
					i.pfxsRecv = n.Value.PrefixesAccepted
					i.pfxsAdv = n.Value.PrefixesAdvertised
					i.maxPrf = n.Value.MaxPrefixLimit
				}
			}
		}
		d.list = append(d.list, i)
	}
	return nil
}

// DisplayTable pretty prints the info populated.
func (d *Bnbrs) DisplayTable(w io.Writer) {
	var data [][]string
	for _, s := range d.list {
		data = append(data, []string{s.hostname, s.remoteAddr.String(), fmt.Sprint(s.remoteAS),
			s.state, fmt.Sprint(s.pfxsRecv), fmt.Sprint(s.pfxsAdv), fmt.Sprint(s.maxPrf)})

	}
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"hostname", "neighbor", "asn", "state", "pfx rcvd", "pfx adv", "max pfx"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
