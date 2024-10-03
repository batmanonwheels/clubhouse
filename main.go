package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// load all html files and components
	router.LoadHTMLFiles("templates/index.html", "templates/components/head.html")

	// load static directory
	router.StaticFS("static", http.Dir("./static"))

	router.SetTrustedProxies(nil)

	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		i := 0
		for {
			i++
			conn.WriteMessage(websocket.TextMessage, []byte("Hello Websocket!"))
			time.Sleep(time.Second)
		}
	})

	// site routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
