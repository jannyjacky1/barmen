package v1

import (
	"context"
	"database/sql/driver"
	"github.com/go-playground/validator"
	"github.com/jannyjacky1/barmen/api/manager/v1/authorization"
	"github.com/jannyjacky1/barmen/api/manager/v1/common"
	complication_levels "github.com/jannyjacky1/barmen/api/manager/v1/complication-levels"
	fortress_levels "github.com/jannyjacky1/barmen/api/manager/v1/fortress-levels"
	"github.com/jannyjacky1/barmen/api/manager/v1/ingredients"
	"github.com/jannyjacky1/barmen/api/manager/v1/instruments"
	"github.com/jannyjacky1/barmen/api/manager/v1/volumes"
	"github.com/jannyjacky1/barmen/tools"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"reflect"
	"strings"
)

type CustomValidator struct {
	validator *validator.Validate
}

func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
	}
	return nil
}

func (cv *CustomValidator) Validate(i interface{}) error {
	cv.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	cv.validator.RegisterCustomTypeFunc(ValidateValuer, common.NullInt64{}, common.NullString{})
	return cv.validator.Struct(i)
}

func App(appConfig tools.App) *echo.Echo {
	app := echo.New()
	app.Validator = &CustomValidator{validator: validator.New()}

	ctx := context.Background()

	app.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", appConfig.Db)
			c.Set("log", appConfig.Log)
			c.Set("config", appConfig.Config)
			c.Set("context", ctx)
			return handlerFunc(c)
		}
	})
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	//app.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	//	KeyLookup:  "header:" + echo.HeaderAuthorization,
	//	AuthScheme: "Bearer",
	//	Validator:  authorization.TokenValidator,
	//	Skipper: func(c echo.Context) bool {
	//		return c.Request().RequestURI == "/" || c.Request().RequestURI == "/auth"
	//	},
	//}))
	app.HTTPErrorHandler = func(e error, c echo.Context) {
		he, _ := e.(*echo.HTTPError)
		app.Logger.Error(c.JSON(he.Code, he))
	}
	app.Debug = appConfig.Config.IsDev == "1"

	router(app, "")

	return app
}

func router(app *echo.Echo, basepath string) {
	app.GET(basepath+"/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, common.Message{Message: "Manager API v1!"})
	})

	app.POST(basepath+"/auth", authorization.Auth)

	app.GET(basepath+"/complication-levels", complication_levels.Index)
	app.POST(basepath+"/complication-levels", complication_levels.Create)
	app.PUT(basepath+"/complication-levels/:id", complication_levels.Update)
	app.DELETE(basepath+"/complication-levels/:id", complication_levels.Delete)

	app.GET(basepath+"/fortress-levels", fortress_levels.Index)
	app.POST(basepath+"/fortress-levels", fortress_levels.Create)
	app.PUT(basepath+"/fortress-levels/:id", fortress_levels.Update)
	app.DELETE(basepath+"/fortress-levels/:id", fortress_levels.Delete)

	app.GET(basepath+"/volumes", volumes.Index)
	app.POST(basepath+"/volumes", volumes.Create)
	app.PUT(basepath+"/volumes/:id", volumes.Update)
	app.DELETE(basepath+"/volumes/:id", volumes.Delete)

	app.GET(basepath+"/ingredients", ingredients.Index)
	app.POST(basepath+"/ingredients", ingredients.Create)
	app.PUT(basepath+"/ingredients/:id", ingredients.Update)
	app.DELETE(basepath+"/ingredients/:id", ingredients.Delete)

	app.GET(basepath+"/instruments", instruments.Index)
	app.POST(basepath+"/instruments", instruments.Create)
	app.PUT(basepath+"/instruments/:id", instruments.Update)
	app.DELETE(basepath+"/instruments/:id", instruments.Delete)
}
