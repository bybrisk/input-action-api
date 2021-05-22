package data

import (
	"encoding/json"
	"io"
)	

func (d *OrderAPIResponse) OrderAPIResponseToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *OrderAPIRequest) FromJSONToOrderAPIRequest (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}