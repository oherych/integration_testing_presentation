package main

import "github.com/oherych/integration_testing_presentation/api"

func main() {
	api, _ := api.NewAPi(api.ApiConfing{
		DatabaseName:           "dev",
		DatabaseUser:           "user",
		DatabaseHost:           "localhost",
		DatabasePassword:       "password",
		StorageEndpoint:        "localhost:9000",
		StorageLocation:        "eu-central-1",
		StorageAccessKeyID:     "BH0K3LEZSZX2KFM53LLS",
		StorageSecretAccessKey: "Pzfn+HrTbw+oPO8Tz5NFnj/1RbWSjH1qQ+cqCJE6",
		StoragePayloadBucket:   "test",
		Port:                   ":3000",
	})

	api.DatabaseMigrate("db")

	api.Run()
}
