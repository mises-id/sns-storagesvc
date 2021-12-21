package env

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var Envs *Env

type Env struct {
	Port            int    `env:"PORT" envDefault:"7070"`
	AppEnv          string `env:"APP_ENV" envDefault:"development"`
	MongoURI        string `env:"MONGO_URI,required"`
	DBUser          string `env:"DB_USER"`
	DBPass          string `env:"DB_PASS"`
	DBName          string `env:"DB_NAME" envDefault:"mises"`
	AWSAccessKeyId  string `env:"AWSAccessKeyId,required"`
	AWSSecretKey    string `env:"AWSSecretKey,required"`
	AWSRegion       string `env:"AWSRegion,required"`
	LocalFilePath   string `env:"LocalFilePath"`
	StorageProvider string `env:"STORAGE_PROVIDER" envDefault:"local"`
	S3Bucket        string `env:"S3Bucket,required"`
	S3Prefix        string `env:"S3Prefix" envDefault:"s3://"`
	SignURL         bool   `env:"SignURL"`
	SignKey         string `env:"SignKey"`
	SignSalt        string `env:"SignSalt"`
	RootPath        string
}

func init() {
	fmt.Println("env initializing...")
	_, b, _, _ := runtime.Caller(0)
	appEnv := os.Getenv("APP_ENV")
	projectRootPath := filepath.Dir(b) + "/../../"
	envPath := projectRootPath + ".env"
	appEnvPath := envPath + "." + appEnv
	localEnvPath := appEnvPath + ".local"
	_ = godotenv.Load(filtePath(localEnvPath, appEnvPath, envPath)...)
	Envs = &Env{}
	err := env.Parse(Envs)
	if err != nil {
		panic(err)
	}
	Envs.RootPath = projectRootPath
	fmt.Println("env loaded...")
}

func filtePath(paths ...string) []string {
	result := make([]string, 0)
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			result = append(result, path)
		}
	}
	return result
}
