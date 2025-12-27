package reader

import (
	"os"
	"testing"
)

func TestReadLines(t *testing.T) {
	t.Run("successfully read lines from file", func(t *testing.T) {
		file, err := os.CreateTemp("", "reader-test-")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}

		t.Cleanup(func() {
			_ = os.Remove(file.Name())
		})

		content := "[INFO] start\n[WARN] slow\n[ERROR] fail"
		_, err = file.WriteString(content)
		if err != nil {
			t.Fatalf("failed to write to temp file: %v", err)
		}

		if err = file.Close(); err != nil {
			t.Fatalf("failed to close temp file: %v", err)
		}

		lines, err := ReadLines(file.Name())
		if err != nil {
			t.Fatalf("ReadLines() returns error: %v", err)
		}

		want := []string{
			"[INFO] start",
			"[WARN] slow",
			"[ERROR] fail",
		}

		if len(lines) != len(want) {
			t.Fatalf(
				"ReadLines() returns %d lines, want %d",
				len(lines),
				len(want),
			)
		}
		for i := range lines {
			if lines[i] != want[i] {
				t.Errorf(
					"ReadLines() returns %v, want %v",
					lines[i],
					want[i],
				)
			}
		}
	})

	t.Run("file does not exist", func(t *testing.T) {
		_, err := ReadLines("this-file-does-not-exist.log")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("file is empty", func(t *testing.T) {
		file, err := os.CreateTemp("", "reader-test-")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}

		t.Cleanup(func() {
			_ = os.Remove(file.Name())
		})

		if err := file.Close(); err != nil {
			t.Fatalf("failed to close temp file: %v", err)
		}

		lines, err := ReadLines(file.Name())
		if err != nil {
			t.Fatalf("ReadLines() returns error: %v", err)
		}

		if len(lines) != 0 {
			t.Fatalf("ReadLines() returns %d lines, want 0", len(lines))
		}
	})
}
