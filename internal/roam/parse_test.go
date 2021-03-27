package roam

import (
	"strings"
	"testing"

	"github.com/tamoore/rr-json-parser/internal/test"
)

func TestRoamParser(t *testing.T) {
	t.Run("Test that the Roam Parser can decode", func(t *testing.T) {
		roamJSON := strings.NewReader(`
		[
			{
				"create-time": 1616493540844,
				"title": "Doctrine bundle analysis",
				":create/user": {
					":user/uid": "foiG9eb60jZkW9UQkes7Vj7ErVm2"
				},
				"children": [],
				"uid": "SJZgptx93",
				"edit-time": 1616493540845,
				":edit/user": {
					":user/uid": "foiG9eb60jZkW9UQkes7Vj7ErVm2"
				}
			}
		]
		`)

		got := RoamNodes{}
		err := got.Decode(roamJSON)

		want := RoamNodes{
			{"Doctrine bundle analysis", "", []RoamNode{}, "SJZgptx93"},
		}

		AssertRoamNodes(t, got, want)

		test.AssertError(t, err, false)
	})
}
