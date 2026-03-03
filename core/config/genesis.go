package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Genesis struct {
	Validators   []string `json:"validators"`
	validatorMap map[string]struct{}
}

func LoadGenesis(path string) (*Genesis, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var g Genesis
	if err := json.Unmarshal(data, &g); err != nil {
		return nil, err
	}

	if len(g.Validators) == 0 {
		return nil, errors.New("genesis must contain at least one validator")
	}

	g.validatorMap = make(map[string]struct{}, len(g.Validators))
	for _, v := range g.Validators {
		g.validatorMap[v] = struct{}{}
	}

	return &g, nil
}

func (g *Genesis) IsValidator(pubKeyBase64 string) bool {
	_, exists := g.validatorMap[pubKeyBase64]
	return exists
}
