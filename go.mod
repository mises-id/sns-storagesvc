module github.com/mises-id/sns-storagesvc

go 1.16

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/aws/aws-sdk-go-v2 v1.11.2
	github.com/aws/aws-sdk-go-v2/credentials v1.6.4
	github.com/aws/aws-sdk-go-v2/service/s3 v1.21.0
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/gavv/httpexpect v2.0.0+incompatible
	github.com/go-kit/kit v0.12.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/h2non/bimg v1.1.5
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/metaverse/truss v0.3.1
	github.com/mises-id/sns v0.0.51
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.5
	github.com/valyala/fasthttp v1.31.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	go.mongodb.org/mongo-driver v1.8.1
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	google.golang.org/genproto v0.0.0-20211129164237-f09f9a12af12 // indirect
	google.golang.org/grpc v1.42.0
)

//replace github.com/metaverse/truss => ../truss
