package health

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestHealthHandler_GetHealthSuccess(t *testing.T) {
	host, _ := os.Hostname()
	var healthHandler = HealthHandlerImpl{startAt: time.Now(), host: host}

	responseData := struct {
		Message string `json:"message"`
		Result  struct {
			Host       string `json:"host"`
			DeployedAt string `json:"deployed_at"`
		} `json:"result"`
	}{}

	e := echo.New()
	r := httptest.NewRequest("GET", "/health", nil)
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	healthHandler.GetHealth(ctx)
	bodyRes, _ := ioutil.ReadAll(w.Result().Body)

	json.Unmarshal(bodyRes, &responseData)

	if assert.NoError(t, healthHandler.GetHealth(ctx)) {
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.Equal(t, responseData.Message, "Success")
		assert.Equal(t, responseData.Result.Host, host)
	} else {
		t.Fatalf("handler error")
	}
}
