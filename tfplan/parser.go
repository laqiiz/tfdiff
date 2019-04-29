package tfplan

import (
	"bufio"
	"bytes"
	"strings"
)

// Parse for console output of terraform plan.
type parser struct{}

func NewParser() *parser {
	return &parser{}
}

type Diff struct {
	Type string // +, -, -/+
	Name string // e.g. google_compute_instance.test-instance-1
}

func (p *parser) Do(planOut []byte) ([]Diff, error) {

	diffs := make([]Diff, 0)

	scanner := bufio.NewScanner(bytes.NewReader(planOut))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Index(line, "+") == 0 || strings.Index(line, "-") == 0 ||
			strings.Index(line, "-/+") == 0 {

			items := strings.Split(line, " ")
			if len(items) == 1 {
				// skip bar "------------------" output
				continue
			}
			diffType := items[0]
			name := strings.Join(items[1:], " ")

			if name == "create" || name == "destroy" || name == "destroy and then create replacement" {
				continue
			}

			diff := Diff{Type: diffType, Name: name}
			diffs = append(diffs, diff)
		}
	}

	return diffs, nil
}
