package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const VERSION = "0.9 ALPHA"

func Run(envConf Configuration, router *mux.Router) chan error {

    errs := make(chan error)

    // Starting HTTP server
    go func() {
        log.Printf("Staring HTTP service on %s ...", envConf.Port)
        log.Fatal( http.ListenAndServe( envConf.Port, router) );

        if err := http.ListenAndServe(envConf.Port, nil); err != nil {
            errs <- err
        }

    }()

    // Starting HTTPS server
    go func() {
        if( envConf.EnableTLS ) {
            log.Printf( "TLS ENABLED.  Listening on port: %s", envConf.TLSPort )
            cfg := &tls.Config{
                MinVersion:               tls.VersionTLS12,
                CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
                PreferServerCipherSuites: true,
                CipherSuites: []uint16{
                    tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
                    tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
                    tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
                    tls.TLS_RSA_WITH_AES_256_CBC_SHA,
                },
            }
            srv := &http.Server{
                Addr:         envConf.TLSPort,
                Handler:      router,
                TLSConfig:    cfg,
                TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
            }
    
            if err := srv.ListenAndServeTLS(envConf.CertsPath + "/bundle.crt", envConf.CertsPath + "/server.key"); err != nil {
               errs <- err
            }
        }
    }()

    return errs
}

func main() {
    envConf := getConfiguration()
    setDBCredentials( envConf.Database.AdminUserName, envConf.Database.AdminPassword, envConf.Database.ServerAddress)
    loadMachines();
    loadRecipes();
    router := NewRouter()

    errs := Run( envConf, router )

    // This will run forever until channel receives error
    select {
    case err := <-errs:
        log.Printf("Could not start serving service due to (error: %s)", err)
    }

}