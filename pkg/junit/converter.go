package junit

import (
	"fmt"
	"strings"
	"time"
)

var (
	propPrefixes   = map[string]bool{"goos": true, "goarch": true, "pkg": true}
	propFieldsFunc = func(r rune) bool { return r == ':' || r == ' ' }
)

// This can be it's own package
// JUnit converts the given report to a collection of JUnit Testsuites.
func convert(report Report) testsuites {
	var suites testsuites
	for _, pkg := range report.Packages {
		var duration time.Duration
		suite := testsuite{Name: pkg.Name}

		if pkg.Coverage > 0 {
			suite.addProperty("coverage.statements.pct", fmt.Sprintf("%.2f", pkg.Coverage))
		}

		for _, line := range pkg.Output {
			if fields := strings.FieldsFunc(line, propFieldsFunc); len(fields) == 2 && propPrefixes[fields[0]] {
				suite.addProperty(fields[0], fields[1])
			}
		}

		for _, test := range pkg.Tests {
			duration += test.Duration

			tc := testcase{
				Classname: pkg.Name,
				Name:      test.Name,
				Time:      formatDuration(test.Duration),
			}

			if test.Result == Fail {
				tc.Failure = &exception{
					Message: "Failed",
					Data:    formatOutput(test.Output, test.Level),
				}
			} else if test.Result == Skip {
				tc.Skipped = &exception{
					Message: formatOutput(test.Output, test.Level),
				}
			}

			suite.addTestcase(tc)
		}

		for _, bm := range mergeBenchmarks(pkg.Benchmarks) {
			tc := testcase{
				Classname: pkg.Name,
				Name:      bm.Name,
				Time:      formatBenchmarkTime(time.Duration(bm.NsPerOp)),
			}

			if bm.Result == Fail {
				tc.Failure = &exception{
					Message: "Failed",
				}
			}

			suite.addTestcase(tc)
		}

		// JUnit doesn't have a good way of dealing with build or runtime
		// errors that happen before a test has started, so we create a single
		// failing test that contains the build error details.
		if pkg.BuildError.Name != "" {
			tc := testcase{
				Classname: pkg.BuildError.Name,
				Name:      pkg.BuildError.Cause,
				Time:      formatDuration(0),
				Failure: &exception{
					Message: "Failed",
					Data:    strings.Join(pkg.BuildError.Output, "\n"),
				},
			}
			suite.addTestcase(tc)
		}

		if pkg.RunError.Name != "" {
			tc := testcase{
				Classname: pkg.RunError.Name,
				Name:      "Failure",
				Time:      formatDuration(0),
				Failure: &exception{
					Message: "Failed",
					Data:    strings.Join(pkg.RunError.Output, "\n"),
				},
			}
			suite.addTestcase(tc)
		}

		if (pkg.Duration) == 0 {
			suite.Time = formatDuration(duration)
		} else {
			suite.Time = formatDuration(pkg.Duration)
		}
		suites.addSuite(suite)
	}
	return suites
}

func formatOutput(output []string, level int) string {
	var lines []string
	for _, line := range output {
		lines = append(lines, trimOutputPrefix(line, level))
	}
	return strings.Join(lines, "\n")
}

func trimOutputPrefix(line string, level int) string {
	// We only want to trim the whitespace prefix if it was part of the test
	// output. Test output is usually prefixed by a series of 4-space indents,
	// so we'll check for that to decide whether this output was likely to be
	// from a test.
	prefixLen := strings.IndexFunc(line, func(r rune) bool { return r != ' ' })
	if prefixLen%4 == 0 {
		// Use the subtest level to trim a consistenly sized prefix from the
		// output lines.
		for i := 0; i <= level; i++ {
			line = strings.TrimPrefix(line, "    ")
		}
	}
	return strings.TrimPrefix(line, "\t")
}

func mergeBenchmarks(benchmarks []Benchmark) []Benchmark {
	var merged []Benchmark

	benchmap := make(map[string][]Benchmark)
	for _, bm := range benchmarks {
		if _, ok := benchmap[bm.Name]; !ok {
			merged = append(merged, Benchmark{Name: bm.Name})
		}
		benchmap[bm.Name] = append(benchmap[bm.Name], bm)
	}

	for i, bm := range merged {
		for _, b := range benchmap[bm.Name] {
			bm.NsPerOp += b.NsPerOp
			bm.MBPerSec += b.MBPerSec
			bm.BytesPerOp += b.BytesPerOp
			bm.AllocsPerOp += b.AllocsPerOp
		}
		n := len(benchmap[bm.Name])
		merged[i].NsPerOp = bm.NsPerOp / float64(n)
		merged[i].MBPerSec = bm.MBPerSec / float64(n)
		merged[i].BytesPerOp = bm.BytesPerOp / int64(n)
		merged[i].AllocsPerOp = bm.AllocsPerOp / int64(n)
	}

	return merged
}
