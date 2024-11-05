package test

import (
	"testing"

	"audio_conversion/internal/audio"
)

func TestConvertWavToFlac(t *testing.T) {
	wavData := []byte{} // Ideally, add a small sample WAV byte slice for testing

	flacData, err := audio.ConvertWavToFlac(wavData)
	if err != nil {
		t.Errorf("Conversion failed: %v", err)
	}

	if len(flacData) == 0 {
		t.Error("Conversion returned empty data")
	}
}
