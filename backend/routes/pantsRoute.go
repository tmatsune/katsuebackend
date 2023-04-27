package routes

import (
	"fmt"
	"net/http"
	"context"
	"katsuebackend/backend/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type Pants struct {
	ID string  `json:"id"`;
	Title string `json:"title"`;
	Cost int `json:"cost"`;
	Quantity int `json:"quantity"`;
	Img string `json:"img"`;
}
var localPants[] Pants = []Pants{
	{ID: "0", Title: "Pants0", Cost: 24, Quantity: 8, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fnike1.png?alt=media&token=13928678-c2b8-4cf7-b30d-16d3a7c725b6"},
	{ID: "1", Title: "Pants1", Cost: 28, Quantity: 8, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fnike1.png?alt=media&token=13928678-c2b8-4cf7-b30d-16d3a7c725b6"},
	{ID: "2", Title: "Pants2", Cost: 22, Quantity: 12, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fnike3.png?alt=media&token=76bda708-1e0a-4e78-a239-d09a7da74b6d"},
	{ID: "3", Title: "Pants3", Cost: 26, Quantity: 10, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fnike3.png?alt=media&token=76bda708-1e0a-4e78-a239-d09a7da74b6d"},
	{ID: "4", Title: "Pants4", Cost: 21, Quantity: 9, Img:"https://firebasestorage.googleapis.com/v0/b/bildapp-ec335.appspot.com/o/test%2Fnike5.png?alt=media&token=0086a261-3e62-4652-8de2-9300ca782ffa"},
}

func addAllPants(c *gin.Context){
	pantsCollection := database.MongoKatuseDb.DB.Collection("pants");
	for _,item := range localPants {
		var cPants *Pants = &item;
		var pantsResult *mongo.InsertOneResult;
		var err error;
		pantsResult, err = pantsCollection.InsertOne(context.TODO(), cPants);
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message":"could not upload to mongodb"});
			return;
		}
		fmt.Print(pantsResult);
	}
	c.IndentedJSON(200, gin.H{"message":"pants added to mongo"});
}
func getOnePants(c *gin.Context){
	var title string = c.Param("title");
	var pantsAdrs Pants = Pants{};
	var pantsPtr *Pants = &pantsAdrs;

	filter := bson.D{{Key: "title", Value: title}};
	pantsCollection := database.MongoKatuseDb.DB.Collection("pants");
	var err error = pantsCollection.FindOne(context.TODO(), filter).Decode(pantsPtr);
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"could not find pants"});
	}
	c.IndentedJSON(200, pantsPtr);
}
func getAllPants(c *gin.Context){
	var allShirts []Shirt = []Shirt{};
	var ptrAllShirts *[]Shirt = &allShirts;

	var err error;
	var cursor *mongo.Cursor;
	pantsCollection := database.MongoKatuseDb.DB.Collection("pants");
	cursor, err = pantsCollection.Find(context.TODO(), bson.D{});

	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"could not get pants data"});
		return;
	}
	err = cursor.All(context.TODO(), ptrAllShirts);
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error senging data to client"});
		return;
	}
	//   |allShirts| acutal item data;
	//   |ptrAllShirts| points to address of all list;
	c.IndentedJSON(200, allShirts);
}
type ReqPantsHandler struct {
	Title string `json:"title"`
	RemAmount int `json:"remAmount"`
}
func updatePants(c *gin.Context){
	var currPants Pants = Pants{};      // pants from mongodb
	var ptrCurrPants *Pants = &currPants;

	var pantsRequest ReqPantsHandler = ReqPantsHandler{};  //getting user data
	var ptrPantsReq *ReqPantsHandler = &pantsRequest;

	var err error = c.BindJSON(ptrPantsReq);
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"could not complete request"});
		return;
	}
	filter := bson.D{{Key: "title", Value: pantsRequest.Title}};
	shirtCollection := database.MongoKatuseDb.DB.Collection("pants");
	var getErr error = shirtCollection.FindOne(context.TODO(), filter).Decode(ptrCurrPants);

	if getErr != nil {
		c.IndentedJSON(400, gin.H{"message":"could not find pants"});
		return;
	}

	var nwQuan int = currPants.Quantity - ptrPantsReq.RemAmount;
	update := bson.D{ { "$set", bson.D{{"quantity", nwQuan}} } };
	result, updErr := shirtCollection.UpdateOne(context.TODO(), filter, update);
	if updErr != nil {
		c.IndentedJSON(400, gin.H{"message":"could not update shirt"});
		return;
	}
	c.IndentedJSON(200, result);
}

func PantsRoute(g *gin.RouterGroup){
	g.POST("/addPants", addAllPants);
	g.GET("/getOnePants/:title", getOnePants);
	g.GET("/getAllPants", getAllPants);
	g.PATCH("/remPants", updatePants)
}