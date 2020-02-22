package main

import (
	"encoding/json"
	"github.com/jannyjacky1/barmen/api/client"
	"github.com/jannyjacky1/barmen/proto"
	"github.com/jannyjacky1/barmen/tools"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

type Config struct {
	host    string
	port    string
	logFile string
}

var appConfig Config

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

	appConfig = Config{
		tools.GetEnv("HOST", ""),
		tools.GetEnv("PORT", "8080"),
		tools.GetEnv("LOG_FILE", "app.log"),
	}
	//TODO: AppConfig errors
}

func main() {

	//TODO: always open
	log.SetFlags(log.Ldate | log.Ltime)
	f, err := os.OpenFile(appConfig.logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	lis, err := net.Listen("tcp", appConfig.host+":"+appConfig.port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterDictionariesServer(grpcServer, &client.DictionariesServer{})
	grpcServer.Serve(lis)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", SomeHttpHandler)
	//mux.HandleFunc("/hello", SomeHttpHandler)
	//
	//server := &http.Server{
	//	Addr:    ":" + AppConfig.port,
	//	Handler: mux,
	//}
	//
	//log.Fatalln(server.ListenAndServe())
}
