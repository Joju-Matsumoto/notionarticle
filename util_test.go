package notionarticle

import (
	"encoding/json"
	"os"
	"testing"
)

func skipLocal(t *testing.T) {
	t.Helper()

	if os.Getenv("LOCAL") != "true" {
		t.Skip("local test")
	}
}

func saveJson(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}

func loadJson(path string, v any) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}
