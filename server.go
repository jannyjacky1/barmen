package main

import (
	"encoding/json"
	"github.com/jannyjacky1/barmen/tools"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Config struct {
	host       string
	port       string
	logFile    string
	dbHost     string
	dbName     string
	dbUser     string
	dbPassword string
}

var config Config

func SomeHttpHandler(w http.ResponseWriter, req *http.Request) {
	url, _ := json.Marshal(req.URL)
	log.Println(req.URL.Path)
	w.WriteHeader(200)
	w.Write(url)
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	config = Config{
		tools.GetEnv("HOST", ""),
		tools.GetEnv("PORT", "8080"),
		tools.GetEnv("LOG_FILE", "app.log"),
		tools.GetEnv("DB_HOST", ""),
		tools.GetEnv("DB_NAME", ""),
		tools.GetEnv("DB_USER", ""),
		tools.GetEnv("DB_PASSWORD", ""),
	}
	//TODO: config errors
}

func main() {

	//TODO: always open
	log.SetFlags(log.Ldate | log.Ltime)
	f, err := os.OpenFile(config.logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	mux := http.NewServeMux()
	mux.HandleFunc("/", SomeHttpHandler)
	mux.HandleFunc("/hello", SomeHttpHandler)

	server := &http.Server{
		Addr:    ":" + config.port,
		Handler: mux,
	}

	log.Fatalln(server.ListenAndServe())
}
