package data

import (
	"context"
	"github.com/Jecosine/alioth-kratos/app/data_service/internal/conf"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

func TestNewData(t *testing.T) {
	// read conf
	c := config.New(
		config.WithSource(
			file.NewSource("../../config/local.yaml"),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	t.Logf("%v", bc)

	db := NewMongoDB(bc.Data, log.With(log.NewStdLogger(os.Stdout)))
	coll := db.Collection("alioth_db")
	//result, err := coll.InsertOne(
	//	context.TODO(),
	//	bson.D{
	//		{"item", "canvas"},
	//		{"qty", 100},
	//		{"tags", bson.A{"cotton"}},
	//		{"size", bson.D{
	//			{"h", 28},
	//			{"w", 35.5},
	//			{"uom", "cm"},
	//		}},
	//	})
	result, err := coll.Find(
		context.TODO(),
		bson.D{{"item", "canvas"}},
	)
	if err != nil {
		t.Fatalf("Failed to fetch, %v", err)
	} else {
		t.Logf("%v", result.ID())
	}
}
