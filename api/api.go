package api

import (
	"fmt"
	"net/http"

	"github.com/minio/minio-go"

	"github.com/mattes/migrate/database/postgres"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // postgress database driver
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/source/file"
)

func NewAPi(conf ApiConfing) (*service, error) {
	postgress, err := gorm.Open("postgres", conf.postgressConection())
	if err != nil {
		return nil, err
	}

	filestorage, err := minio.New(conf.StorageEndpoint, conf.StorageAccessKeyID, conf.StorageSecretAccessKey, false)
	if err != nil {
		return nil, err
	}
	exists, err := filestorage.BucketExists(conf.StoragePayloadBucket)
	if err != nil {
		return nil, err
	}
	if !exists {
		if err = filestorage.MakeBucket(conf.StoragePayloadBucket, conf.StoragePayloadBucket); err != nil {
			return nil, err
		}
	}

	api := service{
		conf:        conf,
		postgress:   postgress,
		fileStorage: filestorage,
	}

	api.server = api.setup()

	return &api, nil
}

type ApiConfing struct {
	Port                   string
	DatabaseHost           string
	DatabaseUser           string
	DatabasePassword       string
	DatabaseName           string
	StorageEndpoint        string
	StorageAccessKeyID     string
	StorageSecretAccessKey string
	StorageLocation        string
	StoragePayloadBucket   string
}

func (a *ApiConfing) postgressConection() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", a.DatabaseHost, a.DatabaseUser, a.DatabasePassword, a.DatabaseName)
}

type service struct {
	conf        ApiConfing
	postgress   *gorm.DB
	fileStorage *minio.Client
	server      http.Handler
}

func (a *service) Run() {
	http.ListenAndServe(a.conf.Port, a.server)
}

func (a *service) setup() http.Handler {
	r := chi.NewRouter()
	r.Get("/book", a.getBookListHandler)
	r.Get("/book/{book_id}", a.getBookHandler)

	return r
}

func (a *service) DatabaseMigrate(migrationsDir string) error {
	db, err := gorm.Open("postgres", a.conf.postgressConection())
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db.DB(), &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationsDir, a.conf.DatabaseName, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	sourceErr, databaseErr := m.Close()
	if sourceErr != nil {
		return sourceErr
	}
	if databaseErr != nil {
		return databaseErr
	}

	return nil
}
