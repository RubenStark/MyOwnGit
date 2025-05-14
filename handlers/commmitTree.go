package handlers

import (
	"bytes"
	"crypto/sha1"
	"fmt"
)

func CommitTree(args []string) string {

	treeSha := args[0]
	commitSha := args[1]
	message := args[2]

	blob := bytes.Buffer{}

	fmt.Fprintf(&blob, "tree %s\n", treeSha)
	fmt.Fprintf(&blob, "parent %s\n", commitSha)
	fmt.Fprint(&blob, "author Jun Lim <rlopez@email.com> 1243040974 -0700\n")
	fmt.Fprint(&blob, "committer Jun Lim <rlopez@email.com> 1243040974 -0700\n\n")
	fmt.Fprintf(&blob, "%s\n", message)

	newBlob := bytes.Buffer{}

	fmt.Fprintf(&newBlob, "commit %d", blob.Len())
	newBlob.WriteByte(0)
	newBlob.Write(blob.Bytes())

	hash := sha1.Sum(newBlob.Bytes())

	WriteObject(hash, newBlob)

	return fmt.Sprintf("%x", hash)
}
