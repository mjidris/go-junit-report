package junit

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarshalUnmarshal(t *testing.T) {
	suites := testsuites{
		Name:     "name",
		Time:     "12.345",
		Tests:    1,
		Errors:   1,
		Failures: 1,
		Disabled: 1,
		Suites: []testsuite{
			{
				Name:       "suite1",
				Tests:      1,
				Errors:     1,
				Failures:   1,
				Hostname:   "localhost",
				ID:         1,
				Package:    "package",
				Skipped:    1,
				Time:       "12.345",
				Timestamp:  "2012-03-09T14:38:06+01:00",
				Properties: []property{{"key", "value"}},
				Testcases: []testcase{
					{
						Name:      "test1",
						Classname: "class",
						Time:      "12.345",
						Status:    "status",
						Skipped:   &exception{Message: "skipped", Type: "type", Data: "data"},
						Error:     &exception{Message: "error", Type: "type", Data: "data"},
						Failure:   &exception{Message: "failure", Type: "type", Data: "data"},
						SystemOut: "system-out",
						SystemErr: "system-err",
					},
				},
				SystemOut: "system-out",
				SystemErr: "system-err",
			},
		},
	}

	data, err := xml.MarshalIndent(suites, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	var unmarshaled testsuites
	if err := xml.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatal(err)
	}

	suites.XMLName.Local = "testsuites"
	if diff := cmp.Diff(suites, unmarshaled); diff != "" {
		t.Errorf("Unmarshal result incorrect, diff (-want +got):\n%s\n", diff)
	}
}
