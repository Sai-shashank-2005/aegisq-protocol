package consensus

/*
ValidatorSet manages authorized block-producing validators.

Security Purpose:
- Controls who is allowed to produce blocks
- Prevents unauthorized block injection
- Enforces governance layer (Layer 6)

Structure:
NodeID → PublicKey
*/

type ValidatorSet struct {
	validators map[string][]byte
}

// NewValidatorSet initializes an empty validator registry.
func NewValidatorSet() *ValidatorSet {
	return &ValidatorSet{
		validators: make(map[string][]byte),
	}
}

// AddValidator registers a new validator.
// If NodeID already exists, it overwrites the public key.
func (v *ValidatorSet) AddValidator(nodeID string, publicKey []byte) {
	v.validators[nodeID] = publicKey
}

// RemoveValidator removes a validator from the set.
func (v *ValidatorSet) RemoveValidator(nodeID string) {
	delete(v.validators, nodeID)
}

// IsAuthorized checks if a validator is registered AND
// that the provided public key matches the registered key.
func (v *ValidatorSet) IsAuthorized(nodeID string, publicKey []byte) bool {
	registeredKey, exists := v.validators[nodeID]
	if !exists {
		return false
	}

	// Constant-time comparison is better in production,
	// but for now simple equality is acceptable.
	if string(registeredKey) != string(publicKey) {
		return false
	}

	return true
}

// GetValidator returns the registered public key and existence flag.
func (v *ValidatorSet) GetValidator(nodeID string) ([]byte, bool) {
	key, exists := v.validators[nodeID]
	return key, exists
}

// Count returns total number of registered validators.
func (v *ValidatorSet) Count() int {
	return len(v.validators)
}