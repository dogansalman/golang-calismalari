package  main

import (
	"net/http"
	"fmt"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.Path)
	fmt.Fprint(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func main()  {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8880",nil))
}
