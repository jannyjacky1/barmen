package client

import (
	"context"
	"database/sql"
	"github.com/jackc/pgtype"
	"github.com/jannyjacky1/barmen/proto"
	"github.com/jannyjacky1/barmen/tools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_ "log"
	"strings"
)

type DrinksServer struct {
	App tools.App
}

type CocktailSearchParams struct {
	Offset       int32
	PerPage      int32
	Fortress     int32
	Complication int32
	Volume       int32
	IsFlacky     bool
	IsFire       bool
	IsIba        bool
	Includes     *pgtype.Int8Array
	Except       *pgtype.Int8Array
	ExceptId     int32
}

func (s *DrinksServer) GetDrinks(ctx context.Context, request *proto.DrinksRequest) (*proto.DrinksResponse, error) {

	dayCocktailId := 283
	filterQuery, searchParams := prepareFilterParams(request)
	response := proto.DrinksResponse{}

	if filterQuery == "" && dayCocktailId > 0 {
		dayDrink, _, err := getDrink(ctx, s, int32(dayCocktailId), "day-drink")
		if err == nil {
			response.DayDrink = dayDrink
			filterQuery += " WHERE tbl_cocktails.id <> :exceptid"
			searchParams.ExceptId = int32(dayCocktailId)
		}
	}

	drinks, err := getDrinks(ctx, s, filterQuery, searchParams)
	if err != nil {
		return &proto.DrinksResponse{}, err
	}

	response.Drinks = drinks
	return &response, nil
}

func (s *DrinksServer) GetDrinkById(ctx context.Context, request *proto.DrinkRequest) (*proto.DrinkResponse, error) {

	_, item, err := getDrink(ctx, s, request.Id, "cocktail-by-id")

	return item, err
}

func getDrinks(ctx context.Context, s *DrinksServer, filterQuery string, searchParams CocktailSearchParams) ([]*proto.DrinkItem, error) {

	var items []*proto.DrinkItem

	query := "SELECT tbl_cocktails.id, tbl_cocktails.name, CONCAT(tbl_complication_levels.name, ' (', tbl_complication_levels.time, ')') AS complication, tbl_fortress_levels.name AS fortress, is_flacky AS isFlacky, is_fire AS isFire, is_iba AS isIba, coalesce(tbl_files.filepath, '') AS preview, coalesce(string_agg(tbl_ingredients.name, ', '), '') AS ingredients FROM tbl_cocktails INNER JOIN tbl_complication_levels ON tbl_complication_levels.id = tbl_cocktails.complication_id INNER JOIN tbl_fortress_levels ON tbl_fortress_levels.id = tbl_cocktails.fortress_id LEFT JOIN tbl_files ON tbl_files.id = tbl_cocktails.preview_id LEFT JOIN tbl_cocktails_to_tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.cocktail_id = tbl_cocktails.id LEFT JOIN tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.ingredient_id = tbl_ingredients.id"
	queryEnd := " GROUP BY tbl_cocktails.id, tbl_cocktails.weight, tbl_complication_levels.name, tbl_complication_levels.time, tbl_fortress_levels.name, tbl_files.filepath ORDER BY weight DESC LIMIT :perpage OFFSET :offset"

	query += filterQuery

	query += queryEnd
	query, args, err := s.App.Db.BindNamed(query, searchParams)
	if err != nil {
		return items, status.Error(codes.Internal, err.Error())
	}

	err = s.App.Db.SelectContext(ctx, &items, query, args...)
	if err != nil {
		return items, status.Error(codes.Internal, err.Error())
	}
	return items, err
}

func getDrink(ctx context.Context, s *DrinksServer, id int32, scenario string) (*proto.DayDrink, *proto.DrinkResponse, error) {

	var dayDrink proto.DayDrink
	var item proto.DrinkResponse
	var query string
	var err error

	if scenario == "day-drink" {
		query = "SELECT tbl_cocktails.id, tbl_cocktails.name, recipe, tbl_cocktails.description, coalesce(mark, 0) AS mark, CONCAT(tbl_complication_levels.name, ' (', tbl_complication_levels.time, ')') AS complication, tbl_fortress_levels.name AS fortress, is_flacky AS isFlacky, is_fire AS isFire, is_iba AS isIba, coalesce(tbl_files.filepath, '') AS preview FROM tbl_cocktails INNER JOIN tbl_complication_levels ON tbl_complication_levels.id = tbl_cocktails.complication_id INNER JOIN tbl_fortress_levels ON tbl_fortress_levels.id = tbl_cocktails.fortress_id LEFT JOIN tbl_files ON tbl_files.id = tbl_cocktails.preview_id LEFT JOIN tbl_cocktails_to_tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.cocktail_id = tbl_cocktails.id LEFT JOIN tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.ingredient_id = tbl_ingredients.id WHERE tbl_cocktails.id = $1"
	} else {
		query = "SELECT tbl_cocktails.id, tbl_cocktails.name, name_en AS nameEn, recipe, tbl_cocktails.description, coalesce(mark, 0) AS mark, CONCAT(tbl_complication_levels.name, ' (', tbl_complication_levels.time, ')') AS complication, tbl_fortress_levels.name AS fortress, is_flacky AS isFlacky, is_fire AS isFire, is_iba AS isIba FROM tbl_cocktails INNER JOIN tbl_complication_levels ON tbl_complication_levels.id = tbl_cocktails.complication_id INNER JOIN tbl_fortress_levels ON tbl_fortress_levels.id = tbl_cocktails.fortress_id LEFT JOIN tbl_files ON tbl_files.id = tbl_cocktails.preview_id LEFT JOIN tbl_cocktails_to_tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.cocktail_id = tbl_cocktails.id LEFT JOIN tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.ingredient_id = tbl_ingredients.id WHERE tbl_cocktails.id = $1"

	}

	if scenario == "day-drink" {
		err = s.App.Db.GetContext(ctx, &dayDrink, query, id)
	} else {
		err = s.App.Db.GetContext(ctx, &item, query, id)
	}

	if err != nil {
		code := codes.Internal
		if err == sql.ErrNoRows {
			code = codes.NotFound
		}
		return &proto.DayDrink{}, &proto.DrinkResponse{}, status.Error(code, err.Error())
	}

	query = "SELECT ingredient_id AS id, name, CONCAT(volume, ' ', unit) AS volume FROM tbl_cocktails_to_tbl_ingredients INNER JOIN tbl_ingredients ON tbl_ingredients.id = tbl_cocktails_to_tbl_ingredients.ingredient_id WHERE cocktail_id = $1"
	if scenario == "day-drink" {
		s.App.Db.SelectContext(ctx, &dayDrink.Ingredients, query, id)
	} else {
		s.App.Db.SelectContext(ctx, &item.Ingredients, query, id)
	}

	query = "SELECT instrument_id AS id, name FROM tbl_cocktails_to_tbl_instruments INNER JOIN tbl_instruments ON tbl_instruments.id = tbl_cocktails_to_tbl_instruments.instrument_id WHERE cocktail_id = $1"
	if scenario == "day-drink" {
		s.App.Db.SelectContext(ctx, &dayDrink.Instruments, query, id)
	} else {
		s.App.Db.SelectContext(ctx, &item.Instruments, query, id)
	}

	return &dayDrink, &item, nil
}

func prepareFilterParams(request *proto.DrinksRequest) (string, CocktailSearchParams) {
	var queryWhere []string

	searchParams := CocktailSearchParams{}

	searchParams.PerPage = 50
	searchParams.Offset = request.Page * searchParams.PerPage

	if request.Fortress > 0 {
		queryWhere = append(queryWhere, "tbl_cocktails.fortress_id = :fortress")
		searchParams.Fortress = request.Fortress
	}
	if request.Complication > 0 {
		queryWhere = append(queryWhere, "tbl_cocktails.complication_id = :complication")
		searchParams.Complication = request.Complication
	}
	if request.Volume > 0 {
		queryWhere = append(queryWhere, "tbl_cocktails.volume_id = :volume")
		searchParams.Volume = request.Volume
	}
	if request.IsFlacky {
		queryWhere = append(queryWhere, "tbl_cocktails.is_flacky = :isflacky")
		searchParams.IsFlacky = true
	}
	if request.IsFire {
		queryWhere = append(queryWhere, "tbl_cocktails.is_fire = :isfire")
		searchParams.IsFire = true
	}
	if request.IsIba {
		queryWhere = append(queryWhere, "tbl_cocktails.is_iba = :isiba")
		searchParams.IsIba = true
	}
	if len(request.Includes) > 0 {
		queryWhere = append(queryWhere, ":includes <@ ARRAY(SELECT ci.ingredient_id FROM tbl_cocktails_to_tbl_ingredients AS ci WHERE ci.cocktail_id = tbl_cocktails.id)")
		searchParams.Includes = &pgtype.Int8Array{}
		searchParams.Includes.Set(request.Includes)
	}
	if len(request.Except) > 0 {
		queryWhere = append(queryWhere, "NOT(:except && ARRAY(SELECT ci.ingredient_id FROM tbl_cocktails_to_tbl_ingredients AS ci WHERE ci.cocktail_id = tbl_cocktails.id))")
		searchParams.Except = &pgtype.Int8Array{}
		searchParams.Except.Set(request.Except)
	}

	if len(queryWhere) > 0 {
		return " WHERE " + strings.Join(queryWhere, " AND "), searchParams
	}

	return "", searchParams
}
