package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type DatabaseConfiguration struct {
	AdminUserName string
	AdminPassword string
	ServerAddress string
}

type Configuration struct {
    SessionPath string
    RecipePath string
    CertsPath string
    EnableTLS bool
    TLSPort string
	Port string
	Database DatabaseConfiguration
}

func getDefaultConfig( inValue string, defaultValue string ) string {

    if( strings.EqualFold( inValue, "" ) ) {
        return defaultValue
    }

    return inValue
}

func getConfiguration() Configuration {
    configuration := Configuration{}

    configuration.SessionPath = getDefaultConfig( os.Getenv( "PB_SESSION_PATH"), "/planbeer/sessions" ) 
    configuration.RecipePath = getDefaultConfig(os.Getenv( "PB_RECIPE_PATH"), "/planbeer/recipes" )
    configuration.CertsPath = getDefaultConfig( os.Getenv( "PB_CERTS_PATH"), "/planbeer/certs" )
    //@TODO: Default is false so this works out.
    configuration.EnableTLS = strings.EqualFold( os.Getenv( "PB_ENABLE_TLS"), "TRUE" ) 
    configuration.Port = getDefaultConfig( os.Getenv( "PB_PORT" ), ":80" )
    configuration.TLSPort = getDefaultConfig( os.Getenv( "PB_TLS_PORT" ), ":443" )
    configuration.Database = DatabaseConfiguration {}
    configuration.Database.AdminUserName = getDefaultConfig( os.Getenv( "PB_DBADMIN"), "admin" )
    configuration.Database.AdminPassword = getDefaultConfig( os.Getenv( "PB_DBPASSWORD"), "password" )
    configuration.Database.ServerAddress = getDefaultConfig( os.Getenv( "PB_DBSERVER"), "localhost" )

    fmt.Println( "--Server Configuration--")
    outConf, _ := json.Marshal(configuration)
    fmt.Println(string(outConf))
    fmt.Println( "--Server Configuration--")

    err := os.MkdirAll(configuration.SessionPath, os.ModePerm)
    if( err != nil ) {
        fmt.Println( "Can't create dir: " + configuration.SessionPath )
        fmt.Println( err )
    }

    err = os.MkdirAll(configuration.RecipePath, os.ModePerm)
    if( err != nil ) {
        fmt.Println( "Can't create dir: " + configuration.RecipePath )
        fmt.Println( err )
    }


    return configuration
}

/*
 * JSON Configuration Alternative
 *
* FORMAT:
{
    "SessionPath" : "/planbeer/sessions",
    "RecipePath" : "/planbeer/recipes",
    "CertsPath" : "",
    "EnableTLS" : false,
    "Port" : ":80",
    "Database" : {
	"AdminUserName" : "admin",
	"AdminPassword" : "password",
	"ServerAddress" : "192.168.0.135"
     }
}

* FUNCTION:
func getConfiguration() Configuration {
    file, _ := os.Open("/planbeer/conf/conf.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Println( "Read Configuration")
        outConf, _ := json.Marshal(configuration)
        fmt.Println(string(outConf))
    }

    return configuration;
}
*/