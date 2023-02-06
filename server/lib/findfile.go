package lib

import (
	"fmt"
	"github.com/chigopher/pathlib"
	"github.com/spf13/afero"
)

// FindFile
// Locate a file in the following places
// - current directory (default)
// - directory of <appname> executable
// - in the directory $HOME/.<appname>
func FindFile(filename string, v ...any) (string, error) {
	var fs afero.Fs = afero.OsFs{}
	if len(v) > 0 {
		if fsa, ok := v[0].(afero.Fs); ok {
			fs = fsa
			v = v[1:]
		}
	}
	filepath := pathlib.NewPathAfero(filename, fs)
	path, err := FindPath(filepath, v...)
	if err != nil {
		return "", err
	}
	return path.String(), err
}

func FindFileExists(path *pathlib.Path) bool {
	if ok, _ := path.Exists(); ok {
		if ok, _ := path.IsFile(); ok {
			return true
		}
	}
	return false
}

// FindPath
// Locate a file in the following places, supporting Afero FS types
// - current directory (default)
// - directory of <appname> executable
// - in the directory $HOME/.<appname>
func FindPath(path *pathlib.Path, search ...any) (*pathlib.Path, error) {
	// try the given file directly, caller may have provided the full path
	if FindFileExists(path) {
		return path, nil
	}

	fs := path.Fs()
	dirs := make([]*pathlib.Path, 0, len(search))
	if len(search) == 0 {
		dirs = append(dirs, pathlib.NewPathAfero(".", fs))
	} else {
		for _, d := range search {
			if s, ok := d.(*pathlib.Path); ok {
				dirs = append(dirs, s)
			} else if s, ok := d.(string); ok {
				dirs = append(dirs, pathlib.NewPathAfero(s, fs))
			}
		}
	}
	for _, dir := range dirs {
		if ok, _ := dir.IsDir(); ok {
			p := dir.JoinPath(path)
			if FindFileExists(p) {
				return p, nil
			}
		}
	}
	return nil, fmt.Errorf("no file found '%s'", path.String())
}
