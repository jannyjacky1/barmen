package ingredients

import (
	"context"
	"database/sql"
	_ "github.com/go-playground/validator"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jannyjacky1/barmen/api/manager/v1/common"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

const tableName = "tbl_ingredients"

func Index(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	ctx := c.Get("context").(context.Context)

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 0
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize == 0 {
		pageSize = 20
	}

	resultModels := []common.IngredientShort{}
	err = db.SelectContext(ctx, &resultModels, "SELECT id, name, name_en FROM "+tableName+" ORDER BY id ASC LIMIT $1 OFFSET $2", pageSize, pageSize*page)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resultModels)
}

func Create(c echo.Context) error {

	var model common.Ingredient

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, common.Error{Message: err.Error(), Code: http.StatusBadRequest})
	}
	if err := c.Validate(&model); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.ParseError(err))
	}

	fileId, _ := common.SaveFile(c, "img", "uploads/")

	db := c.Get("db").(*sqlx.DB)
	var id int
	err := db.QueryRow("INSERT INTO "+tableName+"(name, description, img_id) VALUES($1, $2, $3) RETURNING id", model.Name, model.Description.Value(), fileId.Value()).Scan(&id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Error{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, common.Message{Message: "New ingredient was added successfully"})
}

func Update(c echo.Context) error {

	id := c.Param("id")
	db := c.Get("db").(*sqlx.DB)
	model := new(common.IngredientForm)
	err := db.QueryRow("SELECT name, description, img_id AS file_id FROM "+tableName+" WHERE id=$1", id).Scan(&model.Name, &model.Description, &model.FileId)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, common.Message{Message: "Ingredient not found"})
	}

	if err != nil {
		return err
	}

	if err := c.Bind(model); err != nil {
		return err
	}

	if err := c.Validate(model); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.ParseError(err))
	}

	fileId, _ := common.SaveFile(c, "img", "uploads/")
	if fileId.Int64 > 0 {
		if model.FileId.Value() != nil {
			common.DeleteFile(c, model.FileId.Int64)
		}
		model.FileId = fileId
	}

	_, err = db.Query("UPDATE "+tableName+" SET name=$2, description=$3, img_id=$4 WHERE id=$1", id, model.Name, model.Description.Value(), model.FileId.Value())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, common.Message{Message: "Ingredient was updated successfully"})
}

func Delete(c echo.Context) error {

	id := c.Param("id")
	db := c.Get("db").(*sqlx.DB)
	var dbId int
	var fileId common.NullInt64
	err := db.QueryRow("SELECT id, img_id AS fileId FROM "+tableName+" WHERE id=$1", id).Scan(&dbId, &fileId)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, common.Message{Message: "Ingredient not found"})
	}

	if err != nil {
		return err
	}

	_, err = db.Query("DELETE FROM "+tableName+" WHERE id=$1", id)

	if err != nil {
		return err
	}

	if fileId.Int64 > 0 {
		common.DeleteFile(c, fileId.Int64)
	}
	return c.JSON(http.StatusOK, common.Message{Message: "Ingredient was deleted successfully"})
}
