package api

import (
	"github.com/XiaoMengXinX/otto-go"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	s := r.FormValue("s")

	if s == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var audio []byte
	audio, err := otto.GetOTTO(s)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "audio/wav")
	if _, err := w.Write(audio); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}
