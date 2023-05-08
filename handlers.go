package main

import (
	"encoding/json"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	pageData := PageData{
		Handle:  "",
		Value:   "",
		GetURL:  "/get/",
		PostURL: "/create/",
		PutURL:  "/update/",
	}
	renderTemplate(w, "index", pageData)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	var reqData struct {
		Handle string `json:"handle"`
		Key    string `json:"key"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value, err := dbGet(reqData.Handle, reqData.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resData := map[string]string{"value": value}
	json.NewEncoder(w).Encode(resData)
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
	var reqData struct {
		Handle string `json:"handle"`
		Key    string `json:"key"`
		Value  string `json:"value"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dbSet(reqData.Handle, reqData.Key, reqData.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resData := map[string]string{"message": "success"}
	json.NewEncoder(w).Encode(resData)
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}
	var reqData struct {
		Handle string `json:"handle"`
		Key    string `json:"key"`
		Value  string `json:"value"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dbSet(reqData.Handle, reqData.Key, reqData.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resData := map[string]string{"message": "success"}
	json.NewEncoder(w).Encode(resData)
}

func handleResolveHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	hostname := r.Host

	value, err := dbGetKeyForHandle(hostname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resData := map[string]string{"did": value}
	json.NewEncoder(w).Encode(resData)
}
