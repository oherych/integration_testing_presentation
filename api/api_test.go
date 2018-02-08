package api

import (
	"io/ioutil"
	"net/http/httptest"
	"os"

	"github.com/minio/minio-go"

	"github.com/gavv/httpexpect"
)

const (
	testBookID1 = "913d5f4e-5759-455d-83fe-72939b3ddf3a"
)

// createTestEnv create new test enviroment
func createTestEnv(t httpexpect.LoggerReporter) (*httptest.Server, func()) {
	api, err := NewAPi(ApiConfing{
		DatabaseName:           "dev",
		DatabaseUser:           "user",
		DatabaseHost:           "localhost",
		DatabasePassword:       "password",
		StorageEndpoint:        "localhost:9000",
		StorageLocation:        "eu-central-1",
		StorageAccessKeyID:     "BH0K3LEZSZX2KFM53LLS",
		StorageSecretAccessKey: "Pzfn+HrTbw+oPO8Tz5NFnj/1RbWSjH1qQ+cqCJE6",
		StoragePayloadBucket:   "test",
	})
	if err != nil {
		t.Errorf("%s", err)
	}

	// remove all old tables
	err = api.postgress.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;").Error
	if err != nil {
		t.Errorf("%s", err)
	}

	// create database structure
	err = api.DatabaseMigrate("../db")
	if err != nil {
		t.Errorf("%s", err)
	}

	// load fixtures
	b, err := ioutil.ReadFile("../_testdata/fixtures.sql")
	if err != nil {
		t.Errorf("%s", err)
	}
	err = api.postgress.Exec(string(b)).Error
	if err != nil {
		t.Errorf("%s", err)
	}

	// load files
	files := []string{"logo.png"}
	for _, filePath := range files {
		file, err := os.Open("../_testdata/" + filePath)
		if err != nil {
			t.Errorf("%s", err)
		}

		stat, err := file.Stat()
		if err != nil {
			t.Errorf("%s", err)
		}

		_, err = api.fileStorage.PutObject(api.conf.StoragePayloadBucket, filePath, file, stat.Size(), minio.PutObjectOptions{})
		if err != nil {
			t.Errorf("%s", err)
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

func createTestEnvExpect(t httpexpect.LoggerReporter) (*httpexpect.Expect, func()) {
	server, done := createTestEnv(t)

	// return httpexpect.New(t, server.URL), done
	// or
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
	}), done
}
