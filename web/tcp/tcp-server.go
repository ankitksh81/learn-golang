/*----------------
	TCP SERVER
------------------*/

package main

import (
	"log"
	"net"

	"github.com/spf13/viper"
)

func viperEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error readign from config file: %v", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type")
	}

	return value
}

var HOST = viperEnv("HOST")
var PORT = viperEnv("PORT")
var CONN_TYPE = viperEnv("CONN_TYPE")

func main() {
	listner, err := net.Listen(CONN_TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Error starting tcp server: %v", err)
	}
	defer listner.Close()

	log.Println("Listening on " + HOST + ":" + PORT)

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %v", err.Error())
		}
		log.Println(conn)
	}
}
