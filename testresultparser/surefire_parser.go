package testresultparser

import (
	"encoding/xml"
)

// SurefireParser parses Maven Surefire test results into TestResults structure
type SurefireParser struct {}

// Parse method does the actual Surefire XML Report parsing
func (p *SurefireParser) Parse(filepath string) (*TestResults, error) {
	var t *report
	b := readFile(filepath)
	if err := xml.Unmarshal(b, &t); err != nil {
		return nil, err
	}
	testResult := convertTestResultFromJUnitFormat(t)
	return &testResult, nil
}


