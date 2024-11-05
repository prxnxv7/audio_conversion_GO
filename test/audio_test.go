package test

import (
    "os"
    "testing"

    "audio_conversion/internal/audio"
)

func TestConvertWavToFlac(t *testing.T) {
    // Create a temporary input WAV file with sample data
    tempWavFile, err := os.CreateTemp("", "sample_*.wav")
    if err != nil {
        t.Fatalf("Error creating temporary WAV file: %v", err)
    }
    defer os.Remove(tempWavFile.Name())

    // Write a small sample WAV byte slice to the file
    wavSampleData := []byte{
        // A minimal valid WAV header + some data
        'R', 'I', 'F', 'F', 36, 0, 0, 0, 'W', 'A', 'V', 'E',
        'f', 'm', 't', ' ', 16, 0, 0, 0, 1, 0, 1, 0, 
        44, 0, 0, 0, 88, 0, 0, 0, 16, 0, 'd', 'a', 
        't', 'a', 0, 0, 0, 0, // Add some data here for a valid WAV file
    }
    if _, err := tempWavFile.Write(wavSampleData); err != nil {
        t.Fatalf("Error writing to temporary WAV file: %v", err)
    }
    tempWavFile.Close()

    // Read the WAV data from the temporary file
    wavData, err := os.ReadFile(tempWavFile.Name())
    if err != nil {
        t.Fatalf("Error reading WAV file: %v", err)
    }

    // Run the conversion
    flacData, err := audio.ConvertWAVToFLAC(wavData)
    if err != nil {
        t.Errorf("Conversion failed: %v", err)
    }

    // Verify the output FLAC data is not empty
    if len(flacData) == 0 {
        t.Error("Conversion returned empty data")
    }
}
