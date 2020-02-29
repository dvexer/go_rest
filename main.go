package main

import (
	"encoding/json"
	"fmt"
	"go_rest/logger"
	"go_rest/model"
	"net/http"
	"strings"
)

type contact struct {
	Name  string `json:"name"` // will be extracted using reflection
	Phone string `json:"phone"`
}

type server struct{} // identifier

func getName(req *http.Request) string {
	path := strings.Split(req.URL.Path, "/")
	return path[len(path)-1]
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	name := getName(req)
	if name == "" {
		contacts := model.ReadAll()
		res := make([]contact, len(contacts))
		for i, c := range contacts {
			res[i] = contact{c[0], c[1]}
		}
		jsonRes, err := json.Marshal(res)
		logger.LogErrorIfExist(err)
		fmt.Fprintf(w, string(jsonRes))
	} else {
		phone := model.Read(name)
		if phone != "" {
			jsonRes, err := json.Marshal(&contact{name, phone})
			logger.LogErrorIfExist(err)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, string(jsonRes))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	var c contact
	err := json.NewDecoder(req.Body).Decode(&c)
	logger.LogErrorIfExist(err)
	if c.Name != "" && c.Phone != "" {
		id := model.Create(c.Name, c.Phone)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"id": "%d"}`, id)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handlePut(w http.ResponseWriter, req *http.Request) {
	name := getName(req)
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		var c contact
		json.NewDecoder(req.Body).Decode(&c)
		model.Update(name, c.Name, c.Phone)
		jsonRes, err := json.Marshal(&contact{c.Name, c.Phone})
		logger.LogErrorIfExist(err)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(jsonRes))
	}
}

func handleDelete(w http.ResponseWriter, req *http.Request) {
	name := getName(req)
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		model.Delete(name)
		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		handleGet(w, req)
	case http.MethodPost:
		handlePost(w, req)
	case http.MethodPut:
		handlePut(w, req)
	case http.MethodDelete:
		handleDelete(w, req)
	}
}

func main() {
	defer model.CloseDB()
	http.Handle("/", &server{}) // base multiplexing
	err := http.ListenAndServe(":8000", nil)
	logger.LogErrorIfExist(err)
}
