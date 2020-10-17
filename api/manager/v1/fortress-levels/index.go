package fortress_levels

import (
	"context"
	"database/sql"
	_ "github.com/go-playground/validator"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jannyjacky1/barmen/api/manager/v1/common"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
)

const tableName = "tbl_fortress_levels"

func Index(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	ctx := c.Get("context").(context.Context)
	resultModels := []common.FortressLevel{}
	err := db.SelectContext(ctx, &resultModels, "SELECT * FROM "+tableName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resultModels)
}

func Create(c echo.Context) error {

	var model common.FortressLevel

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, common.Error{Message: err.Error(), Code: http.StatusBadRequest})
	}
	if err := c.Validate(&model); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.ParseError(err))
	}

	db := c.Get("db").(*sqlx.DB)
	ctx := c.Get("context").(context.Context)

	err := db.QueryRowxContext(ctx, "INSERT INTO "+tableName+"(name, fortress_from, fortress_to) VALUES($1, $2, $3) RETURNING *", model.Name, model.FortressFrom, model.FortressTo).StructScan(&model)

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Error{Message: err.Error(), Code: http.StatusBadRequest})
	}

	return c.JSON(http.StatusCreated, model)
}

func Update(c echo.Context) error {

	id := c.Param("id")
	db := c.Get("db").(*sqlx.DB)
	ctx := c.Get("context").(context.Context)

	var model common.FortressLevel
	err := db.QueryRowxContext(ctx, "SELECT * FROM "+tableName+" WHERE id=$1", id).StructScan(&model)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, common.Error{Message: "FortressLevel not found", Code: http.StatusNotFound})
		}
		return c.JSON(http.StatusBadRequest, common.Error{Message: err.Error(), Code: http.StatusBadRequest})
	}

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, common.Error{Message: err.Error(), Code: http.StatusBadRequest})
	}

	if err := c.Validate(&model); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.ParseError(err))
	}

	_, err = db.NamedExecContext(ctx, "UPDATE "+tableName+" SET name=:name, fortress_from=:fortress_from, fortress_to=:fortress_to WHERE id=:id RETURNING *", model)

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Error{Message: err.Error(), Code: http.StatusBadRequest})
	}

	return c.JSON(http.StatusOK, model)
}

func Delete(c echo.Context) error {

	id := c.Param("id")
	db := c.Get("db").(*sqlx.DB)
	ctx := c.Get("context").(context.Context)

	var dbId int
	err := db.QueryRowxContext(ctx, "DELETE FROM "+tableName+" WHERE id=$1 RETURNING id", id).Scan(&dbId)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, common.Error{Message: "FortressLevel not found", Code: http.StatusNotFound})
		}
		return c.JSON(http.StatusBadRequest, common.Error{Message: err.Error(), Code: http.StatusBadRequest})
	}

	return c.NoContent(http.StatusNoContent)
}
