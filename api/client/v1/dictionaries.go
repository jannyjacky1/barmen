package v1

import (
	"context"
	"database/sql"
	"github.com/jannyjacky1/barmen/api/client/v1/protogen"
	"github.com/jannyjacky1/barmen/tools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type DictionariesServer struct {
	App tools.App
}

func (s *DictionariesServer) GetDictionaries(ctx context.Context, _ *protogen.DictionariesRequest) (*protogen.DictionariesResponse, error) {

	complicationLevels, err := getComplicationLevels(ctx, s)
	if err != nil {
		s.App.Log.Error(err.Error())
	}

	fortressLevels, err := getFortressLevels(ctx, s)
	if err != nil {
		s.App.Log.Error(err.Error())
	}

	volumes, err := getVolumes(ctx, s)
	if err != nil {
		s.App.Log.Error(err.Error())
	}

	ingredients, err := getIngredients(ctx, s)
	if err != nil {
		s.App.Log.Error(err.Error())
	}

	other, err := getOther(ctx, s)
	if err != nil {
		s.App.Log.Error(err.Error())
	}

	response := protogen.DictionariesResponse{ComplicationLevels: complicationLevels, FortressLevels: fortressLevels, Volumes: volumes, Ingredients: ingredients, Other: other}
	return &response, nil
}

func (s *DictionariesServer) GetByName(ctx context.Context, request *protogen.NameRequest) (*protogen.NameResponse, error) {
	var items []*protogen.NameItem
	perPage := 100
	offset := perPage * int(request.Page)
	query := "SELECT id, name, 0 AS type FROM tbl_cocktails WHERE LOWER(name) LIKE $1 UNION SELECT id, name, 1 AS type FROM tbl_ingredients WHERE LOWER(name) LIKE $1 UNION SELECT id, name, 2 AS type FROM tbl_instruments WHERE LOWER(name) LIKE $1 LIMIT $2 OFFSET $3"
	err := s.App.Db.SelectContext(ctx, &items, query, "%"+strings.ToLower(request.Name)+"%", perPage, offset)
	if err != nil && err != sql.ErrNoRows {
		s.App.Log.Error(err.Error())
		return &protogen.NameResponse{}, status.Error(codes.Internal, err.Error())
	}

	response := protogen.NameResponse{Items: items}
	return &response, nil
}

func getItemsFromTable(ctx context.Context, s *DictionariesServer, tableName string, sortColumn string) ([]*protogen.Dictionary, error) {
	query := "SELECT id, name FROM " + tableName + " ORDER BY " + sortColumn + " ASC"
	var items []*protogen.Dictionary
	err := s.App.Db.SelectContext(ctx, &items, query)

	return items, err
}

func getComplicationLevels(ctx context.Context, s *DictionariesServer) ([]*protogen.Dictionary, error) {
	return getItemsFromTable(ctx, s, "tbl_complication_levels", "id")
}

func getFortressLevels(ctx context.Context, s *DictionariesServer) ([]*protogen.Dictionary, error) {
	return getItemsFromTable(ctx, s, "tbl_fortress_levels", "id")
}

func getVolumes(ctx context.Context, s *DictionariesServer) ([]*protogen.Dictionary, error) {
	return getItemsFromTable(ctx, s, "tbl_volumes", "id")
}

func getIngredients(ctx context.Context, s *DictionariesServer) ([]*protogen.Dictionary, error) {
	return getItemsFromTable(ctx, s, "tbl_ingredients", "name")
}

func getOther(ctx context.Context, s *DictionariesServer) ([]*protogen.Dictionary, error) {
	return []*protogen.Dictionary{
		&protogen.Dictionary{Id: 1, Name: "Слоеный"},
		&protogen.Dictionary{Id: 2, Name: "Горящий"},
		&protogen.Dictionary{Id: 3, Name: "Коктейль IBA"},
	}, nil
}
