package archiver

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/xi2/xz"
)

func XZDecompress(filePath string, destDir string) error {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Create an xz Reader to read the compressed data
	xr, err := xz.NewReader(bytes.NewReader(f), 0)
	if err != nil {
		return err
	}

	// Open and iterate through the files in the archive.
	tr := tar.NewReader(xr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		destFilePath := filepath.Join(destDir, hdr.Name)
		destFileParentDir := filepath.Dir(destFilePath)

		// Create the parent directory path if it does not exist
		err = os.MkdirAll(destFileParentDir, os.ModePerm)
		if err != nil {
			return err
		}

		if hdr.Typeflag == tar.TypeDir {
			continue
		}

		if hdr.Typeflag == tar.TypeReg {
			destFile, err := os.Create(destFilePath)
			if err != nil {
				return err
			}

			defer destFile.Close()

			if _, err := io.Copy(destFile, tr); err != nil {
				return err
			}
		}
	}

	return nil
}
