package integrations_test

import (
	_ "StockInfoAPIs/inits"
	"StockInfoAPIs/internals"
	"StockInfoAPIs/internals/routers/dailyRouter"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestDailyRouter(t *testing.T) {
	server := internals.NewServer(gin.DebugMode, []func(router *gin.Engine){dailyRouter.Route})

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "http://0.0.0.0:8080/daily_price?stock_code=1101&start_date=2020-01-01&end_date=2020-01-03", nil)

	server.ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Code, http.StatusOK)

	decoder := json.NewDecoder(recorder.Body)
	var data dailyRouter.DailySuccessResponse
	_ = decoder.Decode(&data)
	assert.Equal(t, reflect.TypeOf(data), reflect.TypeOf(dailyRouter.DailySuccessResponse{}))
}
