package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CommandResp struct {
	Response string
}

func sendJson(i interface{}, w http.ResponseWriter, rq *http.Request) {
	rsp, _ := json.Marshal(i)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(rsp)
}

func sendResponse(r CommandResp, w http.ResponseWriter, rq *http.Request) {
	sendJson(r, w, rq)
}

func take_hdmi_f(w http.ResponseWriter, r *http.Request) {
	take_hdmi()
	sendResponse(CommandResp{Response: "HDMI"}, w, r)
}

func restart_f(w http.ResponseWriter, r *http.Request) {
	restart()
	sendResponse(CommandResp{Response: "Restarting"}, w, r)
}
func status_f(w http.ResponseWriter, r *http.Request) {
	sendResponse(
		CommandResp{
			Response: "OK",
		}, w, r)
}

func read_f(w http.ResponseWriter, r *http.Request) {
	sendJson(read(), w, r)
}

func write_f(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	write(r.Form["data"][0])
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Ok")
}

func serve() {
	http.HandleFunc("/hdmi", take_hdmi_f)
	http.HandleFunc("/restart", restart_f)
	http.HandleFunc("/status", status_f)
	http.HandleFunc("/read", read_f)
	http.HandleFunc("/write", write_f)
	http.ListenAndServe(":8080", nil)
}
