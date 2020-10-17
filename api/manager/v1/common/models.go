package common

import (
	"time"
)

type Message struct {
	Message string `json:"message"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type AuthForm struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type AuthData struct {
	Token       string    `json:"token"`
	TokenExpire time.Time `json:"token_expire"`
}

type Admin struct {
	Email        string
	Username     string
	PasswordHash string
	Token        string
	TokenExpire  time.Time
}

type ApiAdmin struct {
	Username string
	Token    string
}

type ComplicationLevel struct {
	Id   int    `json:"id" form:"-" validate:"-"`
	Name string `json:"name" form:"name" validate:"required"`
	Time string `json:"time" form:"time" validate:"required"`
}

type FortressLevel struct {
	Id           int    `json:"id" form:"-" validate:"-"`
	Name         string `json:"name" form:"name" validate:"required"`
	FortressFrom int    `json:"fortress_from" db:"fortress_from" form:"fortress_from" validate:"required,numeric,gt=0,lt=100,ltecsfield=FortressTo"`
	FortressTo   int    `json:"fortress_to" db:"fortress_to" form:"fortress_to" validate:"required,numeric,gt=0,lt=100,gtecsfield=FortressFrom"`
}

type Volume struct {
	Id         int    `json:"id" form:"-" validate:"-"`
	Name       string `json:"name" form:"name" validate:"required"`
	VolumeFrom int    `json:"volume_from" db:"volume_from" form:"volume_from" validate:"required,numeric,gt=0,ltecsfield=VolumeTo"`
	VolumeTo   int    `json:"volume_to" db:"volume_to" form:"volume_to" validate:"required,numeric,gt=0,gtecsfield=VolumeFrom"`
}

type IngredientShort struct {
	Id     int        `json:"id"`
	Name   string     `json:"name"`
	NameEn NullString `json:"name_en" db:"name_en"`
}

type Ingredient struct {
	Id          int        `json:"id" form:"-" validate:"-"`
	Name        string     `json:"name" form:"name" validate:"required"`
	NameEn      NullString `json:"name_en"`
	Fortress    NullInt64  `json:"fortress"`
	Description NullString `json:"description"`
	Filepath    NullString `json:"filepath"`
	FileId      NullInt64  `json:"file_id"`
	Required    bool       `json:"required"`
}

type IngredientForm struct {
	Name        string     `json:"name" form:"name" validate:"required"`
	Description NullString `json:"description" form:"description"`
	FileId      NullInt64  `json:"file_id" form:"file_id"`
}

type Instrument struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description NullString `json:"description"`
	Filepath    NullString `json:"filepath"`
}

type InstrumentForm struct {
	Name        string     `json:"name" form:"name" validate:"required"`
	Description NullString `json:"description" form:"description"`
	FileId      NullInt64  `json:"file_id" form:"file_id"`
}

type ClientCocktail struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	ImagePath     string  `json:"image_path"`
	Ingredients   string  `json:"ingredients"`
	PreparingTime string  `json:"preparing_time"`
	IsShot        bool    `json:"is_shot"`
	Mark          float32 `json:"mark"`
	IsTop         bool    `json:"is_top"`
}

type ClientCocktailDetail struct {
	Id                int               `json:"id"`
	Name              string            `json:"name"`
	ImagesPath        []string          `json:"images_path"`
	TriesCount        int               `json:"tries_count"`
	Mark              float32           `json:"mark"`
	ComplicationLevel ComplicationLevel `json:"complication_level"`
	FortressLevel     FortressLevel     `json:"fortress_level"`
	Ingredients       []Ingredient      `json:"ingredients"`
	Recipe            string            `json:"recipe"`
	Description       string            `json:"history"`
	SimilarCocktails  []ClientCocktail  `json:"similar_cocktails"`
}
