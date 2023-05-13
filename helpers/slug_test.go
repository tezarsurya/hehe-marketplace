package helpers

import "testing"

func TestGenerateSlug(t *testing.T) {
	got := GenerateSlug("Tezar Surya Fernanda")
	want := "tezar_surya_fernanda"

	t.Logf("Result: %s", got)

	if got != want {
		t.Errorf("Expected %s, got %s", want, got)
	}
}
