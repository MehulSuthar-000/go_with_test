// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func Greet(writer io.Writer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// func main() {
// 	Greet(os.Stdout, "Elodie")
// }

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
