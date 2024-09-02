package main

import (
	websocket "botxboom/hi5/pkg"
	"fmt"
	"log"
	"net/http"
)


func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host)

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes(){

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		serveWs(pool, w, r)
	})

	
}

func main(){
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}