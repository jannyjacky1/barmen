package instruments

import (
	"database/sql"
	_ "github.com/go-playground/validator"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jannyjacky1/barmen/api/manager/v1/common"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
)

const tableName = "tbl_instruments"

func Index(c echo.Context) error {
	db := c.Get("db").(*sqlx.DB)
	result, err := db.Query("SELECT " + tableName + ".id, name, description, filepath FROM " + tableName + " LEFT JOIN " + common.FilesTableName + " ON " + tableName + ".img_id=" + common.FilesTableName + ".id")
	if err != nil {
		return err
	}

	resultModels := []common.Instrument{}

	for result.Next() {
		item := common.Instrument{}
		err := result.Scan(&item.Id, &item.Name, &item.Description, &item.Filepath)
		if err != nil {
			return err
		}

		resultModels = append(resultModels, item)
	}

	return c.JSON(http.StatusOK, resultModels)
}

func Create(c echo.Context) error {

	model := new(common.InstrumentForm)

	if err := c.Bind(model); err != nil {
		return err
	}
	if err := c.Validate(model); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.ParseError(err))
	}

	fileId, _ := common.SaveFile(c, "img", "uploads/")

	db := c.Get("db").(*sqlx.DB)
	_, err := db.Query("INSERT INTO "+tableName+"(name, description, img_id) VALUES($1, $2, $3) ", model.Name, model.Description.Value(), fileId.Value())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, common.Message{Message: "New ingredient was added successfully"})
}

func Update(c echo.Context) error {

	id := c.Param("id")
	db := c.Get("db").(*sqlx.DB)
	model := new(common.InstrumentForm)
	err := db.QueryRow("SELECT name, description, img_id AS file_id FROM "+tableName+" WHERE id=$1", id).Scan(&model.Name, &model.Description, &model.FileId)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, common.Message{Message: "Instrument not found"})
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

	return c.JSON(http.StatusOK, common.Message{Message: "Instrument was updated successfully"})
}

func Delete(c echo.Context) error {

	id := c.Param("id")
	db := c.Get("db").(*sqlx.DB)
	var dbId int
	var fileId common.NullInt64
	err := db.QueryRow("SELECT id, img_id AS fileId FROM "+tableName+" WHERE id=$1", id).Scan(&dbId, &fileId)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, common.Message{Message: "Instrument not found"})
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
	return c.JSON(http.StatusOK, common.Message{Message: "Instrument was deleted successfully"})
}
