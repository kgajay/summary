package tests

import (
	"testing"
	"controllers"
	"github.com/labstack/echo"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"fmt"
)

func TestSum(t *testing.T) {
	total := controllers.Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := controllers.Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}

}

func TestStatus(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/status", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.GetStatus(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.Equal(t, resp["status"], "up")
		assert.IsType(t, resp["uptime"], "string")
	}

}

func TestDeepStatus(t *testing.T) {

	setUpTest()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/deepstatus", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.GetDeepStatus(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp map[string]interface{}
		fmt.Println(rec.Body.String())
		json.Unmarshal(rec.Body.Bytes(), &resp)
		dbResp := resp["db"].(map[string]interface{})
		assert.Equal(t, dbResp["status"], "up")
	}
	destroy()
}
