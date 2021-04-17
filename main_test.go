package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetEpochWithoutSavingBefore(t *testing.T) {
	req, err := http.NewRequest("GET", "/getEpoch", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEpoch)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

}

func TestGetEpochWithSavingBefore(t *testing.T) {
	req, err := http.NewRequest("POST", "/saveEpoch", strings.NewReader("1618662464"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveEpoch)
	handler.ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/getEpoch", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = GetEpoch
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `1618662464`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestSaveEpoch(t *testing.T) {
	req, err := http.NewRequest("POST", "/saveEpoch", strings.NewReader("1618662464"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveEpoch)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected := ``
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestSaveEpochWithWrongInput(t *testing.T) {
	req, err := http.NewRequest("POST", "/saveEpoch", strings.NewReader("WRONGINPUT"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveEpoch)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	expected := "strconv.ParseInt: parsing \"WRONGINPUT\": invalid syntax\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestSaveEpochWithNoInput(t *testing.T) {
	req, err := http.NewRequest("POST", "/saveEpoch", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SaveEpoch)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

}
