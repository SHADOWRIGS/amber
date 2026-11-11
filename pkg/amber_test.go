package amber

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestAddCallOver(t *testing.T) {
	pe := &PE{}
	payload := []byte{0x90, 0x90, 0x90}

	result, err := pe.AddCallOver(payload)
	if err != nil {
		t.Fatalf("AddCallOver returned error: %v", err)
	}

	if result[0] != 0xe8 {
		t.Errorf("first byte should be 0xe8 (CALL), got 0x%02x", result[0])
	}

	var size uint32
	if err := binary.Read(bytes.NewReader(result[1:5]), binary.LittleEndian, &size); err != nil {
		t.Fatal(err)
	}
	if size != 3 {
		t.Errorf("encoded size = %d, want 3", size)
	}

	if !bytes.Equal(result[5:], payload) {
		t.Error("payload not appended correctly after call instruction")
	}
}

func TestAddCallOverEmpty(t *testing.T) {
	pe := &PE{}
	result, err := pe.AddCallOver([]byte{})
	if err != nil {
		t.Fatalf("AddCallOver returned error: %v", err)
	}

	if len(result) != 5 {
		t.Errorf("result length = %d, want 5", len(result))
	}
	if result[0] != 0xe8 {
		t.Errorf("first byte should be 0xe8 (CALL), got 0x%02x", result[0])
	}

	var size uint32
	if err := binary.Read(bytes.NewReader(result[1:5]), binary.LittleEndian, &size); err != nil {
		t.Fatal(err)
	}
	if size != 0 {
		t.Errorf("encoded size = %d, want 0", size)
	}
}

func TestAddCallOverLargePayload(t *testing.T) {
	pe := &PE{}
	payload := make([]byte, 65536)
	for i := range payload {
		payload[i] = byte(i % 256)
	}

	result, err := pe.AddCallOver(payload)
	if err != nil {
		t.Fatalf("AddCallOver returned error: %v", err)
	}

	if len(result) != len(payload)+5 {
		t.Errorf("result length = %d, want %d", len(result), len(payload)+5)
	}

	var size uint32
	if err := binary.Read(bytes.NewReader(result[1:5]), binary.LittleEndian, &size); err != nil {
		t.Fatal(err)
	}
	if size != 65536 {
		t.Errorf("encoded size = %d, want 65536", size)
	}
}

func TestOpenNonExistentFile(t *testing.T) {
	_, err := Open("/nonexistent/file.exe")
	if err == nil {
		t.Error("Open should return error for nonexistent file")
	}
}

func TestErrorMessages(t *testing.T) {
	if ErrUnsupportedArch.Error() != "unsupported PE file architecture" {
		t.Error("unexpected ErrUnsupportedArch message")
	}
	if ErrInvalidPeSpecs.Error() != "unsupported PE file specs" {
		t.Error("unexpected ErrInvalidPeSpecs message")
	}
	if ErrInvalidPeHeaders.Error() != "invalid PE headers" {
		t.Error("unexpected ErrInvalidPeHeaders message")
	}
}

func TestLoaderBlobsNotEmpty(t *testing.T) {
	if len(LOADER_64) == 0 {
		t.Error("LOADER_64 should not be empty")
	}
	if len(LOADER_32) == 0 {
		t.Error("LOADER_32 should not be empty")
	}
	if len(SYSCALL_LOADER_64) == 0 {
		t.Error("SYSCALL_LOADER_64 should not be empty")
	}
}

func TestCrcApiBlobs(t *testing.T) {
	if len(CRC_API_32) == 0 {
		t.Error("CRC_API_32 should not be empty")
	}
	if len(CRC_API_64) == 0 {
		t.Error("CRC_API_64 should not be empty")
	}
	if len(IAT_API_32) == 0 {
		t.Error("IAT_API_32 should not be empty")
	}
	if len(IAT_API_64) == 0 {
		t.Error("IAT_API_64 should not be empty")
	}
}

func TestPEStructDefaults(t *testing.T) {
	pe := &PE{}
	if pe.Architecture != 0 {
		t.Error("default architecture should be 0")
	}
	if pe.SyscallLoader {
		t.Error("default SyscallLoader should be false")
	}
	if pe.IatResolver {
		t.Error("default IatResolver should be false")
	}
	if pe.ScrapeHeaders {
		t.Error("default ScrapeHeaders should be false")
	}
	if pe.IsDLL {
		t.Error("default IsDLL should be false")
	}
}
