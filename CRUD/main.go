package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Y-elv/gop-bn-usefull.git/routes"
	"github.com/Y-elv/gop-bn-usefull.git/common"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to My Blog application!")
}
func main() {
	// then time for using routers
	
	r := routes.NewRouter()

	r.HandleFunc("/", WelcomeHandler)
    common.InitMongoDB()
	port := "8000"
	fmt.Printf("Server running on port %s ...\n", port)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
