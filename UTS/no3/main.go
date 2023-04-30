package main

import (
	"dataDiri"
	"encoding/json"
	"net/http"
)

func InfotHandler(w http.ResponseWriter, r *http.Request) {

	var data dataDiri.Informasi
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newData := dataDiri.NewInfo(data.Nama, data.NIM, data.Alamat)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newData)

}

func main() {
	http.HandleFunc("/nama", InfotHandler)
	http.ListenAndServe(":8080", nil)

}
