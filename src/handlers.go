package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/build/version.html")
}

func UpdateHeader( w http.ResponseWriter,  retcode int ) http.ResponseWriter {
	w.Header().Set("Cache-Control", "private")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Server", "Microsoft-IIS/10.0")
	w.Header().Set("X-AspNetWebPages-Version", "3.0")
	w.Header().Set("X-AspNet-Version", "4.0.30319")
	w.Header().Set("Request-Context", "appId=cid-v1:b51654f1-5f75-4752-82da-97ecb1c4d76d")
	w.Header().Set("Access-Control-Expose-Headers", "Request-Context")
	w.Header().Set("X-Powered-By", "ASP.NET")
	
	w.WriteHeader(retcode) // unprocessable entity

	return w;
}

func PicoBrewZState(w http.ResponseWriter, r *http.Request) {
	var zsess ZState
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &zsess); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	wRet := UpdateHeader( w, 200 );



	values, _ := url.ParseQuery( r.URL.RawQuery )

	token := values.Get( "token" )

	var resp ZStateResponse

	t := createZStateResponse(resp, token)

	updateMachineStatusIdle( token )

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}

	log.Printf("%s\t%d", zsess.CurrentFirmware, zsess.BoilerType)
}


func PicoBrewRecipeRefListController(w http.ResponseWriter, r *http.Request) {
	var inmsg RecipeRefListControllerMsg
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inmsg); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	wRet := UpdateHeader( w, 200 );
	
	t := createRecipeRefListControllerRespMsg(inmsg)

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}

	log.Printf("%d\t%d", inmsg.Kind, inmsg.MaxCount)
}

func PicoBrewRecipeController(w http.ResponseWriter, r *http.Request) {
    IDs, ok := r.URL.Query()["id"]
    
    if !ok || len(IDs[0]) < 1 {
        log.Println("Url Param 'id' is missing")
        return
    }

	wRet := UpdateHeader( w, 200 );
	
	id, err := strconv.Atoi( IDs[0] )
	
	if( err != nil ) {
		panic( err )
	}

	t:= createRecipeControllerRespMsg( id )

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}

	log.Printf("%d\t%s", t.ID, t.Name )
}

func PicoBrewSessionController(w http.ResponseWriter, r *http.Request) {
	var inmsg SessionControllerMsg
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inmsg); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	wRet := UpdateHeader( w, 200 );

	values, _ := url.ParseQuery( r.URL.RawQuery )

	token := values.Get( "token" )
	
	t := createSessionControllerRespMsg(inmsg, token)

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}

func PicoBrewSessionLogController(w http.ResponseWriter, r *http.Request) {
	var inmsg SessionLogMsg
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &inmsg); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	wRet := UpdateHeader( w, 200 );

	values, _ := url.ParseQuery( r.URL.RawQuery )

	token := values.Get( "token" )

	t := createSessionLogRespMsg(inmsg, token)

	if err := json.NewEncoder(wRet).Encode(t); err != nil {
		panic(err)
	}
}



