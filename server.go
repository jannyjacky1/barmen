package main

import (
	client "github.com/jannyjacky1/barmen/api/client/v1"
	"github.com/jannyjacky1/barmen/api/client/v1/protogen"
	manager "github.com/jannyjacky1/barmen/api/manager/v1"
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
	c.AddFunc("@daily * * * *", func() {
		tools.SetCocktailOfDay(app)
	})
	c.Start()

	managerApp := manager.App(app)
	go func() {
		app.Log.Info("start rest server port: " + app.Config.ManagerPort)
		err := managerApp.Start(app.Config.Host + ":" + app.Config.ManagerPort)
		if err != nil {
			app.Log.Panic(err.Error())
		}
	}()

	lis, err := net.Listen("tcp", app.Config.Host+":"+app.Config.Port)
	if err != nil {
		app.Log.Panic("failed to listen " + err.Error())
	}

	grpcServer := grpc.NewServer()
	protogen.RegisterDictionariesServer(grpcServer, &client.DictionariesServer{App: app})
	protogen.RegisterDrinksServer(grpcServer, &client.DrinksServer{App: app})
	protogen.RegisterIngredientsServer(grpcServer, &client.IngredientsServer{App: app})
	protogen.RegisterInstrumentsServer(grpcServer, &client.InstrumentsServer{App: app})

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
