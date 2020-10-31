package testrun

import "time"

type Result int

const (
	Unknown Result = iota
	Pass
	Fail
	Skip
)

// Expose this type.
type Report struct {
	Packages []Package
}

// This function is used to set an error exit code. Unsure if we actually need this.
func (r *Report) HasFailures() bool {
	for _, pkg := range r.Packages {
		for _, t := range pkg.Tests {
			if t.Result == Fail {
				return true
			}
		}
	}
	return false
}

type Package struct {
	Name     string
	Duration time.Duration
	Coverage float64
	Output   []string

	Tests      []Test
	Benchmarks []Benchmark

	BuildError Error
	RunError   Error
}

type Test struct {
	Name     string
	Duration time.Duration
	Result   Result
	Level    int
	Output   []string
}

// Maybe let's get rid of this???
type Benchmark struct {
	Name        string
	Result      Result
	Output      []string
	Iterations  int64
	NsPerOp     float64
	MBPerSec    float64
	BytesPerOp  int64
	AllocsPerOp int64
}

type Error struct {
	Name     string
	Duration time.Duration
	Cause    string
	Output   []string
}
