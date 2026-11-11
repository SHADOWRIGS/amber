package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"zero length", 0},
		{"short string", 5},
		{"medium string", 16},
		{"long string", 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomString(tt.length)
			if len(result) != tt.length {
				t.Errorf("RandomString(%d) returned string of length %d", tt.length, len(result))
			}
		})
	}

	s1 := RandomString(32)
	s2 := RandomString(32)
	if s1 == s2 {
		t.Error("RandomString produced identical strings on consecutive calls")
	}
}

func TestRandomStringCharset(t *testing.T) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := RandomString(1000)
	for i, c := range result {
		found := false
		for _, valid := range charset {
			if c == valid {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("RandomString produced invalid character %q at index %d", c, i)
		}
	}
}

func TestGetFileSize(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.bin")

	content := []byte("hello world")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}

	size, err := GetFileSize(path)
	if err != nil {
		t.Fatalf("GetFileSize returned error: %v", err)
	}
	if size != len(content) {
		t.Errorf("GetFileSize = %d, want %d", size, len(content))
	}
}

func TestGetFileSizeEmpty(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.bin")

	if err := os.WriteFile(path, []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

	size, err := GetFileSize(path)
	if err != nil {
		t.Fatalf("GetFileSize returned error: %v", err)
	}
	if size != 0 {
		t.Errorf("GetFileSize = %d, want 0", size)
	}
}

func TestGetFileSizeNotFound(t *testing.T) {
	_, err := GetFileSize("/nonexistent/file.bin")
	if err == nil {
		t.Error("GetFileSize should return error for nonexistent file")
	}
}
