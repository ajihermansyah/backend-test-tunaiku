package main

import (
	"fmt"
	"net/http"
)

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

const helloTunaiKu = "<html><body><h1><center>Hello <font color=green>Tunai</font><font color=blue>ku!</font></center></h1></body></html>"

func main() {
	fmt.Println("Run http://localhost:8080 on browser")
	http.HandleFunc("/add", Add)
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)

}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	if url == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, AddForm)
		return
	}
	fmt.Fprintf(w, "http://localhost:8080/%s", url)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, helloTunaiKu)
	return
}
