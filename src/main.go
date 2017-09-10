package main
import (
	"encoding/json"
	"fmt"
	"time"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Todo struct {
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}

type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	

	fmt.Println("Starting up on 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(todos)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo index")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	fmt.Fprintln(w, "Todo show:", todoId)
}