package vfs

import "strings"

const (
	DEFAULT_PATH_DELIMITER = "/"
)
type Path struct {
	paths []string
	filename string
	delimiter string
}

type Iterator interface {
	Value() string
	HasNext() bool
}

type pathIterator struct {
	path *Path
	current int
}

func NewPath(path string) *Path {
	newPath := &Path{
		paths: make([]string, 0),
		delimiter: DEFAULT_PATH_DELIMITER,
	}
	splits := strings.Split(path, DEFAULT_PATH_DELIMITER)

	for _, p := range splits {
		if p != "" {
			newPath.paths = append(newPath.paths, p)
		}
	}

	if strings.HasSuffix(path, newPath.delimiter) {
		newPath.filename = ""
	} else {
		newPath.filename = newPath.paths[len(newPath.paths) - 1]
	}

	return newPath
}

func NewPathWithDelimiter(path string, delimiter string) *Path {
	newPath := &Path{
		paths: make([]string, 0),
		delimiter: delimiter,
	}
	splits := strings.Split(path, delimiter)

	for _, p := range splits {
		if p != "" {
			newPath.paths = append(newPath.paths, p)
		}
	}

	if strings.HasSuffix(path, delimiter) || len(newPath.paths) == 0 {
		newPath.filename = ""
	} else {
		newPath.filename = newPath.paths[len(newPath.paths) - 1]
	}

	return newPath
}

func (p *Path) Len() int {
	return len(p.paths)
}

func (p *Path) NthPath(i int) string {
	return p.paths[i]
}

func (p *Path) FileName() string {
	return p.filename
}

func (p *Path) Iterator() Iterator {
	return newPathIterator(p)
}

func newPathIterator(path *Path) *pathIterator {
	return &pathIterator{
		path: path,
		current: -1,
	}
}

func (pi *pathIterator) Value() string {
	return pi.path.paths[pi.current]
}

func (pi *pathIterator) HasNext() bool {
	pi.current++
	return len(pi.path.paths) > pi.current
}
