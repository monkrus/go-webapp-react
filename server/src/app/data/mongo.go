package mongo

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client, err = mongo.Connect(context.Background(), "Insert your MongoDB URI here!", nil)
