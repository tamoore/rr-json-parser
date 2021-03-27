package roam

import (
	"encoding/json"
	"io"

	"github.com/tamoore/rr-json-parser/internal/helpers"
)

type RoamNodes []RoamNode

func (r *RoamNodes) Decode(file io.ReadSeeker) (err error) {
	err = json.NewDecoder(file).Decode(r)

	if helpers.AssertNoErrors(err) {
		return nil
	}

	return err
}

type RoamNode struct {
	Title    string
	String   string
	Children []RoamNode
	Uid      string
}
