package data

import (
	"fmt"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

// ProviderSet is data providers.
var (
	ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewMongoDB, NewAuthRepo)
)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *mongo.Database
	log *log.Helper
}

// NewMongoDB new a mongodb client
func NewMongoDB(c *conf.Data, logger log.Logger) *mongo.Database {
	ctx, _ := context.WithCancel(context.Background())
	logHelper := log.NewHelper(log.With(logger, "module", "data_service"))
	uri := fmt.Sprintf("mongodb://%s:%d", c.Mongodb.Host, c.Mongodb.Port)
	logHelper.Infof("mongodb uri: %s", uri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(options.Credential{
		AuthSource: c.Mongodb.Database,
		Username:   c.Mongodb.Username,
		Password:   c.Mongodb.Password,
	}))
	if err != nil {
		logHelper.Infof("Failed to connect mongodb: %v", err)
		panic(err)
	}
	return client.Database(c.Mongodb.Database)
}

// NewData .
func NewData(mongodb *mongo.Database, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(log.With(logger, "module", "data_service"))

	cleanup := func() {
		logHelper.Info("closing the data resources")
	}

	// init mongodb
	return &Data{
		db:  mongodb,
		log: logHelper,
	}, cleanup, nil
}
