package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetchURL(url string, wg *sync.WaitGroup) {
	fmt.Println("now running for url ", url)
	out, error := http.Get(url)
	defer wg.Done()
	if error != nil {
		fmt.Println("error foudn in ", error)
		return
	}
	fmt.Println(out.Status)

}

func homePage(w http.ResponseWriter, req *http.Request) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}
	wg.Wait()
	fmt.Println("now all fetch is onde")
	fmt.Fprintln(w, "now all responses are done")
}

func rootPage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("now running root request")
	fmt.Fprintln(w, "now all responses are done from root")
}

func handleRequests() {
	fmt.Println("now intializing handlerequest")
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/home", homePage)
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func main() {
	handleRequests()
}

// package main

// import (
// 	"fmt"

// 	"github.com/rajeshpachar/hellomod/child"
// 	"github.com/rajeshpachar/hellomod/hellotest"
// )

// func main() {
// 	hellotest.SayHello("client calling")
// 	hellotest.DepHello()
// 	child.HelloChild()
// 	fmt.Println("now calling private child ")
// 	fmt.Println("time to say hello to internal module ")
// }
