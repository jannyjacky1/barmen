package client

import (
	"context"
	"github.com/jannyjacky1/barmen/proto"
	"github.com/jannyjacky1/barmen/tools"
	"log"
)

type DictionariesServer struct {
	App tools.App
}

func (s *DictionariesServer) GetDictionaries(ctx context.Context, request *proto.DictionariesRequest) (*proto.DictionariesResponse, error) {

	complicationLevels, err := getComplicationLevels(ctx, s)
	if err != nil {
		log.Fatalln(err)
	}

	fortressLevels, err := getFortressLevels(ctx, s)
	if err != nil {
		log.Fatalln(err)
	}

	volumes, err := getVolumes(ctx, s)
	if err != nil {
		log.Fatalln(err)
	}

	ingredients, err := getIngredients(ctx, s)
	if err != nil {
		log.Fatalln(err)
	}

	other, err := getOther(ctx, s)
	if err != nil {
		log.Fatalln(err)
	}

	response := proto.DictionariesResponse{ComplicationLevels: complicationLevels, FortressLevels: fortressLevels, Volumes: volumes, Ingredients: ingredients, Other: other}
	return &response, nil
}

func getItemsFromTable(ctx context.Context, s *DictionariesServer, tableName string) ([]*proto.Dictionary, error) {
	query := "SELECT id, name FROM " + tableName
	var items []*proto.Dictionary
	err := s.App.Db.SelectContext(ctx, &items, query)

	return items, err
}

func getComplicationLevels(ctx context.Context, s *DictionariesServer) ([]*proto.Dictionary, error) {
	return getItemsFromTable(ctx, s, "tbl_complication_levels")
}

func getFortressLevels(ctx context.Context, s *DictionariesServer) ([]*proto.Dictionary, error) {
	return getItemsFromTable(ctx, s, "tbl_fortress_levels")
}

func getVolumes(ctx context.Context, s *DictionariesServer) ([]*proto.Dictionary, error) {
	return getItemsFromTable(ctx, s, "tbl_volumes")
}

func getIngredients(ctx context.Context, s *DictionariesServer) ([]*proto.Dictionary, error) {
	return getItemsFromTable(ctx, s, "tbl_ingredients")
}

func getOther(ctx context.Context, s *DictionariesServer) ([]*proto.Dictionary, error) {
	return []*proto.Dictionary{
		&proto.Dictionary{Id: 1, Name: "Слоеный"},
		&proto.Dictionary{Id: 2, Name: "Горящий"},
		&proto.Dictionary{Id: 1, Name: "Коктейль IBA"},
	}, nil
}
