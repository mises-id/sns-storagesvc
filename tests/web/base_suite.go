package web

import (
	"fmt"
	"net/http"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo"
	"github.com/mises-id/sns-storagesvc/config/router"
	"github.com/mises-id/sns-storagesvc/lib/middleware"
	"github.com/mises-id/sns-storagesvc/tests"
)

func init() {

	fmt.Println("test web init")
}

type WebBaseTestSuite struct {
	tests.BaseTestSuite
	Handler http.Handler
	Expect  *httpexpect.Expect
}

func (suite *WebBaseTestSuite) SetupSuite() {
	suite.BaseTestSuite.SetupSuite()
	suite.SetWebHandler()
	suite.InitExpect()
}

func (suite *WebBaseTestSuite) SetWebHandler() {
	e := echo.New()
	e.Use(middleware.ErrorResponseMiddleware)
	//router
	router.SetRouter(e)
	suite.Handler = e
}

func (suite *WebBaseTestSuite) TearDownSuite() {
	suite.BaseTestSuite.TearDownSuite()
}
func (suite *WebBaseTestSuite) InitExpect() {
	suite.Expect = httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(suite.Handler),
		},
		Reporter: httpexpect.NewRequireReporter(suite.T()),
	})
}
