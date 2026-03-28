package validation

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"runtime"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

/*
Package validation provides the approved Go schema-validation entrypoint for
brikbyteos-schema.

Phase 0 goals:
  - validate owned JSON fixtures against owned JSON schemas
  - provide one stable validator surface for tests
  - avoid duplicated validation logic across the suite
  - keep behavior deterministic and easy to extend
*/

// ValidationIssue is a stable, test-facing representation of one validation issue.
type ValidationIssue struct {
	// InstanceLocation points to the failing location in the JSON instance.
	InstanceLocation string

	// Message is the human-readable validation message.
	Message string
}

// Result is the normalized output of one validation run.
type Result struct {
	// Valid reports whether validation succeeded.
	Valid bool

	// Issues contains normalized validation issues when validation fails.
	Issues []ValidationIssue
}

// Validator validates JSON payloads against schema files.
type Validator struct{}

// New creates a new stateless Validator.
func New() *Validator {
	return &Validator{}
}

// ValidateFile validates a JSON file against a JSON Schema file.
func (v *Validator) ValidateFile(schemaPath, payloadPath string) (Result, error) {
	payloadBytes, err := os.ReadFile(payloadPath)
	if err != nil {
		return Result{}, fmt.Errorf("read payload file %q: %w", payloadPath, err)
	}

	return v.ValidateBytes(schemaPath, payloadBytes)
}

// ValidateBytes validates raw JSON bytes against a JSON Schema file.
func (v *Validator) ValidateBytes(schemaPath string, payload []byte) (Result, error) {
	schemaAbs, err := filepath.Abs(schemaPath)
	if err != nil {
		return Result{}, fmt.Errorf("resolve schema path %q: %w", schemaPath, err)
	}

	compiler := jsonschema.NewCompiler()

	// Compile from the actual schema file location so the validator has a stable
	// base URI for root/fragments and any future relative $ref resolution.
	schema, err := compiler.Compile(schemaAbs)
	if err != nil {
		return Result{}, fmt.Errorf("compile schema %q: %w", schemaPath, err)
	}

	var document any
	if err := json.Unmarshal(payload, &document); err != nil {
		return Result{}, fmt.Errorf("decode payload json: %w", err)
	}

	if err := schema.Validate(document); err != nil {
		return Result{
			Valid:  false,
			Issues: extractIssues(err),
		}, nil
	}

	return Result{
		Valid:  true,
		Issues: nil,
	}, nil
}

// extractIssues normalizes validator errors into a stable test-facing shape.
func extractIssues(err error) []ValidationIssue {
	ve, ok := err.(*jsonschema.ValidationError)
	if !ok {
		return []ValidationIssue{
			{
				InstanceLocation: "<root>",
				Message:          err.Error(),
			},
		}
	}

	var issues []ValidationIssue
	flattenValidationError(ve, &issues)

	if len(issues) == 0 {
		issues = append(issues, ValidationIssue{
			InstanceLocation: "<root>",
			Message:          ve.Error(),
		})
	}

	return issues
}

// flattenValidationError recursively flattens nested validator errors.
func flattenValidationError(err *jsonschema.ValidationError, issues *[]ValidationIssue) {
	if err == nil {
		return
	}

	location := formatInstanceLocation(err.InstanceLocation)

	*issues = append(*issues, ValidationIssue{
		InstanceLocation: location,
		Message:          err.Error(),
	})

	for _, cause := range err.Causes {
		flattenValidationError(cause, issues)
	}
}

// formatInstanceLocation converts the validator's instance-location slice into a
// stable, human-readable path string for tests and diagnostics.
func formatInstanceLocation(parts []string) string {
	if len(parts) == 0 {
		return "<root>"
	}

	cleaned := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		cleaned = append(cleaned, part)
	}

	if len(cleaned) == 0 {
		return "<root>"
	}

	return "/" + strings.Join(cleaned, "/")
}

// ValidateNormalizedResultV01 validates raw JSON bytes against the canonical
// normalized-result v0.1 schema.
//
// This helper gives tests and future consumers a stable, intention-revealing
// entrypoint instead of requiring each caller to hard-code schema paths.
func (v *Validator) ValidateNormalizedResultV01(payload []byte) error {
	result, err := v.ValidateBytes(normalizedResultV01SchemaPath(), payload)
	if err != nil {
		return err
	}

	if result.Valid {
		return nil
	}

	return validationIssuesError(
		"normalized-result v0.1 validation failed",
		result.Issues,
	)
}

// ValidateNormalizedResultV01File validates a JSON file against the canonical
// normalized-result v0.1 schema.
func (v *Validator) ValidateNormalizedResultV01File(payloadPath string) (Result, error) {
	return v.ValidateFile(normalizedResultV01SchemaPath(), payloadPath)
}

// normalizedResultV01SchemaPath returns the absolute filesystem path to the
// canonical normalized-result v0.1 schema.
//
// The path is resolved relative to this source file rather than the current
// working directory so tests remain stable across packages.
func normalizedResultV01SchemaPath() string {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to resolve validator source path")
	}

	validationDir := filepath.Dir(thisFile)
	repoRoot := filepath.Clean(filepath.Join(validationDir, "..", ".."))

	return filepath.Join(
		repoRoot,
		"schemas",
		"normalized-result",
		"v0.1",
		"schema.json",
	)
}

// validationIssuesError turns normalized validation issues into one stable error.
// This keeps call sites simple while preserving useful diagnostics.
func validationIssuesError(prefix string, issues []ValidationIssue) error {
	if len(issues) == 0 {
		return fmt.Errorf("%s", prefix)
	}

	first := issues[0]
	return fmt.Errorf("%s: %s at %s", prefix, first.Message, first.InstanceLocation)
}