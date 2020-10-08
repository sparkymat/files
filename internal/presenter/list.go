package presenter

import "github.com/sparkymat/files/config"

type List struct {
	CurrentPath    string
	PathSegments   []PathSegment
	Entries        []Entry
	ViewType       config.ViewType
	ShowGridButton bool
	ShowListButton bool
}
