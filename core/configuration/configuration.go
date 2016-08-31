package configuration

import (
    "gopkg.in/yaml.v2"
    "github.com/ecooper/combinatoric"
    "sort"
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

func ProductFor(s VerticalSelection, v EnvironmentVerticals) ([][]string, error) {
    v.Verticals[s.Vertical] = []string{s.Selection}
    return Product(v)
}

func Product(v EnvironmentVerticals) ([][]string, error) {
    values := make([][]interface{}, 0, len(v.Verticals))
    for key, value := range v.Verticals {
        values = append(values, verticalEntryOf(key, value))
    }

    pIter, err := combinatoric.Product(values);

    if err != nil {
        return [][]string{}, err
    }

    products := make([][]string, 0, len(values))
    for pIter.HasNext() {
        products = append(products, flatten(pIter.Next()))
    }
    return products, nil
}

func Read(configuration string) (TestRunnerConfiguration, error) {
    c := TestRunnerConfiguration{}
    err := yaml.Unmarshal([]byte(configuration), &c)
    return c, err
}

func verticalEntryOf(key string, s []string) []interface{} {
    result := make([]interface{}, len(s))
    for i, v := range s {
        result[i] = verticalEntry{key, v}
    }
    return result
}


func flatten(s []interface{}) []string {
    entries := make([]verticalEntry, len(s))
    for i, v := range s {
        entries[i] = v.(verticalEntry)
    }
    sort.Sort(byVerticalName(entries))
    result := make([]string, len(entries))
    for i, v := range entries {
        result[i] = v.value;
    }
    return result
}

type verticalEntry struct {
    vertical, value string
}

type byVerticalName []verticalEntry

func (slice byVerticalName) Len() int {
    return len(slice)
}

func (slice byVerticalName) Less(i, j int) bool {
    return slice[i].vertical < slice[j].vertical
}

func (slice byVerticalName) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}