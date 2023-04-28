package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/joho/godotenv"
	"os"
)



//context that defines how long you can make a request;
type MongoDatabaseInstance struct {
	DB *mongo.Database
	UserCollection *mongo.Collection
	ShirtsCollection *mongo.Collection
	PantsCollection *mongo.Collection
}
var MongoKatuseDb MongoDatabaseInstance;
//var collection mongo.Collection; //pointer to mongo collection
func Init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
    	panic(envErr);
	};
	var connection string = os.Getenv("MONGO_URL");


	clientOption := options.Client().ApplyURI(connection);

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption); //background keeps happening in background
	if err != nil {
		panic(err);
	}
	fmt.Println("connected to backend");
	Database := client.Database("katsuedb")
	UserCollection := Database.Collection("users");
	ShirtsCollection := Database.Collection("shirts");
	PantsCollection := Database.Collection("pants");

	MongoKatuseDb = MongoDatabaseInstance{
		DB: Database,
		UserCollection: UserCollection,
		ShirtsCollection: ShirtsCollection,
		PantsCollection: PantsCollection,
	}
	//shirtsCollection = Database.Collection("shirts");

}
