package csiem

import "encoding/json"

type Rule struct {
	Name string

	Tags []string `json:",omitempty"`
}

func ParseRule(input []byte) (Rule, error) {

	rule := Rule{}

	err := json.Unmarshal(input, &rule)

	return rule, err
}
