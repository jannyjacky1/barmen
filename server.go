package main

import (
	"github.com/jannyjacky1/barmen/api/client/v1"
	"github.com/jannyjacky1/barmen/api/client/v1/protogen"
	"github.com/jannyjacky1/barmen/tools"
	"github.com/robfig/cron"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

var app tools.App

func init() {
	app = tools.GetApp()
}

func main() {

	c := cron.New()
	c.AddFunc("1 * * * *", func() {
		tools.SetCocktailOfDay(app)
	})
	c.Start()

	lis, err := net.Listen("tcp", app.Config.Host+":"+app.Config.Port)
	if err != nil {
		app.Log.Panic("failed to listen " + err.Error())
	}

	grpcServer := grpc.NewServer()
	protogen.RegisterDictionariesServer(grpcServer, &v1.DictionariesServer{App: app})
	protogen.RegisterDrinksServer(grpcServer, &v1.DrinksServer{App: app})
	protogen.RegisterIngredientsServer(grpcServer, &v1.IngredientsServer{App: app})
	protogen.RegisterInstrumentsServer(grpcServer, &v1.InstrumentsServer{App: app})

	go func() {
		app.Log.Info("start grpc server port: " + app.Config.Port)
		err = grpcServer.Serve(lis)
		if err != nil {
			app.Log.Panic(err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	app.Log.Info("stopping grpc server...")
	grpcServer.GracefulStop()
	c.Stop()
}
