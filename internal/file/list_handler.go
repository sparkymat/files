package file

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/files/internal/presenter"
	"github.com/sparkymat/files/view"
)

type ListHandlerConfig interface {
	RootFolder() string
}

func ListHandler(cfg ListHandlerConfig) func(echo.Context) error {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		folderPath := filepath.Join(cfg.RootFolder(), path)

		var folderInfo os.FileInfo
		var err error
		if folderInfo, err = os.Stat(folderPath); os.IsNotExist(err) {
			return c.String(http.StatusNotFound, "no such file or directory")
		}
		if !folderInfo.IsDir() {
			return c.File(folderPath)
		}

		folderEntries := []presenter.Entry{}
		if path != "" && path != "/" {
			parentPath := strings.TrimSuffix(path, filepath.Base(path))
			parentPath = strings.TrimSuffix(parentPath, "/")
			if parentPath == "" {
				parentPath = "/"
			}
			folderEntries = append(folderEntries, presenter.Entry{
				Label:        "..",
				Type:         presenter.EntryFolder,
				MaterialIcon: "folder",
				Path:         parentPath,
			})
		}

		fileEntries := []presenter.Entry{}

		fileInfos, err := ioutil.ReadDir(folderPath)
		if err != nil {
			return c.String(http.StatusNotFound, "unable to read folder")
		}

		for _, fileInfo := range fileInfos {
			var entryPath string
			if path == "/" || path == "" {
				entryPath = fmt.Sprintf("/%s", fileInfo.Name())
			} else {
				entryPath = fmt.Sprintf("%s/%s", path, fileInfo.Name())
			}
			if fileInfo.IsDir() {
				folderEntries = append(folderEntries, entryForFolder(fileInfo, entryPath))
			} else {
				fileEntries = append(fileEntries, entryForFile(fileInfo, entryPath))
			}
		}

		entries := append(folderEntries, fileEntries...)

		listPresenter := presenter.List{
			PathSegments: presenter.PathSegmentsFromPath(path),
			Entries:      entries,
		}

		html := view.Layout("files", view.List(listPresenter))
		return c.HTML(http.StatusOK, html)
	}
}

func entryForFolder(fileInfo os.FileInfo, entryPath string) presenter.Entry {
	return presenter.Entry{
		Label:        fileInfo.Name(),
		Type:         presenter.EntryFolder,
		MaterialIcon: "folder",
		Path:         entryPath,
	}
}

func entryForFile(fileInfo os.FileInfo, entryPath string) presenter.Entry {
	extension := filepath.Ext(fileInfo.Name())
	entryType, icon := presenter.EntryTypeAndIconFromExtension(extension)
	return presenter.Entry{
		Label:        fileInfo.Name(),
		Type:         entryType,
		MaterialIcon: icon,
		Path:         entryPath,
		Size:         renderSize(fileInfo.Size()),
	}
}

func renderSize(size int64) string {
	const unit = 1000
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "kMGTPE"[exp])
}
