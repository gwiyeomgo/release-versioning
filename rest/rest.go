package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type url string

//const port string = ":4000"
var port string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func versions(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//version , err := release.Find()
		//block, err := relase.FindBlock(hash)
		//json.NewEncoder(rw).Encode(blockchain.Blocks(blockchain.Blockchain()))
	case "POST":

		//newBlock := relase.Blockchain().AddBlock()

		//rw.WriteHeader(http.StatusCreated)

	}
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	// data 는 Go의 세계에 있는 slice
	//struct 의 slice
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "test",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
	}

	json.NewEncoder(rw).Encode(data)

}

//모든 request에 content-type을 설정하는 middlewares 추가하기
//middleware 는 function 인데 먼저 호출되고
//다음 function을 부르고 그럼 거기서 또 다음 function을 부른다

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		//response
		rw.Header().Add("Content-Type", "application/json")
		//(2) 다음 handlerFunc 호출
		next.ServeHTTP(rw, request)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL)
		next.ServeHTTP(rw, request)
	})
}

func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.Use(jsonContentTypeMiddleware, loggerMiddleware)
	router.HandleFunc("/release/versioning", versions).Methods("GET", "POST")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
