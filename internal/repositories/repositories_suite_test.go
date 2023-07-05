package repositories_test

import (
	"log"
	"testing"
	"time"

	"github.com/diazharizky/go-mongodb-with-tests/pkg/mongodb"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbc *mongo.Client
	err error
	db  *mongo.Database
	now time.Time
)

func TestRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repositories Suite")
}

var _ = BeforeSuite(func() {
	dbc, err = mongodb.GetClient()
	if err != nil {
		log.Fatalf("Error unable to get DB client: %v", err)
	}

	db = dbc.Database("go_mongodb_test")
})

func getNow() time.Time {
	return time.Now().UTC().Truncate(time.Millisecond)
}
