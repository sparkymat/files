package presenter

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
	Size         string
}

func EntryTypeAndIconFromExtension(extension string) (string, string) {
	switch extension {
	case ".jpg", ".jpeg", ".gif", ".png", ".bmp":
		return EntryImageFile, "image"
	case ".mov", ".mpg", ".mpeg", ".mp4", ".mkv", ".flv":
		return EntryVideoFile, "movie"
	case ".mp3", ".m4a", ".aac", ".ac3", ".wav", ".flac":
		return EntryMusicFile, "music_note"
	case ".pdf":
		return EntryPDFFile, "insert_drive_file"
	case ".iso", ".dmg", ".img":
		return EntryDiscImageFile, "album"
	case ".csv":
		return EntryTextFile, "insert_drive_file"
	default:
		return EntryUnknown, "info"
	}
}
