package nettable

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// Parser  is the interface that wraps the methods to read a
// Telemetry message and transform it to a summary table.
type Parser interface {
	Read(r io.Reader) error
	DisplayTable(w io.Writer)
}

// ParseTelemetry reads the Telemetry input and outputs to a
// table
func ParseTelemetry(p Parser, r io.Reader, w io.Writer) error {
	err := p.Read(r)
	if err != nil {
		return errors.Wrap(err, "error reading file")
	}
	p.DisplayTable(w)
	return nil
}

func decodeTelemetry(v interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(v)
}
