package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("[2023-06-14T06:07:11.817Z]' 'dev.e1-us-east-azure.preview-dv.choreoapis.dev' 'golang-service-new-2510529736.dp-development-default-project-224-321563503.svc.cluster.local' 'GET' '/90bce838-056d-46df-8a0d-0fbaac6d3499/g46r/golang-service-new/service-2-c5d/2.0.0/hello-world' '/hello-world' 'HTTP/1.1' '200' 'via_upstream' '-' 'Mozilla/5.0 (Macintosh  Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36' '7f4cbd2f66b52e957ae7d9c6dba97e2f' '-' '172.19.216.59:8081' '0' '30' '3' '-' '2' '0' '0' '2' '")
	log.Println("Other log")
	fmt.Println("[2023-06-14T06:07:11.817Z]' 'dev.e1-us-east-azure.preview-dv.choreoapis.dev' 'golang-service-new-2510529736.dp-development-default-project-224-321563503.svc.cluster.local' 'PUT' '/834jwfwe-u04r-jg49-rg9s-g9e4jgso9u3r/g46r/golang-service-new/service-2-c5d/2.0.0/hello-world' '/hello-world' 'HTTP/1.1' '200' 'via_upstream' '-' 'Mozilla/5.0 (Macintosh  Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36' '7f4cbd2f66b52e957ae7d9c6dba97e2f' '-' '172.19.216.59:8081' '0' '30' '3' '-' '2' '0' '0' '2' '")
	log.Println("Other log")
	fmt.Println("") // empty line

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(string(b))
	w.WriteHeader(http.StatusOK)
}
