package audio

import (
    "bytes"
    "os/exec"
)

func ConvertWavToFlac(wavData []byte) ([]byte, error) {
    cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "flac", "pipe:1")

    cmd.Stdin = bytes.NewReader(wavData)
    var outBuf bytes.Buffer
    cmd.Stdout = &outBuf

    if err := cmd.Run(); err != nil {
        print(err)
        return nil, err
    }

    return outBuf.Bytes(), nil
}
