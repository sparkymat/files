package presenter

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	EntryFolder        string = "folder"
	EntryImageFile     string = "image"
	EntryVideoFile     string = "video"
	EntryMusicFile     string = "music"
	EntryDiscImageFile string = "disc_image"
	EntryTextFile      string = "text"
	EntryPDFFile       string = "pdf"
	EntryUnknown       string = "unknown"
)

type Entry struct {
	Type         string
	Label        string
	Path         string
	MaterialIcon string
	IconClass    string
	Size         string
	Linkable     bool
	LinkClass    string
}

func EntryFromFileInfo(fileInfo os.FileInfo, entryPath string) Entry {
	if fileInfo.IsDir() {
		return entryForFolder(fileInfo, entryPath)
	}

	return entryForFile(fileInfo, entryPath)
}

//nolint:goconst
func EntryTypeAndIconDetailsFromExtension(extension string) (string, string, string) { //nolint:revive
	switch extension {
	case ".jpg", ".jpeg", ".gif", ".png", ".bmp":
		return EntryImageFile, "image", "green-text"
	case ".mov", ".mpg", ".mpeg", ".mp4", ".mkv", ".flv":
		return EntryVideoFile, "movie", "red-text text-accent-4"
	case ".mp3", ".m4a", ".aac", ".ac3", ".wav", ".flac":
		return EntryMusicFile, "music_note", "blue-text text-accent-3"
	case ".pdf":
		return EntryPDFFile, "insert_drive_file", "red-text text-accent-4"
	case ".iso", ".dmg", ".img":
		return EntryDiscImageFile, "album", "grey-text text-darken-1"
	case ".csv":
		return EntryTextFile, "insert_drive_file", "grey-text text-darken-1"
	default:
		return EntryUnknown, "info", "grey-text text-darken-1"
	}
}

func EntryForParentFolder(entryPath string) Entry {
	return Entry{
		Label:        "..",
		Type:         EntryFolder,
		MaterialIcon: "folder",
		Path:         entryPath,
		Linkable:     true,
	}
}

func entryForFolder(fileInfo os.FileInfo, entryPath string) Entry {
	return Entry{
		Label:        fileInfo.Name(),
		Type:         EntryFolder,
		MaterialIcon: "folder",
		IconClass:    "yellow-text text-accent-4",
		Path:         entryPath,
		Linkable:     true,
	}
}

func entryForFile(fileInfo os.FileInfo, entryPath string) Entry {
	extension := filepath.Ext(fileInfo.Name())
	entryType, icon, iconClass := EntryTypeAndIconDetailsFromExtension(extension)

	var linkClass string
	if entryType == EntryImageFile {
		linkClass = "image-popup"
	}

	return Entry{
		Label:        fileInfo.Name(),
		Type:         entryType,
		MaterialIcon: icon,
		IconClass:    iconClass,
		Path:         entryPath,
		Size:         renderSize(fileInfo.Size()),
		LinkClass:    linkClass,
		Linkable:     entryType == EntryImageFile,
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
