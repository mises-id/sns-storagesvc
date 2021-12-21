package imgae

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mises-id/sns-storagesvc/app/services/image/imagemeta"
	"github.com/mises-id/sns-storagesvc/app/services/image/imagetype"
	"github.com/mises-id/sns-storagesvc/config/env"
	"github.com/mises-id/sns-storagesvc/tests/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type (
	ImageSuite struct {
		web.WebBaseTestSuite
	}
)

func (suite *ImageSuite) SetupSuite() {
	suite.WebBaseTestSuite.SetupSuite()
}

func (suite *ImageSuite) TearDownSuite() {
	suite.WebBaseTestSuite.TearDownSuite()
}
func (suite *ImageSuite) SetupTest() {

}

func (suite *ImageSuite) send(path string) *httptest.ResponseRecorder {

	req := httptest.NewRequest(http.MethodGet, path, nil)
	rw := httptest.NewRecorder()
	suite.Handler.ServeHTTP(rw, req)
	return rw
}

func TestImagesServer(t *testing.T) {
	suite.Run(t, &ImageSuite{})
}

func (suite *ImageSuite) TestImageSource() {
	suite.T().Run("find image source", func(t *testing.T) {
		resp := suite.Expect.GET("/test.jpeg").Expect()
		assert.Equal(suite.T(), http.StatusOK, resp.Raw().StatusCode)

	})
	suite.T().Run("not found image source", func(t *testing.T) {
		resp := suite.Expect.GET("/nothing.jpeg").Expect()
		assert.Equal(suite.T(), http.StatusNotFound, resp.Raw().StatusCode)
	})
}

func (suite *ImageSuite) TestImageOptions() {

	//resize
	suite.T().Run("image resize", func(t *testing.T) {
		options_str := "resize:fit:200:200"
		rw := suite.send("/test.jpeg?" + options_str)
		res := rw.Result()
		assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
		assert.Equal(suite.T(), "image/jpeg", res.Header.Get("Content-Type"))
		meta, err := imagemeta.DecodeMeta(res.Body)
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), 200, meta.Width())
		assert.Equal(suite.T(), 200, meta.Height())
	})

	//crop
	suite.T().Run("image crop", func(t *testing.T) {
		options_str := "crop:200:110"
		rw := suite.send("/test.jpeg?" + options_str)
		res := rw.Result()
		assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
		meta, err := imagemeta.DecodeMeta(res.Body)
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), 200, meta.Width())
		assert.Equal(suite.T(), 110, meta.Height())
	})

	//format
	suite.T().Run("image format", func(t *testing.T) {
		options_str := "format:png"
		rw := suite.send("/test.jpeg?" + options_str)
		res := rw.Result()
		assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
		assert.Equal(suite.T(), "image/png", res.Header.Get("Content-Type"))
		meta, err := imagemeta.DecodeMeta(res.Body)
		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), imagetype.PNG, meta.Format())
	})
	//watermark text
	suite.T().Run("image text watermark", func(t *testing.T) {
		options_str := "watermark:" + base64.RawURLEncoding.EncodeToString([]byte("mises"))
		rw := suite.send("/test.jpeg?" + options_str)
		res := rw.Result()
		assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
		//hash
		b, _ := os.ReadFile("./test_watermark.jpeg")
		expected := Sha256(b)
		fmt.Println("expected ", expected)
		a, err := ioutil.ReadAll(res.Body)
		assert.Nil(suite.T(), err)
		writeLocalFile("./test_watermark_res.jpeg", a)
		actual := Sha256(a)
		fmt.Println("actual ", actual)
		assert.Equal(suite.T(), expected, actual)
	})

}

func Sha256(b []byte) string {
	m := sha256.New()
	m.Write(b)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func writeLocalFile(localfile string, data []byte) (err error) {
	dst, err := os.Create(localfile)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = dst.Write(data)
	if err != nil {
		fmt.Println("write file err:", err.Error())
		return err
	}
	return nil
}

func (suite *ImageSuite) TestImageUri() {
	env.Envs = &env.Env{
		LocalFilePath: "./",
		SignURL:       true,
		SignKey:       "736563726574",
		SignSalt:      "68656C6C6F",
	}
	suite.T().Run("image uri signature invalid", func(t *testing.T) {
		resp := suite.Expect.GET("/test.jpeg?sssss").Expect()
		assert.Equal(suite.T(), http.StatusForbidden, resp.Raw().StatusCode)
	})
	suite.T().Run("image uri signature valid", func(t *testing.T) {
		resp := suite.Expect.GET("/test.jpeg?uCeERkq-3089BNBnBHOXz6NS0ukUQ4d0FiQjV5xSPzg").Expect()
		assert.Equal(suite.T(), http.StatusOK, resp.Raw().StatusCode)
	})
}
