package presenter

const (
	EntryFolder   string = "folder"
	EntryJPEGFile string = "jpeg"
	EntryUnknown  string = "unknown"
)

type Entry struct {
	Type         string
	Label        string
	Path         string
	MaterialIcon string
}

func EntryTypeAndIconFromExtension(extension string) (string, string) {
	switch extension {
	case ".jpg", ".jpeg":
		return EntryJPEGFile, "image"
	default:
		return EntryUnknown, "info"
	}
}
