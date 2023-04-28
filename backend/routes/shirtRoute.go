package routes

import (
	"context"
	"fmt"
	"katsuebackend/backend/database"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shirt struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title string `json:"title"`;
	Cost int `json:"cost"`;
	Quantity int `json:"quantity"`;
	Img string `json:"img"`;
}
var shirts[] Shirt = []Shirt{
	{Title: "Shirt0", Cost: 24, Quantity: 14, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fnike1.png?alt=media&token=13928678-c2b8-4cf7-b30d-16d3a7c725b6"},
	{Title: "Shirt1", Cost: 28, Quantity: 15, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fnike2.png?alt=media&token=156768ba-5505-4d40-ae6f-b39c3f78a630"},
	{Title: "Shirt2", Cost: 22, Quantity: 16, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fnike3.png?alt=media&token=76bda708-1e0a-4e78-a239-d09a7da74b6d"},
	{Title: "Shirt3", Cost: 25, Quantity: 18, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fnike5.png?alt=media&token=0086a261-3e62-4652-8de2-9300ca782ffa"},
	{Title: "Shirt4", Cost: 24, Quantity: 18, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fsweater2.png?alt=media&token=2c23f784-82ae-4b54-a41e-5b53e4e259fd"},
	{Title: "Shirt5", Cost: 21, Quantity: 18, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fnike1.png?alt=media&token=13928678-c2b8-4cf7-b30d-16d3a7c725b6"},
}

func GetAllShirts(c *gin.Context) {
	var mongoShirts[] Shirt; // = []Shirt{}
	var err error;
	var cursor *mongo.Cursor
	cursor, err = database.MongoKatuseDb.ShirtsCollection.Find(context.TODO(), bson.D{});
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"err with getting data from mongo"});
		return;
	}
	err = cursor.All(context.TODO(), &mongoShirts);
	if err != nil {
		panic(err);
	}
	
	c.IndentedJSON(200, mongoShirts);
}

type Data struct {
	Name string `json:"name"`;
	Team string `json:"team"`;
} 
func Test(c *gin.Context){
	var userData Data;
	var err error = c.BindJSON(&userData)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"item not found"})
		return 
	}
	fmt.Println(userData);
	fmt.Println(userData.Team);
	fmt.Println(userData.Name)
	c.IndentedJSON(http.StatusOK, userData)
}
// ------------ reomve from invetnroy handler ---------//

type ShirtHandler struct {
	Title string `json:"title"`;
	RemAmount int `json:"remAmount"`;
}
func removeFromDb(c *gin.Context) {
	var shirtPtr *Shirt = &Shirt{};
	
	var shirtRequest ShirtHandler; //remShirt.Id, remShirt.RemAmount
	var err error = c.BindJSON(&shirtRequest);
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"request unsuccessful..."});
		return;
	}
	fmt.Println(shirtRequest);
	filter := bson.D{{Key: "title", Value: shirtRequest.Title}};
	shirtCollection := database.MongoKatuseDb.DB.Collection("shirts");
	var getErr error = shirtCollection.FindOne(context.TODO(), filter).Decode(shirtPtr);
	if getErr != nil {
		c.IndentedJSON(400, gin.H{"message":"shirt not found"})
		return;
	}

	var nwQuan int = shirtPtr.Quantity - shirtRequest.RemAmount;
	update := bson.D{ { "$set", bson.D{{"quantity", nwQuan}} } };
	result, updErr := shirtCollection.UpdateOne(context.TODO(), filter, update);
	if updErr != nil {
		c.IndentedJSON(400, gin.H{"message":"could not update shirt"});
		return;
	}
	c.IndentedJSON(400, result);
}

func addShirtMongo(c *gin.Context){
	shirtsCollection := database.MongoKatuseDb.DB.Collection("shirts");
	for _,item := range shirts {
		var cShirt *Shirt = &item;
		userResult, err := shirtsCollection.InsertOne(context.TODO(), cShirt);
		if err != nil{
			c.IndentedJSON(400, gin.H{"message":"erroor"});
		}
		fmt.Println(userResult);
	}
}


func getOneShirt(c *gin.Context) {
	var title string = c.Param("title")
	var shirtAddress Shirt= Shirt{};
	var shirtPtr *Shirt = &shirtAddress;

	filter := bson.D{{Key: "title", Value: title}};
	shirtCollection := database.MongoKatuseDb.DB.Collection("shirts");
	var err error = shirtCollection.FindOne(context.TODO(), filter).Decode(shirtPtr);
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"could not get data from mongoDB"});
		fmt.Println(err);
		return;
	}
	c.IndentedJSON(200, shirtAddress);

}

func ShirtsRoute(g *gin.RouterGroup){
	g.GET("/getAllShirts", GetAllShirts);
	g.GET("/getOneShirt/:title", getOneShirt);
	g.GET("/test", Test)
	g.PATCH("/remShirt", removeFromDb);
	g.POST("/addShirts" ,addShirtMongo);
}
