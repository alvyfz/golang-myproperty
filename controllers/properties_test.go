package controller

// import (
// 	"ca-myproperty/config"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// )

// func TestGetAllProperties(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "/properties", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := GetAllProperties(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }

// func TestCreatProperty(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodPost, "/properties", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := CreateProperty(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }

// func TestUpdateProperty(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodPut, "/properties:id", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := UpdateProperty(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }

// func TestGetPropertyByID(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodPut, "/properties/:id", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := GetPropertyByID(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }
// func TestDeletePropertyByID(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodDelete, "/properties/:id", nil)
// 	rec := httptest.NewRecorder()

// 	config.InitDB()
// 	app := echo.New()
// 	c := app.NewContext(req, rec)

// 	err := DeletePropertyByID(c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	if rec.Result().StatusCode != 200 {
// 		t.Error(err.Error())
// 	}
// }
