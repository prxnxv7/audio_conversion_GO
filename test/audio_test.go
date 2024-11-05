package test

import (
    "os"
    "testing"
    "audio_conversion/internal/audio"
)

func TestConvertWavToFlac(t *testing.T) {
    tempWavFile, err := os.CreateTemp("", "sample_*.wav")
    if err != nil {
        t.Fatalf("Error creating temporary WAV file: %v", err)
    }
    defer os.Remove(tempWavFile.Name())

    wavSampleData := []byte{
        'R', 'I', 'F', 'F', 36, 0, 0, 0, 'W', 'A', 'V', 'E',
        'f', 'm', 't', ' ', 16, 0, 0, 0, 1, 0, 1, 0, 
        44, 0, 0, 0, 88, 0, 0, 0, 16, 0, 'd', 'a', 
        't', 'a', 0, 0, 0, 0,
    }
    if _, err := tempWavFile.Write(wavSampleData); err != nil {
        t.Fatalf("Error writing to temporary WAV file: %v", err)
    }
    tempWavFile.Close()

    wavData, err := os.ReadFile(tempWavFile.Name())
    if err != nil {
        t.Fatalf("Error reading WAV file: %v", err)
    }

    flacData, err := audio.ConvertWAVToFLAC(wavData)
    if err != nil {
        t.Errorf("Conversion failed: %v", err)
    }
    
    if len(flacData) == 0 {
        t.Error("Conversion returned empty data")
    }
}
