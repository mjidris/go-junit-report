package junit

import "time"

// Result represents the outcome of a test (eg. Pass, Fail, etc.).
type Result int

const (
	// Unknown represents a result which is unknown.
	Unknown Result = iota
	// Pass represents a successful test result.
	Pass
	// Fail represents an unsuccessful test result.
	Fail
	// Skip represents a test which was not executed.
	Skip
)

// Report represents a collection of packages containing tests.
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

// AddPackage adds a package of tests to the report.
func (r *Report) AddPackage(p Package) {
	r.Packages = append(r.Packages, p)
}

// Package is a collection of tests.
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

// AddTest adds a single test to the package.
func (p *Package) AddTest(t Test) {
	p.Duration += t.Duration
	p.Tests = append(p.Tests, t)
}

// AddTests add multiple tests to the package.
func (p *Package) AddTests(tests ...Test) {
	for _, t := range tests {
		p.AddTest(t)
	}
}

// Test represents a single test case.
type Test struct {
	Name     string
	Duration time.Duration
	Result   Result
	Level    int
	Output   string
}

// Benchmark represents a single benchmark test case.
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

// Error represents an error that occured during the test.
type Error struct {
	Name     string
	Duration time.Duration
	Cause    string
	Output   []string
}
