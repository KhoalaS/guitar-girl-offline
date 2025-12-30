package rpc

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os/exec"
)

func encodeBase64(input []byte) string {
	encodedString := base64.StdEncoding.EncodeToString(input)

	return encodedString
}

func compressBz2(data []byte) ([]byte, error) {
	cmd := exec.Command("bzip2", "-c") // -c = output to stdout

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	// Write input and close stdin to signal EOF
	go func() {
		defer stdin.Close()
		stdin.Write(data)
	}()

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("bzip2 error: %v: %s", err, out.String())
	}

	return out.Bytes(), nil
}
