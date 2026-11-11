package config

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}
}

func TestConfigDefaults(t *testing.T) {
	cfg := &Config{}
	if cfg.EncodeCount != 0 {
		t.Errorf("default EncodeCount = %d, want 0", cfg.EncodeCount)
	}
	if cfg.ObfuscationLimit != 0 {
		t.Errorf("default ObfuscationLimit = %d, want 0", cfg.ObfuscationLimit)
	}
	if cfg.UseIAT {
		t.Error("default UseIAT should be false")
	}
	if cfg.UseSyscalls {
		t.Error("default UseSyscalls should be false")
	}
	if cfg.ScrapePeHeaders {
		t.Error("default ScrapePeHeaders should be false")
	}
	if cfg.Verbose {
		t.Error("default Verbose should be false")
	}
}

func TestPrintSummary(t *testing.T) {
	cfg := &Config{
		FileName:         "test.exe",
		EncodeCount:      3,
		ObfuscationLimit: 10,
		UseIAT:           false,
		UseSyscalls:      false,
	}
	cfg.PrintSummary()

	cfg.UseIAT = true
	cfg.UseSyscalls = true
	cfg.PrintSummary()
}
