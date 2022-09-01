package go_server

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("./Static"))
	http.Handle("/", fileServer)
	http.HandleFunc()
}
