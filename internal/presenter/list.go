package presenter

type List struct {
	CurrentPath    string
	PathSegments   []PathSegment
	Entries        []Entry
	ShowGridButton bool
	ShowListButton bool
}
