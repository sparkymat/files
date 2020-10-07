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
			folderEntries = append(folderEntries, presenter.EntryForParentFolder(parentPath))
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
				folderEntries = append(folderEntries, presenter.EntryFromFileInfo(fileInfo, entryPath))
			} else {
				fileEntries = append(fileEntries, presenter.EntryFromFileInfo(fileInfo, entryPath))
			}
		}

		entries := append(folderEntries, fileEntries...)

		listPresenter := presenter.List{
			PathSegments:   presenter.PathSegmentsFromPath(path),
			Entries:        entries,
			ShowGridButton: true,
			ShowListButton: true,
		}

		html := view.Layout("files", view.List(listPresenter))
		return c.HTML(http.StatusOK, html)
	}
}
