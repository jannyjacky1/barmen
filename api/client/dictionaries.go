package client

import (
	"context"
	"github.com/jannyjacky1/barmen/db"
	"github.com/jannyjacky1/barmen/proto"
	"log"
)

type DictionariesServer struct {
}

func (s *DictionariesServer) GetDictionaries(ctx context.Context, request *proto.DictionariesRequest) (*proto.DictionariesResponse, error) {

	complicationLevels, err := getComplicationLevels(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	fortressLevels, err := getFortressLevels(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	volumes, err := getVolumes(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	response := proto.DictionariesResponse{ComplicationLevels: complicationLevels, FortressLevels: fortressLevels, Volumes: volumes}
	return &response, nil
}

func getItemsFromTable(ctx context.Context, tableName string) ([]*proto.Dictionary, error) {
	database := db.Database()
	query := "SELECT id, name FROM " + tableName
	var items []*proto.Dictionary
	err := database.SelectContext(ctx, &items, query)

	return items, err
}

func getComplicationLevels(ctx context.Context) ([]*proto.Dictionary, error) {
	return getItemsFromTable(ctx, "tbl_complication_levels")
}

func getFortressLevels(ctx context.Context) ([]*proto.Dictionary, error) {
	return getItemsFromTable(ctx, "tbl_fortress_levels")
}

func getVolumes(ctx context.Context) ([]*proto.Dictionary, error) {
	return getItemsFromTable(ctx, "tbl_volumes")
}
