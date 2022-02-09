package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type configDB struct {
	host     string
	port     string
	user     string
	password string
}

func NewConfigDB(
	host string,
	port string,
	user string,
	password string,
) *configDB {
	return &configDB{
		host,
		port,
		user,
		password,
	}
}

func (c *configDB) Init(ctx context.Context) *mongo.Client {
	switch {
	case c.host == "":
		log.Fatal("not configure the environment variables for host")
	case c.port == "":
		log.Fatal("not configure the environment variables for port")
	case c.user == "":
		log.Fatal("not configure the environment variables for user")
	case c.password == "":
		log.Fatal("not configure the environment variables for password")
	}

	credential := options.Credential{
		Username: c.user,
		Password: c.password,
	}

	cl, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+c.host+":"+c.port).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	return cl
}
