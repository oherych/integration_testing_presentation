package api

import (
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/minio/minio-go"

	"github.com/gavv/httpexpect"
)

const (
	testBookID1 = "913d5f4e-5759-455d-83fe-72939b3ddf3a"
)

var (
	defaultTestEnv = map[string]string{
		"DATABASE_HOST":          "localhost",
		"DATABASE_USER":          "user",
		"DATABASE_NAME":          "dev",
		"DATABASE_PASSWORD":      "password",
		"STORAGE_ENDPOINT":       "localhost:9000",
		"STORAGE_ACCESS_KEY":     "BH0K3LEZSZX2KFM53LLS",
		"STORAGE_SECRET_KEY":     "Pzfn+HrTbw+oPO8Tz5NFnj/1RbWSjH1qQ+cqCJE6",
		"STORAGE_LOCATION":       "eu-central-1",
		"STORAGE_PAYLOAD_BUCKET": "test",
	}
)

// createTestEnv create new test enviroment
func createTestEnv(t testing.TB) (*httptest.Server, func()) {
	for name, val := range defaultTestEnv {
		if _, ok := os.LookupEnv(name); !ok {
			os.Setenv(name, val)
		}
	}

	cfg, err := ParseConfig()
	if err != nil {
		t.Fatal(err)
	}

	api, err := NewAPi(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// remove all old tables
	err = api.postgress.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;").Error
	if err != nil {
		t.Fatal(err)
	}

	// create database structure
	err = api.DatabaseMigrate("../db")
	if err != nil {
		t.Fatal(err)
	}

	// load fixtures
	b, err := ioutil.ReadFile("../_testdata/fixtures.sql")
	if err != nil {
		t.Fatal(err)
	}
	err = api.postgress.Exec(string(b)).Error
	if err != nil {
		t.Fatal(err)
	}

	// load files
	files := []string{"logo.png"}
	for _, filePath := range files {
		file, err := os.Open("../_testdata/" + filePath)
		if err != nil {
			t.Fatal(err)
		}

		stat, err := file.Stat()
		if err != nil {
			t.Fatal(err)
		}

		_, err = api.fileStorage.PutObject(api.conf.StoragePayloadBucket, filePath, file, stat.Size(), minio.PutObjectOptions{})
		if err != nil {
			t.Fatal(err)
		}
	}

	server := httptest.NewServer(api.server)

	done := func() {
		server.Close()
		api.Close()
	}

	// return test server
	return server, done
}

func createTestEnvExpect(t testing.TB) (*httpexpect.Expect, func()) {
	server, done := createTestEnv(t)

	// return httpexpect.New(t, server.URL), done
	// or
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
	}), done
}
