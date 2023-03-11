package gh

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type jsonBody struct {
	Name string `json:"name"`
}

func bindJSONBody(jq *jsonBody) *jsonBody {
	return jq
}

func gotErr() error {
	return errors.New("hello")
}

type uriErr struct {
	Set string `uri:"set"`
}

func uriControlErr(u *uriErr) (*jsonBody, error) {
	if u.Set == "1" {
		return nil, errors.New("setted")
	}
	return &jsonBody{Name: "haha"}, nil
}

type service struct {
}

func (s *service) json(jq *jsonBody) *jsonBody {
	return jq
}

func setupRouter() *gin.Engine {
	s := &service{}
	r := gin.Default()
	r.POST("/json", MustBind(bindJSONBody).HandlerFunc())
	r.POST("/service/json", MustBind(s.json).HandlerFunc())
	r.GET("/err", MustBind(gotErr).HandlerFunc())
	r.GET("/err/:set", MustBind(uriControlErr).HandlerFunc())
	return r
}

func TestMustBind(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/json", bytes.NewBufferString(`{"name": "h"}`))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"name":"h"}`, w.Body.String())
}

func TestGotError(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/err", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
	assert.Equal(t, `hello`, w.Body.String())
}

func TestURI(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/err/0", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"name":"haha"}`, w.Body.String())

	w = httptest.NewRecorder()

	req, err = http.NewRequest("GET", "/err/1", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "setted", w.Body.String())
}

func TestServiceStruct(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/service/json", bytes.NewBufferString(`{"name": "h"}`))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"name":"h"}`, w.Body.String())
}
