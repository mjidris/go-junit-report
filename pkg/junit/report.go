package junit

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

// HasFailures returns true if the report contains any failures.
func (r *Report) HasFailures() bool {
	for _, p := range r.Packages {
		if p.HasFailures() {
			return true
		}
	}
	return false
}

func (r *Report) AddPackage(p Package) {
	r.Packages = append(r.Packages, p)
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

// HasFailures returns true if the package contains any failures.
func (p *Package) HasFailures() bool {
	for _, t := range p.Tests {
		if t.Result == Fail {
			return true
		}
	}
	return false
}

func (p *Package) AddTest(t Test) {
	p.Duration += t.Duration
	p.Tests = append(p.Tests, t)
}

func (p *Package) AddTests(tests ...Test) {
	for _, t := range tests {
		p.AddTest(t)
	}
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
