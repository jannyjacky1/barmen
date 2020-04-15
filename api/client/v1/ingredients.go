package v1

import (
	"context"
	"database/sql"
	"github.com/jannyjacky1/barmen/api/client/v1/protogen"
	"github.com/jannyjacky1/barmen/tools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type IngredientsServer struct {
	App tools.App
}

func (s *IngredientsServer) GetIngredientById(ctx context.Context, request *protogen.IngredientRequest) (*protogen.IngredientResponse, error) {

	var ingredient protogen.IngredientResponse

	err := s.App.Db.GetContext(ctx, &ingredient, "SELECT tbl_ingredients.id, name, coalesce(name_en, '') AS nameEn, COUNT(DISTINCT tbl_cocktails_to_tbl_ingredients.cocktail_id) AS info, coalesce(description, '') AS description, coalesce(tbl_files.filepath, '') AS icon FROM tbl_ingredients LEFT JOIN tbl_files ON tbl_files.id = tbl_ingredients.img_id LEFT JOIN tbl_cocktails_to_tbl_ingredients ON tbl_cocktails_to_tbl_ingredients.ingredient_id = tbl_ingredients.id WHERE tbl_ingredients.id = $1 GROUP BY tbl_ingredients.id, tbl_files.filepath", request.Id)

	if err != nil {
		code := codes.Internal
		if err == sql.ErrNoRows {
			code = codes.NotFound
		} else {
			s.App.Log.Error(err.Error())
		}
		return &ingredient, status.Error(code, err.Error())
	}

	cnt, err := strconv.Atoi(ingredient.Info)
	if err != nil {
		s.App.Log.Error(err.Error())
	}
	ingredient.Info = "Используется в " + ingredient.Info + " " + tools.GetWord(cnt, "коктейле", "коктейлях", "коктейлях")

	return &ingredient, nil
}
