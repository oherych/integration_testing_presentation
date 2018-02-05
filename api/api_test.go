package api

import (
	"github.com/minio/minio-go"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
)

// createTestEnv create new test enviroment
func createTestEnv(t *testing.T) *httptest.Server {
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

	// return test server
	return httptest.NewServer(api.server)
}

func createTestEnvExpect(t *testing.T) *httpexpect.Expect {
	server := createTestEnv(t)
	return httpexpect.New(t, server.URL)
}
