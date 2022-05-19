package presenter

import (
	"fmt"
	"strings"
)

type PathSegment struct {
	Label string
	Path  string
}

func PathSegmentsFromPath(path string) []PathSegment {
	pathSegments := []PathSegment{
		{
			Label: "root",
			Path:  "/",
		},
	}

	pathSoFar := "/"
	segments := strings.FieldsFunc(path, func(c rune) bool { return c == '/' })

	for _, segment := range segments {
		if pathSoFar == "/" {
			pathSoFar = fmt.Sprintf("/%s", segment)
		} else {
			pathSoFar = fmt.Sprintf("%s/%s", pathSoFar, segment)
		}

		pathSegments = append(pathSegments, PathSegment{
			Label: segment,
			Path:  pathSoFar,
		})
	}

	return pathSegments
}
