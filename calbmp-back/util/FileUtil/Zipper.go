package FileUtil

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func GetFilesToZip(root, targetString string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, targetString) {
			if info.IsDir() {
				filesInDir, err := filepath.Glob(filepath.Join(path, "*"))
				if err != nil {
					return err
				}

				for _, file := range filesInDir {
					fInfo, err := os.Stat(file)
					if err != nil {
						return err
					}

					if !fInfo.IsDir() {
						file = strings.Replace(file, "\\", "/", -1)
						files = append(files, file)
					}
				}
			} else {
				path = strings.Replace(path, "\\", "/", -1)
				files = append(files, path)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func CreateZipFile(zipFileName string, files []string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if strings.HasSuffix(file, ".zip") {
			continue
		}
		err = addToZip(zipWriter, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func addToZip(zipWriter *zip.Writer, file string) error {
	srcFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	fileInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	fileHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	fileHeader.Name = file
	zipEntryWriter, err := zipWriter.CreateHeader(fileHeader)
	if err != nil {
		return err
	}

	_, err = io.Copy(zipEntryWriter, srcFile)
	if err != nil {
		return err
	}

	return nil
}
