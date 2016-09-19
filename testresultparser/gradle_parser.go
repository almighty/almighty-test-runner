package testresultparser

import "encoding/xml"

// GradleParser parses Gradle test results into TestResults structure
type GradleParser struct {}

// Parse method does the actual gradle XML Report parsing
func (p *GradleParser) Parse(filepath string) (*TestResults, error) {
    var t *report
    b := readFile(filepath)
    if err := xml.Unmarshal(b, &t); err != nil {
        return nil, err
    }
    testResult := convertTestResultFromJUnitFormat(t)
    return &testResult, nil
}


