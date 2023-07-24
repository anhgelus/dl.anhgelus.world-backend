package src

import (
	"errors"
	"io/fs"
	"os"
	"regexp"
)

type FileInfo struct {
	Folder bool   `json:"folder"`
	Path   string `json:"path"`
}

type FSInfo struct {
	valid   bool
	Folder  bool
	Files   []*FileInfo `json:"files"`
	Content []byte
}

func getInfos(info *UriInfo) (*FSInfo, error) {
	if info.File {
		return getFile(info)
	} else {
		return getDir(info)
	}
}

func getFile(info *UriInfo) (*FSInfo, error) {
	uri := info.genUri()
	c, err := os.ReadFile(uri)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	} else if errors.Is(err, os.ErrNotExist) {
		return &FSInfo{
			valid:   false,
			Files:   nil,
			Folder:  false,
			Content: c,
		}, nil
	}
	return &FSInfo{
		valid:   true,
		Files:   nil,
		Folder:  false,
		Content: c,
	}, nil
}

func getDir(info *UriInfo) (*FSInfo, error) {
	uri := info.genUri()
	folder := os.DirFS(uri)
	entries, err := fs.ReadDir(folder, ".")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	} else if errors.Is(err, fs.ErrNotExist) {
		return &FSInfo{
			valid:   false,
			Folder:  true,
			Files:   nil,
			Content: nil,
		}, nil
	}
	var infos []*FileInfo
	for _, entry := range entries {
		hide, err := regexp.Compile("^[.].+$")
		if err != nil {
			return nil, err
		}
		if hide.Match([]byte(entry.Name())) {
			continue
		}
		infos = append(infos, &FileInfo{
			Folder: entry.IsDir(),
			Path:   info.Uri + entry.Name(),
		})
	}
	return &FSInfo{
		valid:   true,
		Folder:  true,
		Files:   infos,
		Content: nil,
	}, nil
}
