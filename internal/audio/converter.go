package audio

import (
    "log"
    "os"
    "os/exec"
)

// ConvertWAVToFLAC converts a byte slice containing WAV audio data to FLAC format.
func ConvertWAVToFLAC(wavData []byte) ([]byte, error) {
    // Create a temporary WAV file
    tempWavFile, err := os.CreateTemp("", "*.wav")
    if err != nil {
        log.Println("Error creating temporary WAV file:", err)
        return nil, err
    }
    defer os.Remove(tempWavFile.Name()) // Clean up the temporary file after the function exits

    // Write the WAV data to the temporary file
    if _, err := tempWavFile.Write(wavData); err != nil {
        log.Println("Error writing to temporary WAV file:", err)
        return nil, err
    }
    tempWavFile.Close() // Close the file to flush data

    // Create a temporary FLAC file
    tempFlacFile, err := os.CreateTemp("", "*.flac")
    if err != nil {
        log.Println("Error creating temporary FLAC file:", err)
        return nil, err
    }
    defer os.Remove(tempFlacFile.Name()) // Clean up the temporary file after the function exits

    // Run FFmpeg to convert the WAV file to FLAC
    cmd := exec.Command("ffmpeg", "-i", tempWavFile.Name(), tempFlacFile.Name())
    output, err := cmd.CombinedOutput() // Capture stdout and stderr
    if err != nil {
        log.Printf("Error converting WAV to FLAC: %v\nOutput: %s", err, output)
        return nil, err
    }

    // Read the FLAC data from the temporary file
    flacData, err := os.ReadFile(tempFlacFile.Name())
    if err != nil {
        log.Println("Error reading temporary FLAC file:", err)
        return nil, err
    }

    // Log sizes for debugging
    log.Printf("Temporary WAV file created: %s", tempWavFile.Name())
    wavFileInfo, _ := os.Stat(tempWavFile.Name())
    log.Printf("Temporary WAV file size: %d bytes", wavFileInfo.Size())
    
    log.Printf("Converted FLAC file size: %d bytes", len(flacData))

    return flacData, nil
}
