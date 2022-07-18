package httpjson

import (
	"log"
	"net/http"

	"github.com/mailru/easyjson"
)

// Start entrypoint
func Start() {
	http.HandleFunc("/", CreateKrs)
	log.Println(http.ListenAndServe(":60001", nil))
}

func CreateKrs(w http.ResponseWriter, r *http.Request) {
	var krs Krs
	easyjson.UnmarshalFromReader(r.Body, &krs)
	defer r.Body.Close()

	json, _ := easyjson.Marshal(Response{
		Code:        200,
		Message:     "OK",
		MataKuliahs: krs.MataKuliahs,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
