package network

import (
	"emp-parser/server/globals"
	"fmt"
	"log"
	"net/http"
	"os"
)

var networkLogger = log.New(os.Stdout, "[network] ", 0)

// StartWebServer Démarrage du serveur web (http + websockets)
func StartWebServer(port int) {
	networkLogger.Printf("Serveur web en écoute sur le port %d.\n", port)
	http.Handle("/", http.FileServer(http.Dir("../client")))
	globals.Ch <- 1
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		globals.ErrLogger.Println("Erreur à la création du serveur HTTP : ", err.Error())
	}
	globals.Ch <- 1
}
