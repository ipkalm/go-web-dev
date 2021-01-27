package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/ipkalm/go-web-dev/015-go-mongo/117-connect-to-mongo/001/controller"
	"github.com/julienschmidt/httprouter"
)

// Run launch app
func Run() {
	uc := controller.NewUserController(getClient())
	r := httprouter.New()

	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.DeleteUser)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func getClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI("mongodb://root:WTzHXg80bRKUAOQC2mhF@127.0.0.1:27017")

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client
}
