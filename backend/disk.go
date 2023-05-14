package backend

import (
	"os"
	"path/filepath"

	"github.com/shirou/gopsutil/disk"
)

type DiskUsage struct {
	Total uint64
	Used  uint64
}

func (b *Backend) GetDiskUsage() (DiskUsage, error) {
	usageStat, err := disk.Usage("/")
	if err != nil {
		return DiskUsage{}, err
	}

	return DiskUsage{Total: usageStat.Total, Used: usageStat.Used}, nil
}

type FolderSize struct {
	Path string
	Size int64
}

func (b *Backend) GetFolderSizes(path string) ([]FolderSize, error) {
	var sizes []FolderSize
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			size := int64(0)
			filepath.Walk(info.Name(), func(_ string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if !info.IsDir() {
					size += info.Size()
				}

				return nil
			})

			sizes = append(sizes, FolderSize{Path: info.Name(), Size: size})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return sizes, nil
}
