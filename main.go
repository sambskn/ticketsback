package main

import (
	"net/http"

	"github.com/gorilla/websocket"

	"flag"
	"log"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// // seed data for albums (cuz no db)
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use defaults lol

func ticketDispenser(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ticket", ticketDispenser)
	log.Fatal(http.ListenAndServe(*addr, nil))

	// router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.POST("/albums", postAlbums)

	// router.Run("localhost:8080")

}

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	// use BindJSON to bind the post data to newAlbum
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }
