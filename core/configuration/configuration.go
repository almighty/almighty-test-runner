package configuration

import (
    "gopkg.in/yaml.v2"
    "github.com/ecooper/combinatoric"
)

type TestRunnerConfiguration struct {
    Description, Path string
    BuildToolConfiguration `yaml:"build_tool,flow"`
    EnvironmentVerticals `yaml:"verticals,inline"`
}

type BuildToolConfiguration struct {
    Command string `yaml:"cmd"`
}

type EnvironmentVerticals struct {
    Verticals map[string][]string `yaml:"verticals,flow"`
}

type VerticalSelection struct {
    Vertical, Selection string
}

func ProductFor(s VerticalSelection, v EnvironmentVerticals) [][]string {
    v.Verticals[s.Vertical] = []string{s.Selection}
    return Product(v)
}

func Product(v EnvironmentVerticals) [][]string {
    values := make([][]interface{}, 0, len(v.Verticals))
    for _, value := range v.Verticals {
        values = append(values, sliceOfInterfaces(value))
    }
    pIter, _ := combinatoric.Product(values)
    products := make([][]string, 0, len(values))
    for pIter.HasNext() {
        products = append(products, sliceOfStrings(pIter.Next()))
    }
    return products
}

func Read(configuration string) (TestRunnerConfiguration, error) {
    var c TestRunnerConfiguration = TestRunnerConfiguration{}
    if err := yaml.Unmarshal([]byte(configuration), &c); err != nil {
        return c, err
    }
    return c, nil
}

func sliceOfInterfaces(s []string) []interface{} {
    result := make([]interface{}, len(s))
    for i, v := range s {
        result[i] = v
    }
    return result
}

func sliceOfStrings(s []interface{}) []string {
    result := make([]string, len(s))
    for i, v := range s {
        result[i] = v.(string)
    }
    return result
}
