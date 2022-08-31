package controller

import (
	"encoding/json"
	"net/http"

	"timbul.io/lib/mongodb"
	"timbul.io/service/backend"
)

func GetMovie(env *backend.Service, w http.ResponseWriter, r *http.Request) {
	db := mongodb.New()
	result := db.QueryOne("sample_mflix", "movies", "title", "Back to the Future")
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(jsonData))

}
