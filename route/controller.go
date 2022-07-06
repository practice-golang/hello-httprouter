package route

import (
	"fmt"
	"hello-httprouter/model"
	"log"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println(r.URL.Query().Get("age"))

	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()

	result, err := json.Marshal(r.Form)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(r.FormValue("name") + "\n"))
	w.Write([]byte(r.FormValue("password") + "\n"))

	for k, vs := range r.PostForm {
		fmt.Fprintf(w, "%s: %s\n", k, vs)
		for _, v := range vs {
			fmt.Fprintf(w, "%s: %s\n", k, v)
		}
	}

	w.Write(result)
}

func User(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := model.UserInfo{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	result, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
