package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"devx/iferr"

	"github.com/julienschmidt/httprouter"
)

var id uint64 = 100

func apiCreateObjective(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Validate Content-Type
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		http.Error(w, "Expect Content-Type: application/json", http.StatusBadRequest)
		return
	}

	// Read json data
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Read: [%s] %s\n", ct, b)
	var data struct {
		Value string
		Id    uint64
	}
	if err := json.Unmarshal(b, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process data
	time.Sleep(time.Second * 5)
	id++
	data.Id = id

	// Response
	ct = "application/json"
	w.Header().Set("Content-Type", ct)
	b, err = json.Marshal(&data)
	iferr.Panic(err)
	fmt.Printf("Response: [%s] %s\n", ct, b)
	w.Write(b)
}

func apiUpdateObjective(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Validate Content-Type
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		http.Error(w, "Expect Content-Type: application/json", http.StatusBadRequest)
		return
	}

	// Read json data
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Read: [%s] %s\n", ct, b)
	var data struct {
		Value string
		Id    uint64
	}
	if err := json.Unmarshal(b, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process data
	time.Sleep(time.Second * 5)
	data.Id, err = paramId(p, "id")
	iferr.Panic(err)

	// Response
	ct = "application/json"
	w.Header().Set("Content-Type", ct)
	b, err = json.Marshal(&data)
	iferr.Panic(err)
	fmt.Printf("Response: [%s] %s\n", ct, b)
	w.Write(b)
}

func apiDeleteObjective(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// id, err := paramId(p, "id")
	// iferr.Panic(err)

	// Response
	ct := "application/json"
	w.Header().Set("Content-Type", ct)
	w.Write([]byte("{}"))
}

func apiUpdateObjectiveParent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := paramId(p, "id")
	iferr.Panic(err)
	pid, err := paramId(p, "pid")
	iferr.Panic(err)

	if rand.Float64() < 0.5 {
		http.Error(w, "You are unlucky", http.StatusInternalServerError)
		return
	}

	fmt.Println("Update objective parent:", id, pid)
	var data struct {
		HTML string
	}
	data.HTML = genHtmlObjParent(p.ByName("pid"))
	b, err := json.Marshal(&data)
	iferr.Panic(err)
	ct := "application/json"
	w.Header().Set("Content-Type", ct)
	w.Write(b)
}

func apiDeleteKeyResult(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// id, err := paramId(p, "id")
	// iferr.Panic(err)

	// Response
	ct := "application/json"
	w.Header().Set("Content-Type", ct)
	w.Write([]byte("{}"))
}
