package validation

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

/*
testdata.go provides small helpers for loading canonical fixtures and creating
minimal negative-case mutations without duplicating large fixture files.
*/

// LoadFixtureMap loads a canonical JSON fixture into a mutable map.
func LoadFixtureMap(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read fixture %q: %w", path, err)
	}

	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("decode fixture %q: %w", path, err)
	}

	return payload, nil
}

// DeepCopyMap performs a JSON round-trip deep copy suitable for test mutation.
func DeepCopyMap(in map[string]any) (map[string]any, error) {
	data, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("marshal deep copy source: %w", err)
	}

	var out map[string]any
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("unmarshal deep copy destination: %w", err)
	}

	return out, nil
}

// MustWriteTempJSON writes a payload to a temp JSON file for validator tests.
func MustWriteTempJSON(dir, name string, payload map[string]any) string {
	path := filepath.Join(dir, name)

	data, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		panic(err)
	}

	return path
}