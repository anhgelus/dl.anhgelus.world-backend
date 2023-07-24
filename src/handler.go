package src

import (
	"net/http"
	"regexp"
)

type UriInfo struct {
	Folder bool `json:"folder"`
	File   bool `json:"file"`
	Hidden bool `json:"hidden"`
	Valid  bool `json:"valid"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	info, err := parseUri(r.RequestURI)
	if err != nil {
		internalError(err, w)
		return
	}
	if !info.Valid {
		badRequest("Invalid uri", w)
		return
	}
	respondWithData(info, "test", w)
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
	info := UriInfo{false, true, false, true}
	if hide.Match(bUri) {
		info.Hidden = true
	}
	if folder.Match(bUri) {
		info.Folder = true
		info.File = false
	}
	return &info, nil
}
