package fileutil

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type fileInfo struct {
	path  string
	size  int64
	mtime time.Time
}

func DirSizeBytes(dir string) (int64, error) {
	var total int64
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		total += info.Size()
		return nil
	})
	return total, err
}

// CleanupDirByOldest deletes files under dir until the total size is <= thresholdBytes.
// It deletes regular files only, sorted by mod time (oldest first).
func CleanupDirByOldest(dir string, thresholdBytes int64) (deleted int, freedBytes int64, err error) {
	if thresholdBytes <= 0 {
		return 0, 0, errors.New("thresholdBytes must be positive")
	}
	if _, statErr := os.Stat(dir); statErr != nil {
		if os.IsNotExist(statErr) {
			return 0, 0, nil
		}
		return 0, 0, statErr
	}

	var files []fileInfo
	var total int64
	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		total += info.Size()
		files = append(files, fileInfo{path: path, size: info.Size(), mtime: info.ModTime()})
		return nil
	})
	if err != nil {
		return 0, 0, err
	}
	if total <= thresholdBytes {
		return 0, 0, nil
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].mtime.Equal(files[j].mtime) {
			return files[i].path < files[j].path
		}
		return files[i].mtime.Before(files[j].mtime)
	})

	for _, f := range files {
		if total <= thresholdBytes {
			break
		}
		if rmErr := os.Remove(f.path); rmErr != nil && !os.IsNotExist(rmErr) {
			// Continue deleting others, return last error.
			err = rmErr
			continue
		}
		deleted++
		freedBytes += f.size
		total -= f.size
	}
	return deleted, freedBytes, err
}
