package validation

import "testing"

func TestNew(t *testing.T) {
	t.Parallel()

	v := New()
	if v == nil {
		t.Fatal("expected validator")
	}
}

func TestDeepCopyMap(t *testing.T) {
	t.Parallel()

	original := map[string]any{
		"name": "example",
		"nested": map[string]any{
			"value": float64(1),
		},
	}

	cloned, err := DeepCopyMap(original)
	if err != nil {
		t.Fatalf("expected deep copy success, got error: %v", err)
	}

	nested := cloned["nested"].(map[string]any)
	nested["value"] = float64(2)

	originalNested := original["nested"].(map[string]any)
	if originalNested["value"] == float64(2) {
		t.Fatal("expected deep copy to avoid aliasing original nested map")
	}
}