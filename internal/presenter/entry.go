package presenter

const (
	EntryFolder        string = "folder"
	EntryImageFile     string = "image"
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
