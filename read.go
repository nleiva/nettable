package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func decodeTelemetry(v interface{}, f string) error {
	file, err := os.Open(f)
	if err != nil {
		return fmt.Errorf("could not open the file: %s; %v", f, err)
	}
	return json.NewDecoder(file).Decode(v)
}
