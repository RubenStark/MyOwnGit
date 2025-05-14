package handlers

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"sort"
)

func readTreeEntry(data []byte) (used int, mode string, name string, hash [20]byte) {
	idx := FindNull(data)

	entry := string(data[:idx])
	fmt.Sscanf(entry, "%s %s", &mode, &name)
	copy(hash[:], data[idx+1:idx+21])

	// fmt.Printf("Entry: %s, Mode: %s, Name: %s, Hash: %x\n", entry, mode, name, hash)

	return idx, mode, name, hash
}

func LsTree(args []string) string {

	var result string

	sha := args[2]

	if file, err := os.Open(fmt.Sprintf(".git/objects/%s/%s", sha[:2], sha[2:])); err == nil {
		if reader, err := zlib.NewReader(file); err == nil {
			if data, err := io.ReadAll(reader); err == nil {
				idx := FindNull(data)
				data = data[idx+1:]

				entries := []string{}
				for len(data) != 0 {
					used, _, name, _ := readTreeEntry(data)
					data = data[used:]
					entries = append(entries, name)
				}

				sort.Strings(entries)
				fmt.Printf("Entries: %v\n", entries)

				for _, entry := range entries {
					result += fmt.Sprintf("%s\n", entry)
				}
			}
		}
	} else {
		result = fmt.Sprintf("Error opening file: %s\n", err)
	}

	return result
}
