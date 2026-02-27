package config

import (
	"encoding/json"
	"errors"
	"os"
)

// Genesis defines initial validator trust root.
type Genesis struct {
	Validators []string `json:"validators"`
}

// LoadGenesis loads genesis configuration from file.
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

	return &g, nil
}

// IsValidator checks if provided public key (base64) is authorized.
func (g *Genesis) IsValidator(pubKeyBase64 string) bool {
	for _, v := range g.Validators {
		if v == pubKeyBase64 {
			return true
		}
	}
	return false
}