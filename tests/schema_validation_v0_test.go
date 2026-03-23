package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/BrikByte-Studios/brikbyteos-schema/internal/validation"
)

/*
schema_validation_v0_test.go defines the Phase 0 Go validation baseline for the
current schema families.

Coverage rules:
  - one positive case per schema family
  - one negative case per schema family
  - stable failure assertions using useful characteristics rather than brittle
    full-message equality
*/

type schemaCase struct {
	name             string
	schemaPath       string
	fixturePath      string
	invalidMutator   func(map[string]any)
	expectedHint     string
	expectedPathHint string
}

func TestSchemaValidationV0_PositiveFixturesPass(t *testing.T) {
	t.Parallel()

	v := validation.New()

	for _, tc := range schemaCases(t) {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := v.ValidateFile(tc.schemaPath, tc.fixturePath)
			if err != nil {
				t.Fatalf("expected validator execution success, got error: %v", err)
			}

			if !result.Valid {
				t.Fatalf("expected fixture to validate, got issues: %+v", result.Issues)
			}
		})
	}
}

func TestSchemaValidationV0_NegativeCasesFailPredictably(t *testing.T) {
	t.Parallel()

	v := validation.New()

	for _, tc := range schemaCases(t) {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			fixture, err := validation.LoadFixtureMap(tc.fixturePath)
			if err != nil {
				t.Fatalf("expected fixture load success, got error: %v", err)
			}

			mutated, err := validation.DeepCopyMap(fixture)
			if err != nil {
				t.Fatalf("expected deep copy success, got error: %v", err)
			}

			tc.invalidMutator(mutated)

			tempDir := t.TempDir()
			tempPayloadPath := validation.MustWriteTempJSON(tempDir, "invalid.json", mutated)

			result, err := v.ValidateFile(tc.schemaPath, tempPayloadPath)
			if err != nil {
				t.Fatalf("expected validator execution success, got error: %v", err)
			}

			if result.Valid {
				t.Fatal("expected invalid payload to fail validation")
			}

			if len(result.Issues) == 0 {
				t.Fatal("expected one or more validation issues")
			}

			assertIssueContains(t, result.Issues, tc.expectedHint, tc.expectedPathHint)
		})
	}
}

// schemaCases defines the current in-scope v0 schema families and their minimum
// negative-case mutation strategy.
func schemaCases(t *testing.T) []schemaCase {
	t.Helper()

	return []schemaCase{
		{
			name:       "raw-result",
			schemaPath: repoPath(t, "schemas/raw-result/v0/schema.json"),
			fixturePath: repoPath(t, "schemas/raw-result/v0/examples/valid.minimal.json"),
			invalidMutator: func(m map[string]any) {
				delete(m, "schema_family")
			},
			expectedHint:     "required",
			expectedPathHint: "<root>",
		},
		{
			name:       "normalized-result",
			schemaPath: repoPath(t, "schemas/normalized-result/v0/schema.json"),
			fixturePath: repoPath(t, "schemas/normalized-result/v0/examples/valid.minimal.json"),
			invalidMutator: func(m map[string]any) {
				delete(m, "schema_family")
			},
			expectedHint:     "required",
			expectedPathHint: "<root>",
		},
		{
			name:       "policy-evaluation-result",
			schemaPath: repoPath(t, "schemas/policy-evaluation-result/v0/schema.json"),
			fixturePath: repoPath(t, "schemas/policy-evaluation-result/v0/examples/valid.minimal.json"),
			invalidMutator: func(m map[string]any) {
				delete(m, "schema_family")
			},
			expectedHint:     "required",
			expectedPathHint: "<root>",
		},
		{
			name:       "verdict",
			schemaPath: repoPath(t, "schemas/verdict/v0/schema.json"),
			fixturePath: repoPath(t, "schemas/verdict/v0/examples/valid.minimal.json"),
			invalidMutator: func(m map[string]any) {
				delete(m, "schema_family")
			},
			expectedHint:     "required",
			expectedPathHint: "<root>",
		},
		{
			name:       "audit-manifest",
			schemaPath: repoPath(t, "schemas/audit-manifest/v0/schema.json"),
			fixturePath: repoPath(t, "schemas/audit-manifest/v0/examples/valid.minimal.json"),
			invalidMutator: func(m map[string]any) {
				delete(m, "artifacts")
			},
			expectedHint:     "required",
			expectedPathHint: "<root>",
		},
	}
}

// repoPath resolves a repository-relative path into an absolute filesystem path
// so tests remain stable regardless of the package working directory.
func repoPath(t *testing.T, rel string) string {
	t.Helper()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}

	// When running tests for the `tests` package, cwd is typically repo/tests.
	// Stepping one level up returns the repository root.
	root := filepath.Clean(filepath.Join(wd, ".."))
	return filepath.Join(root, filepath.FromSlash(rel))
}

// assertIssueContains verifies that at least one issue contains the expected
// message hint and/or path hint.
func assertIssueContains(t *testing.T, issues []validation.ValidationIssue, wantMessage, wantPath string) {
	t.Helper()

	for _, issue := range issues {
		messageOK := strings.Contains(strings.ToLower(issue.Message), strings.ToLower(wantMessage))
		pathOK := strings.Contains(issue.InstanceLocation, wantPath)

		if messageOK || pathOK {
			return
		}
	}

	t.Fatalf(
		"expected one issue to contain message hint %q or path hint %q, got: %+v",
		wantMessage,
		wantPath,
		issues,
	)
}