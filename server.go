package main

import (
	"github.com/jannyjacky1/barmen/api/client"
	"github.com/jannyjacky1/barmen/protogen"
	"github.com/jannyjacky1/barmen/tools"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var app tools.App

func init() {
	app = tools.GetApp()
}

func main() {

	//TODO: always open
	log.SetFlags(log.Ldate | log.Ltime)
	f, err := os.OpenFile(app.Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	lis, err := net.Listen("tcp", app.Config.Host+":"+app.Config.Port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	protogen.RegisterDictionariesServer(grpcServer, &client.DictionariesServer{app})
	protogen.RegisterDrinksServer(grpcServer, &client.DrinksServer{app})
	protogen.RegisterIngredientsServer(grpcServer, &client.IngredientsServer{app})
	protogen.RegisterInstrumentsServer(grpcServer, &client.InstrumentsServer{app})
	grpcServer.Serve(lis)
}
