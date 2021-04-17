package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var Timestamp time.Time

func GetEpoch(w http.ResponseWriter, _ *http.Request) {
	if Timestamp.Unix() < 0 {
		http.Error(w, "No data available.", 404)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	logErr(w.Write([]byte(strconv.FormatInt(Timestamp.Unix(), 10))))
}

func SaveEpoch(w http.ResponseWriter, r *http.Request) {
	done := make(chan struct{})

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(b) > 0 {
		handleEpochSave(w, b, done)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	logErr(w.Write([]byte(`{"error": "not created"}`)))
}

func handleEpochSave(w http.ResponseWriter, b []byte, done chan struct{}) {
	stringBody := string(b)
	epoch, err := strconv.ParseInt(stringBody, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	go func() {
		pointerToTimestamp := &Timestamp
		*pointerToTimestamp = time.Unix(epoch, 0)
		done <- struct{}{}
	}()

	<-done
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(Timestamp.Unix())
}

func logErr(_ int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/getEpoch", GetEpoch).Methods(http.MethodGet)
	api.HandleFunc("/saveEpoch", SaveEpoch).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}
