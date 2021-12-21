package tests

import (
	"fmt"

	"github.com/mises-id/sns-storagesvc/config/env"
	"github.com/stretchr/testify/suite"
)

func init() {
	fmt.Println("test init")

	env.Envs = &env.Env{
		LocalFilePath: "./",
		SignURL:       false,
		SignKey:       "736563726574",
		SignSalt:      "68656C6C6F",
	}
	fmt.Println("test env init")
}

type BaseTestSuite struct {
	suite.Suite
}

func (suite *BaseTestSuite) SetupSuite() {

}

func (suite *BaseTestSuite) TearDownSuite() {

}
