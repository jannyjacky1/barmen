package v1

import (
	"context"
	"database/sql"
	"github.com/jackc/pgtype"
	"github.com/jannyjacky1/barmen/api/client/v1/protogen"
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
	Fortress     *pgtype.Int8Array
	Complication *pgtype.Int8Array
	Volume       *pgtype.Int8Array
	IsFlacky     bool
	IsFire       bool
	IsIba        bool
	Includes     *pgtype.Int8Array
	Except       *pgtype.Int8Array
	Instruments  *pgtype.Int8Array
	Similar      *pgtype.Int8Array
	ExceptId     int32
}

func (s *DrinksServer) GetDrinks(ctx context.Context, request *protogen.DrinksRequest) (*protogen.DrinksResponse, error) {

	filterQuery, searchParams, errors := prepareFilterParams(request)
	for i := 0; i < len(errors); i++ {
		s.App.Log.Warn(errors[i].Error())
	}

	response := protogen.DrinksResponse{}

	drinks, err := getDrinks(ctx, s, filterQuery, searchParams)
	if err != nil {
		s.App.Log.Error(err.Error())
		return &protogen.DrinksResponse{}, err
	}

	response.Drinks = drinks
	return &response, nil
}

func (s *DrinksServer) GetDrinkOfDay(ctx context.Context, _ *protogen.Empty) (*protogen.DrinkOfDayResponse, error) {

	dayCocktailId := 0
	response := &protogen.DrinkOfDayResponse{}

	err := s.App.Db.GetContext(ctx, &dayCocktailId, "SELECT CAST(value AS integer) FROM tbl_settings WHERE alias = 'day_cocktail'")
	if err != nil {
		s.App.Log.Error(err.Error())
		return response, nil
	}

	if dayCocktailId > 0 {
		response, _, err = getDrink(ctx, s, int32(dayCocktailId), "day-drink")
		if err != nil {
			s.App.Log.Error(err.Error())
			return response, err
		}
	}

	return response, nil
}

func (s *DrinksServer) GetDrinkById(ctx context.Context, request *protogen.DrinkRequest) (*protogen.DrinkResponse, error) {

	_, item, err := getDrink(ctx, s, request.Id, "cocktail-by-id")

	if err != nil {
		s.App.Log.Error(err.Error())
	}

	return item, err
}

func (s *DrinksServer) SetDrinkTried(ctx context.Context, request *protogen.DrinkTryRequest) (*protogen.Empty, error) {
	id := 0
	err := s.App.Db.GetContext(ctx, &id, "SELECT id FROM tbl_users WHERE device_id = $1", request.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			stmt, err := s.App.Db.PrepareNamed("INSERT INTO tbl_users (device_id) VALUES (:deviceid) RETURNING id")
			if err != nil {
				s.App.Log.Error(err.Error())
				return &protogen.Empty{}, status.Error(codes.Internal, err.Error())
			}
			err = stmt.Get(&id, struct {
				DeviceId string
			}{DeviceId: request.UserId})
			if err != nil {
				s.App.Log.Error(err.Error())
				return &protogen.Empty{}, status.Error(codes.Internal, err.Error())
			}
		} else {
			s.App.Log.Error(err.Error())
			return &protogen.Empty{}, status.Error(codes.Internal, err.Error())
		}
	}

	_, err = s.App.Db.Exec("INSERT INTO tbl_tries (user_id, cocktail_id) VALUES ($1, $2)", id, request.Id)

	if err != nil {
		s.App.Log.Error(err.Error())
		return &protogen.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &protogen.Empty{}, nil
}

func (s *DrinksServer) SetDrinkMark(ctx context.Context, request *protogen.DrinkMarkRequest) (*protogen.DrinkMarkResponse, error) {

	var result protogen.DrinkMarkResponse
	var mark struct {
		Mark    int
		MarkCnt int
	}
	err := s.App.Db.GetContext(ctx, &mark, "SELECT mark, mark_cnt AS markCnt FROM tbl_cocktails WHERE id = $1", request.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &result, status.Error(codes.NotFound, err.Error())
		}
		return &result, status.Error(codes.Internal, err.Error())
	}
	_, err = s.App.Db.Exec("UPDATE tbl_cocktails SET mark = $1, mark_cnt = $2", mark.Mark+int(request.Mark), mark.MarkCnt+1)
	if err != nil {
		s.App.Log.Error(err.Error())
		return &result, status.Error(codes.Internal, err.Error())
	}

	err = s.App.Db.GetContext(ctx, &result, "SELECT ROUND(CAST(mark AS decimal)/GREATEST(mark_cnt,1)) AS mark, CONCAT(ROUND(CAST(mark AS decimal)/GREATEST(mark_cnt,1), 1), ' (по ', mark_cnt, ' оценкам)') AS markDescription FROM tbl_cocktails WHERE id = $1", request.Id)
	if err != nil {
		s.App.Log.Error(err.Error())
		return &result, status.Error(codes.Internal, err.Error())
	}

	return &result, nil
}

func getDrinks(ctx context.Context, s *DrinksServer, filterQuery string, searchParams CocktailSearchParams) ([]*protogen.DrinkItem, error) {

	var items []*protogen.DrinkItem
	query := "SELECT tbl_cocktails.id, tbl_cocktails.name, CONCAT(tbl_fortress_levels.name, ', ', tbl_complication_levels.name) AS properties, ROUND(CAST(mark AS decimal)/GREATEST(mark_cnt,1)) AS mark, is_flacky AS isFlacky, is_fire AS isFire, is_iba AS isIba, CONCAT('" + s.App.Config.MediaUrl + "', coalesce(tbl_files.filepath, '')) AS icon, coalesce(string_agg(tbl_ingredients.name, ', '), '') AS ingredients FROM tbl_cocktails INNER JOIN tbl_complication_levels ON tbl_complication_levels.id = tbl_cocktails.complication_id INNER JOIN tbl_fortress_levels ON tbl_fortress_levels.id = tbl_cocktails.fortress_id LEFT JOIN tbl_cocktails_to_tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.cocktail_id = tbl_cocktails.id LEFT JOIN tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.ingredient_id = tbl_ingredients.id LEFT JOIN tbl_files ON tbl_files.id = tbl_cocktails.icon_id"

	filterQuery, args, err := s.App.Db.BindNamed(filterQuery, searchParams)
	query = query + filterQuery

	if err != nil {
		return items, status.Error(codes.Internal, err.Error())
	}

	err = s.App.Db.SelectContext(ctx, &items, query, args...)
	if err != nil {
		return items, status.Error(codes.Internal, err.Error())
	}
	return items, err
}

func getDrink(ctx context.Context, s *DrinksServer, id int32, scenario string) (*protogen.DrinkOfDayResponse, *protogen.DrinkResponse, error) {

	var dayDrink protogen.DrinkOfDayResponse
	var item protogen.DrinkResponse
	var query string
	var err error

	if scenario == "day-drink" {
		query = "SELECT tbl_cocktails.id, tbl_cocktails.name, recipe, tbl_cocktails.description, ROUND(CAST(mark AS decimal)/GREATEST(mark_cnt,1)) AS mark, CONCAT(tbl_complication_levels.name, ' (', tbl_complication_levels.time, ')') AS complication, tbl_fortress_levels.name AS fortress, is_flacky AS isFlacky, is_fire AS isFire, is_iba AS isIba, coalesce(NULLIF(CONCAT('" + s.App.Config.MediaUrl + "', tbl_files.filepath), '" + s.App.Config.MediaUrl + "'), '" + s.App.Config.MediaUrl + "/drink.jpg') AS preview, CONCAT(COUNT(DISTINCT tbl_tries.user_id), ' человек') AS triedBy FROM tbl_cocktails INNER JOIN tbl_complication_levels ON tbl_complication_levels.id = tbl_cocktails.complication_id INNER JOIN tbl_fortress_levels ON tbl_fortress_levels.id = tbl_cocktails.fortress_id LEFT JOIN tbl_files ON tbl_files.id = tbl_cocktails.preview_id LEFT JOIN tbl_cocktails_to_tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.cocktail_id = tbl_cocktails.id LEFT JOIN tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.ingredient_id = tbl_ingredients.id LEFT JOIN tbl_tries ON tbl_tries.cocktail_id = tbl_cocktails.id WHERE tbl_cocktails.id = $1 GROUP BY tbl_cocktails.id, tbl_complication_levels.name, tbl_complication_levels.time, tbl_fortress_levels.name, tbl_files.filepath"
	} else {
		query = "SELECT tbl_cocktails.id, tbl_cocktails.name, tbl_cocktails.name_en AS nameEn, recipe, tbl_cocktails.description, ROUND(CAST(mark AS decimal)/GREATEST(mark_cnt,1)) AS mark, CONCAT(ROUND(CAST(mark AS decimal)/GREATEST(mark_cnt,1), 1), ' (по ', mark_cnt, ' оценкам)') AS markDescription, CONCAT(tbl_complication_levels.name, ' (', tbl_complication_levels.time, ')') AS complication, tbl_fortress_levels.name AS fortress, is_flacky AS isFlacky, is_fire AS isFire, is_iba AS isIba, CONCAT('" + s.App.Config.MediaUrl + "', coalesce(tbl_files.filepath, '')) AS icon FROM tbl_cocktails INNER JOIN tbl_complication_levels ON tbl_complication_levels.id = tbl_cocktails.complication_id INNER JOIN tbl_fortress_levels ON tbl_fortress_levels.id = tbl_cocktails.fortress_id LEFT JOIN tbl_cocktails_to_tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.cocktail_id = tbl_cocktails.id LEFT JOIN tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.ingredient_id = tbl_ingredients.id LEFT JOIN tbl_files ON tbl_files.id = tbl_cocktails.icon_id WHERE tbl_cocktails.id = $1"

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
		return &dayDrink, &item, status.Error(code, err.Error())
	}

	query = "SELECT ingredient_id AS id, name, CONCAT(volume, ' ', unit) AS volume FROM tbl_cocktails_to_tbl_ingredients INNER JOIN tbl_ingredients ON tbl_ingredients.id = tbl_cocktails_to_tbl_ingredients.ingredient_id WHERE cocktail_id = $1"
	if scenario == "day-drink" {
		err = s.App.Db.SelectContext(ctx, &dayDrink.Ingredients, query, id)
	} else {
		err = s.App.Db.SelectContext(ctx, &item.Ingredients, query, id)
	}

	if err != nil && err != sql.ErrNoRows {
		return &dayDrink, &item, status.Error(codes.Internal, err.Error())
	}

	query = "SELECT instrument_id AS id, name FROM tbl_cocktails_to_tbl_instruments INNER JOIN tbl_instruments ON tbl_instruments.id = tbl_cocktails_to_tbl_instruments.instrument_id WHERE cocktail_id = $1"
	if scenario == "day-drink" {
		err = s.App.Db.SelectContext(ctx, &dayDrink.Instruments, query, id)
	} else {
		err = s.App.Db.SelectContext(ctx, &item.Instruments, query, id)
	}

	if err != nil && err != sql.ErrNoRows {
		return &dayDrink, &item, status.Error(codes.Internal, err.Error())
	}

	return &dayDrink, &item, nil
}

func prepareFilterParams(request *protogen.DrinksRequest) (string, CocktailSearchParams, []error) {
	var queryWhere []string
	var errors []error

	resultQuery := ""
	querySimilar := "(SELECT COUNT(*) FROM (SELECT ci.ingredient_id FROM tbl_cocktails_to_tbl_ingredients AS ci INNER JOIN tbl_ingredients AS ti ON ti.id = ci.ingredient_id WHERE required = true AND ci.cocktail_id = tbl_cocktails.id INTERSECT SELECT ci2.ingredient_id FROM tbl_cocktails_to_tbl_ingredients AS ci2 WHERE ci2.cocktail_id = ANY(:similar)) AS tmp) DESC, "
	queryGroupBy := " GROUP BY tbl_cocktails.id, tbl_cocktails.weight, tbl_complication_levels.name, tbl_complication_levels.time, tbl_fortress_levels.name, tbl_files.filepath"

	searchParams := CocktailSearchParams{}

	searchParams.PerPage = 50
	searchParams.Offset = request.Page * searchParams.PerPage

	if len(request.Fortress) > 0 {
		queryWhere = append(queryWhere, "tbl_cocktails.fortress_id = ANY(:fortress)")
		searchParams.Fortress = &pgtype.Int8Array{}
		err := searchParams.Fortress.Set(request.Fortress)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(request.Complication) > 0 {
		queryWhere = append(queryWhere, "tbl_cocktails.complication_id = ANY(:complication)")
		searchParams.Complication = &pgtype.Int8Array{}
		err := searchParams.Complication.Set(request.Complication)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(request.Volume) > 0 {
		queryWhere = append(queryWhere, "tbl_cocktails.volume_id = ANY(:volume)")
		searchParams.Volume = &pgtype.Int8Array{}
		err := searchParams.Volume.Set(request.Volume)
		if err != nil {
			errors = append(errors, err)
		}
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
		err := searchParams.Includes.Set(request.Includes)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(request.Except) > 0 {
		queryWhere = append(queryWhere, "NOT(:except && ARRAY(SELECT ci.ingredient_id FROM tbl_cocktails_to_tbl_ingredients AS ci WHERE ci.cocktail_id = tbl_cocktails.id))")
		searchParams.Except = &pgtype.Int8Array{}
		err := searchParams.Except.Set(request.Except)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(request.Instruments) > 0 {
		queryWhere = append(queryWhere, ":instruments <@ ARRAY(SELECT ci.instrument_id FROM tbl_cocktails_to_tbl_instruments AS ci WHERE ci.cocktail_id = tbl_cocktails.id)")
		searchParams.Instruments = &pgtype.Int8Array{}
		err := searchParams.Instruments.Set(request.Instruments)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(request.Similar) > 0 {
		queryWhere = append(queryWhere, "ARRAY(SELECT ci.ingredient_id FROM tbl_cocktails_to_tbl_ingredients AS ci INNER JOIN tbl_ingredients AS ti ON ti.id = ci.ingredient_id WHERE required = true AND ci.cocktail_id = tbl_cocktails.id) && ARRAY(SELECT ci2.ingredient_id FROM tbl_cocktails_to_tbl_ingredients AS ci2 WHERE ci2.cocktail_id = ANY(:similar))")
		queryWhere = append(queryWhere, "NOT(tbl_cocktails.id = ANY(:similar))")
		searchParams.Similar = &pgtype.Int8Array{}
		err := searchParams.Similar.Set(request.Similar)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(queryWhere) > 0 {
		resultQuery += " WHERE " + strings.Join(queryWhere, " AND ")
	}

	resultQuery += queryGroupBy

	if len(request.Similar) > 0 {
		resultQuery += " ORDER BY " + querySimilar + "weight DESC, tbl_cocktails.id ASC LIMIT :perpage OFFSET :offset"
	} else {
		resultQuery += " ORDER BY weight DESC, tbl_cocktails.id ASC LIMIT :perpage OFFSET :offset"
	}

	return resultQuery, searchParams, errors
}
