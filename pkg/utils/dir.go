package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

// CopyDirectory copies content of one directory to another
func CopyDirectory(scrDir, dest string) error {
	entries, err := ioutil.ReadDir(scrDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		// get src entry details
		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			CheckErrNonFatal(err)
			continue
		}

		// check if dest already exists
		_, err = os.Stat(destPath)
		if err != nil && !os.IsNotExist(err) {
			CheckErrNonFatal(err)
			continue
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			CheckErrNonFatal(fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath))
			continue
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := createIfNotExists(destPath, 0755); err != nil {
				CheckErrNonFatal(err)
				continue
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				CheckErrNonFatal(err)
				continue
			}
		case os.ModeSymlink:
			if err := copySymLink(sourcePath, destPath); err != nil {
				CheckErrNonFatal(err)
				continue
			}
		default:
			if err := copy(sourcePath, destPath); err != nil {
				CheckErrNonFatal(err)
				continue
			}
		}

		if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
			CheckErrNonFatal(err)
			continue
		}

		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				CheckErrNonFatal(err)
				continue
			}
		}
	}
	return nil
}

func copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer func() {
		err = in.Close()
	}()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func checkExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func createIfNotExists(dir string, perm os.FileMode) error {
	if checkExists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func copySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}
