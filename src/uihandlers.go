package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListRecipes(w http.ResponseWriter, r *http.Request) {
	
	wRet := UpdateHeader( w, 200 );

	t := ListRecipesController();

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func ListMachines(w http.ResponseWriter, r *http.Request) {
	
	wRet := UpdateHeader( w, 200 );

	t := ListMachinesController();

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func ListSessions(w http.ResponseWriter, r *http.Request) {
	log.Printf( "ListSessions")
	wRet := UpdateHeader( w, 200 );

	t := ListSessionsController();

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func ImportRecipe(w http.ResponseWriter, r *http.Request) {
	
	vars :=mux.Vars( r )
	name, _ := ( vars["name"] )

	if( !ImportRecipeController( name ) ) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
	}
}

func ImportSession(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var inmsg ImportSessionData
	if err := json.Unmarshal(body, &inmsg); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	if( !ImportSessionController( inmsg ) ) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
	}
}

func UploadFiles(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var inmsg UploadFilesData
	if err := json.Unmarshal(body, &inmsg); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	if( !UploadFilesController( inmsg ) ) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
	}

}

// Determine datastructure coming over
// Propogate boolean through the function calls instead of just err
func UpdateRecipeFromHeader(w http.ResponseWriter, r *http.Request) {

	wRet := UpdateHeader( w, 200 );

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var inmsg UpdateRecipeList
	if err := json.Unmarshal(body, &inmsg); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	vars :=mux.Vars( r )
	id, _ := strconv.Atoi( vars["id"] )

	t := UpdateUIRecipe( id, inmsg ) 

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
	}
}


func GetRecipe(w http.ResponseWriter, r *http.Request) {
	
	wRet := UpdateHeader( w, 200 );

	vars :=mux.Vars( r )
	id, _ := strconv.Atoi( vars["id"] )

	t := ViewRecipe( id );

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func DeleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	
	wRet := UpdateHeader( w, 200 );

	vars :=mux.Vars( r )
	id, _ := strconv.Atoi( vars["id"] )

	t := DeleteUIRecipe( id );

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func GetSessionHandler(w http.ResponseWriter, r *http.Request) {
	
	wRet := UpdateHeader( w, 200 );

	vars :=mux.Vars( r )
	id, _ := strconv.Atoi( vars["id"] )

	t := GetUISession( id );

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func DeleteImportableRecipeHandler(w http.ResponseWriter, r *http.Request) {
	wRet := UpdateHeader( w, 200 );

	vars :=mux.Vars( r )
	id, _ := strconv.Atoi( vars["id"] )

	t := DeleteImportableRecipe( id );

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func DeleteMachineHandler(w http.ResponseWriter, r *http.Request) {

	wRet := UpdateHeader( w, 200 );

	vars :=mux.Vars( r )
	token := vars["token"]

	t := DeleteUIMachine( token );

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func DeleteImportableSessionHandler(w http.ResponseWriter, r *http.Request) {
	wRet := UpdateHeader( w, 200 );

	vars :=mux.Vars( r )
	id, _ := strconv.Atoi( vars["id"] )

	t := DeleteImportableSession( id );

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func DeleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	
	wRet := UpdateHeader( w, 200 );

	vars :=mux.Vars( r )
	id, _ := strconv.Atoi( vars["id"] )

	t := DeleteUISession( id );

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}


func GetLogs(w http.ResponseWriter, r *http.Request) {
	
	wRet := UpdateHeader( w, 200 );

	vars :=mux.Vars( r )

	id, _ := strconv.Atoi( vars["id"] )

	t := GetLogsForSession( id );

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func GetImportableItems(w http.ResponseWriter, r *http.Request) {

	wRet := UpdateHeader( w, 200 );

	t := ListImportablesController();

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

type DefaultRespMsg struct {
	Value      string    `json:"value"`
	Err      string    `json:"error"`
}