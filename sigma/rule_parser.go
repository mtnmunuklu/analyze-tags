package sigma

import (
	"gopkg.in/yaml.v3"
)

type Rule struct {
	Title string

	Tags []string `yaml:",omitempty" json:",omitempty"`
}

func ParseRule(input []byte) (Rule, error) {

	rule := Rule{}

	err := yaml.Unmarshal(input, &rule)

	return rule, err
}
