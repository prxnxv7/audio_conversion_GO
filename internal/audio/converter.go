package audio

import (
	"bytes"
	"fmt"
	// "log"
	// "os"
	"os/exec"
)

// func ConvertWAVToFLAC(wavData []byte) ([]byte, error) {
//     tempWavFile, err := os.CreateTemp("", "*.wav")
//     if err != nil {
//         log.Println("Error creating temporary WAV file:", err)
//         return nil, err
//     }
//     defer os.Remove(tempWavFile.Name())

//     if _, err := tempWavFile.Write(wavData); err != nil {
//         log.Println("Error writing to temporary WAV file:", err)
//         return nil, err
//     }
//     tempWavFile.Close()
//     tempFlacFile, err := os.CreateTemp("", "*.flac")
//     if err != nil {
//         log.Println("Error creating temporary FLAC file:", err)
//         return nil, err
//     }
//     defer os.Remove(tempFlacFile.Name())

//     cmd := exec.Command("ffmpeg", "-i", tempWavFile.Name(), tempFlacFile.Name())
//     output, err := cmd.CombinedOutput()
//     if err != nil {
//         log.Printf("Error converting WAV to FLAC: %v\nOutput: %s", err, output)
//         return nil, err
//     }

//     flacData, err := os.ReadFile(tempFlacFile.Name())
//     if err != nil {
//         log.Println("Error reading temporary FLAC file:", err)
//         return nil, err
//     }

//     log.Printf("Temporary WAV file created: %s", tempWavFile.Name())
//     wavFileInfo, _ := os.Stat(tempWavFile.Name())
//     log.Printf("Temporary WAV file size: %d bytes", wavFileInfo.Size())

//     log.Printf("Converted FLAC file size: %d bytes", len(flacData))

//     return flacData, nil
// }

// @Summary Convert WAV to FLAC
// @Description Convert incoming WAV audio stream to FLAC format
// @Accept  audio/wav
// @Produce audio/flac
// @Success 200 {file} []byte
// @Router /convert [post]
func ConvertWAVToFLAC(wavData []byte) ([]byte, error) {
    cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "flac", "pipe:1")
    cmd.Stdin = bytes.NewReader(wavData)
    var flacBuffer bytes.Buffer
    cmd.Stdout = &flacBuffer
  
    err := cmd.Run()
    if err != nil {
      return nil, fmt.Errorf("ffmpeg error: %w", err)
    }
  
    return flacBuffer.Bytes(), nil
  }
