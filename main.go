package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/tj/docopt"
)

type Error struct {
	LineNumber  int    `json:"line_number"`
	ErrorString string `json:"error_string"`
}

type Summary struct {
	Name        string  `json:"string"`
	Description string  `json:"description"`
	Errors      []Error `json:"errors"`
}

type Metric struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Summaries   map[string]Summary `json:"summaries"`
	Weight      float32            `json:"weight"`
	Percentage  float32            `json:"percentage"`
	Error       string             `json:"error"`
}

type Report struct {
	Project string            `json:"project"`
	Score   int               `json:"score"`
	Grade   int               `json:"grade"`
	Metrics map[string]Metric `json:"metrics"`
}

type Config struct {
	Exclude  []string
	Filename string
}

func main() {
	usage := `Scrivener - Pretty print goreport results

Usage:
  scrivener <json-file>  [--except=<package>...]
  scrivener -h | --help
  scrivener --version

Options:
  -h --help          Show this screen.
  --version          Show version.
  --except=<name>... Exclude packages.`

	arguments, _ := docopt.Parse(usage, nil, true, "Scrivener", false)

	config := &Config{}
	config.Filename = arguments["<json-file>"].(string)
	config.Exclude = arguments["--except"].([]string)

	report, err := unmarshalReport(config)

	if err != nil {
		fmt.Printf("Error reading configuration: %v\n", err)
		os.Exit(1)
	}

	writeTable(report, config)
}

func unmarshalReport(config *Config) (*Report, error) {
	bytes, err := ioutil.ReadFile(config.Filename)

	if err != nil {
		return nil, err
	}

	results := Report{}
	err = json.Unmarshal(bytes, &results)

	if err != nil {
		return nil, err
	}

	return &results, nil
}

func writeTable(report *Report, config *Config) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Test", "Summary", "Error"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)

External:
	for metric, result := range report.Metrics {
		for _, exclusion := range config.Exclude {
			if exclusion == metric {
				continue External
			}
		}

		for test, summary := range result.Summaries {
			for _, err := range summary.Errors {
				table.Append([]string{metric, test, err.ErrorString})
			}
		}
	}

	table.Render()
}
