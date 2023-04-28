package routes

import (
	"context"
	"katsuebackend/backend/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
}
func displayUsers(c *gin.Context){
	var mongoUsers[] User = []User {};
	usersCollection := database.MongoKatuseDb.DB.Collection("users");

	var err error;
	var cursor *mongo.Cursor

	cursor, err = usersCollection.Find(context.TODO(), bson.D{});
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"coud not get data"});
		return;
	}
	err = cursor.All(context.TODO(), &mongoUsers);
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"coud not get data"});
		return;
	}
	c.IndentedJSON(200, mongoUsers)
}
type GetUserHandler struct {
	Email string `json:"email"`;
}

func adminOneUser(c *gin.Context){
	//getting info from frontend
	var usersDetails GetUserHandler;
	var err error = c.BindJSON(&usersDetails);
	//converts to josn data 
	var jsonUser User = User{};
	var ptrJsonUser *User = &jsonUser;

	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"could not get data"});
		return;
	}
	usersCollection := database.MongoKatuseDb.DB.Collection("users");
	filter := bson.D{{Key:"email", Value: usersDetails.Email}};
	var dbErr error = usersCollection.FindOne(context.TODO(), filter).Decode(ptrJsonUser);
	if dbErr != nil {
		c.IndentedJSON(400, gin.H{"message":"could not find user"})
		return;
	}
	c.IndentedJSON(200, jsonUser);
}
func adminDeleteUser(c *gin.Context){
	var usersDetails GetUserHandler;
	var err error = c.BindJSON(&usersDetails);
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"could not get data"});
		return;
	}
	usersCollection := database.MongoKatuseDb.DB.Collection("users");
	filter := bson.D{{Key:"email", Value: usersDetails.Email}};
	//var mongoErr *mongo.DeleteResult
	var dbErr error;
	_, dbErr = usersCollection.DeleteOne(context.TODO(), filter);
	if dbErr != nil {
		c.IndentedJSON(400, gin.H{"message":"could not find user"})
		return;
	}
	c.IndentedJSON(200, gin.H{"message":"user deleted"});
}

func ManageUsrsRoute(g *gin.RouterGroup){
	g.GET("/getAllUsers", displayUsers);
	g.GET("/adimGetUser", adminOneUser);
	g.POST("/adminRemUser", adminDeleteUser);
	
}
