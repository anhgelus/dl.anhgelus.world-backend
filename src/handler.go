package src

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type UriInfo struct {
	Folder bool   `json:"folder"`
	File   bool   `json:"file"`
	Hidden bool   `json:"hidden"`
	Valid  bool   `json:"valid"`
	Uri    string `json:"uri"`
}

func (info *UriInfo) genUri() string {
	return fmt.Sprintf("/data%s", info.Uri)
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("Called %s", r.RequestURI)
	info, err := parseUri(r.RequestURI)
	if err != nil {
		internalError(err, w)
		return
	}
	if !info.Valid {
		badRequest("Invalid uri", w)
		return
	}
	badFile, err := regexp.Compile("^read .+: is a directory$")
	if err != nil {
		internalError(err, w)
		return
	}
	infos, err := getInfos(info)
	if err != nil {
		if badFile.Match([]byte(err.Error())) {
			badRequest("The file requested is a folder", w)
		} else {
			internalError(err, w)
		}
		return
	}
	if !infos.valid {
		notFound("Not found", w)
		return
	}
	if infos.Folder {
		respondWithData(infos.Files, "Fetched", w)
		return
	}
	_, err = w.Write(infos.Content)
	if err != nil {
		internalError(err, w)
	}
}

func HandleNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Header().Add("Allow", http.MethodGet)
}

func parseUri(uri string) (*UriInfo, error) {
	hide, err := regexp.Compile("^[.].+$")
	if err != nil {
		return nil, err
	}
	folder, err := regexp.Compile("/$")
	if err != nil {
		return nil, err
	}
	valid, err := regexp.Compile("/[.]*/")
	if err != nil {
		return nil, err
	}
	bUri := []byte(uri)
	if valid.Match(bUri) {
		info := UriInfo{Valid: false}
		return &info, nil
	}
	info := UriInfo{false, true, false, true, uri}
	if hide.Match(bUri) {
		info.Hidden = true
	}
	if folder.Match(bUri) {
		info.Folder = true
		info.File = false
	}
	return &info, nil
}
