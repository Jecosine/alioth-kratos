package data

import (
	"database/sql"
	"fmt"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

// ProviderSet is data providers.
var (
	ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewMongoDB, NewPostgres, NewAuthRepo)
)

// Data .
type Data struct {
	// TODO wrapped database client
	db       *mongo.Database
	postgres *sql.DB
	log      *log.Helper
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

// NewPostgres new a postgresql client
func NewPostgres(c *conf.Data, logger log.Logger) *sql.DB {
	logHelper := log.NewHelper(log.With(logger, "module", "data_service"))
	uri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", c.Postgres.Host, c.Postgres.Port, c.Postgres.Username, c.Postgres.Password, c.Postgres.Database)
	db, err := sql.Open("pgx", uri)
	if err != nil {
		panic(err)
	}
	logHelper.Info("Successfully connect to postgresql")
	return db
	//drv := entsql.OpenDB(dialect.Postgres, db)
	//client := ent.NewClient(ent.Driver(drv))
	//err = client.Schema.Create(context.Background())
	//if err != nil {
	//	logHelper.Warn("cannot connect to postgresql db")
	//	panic(err)
	//}
	//return client
}

// NewData .
func NewData(mongodb *mongo.Database, postgres *sql.DB, logger log.Logger) (*Data, func(), error) {
	logHelper := log.NewHelper(log.With(logger, "module", "data_service"))

	cleanup := func() {
		logHelper.Info("closing the data resources")
	}

	// init mongodb
	return &Data{
		db:       mongodb,
		postgres: postgres,
		log:      logHelper,
	}, cleanup, nil
}
