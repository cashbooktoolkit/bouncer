/*
 * bouncer.go - Main entry point
 *
 * Bouncer is (c) 2014 Sourdough Labs Research and Development Corp.
 *
 * License: MIT (See LICENSE for details)
 */

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
	"log"
	"flag"
	"net/http"
	"html/template"
	"net/url"
	"github.com/RangelReale/osin"
)

func main() {

	port := flag.String("port", "14000", "Port number to listen on")
	backend_url := flag.String("backend", "http://localhost:14001/authenticate", "Address of the authentication backend")

    flag.Parse()

	config := osin.NewServerConfig()
	config.AllowGetAccessRequest = true
	config.AllowClientSecretInParams = true

	storage := NewInMemoryStorage()

	load_clients(storage)

	server := osin.NewServer(config, storage)

	// Authorization code endpoint
	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {
			if !HandleLoginPage(*backend_url, resp, ar, w, r) {
				return
			}
			ar.Authorized = true
			server.FinishAuthorizeRequest(resp, r, ar)
		}
		if resp.IsError && resp.InternalError != nil {
			fmt.Printf("ERROR: %s\n", resp.InternalError)
		}
		osin.OutputJSON(resp, w, r)
	})

	// Access token endpoint
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		if ar := server.HandleAccessRequest(resp, r); ar != nil {
			ar.Authorized = true
			server.FinishAccessRequest(resp, r, ar)
		}
		if resp.IsError && resp.InternalError != nil {
			fmt.Printf("ERROR: (internal) %s\n", resp.InternalError)
		}
		osin.OutputJSON(resp, w, r)
	})

	// Information endpoint
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		if ir := server.HandleInfoRequest(resp, r); ir != nil {
			server.FinishInfoRequest(resp, r, ir)
		}
		osin.OutputJSON(resp, w, r)
	})


	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.ListenAndServe(":" + *port, nil)
}

type RemoteReponse struct {
	Uid string
	Error string
	Success bool
}

type PageContext struct {
	Error string
	PostUrl string
}

func HandleLoginPage(backend_url string, resp *osin.Response, ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) bool {

	r.ParseForm()

	error_message := ""

	if r.Method == "POST" {

		res, err := http.PostForm(backend_url, r.Form)

		if err != nil {
			error_message = "Service is currently unavailable - Code: 001"
			log.Printf("Error connecting with host: %v", err)
		} else {
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			
			var backend_response RemoteReponse
			err = json.Unmarshal(body, &backend_response)
			
			if err != nil {
				error_message = "Service is currently unavailable - Code: 002"
				log.Printf("Error reading json %v", err)
			}
			
			if backend_response.Success == true {
				resp.Output["uid"] = backend_response.Uid
				return true
			} else {
				error_message = backend_response.Error
			}
		}
	}

	post_url := fmt.Sprintf("/authorize?response_type=%s&client_id=%s&state=%s&redirect_uri=%s", ar.Type, ar.Client.Id, ar.State, url.QueryEscape(ar.RedirectUri))

	var t = template.Must(template.New("login.html").ParseFiles("login.html"))
	t.Execute(w, &PageContext{Error: error_message, PostUrl: post_url})

	return false
}

func load_clients(storage *InMemoryStorage) {

	file, e := ioutil.ReadFile("clients.json")

    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

	var clients []osin.Client
	err := json.Unmarshal(file, &clients)
	
	if err != nil {
		log.Printf("Error reading json %v", err)
		os.Exit(1)
	}

	for _, client := range clients {
		storage.SetClient(client.Id, &client)
	}
}
