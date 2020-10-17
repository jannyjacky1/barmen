package common

import (
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/random"
	"io"
	"os"
)

const FilesTableName = "tbl_files"

func SaveFile(c echo.Context, field string, path string) (NullInt64, error) {

	file, err := c.FormFile(field)
	if err != nil {
		return NullInt64{Int64: 0, Valid: false}, err
	}
	src, err := file.Open()
	if err != nil {
		return NullInt64{Int64: 0, Valid: false}, err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(path + random.String(8) + file.Filename)
	if err != nil {
		return NullInt64{Int64: 0, Valid: false}, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return NullInt64{Int64: 0, Valid: false}, err
	}

	db := c.Get("db").(*sql.DB)
	var fileId int

	_, err = db.Query("INSERT INTO "+FilesTableName+" (filepath) VALUES($1)", dst.Name())
	if err != nil {
		return NullInt64{Int64: 0, Valid: false}, err
	}

	err = db.QueryRow("SELECT id AS fileId FROM "+FilesTableName+" WHERE filepath=$1", dst.Name()).Scan(&fileId)

	if err != nil {
		return NullInt64{Int64: 0, Valid: false}, err
	}

	return NullInt64{Int64: int64(fileId), Valid: true}, nil
}

func DeleteFile(c echo.Context, id int64) error {
	db := c.Get("db").(*sql.DB)
	var filepath string
	err := db.QueryRow("SELECT filepath FROM "+FilesTableName+" WHERE id=$1", id).Scan(&filepath)

	if err != sql.ErrNoRows {
		db.Query("DELETE FROM "+FilesTableName+" WHERE id=$1", id)
		os.Remove(filepath)
	}

	return err
}

func ParseError(err error) []ValidationError {
	var errors []ValidationError
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, ValidationError{
			Field:   e.Field(),
			Message: "Field should be " + e.ActualTag(),
		})
	}
	return errors
}
