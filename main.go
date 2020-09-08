package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	resp "github.com/nicklaw5/go-respond"
)

func main() {
	serverAddress := ":3000"
	createJobList()
	router := mux.NewRouter()

	router.HandleFunc("/ws", wsHandler)
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/update.json", updateHandler).Methods("POST")
	router.HandleFunc("/list.json", listHandler).Methods("GET")
	log.Printf("server started at %s\n", serverAddress)
	go echo()
	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// register client
	clients[ws] = true
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var job Job
	render := resp.NewResponse(w)
	err := json.NewDecoder(r.Body).Decode(&job)

	if err != nil {
		render.BadRequest(err.Error())
	}
	for _, j := range listJobs {
		if job.ID == j.ID {
			j.Status = job.Status
		}
	}

	go writer(&job)
	listJobsjson, _ := json.Marshal(listJobs)
	render.Ok(string(listJobsjson))
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	render := resp.NewResponse(w)
	listJobsjson, _ := json.Marshal(listJobs)
	render.Ok(string(listJobsjson))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.tmpl"))
	data := IndexData{Jobs: listJobs}
	tmpl.Execute(w, data)
}
