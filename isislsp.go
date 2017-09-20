package nettable

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/nleiva/tlv"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// ILSPs contains a list of IS-IS LSP info
type ILSPs struct {
	list []*iLSP
}

type iLSP struct {
	localID   string
	area      string
	hostname  string
	teID      net.IP
	ipAddr    net.IP
	neighbors []neighbor
	prefixes  []prefix
}

func (d *ILSPs) Read(r io.Reader) error {
	data := new(ISISLSP)
	err := decodeTelemetry(data, r)
	if err != nil {
		return errors.Wrap(err, "error decoding JSON file")
	}
	for _, b := range data.Rows {
		i := new(iLSP)
		err = readBytes([]byte(b.Content.LspBody), i)
		if err != nil {
			return errors.Wrap(err, "error reading bytes")
		}
		d.list = append(d.list, i)
	}
	return nil
}

// DisplayTable pretty prints the info populated.
func (d *ILSPs) DisplayTable(w io.Writer) {
	var data [][]string
	for _, s := range d.list {
		for _, n := range s.neighbors {
			data = append(data, []string{s.localID, s.hostname, n.remoteID, fmt.Sprint(n.metric)})
		}
	}
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Local ID", "Hostname", "Remote ID", "Metric"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func readBytes(src []byte, d *iLSP) error {
	// Decode base64 data into bytes
	var enc = base64.NewEncoding(encodeStd)
	dst := make([]byte, enc.DecodedLen(len(src)))
	n, err := enc.Decode(dst, src)
	if err != nil {
		return errors.Wrap(err, "error decoding file")
	}
	// Read PDU info
	r, err := readHeader(dst, n, d)
	if err != nil {
		return errors.Wrap(err, "error reading header")
	}
	tlvs, err := tlv.Read(r)
	if err != nil {
		return errors.Wrap(err, "failed to read TLVs: ")
	}
	err = exploreTLV(tlvs.GetThemAll(), d)
	if err != nil {
		return errors.Wrap(err, "failed to read TLV details: ")
	}
	return nil
}

func read222(t []byte, d *iLSP) error {
	if len(t) < 13 {
		return fmt.Errorf("not a valid TLV, lenght: %v", len(t))
	}
	var mt uint16
	var subtlv uint8
	var nsel uint8
	s := make([]byte, 6)
	m := make([]byte, 3)

	buf := bytes.NewReader(t)
	err := binary.Read(buf, binary.BigEndian, &mt)
	if err != nil {
		return errors.Wrap(err, "failed to read MT ID")
	}
	err = binary.Read(buf, binary.BigEndian, &s)
	if err != nil {
		return errors.Wrap(err, "failed to read System ID")
	}
	err = binary.Read(buf, binary.BigEndian, &nsel)
	if err != nil {
		return errors.Wrap(err, "failed to read NSAP selector")
	}
	err = binary.Read(buf, binary.BigEndian, &m)
	if err != nil {
		return errors.Wrap(err, "failed to read Metric")
	}
	err = binary.Read(buf, binary.BigEndian, &subtlv)
	if err != nil {
		return errors.Wrap(err, "failed to read subTLV")
	}
	mtid := "Unknown"
	switch mt {
	case 2:
		mtid = "IPv6 Unicast"
	default:
	}
	// Metric has three bytes
	met := uint32(m[0])*65536 + uint32(m[1])*256 + uint32(m[2])
	// Format the System ID
	sysid := fmt.Sprintf("%x", (s[0:2])) + "." + fmt.Sprintf("%x", (s[2:4])) + "." + fmt.Sprintf("%x", (s[4:6]))

	n := neighbor{
		remoteID: sysid,
		metric:   met,
		mtid:     mtid,
	}
	d.neighbors = append(d.neighbors, n)
	if subtlv != 0 {
		// TODO Process the sub-TLV
		// return mtid, fmt.Errorf("missed a subTLV")
	}
	return nil
}

func read237(t []byte, d *iLSP) error {
	if len(t) < 2 {
		return fmt.Errorf("not a valid TLV, lenght: %v", len(t))
	}
	// switch binary.BigEndian.Uint16(t[0:2]) {
	// case 2:
	// 	mtid := "IPv6 Unicast"
	// default:
	// 	mtid := "Unknown"
	// }
	return readPrefix(bytes.NewReader(t[2:]), d)
}

func readPrefix(buf *bytes.Reader, d *iLSP) (err error) {
	if buf.Len() == 0 {
		return err
	}
	if buf.Len() <= 6 {
		return fmt.Errorf("not a valid Prefix, lenght: %v", buf.Len())
	}
	var mask, flags, slen uint8
	var metric uint32

	err = binary.Read(buf, binary.BigEndian, &metric)
	if err != nil {
		return errors.Wrap(err, "failed to read Metric")
	}
	err = binary.Read(buf, binary.BigEndian, &flags)
	if err != nil {
		return errors.Wrap(err, "failed to read sub-TLV")
	}
	err = binary.Read(buf, binary.BigEndian, &mask)
	if err != nil {
		return errors.Wrap(err, "failed to read MAsk")
	}
	prfx := make([]byte, mask/8)
	err = binary.Read(buf, binary.BigEndian, &prfx)
	if err != nil {
		return errors.Wrap(err, "failed to read Prefix")
	}
	// Pad with additional bytes for IPv6 address compliance
	pad := make([]byte, 16-mask/8)
	prfx = append(prfx, pad...)

	// Check if subtlv present flag is on
	if flags&(1<<5) != 0 {
		// TODO Process the sub-TLV
		err = binary.Read(buf, binary.BigEndian, &slen)
		subtlv := make([]byte, int(slen))
		err = binary.Read(buf, binary.BigEndian, &subtlv)
		if err != nil {
			return errors.Wrap(err, "failed to read sub-TLV")
		}
	}
	d.prefixes = append(d.prefixes, prefix{
		ip:     string(prfx),
		mask:   mask,
		metric: metric,
	})
	return readPrefix(buf, d)
}

func readHeader(h []byte, n int, d *iLSP) (buf *bytes.Reader, err error) {
	if len(h) < 15 {
		return buf, fmt.Errorf("not a valid Header, lenght: %v", len(h))
	}
	sysid := fmt.Sprintf("%x", (h[0:2])) + "." + fmt.Sprintf("%x", (h[2:4])) + "." + fmt.Sprintf("%x", (h[4:6]))
	d.localID = sysid
	// Get a io.Reader from a []byte slice
	buf = bytes.NewReader(h[15:])
	return buf, err
}

func exploreTLV(ts []tlv.TLV, d *iLSP) error {
	for _, tl := range ts {
		switch tl.Type() {
		case 1:
			a := fmt.Sprintf("%x.%x.%x", tl.Value()[1:2], tl.Value()[2:4], tl.Value()[4:6])
			d.area = a
		case 137:
			d.hostname = string(tl.Value())
		case 140:
			d.teID = net.IP(tl.Value())
		case 232:
			d.ipAddr = net.IP(tl.Value())
		case 222:
			err := read222(tl.Value()[:tl.Length()], d)
			if err != nil {
				return errors.Wrap(err, "failed to read TLV 222")
			}
		case 237:
			err := read237(tl.Value()[:tl.Length()], d)
			if err != nil {
				return errors.Wrap(err, "failed to read TLV 237")
			}
		default:
			// fmt.Printf("Type%03d,  L%03d: %#x\n", tl.Type(), tl.Length(), tl.Value())
		}
	}
	return nil
}
