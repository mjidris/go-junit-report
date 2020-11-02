package junit

import (
	"encoding/xml"
	"fmt"
	"time"
)

// testsuites is a collection of JUnit testsuites.
type testsuites struct {
	XMLName xml.Name `xml:"testsuites"`

	Name     string `xml:"name,attr,omitempty"`
	Time     string `xml:"time,attr,omitempty"` // total duration in seconds
	Tests    int    `xml:"tests,attr,omitempty"`
	Errors   int    `xml:"errors,attr,omitempty"`
	Failures int    `xml:"failures,attr,omitempty"`
	Disabled int    `xml:"disabled,attr,omitempty"`

	Suites []testsuite `xml:"testsuite,omitempty"`
}

// addSuite adds a Testsuite and updates this testssuites' totals.
func (t *testsuites) addSuite(ts testsuite) {
	t.Suites = append(t.Suites, ts)
	t.Tests += ts.Tests
	t.Errors += ts.Errors
	t.Failures += ts.Failures
	t.Disabled += ts.Disabled
}

// testsuite is a single JUnit testsuite containing testcases.
type testsuite struct {
	// required attributes
	Name  string `xml:"name,attr"`
	Tests int    `xml:"tests,attr"`

	// optional attributes
	Disabled  int    `xml:"disabled,attr,omitempty"`
	Errors    int    `xml:"errors,attr"`
	Failures  int    `xml:"failures,attr"`
	Hostname  string `xml:"hostname,attr,omitempty"`
	ID        int    `xml:"id,attr"`
	Package   string `xml:"package,attr,omitempty"`
	Skipped   int    `xml:"skipped,attr,omitempty"`
	Time      string `xml:"time,attr"`                // duration in seconds
	Timestamp string `xml:"timestamp,attr,omitempty"` // date and time in ISO8601

	Properties []property `xml:"properties>property,omitempty"`
	Testcases  []testcase `xml:"testcase,omitempty"`
	SystemOut  string     `xml:"system-out,omitempty"`
	SystemErr  string     `xml:"system-err,omitempty"`
}

func (t *testsuite) addProperty(name, value string) {
	t.Properties = append(t.Properties, property{Name: name, Value: value})
}

func (t *testsuite) addTestcase(tc testcase) {
	t.Testcases = append(t.Testcases, tc)
	t.Tests++

	if tc.Error != nil {
		t.Errors++
	}

	if tc.Failure != nil {
		t.Failures++
	}
}

func (t *testsuite) setTimestamp(u time.Time) {
	t.Timestamp = u.Format(time.RFC3339)
}

// testcase represents a single test with its results.
type testcase struct {
	// required attributes
	Name      string `xml:"name,attr"`
	Classname string `xml:"classname,attr"`

	// optional attributes
	Time   string `xml:"time,attr,omitempty"` // duration in seconds
	Status string `xml:"status,attr,omitempty"`

	Skipped   *exception `xml:"skipped,omitempty"`
	Error     *exception `xml:"error,omitempty"`
	Failure   *exception `xml:"failure,omitempty"`
	SystemOut string     `xml:"system-out,omitempty"`
	SystemErr string     `xml:"system-err,omitempty"`
}

// property represents a key/value pair.
type property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// exception represents the results of a test that was failed, was skipped
// or had an error.
type exception struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr,omitempty"`
	Data    string `xml:",chardata"`
}

// formatDuration returns the JUnit string representation of the given
// duration.
func formatDuration(d time.Duration) string {
	return fmt.Sprintf("%.3f", d.Seconds())
}

// formatBenchmarkTime returns the JUnit string representation of the given
// benchmark time.
func formatBenchmarkTime(d time.Duration) string {
	return fmt.Sprintf("%.9f", d.Seconds())
}
