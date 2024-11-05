package ffmpeg

import (
    "bytes"
    "os/exec"
)

func RunCommand(input []byte, args ...string) ([]byte, error) {
    cmd := exec.Command("ffmpeg", args...)

    cmd.Stdin = bytes.NewReader(input)
    var outBuf bytes.Buffer
    cmd.Stdout = &outBuf

    if err := cmd.Run(); err != nil {
        return nil, err
    }

    return outBuf.Bytes(), nil
}
