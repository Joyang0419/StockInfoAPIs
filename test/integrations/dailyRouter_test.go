package integrations_test

import (
	"StockInfoAPIs"
	_ "StockInfoAPIs/inits"
	"StockInfoAPIs/internals/routers"
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
	server := StockInfoAPIs.NewServer(gin.DebugMode, []func(router *gin.Engine){dailyRouter.Route})

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "http://0.0.0.0:8080/daily_price?stock_code=1101&start_date=2020-01-01&end_date=2020-01-03", nil)

	server.ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Code, http.StatusOK)

	decoder := json.NewDecoder(recorder.Body)
	var data routers.RouteResponse
	_ = decoder.Decode(&data)
	assert.Equal(t, reflect.TypeOf(data), reflect.TypeOf(routers.RouteResponse{}))
}
