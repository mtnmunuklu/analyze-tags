package sigma

import (
	"gopkg.in/yaml.v3"
)

type Rule struct {
	// Required fields
	Title string // The title of the Sigma rule

	// Optional fields
	Tags []string `yaml:",omitempty" json:",omitempty"` // Tags that can be used to organize the rules

}

// ParseRule reads a byte slice and returns a parsed Rule object and an error (if any)
func ParseRule(input []byte) (Rule, error) {
	// Create a Rule instance to hold the parsed YAML data
	rule := Rule{}

	// Unmarshal the input YAML data into the Rule instance
	err := yaml.Unmarshal(input, &rule)

	// Return the Rule instance and error (if any)
	return rule, err
}
