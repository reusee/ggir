package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGirParser(t *testing.T) {
	girDir := "/usr/share/gir-1.0"
	tmpDir := os.TempDir()
	filepath.Walk(girDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !strings.HasSuffix(path, ".gir") {
			return nil
		}
		p("%s\n", path)
		Gen(tmpDir, path)
		return nil
	})
}
