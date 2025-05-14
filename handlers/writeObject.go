package handlers

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"os"
)

func WriteObject(hash [20]byte, blob bytes.Buffer) error {
	err := os.MkdirAll(fmt.Sprintf(".git/objects/%x/", hash[:1]), 0755)
	if err != nil {
		return err
	}

	compressed := bytes.Buffer{}
	writer := zlib.NewWriter(&compressed)
	writer.Write(blob.Bytes())
	writer.Close()

	err = os.WriteFile(fmt.Sprintf(".git/objects/%x/%x", hash[:1], hash[1:]), compressed.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
