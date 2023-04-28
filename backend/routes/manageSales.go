package routes

import (
	"context"
	"fmt"
	//"fmt"
	"katsuebackend/backend/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"katsuebackend/backend/models"
)

type Sale struct{
	Amount int  `json:"amount"`
	Items map[string]int `json:"items"`
}

func addSaleToDb(c *gin.Context){
	var currSale Sale = Sale{};
	var currSalePtr *Sale = &currSale;
	var err error = c.BindJSON(currSalePtr);
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"could not add to database"});
		return;
	}
	var salesCollection *mongo.Collection = database.MongoKatuseDb.DB.Collection("sales");

	res, dberr := salesCollection.InsertOne(context.TODO(), currSale);
	if dberr != nil {
		c.IndentedJSON(400, gin.H{"msg":"could not add to db"});
		return;
	}
	c.IndentedJSON(200, res);
}
func updateInvetory(c *gin.Context){
	//var currItem Shirt = Shirt{};
	//var ptrCurrItem *Shirt = &currItem;

	var currSale = Sale{};
	var currSalePtr *Sale = &currSale;

	var err error = c.BindJSON(currSalePtr);
	if err != nil {
		c.IndentedJSON(400, gin.H{"message":"could not add to database"});
		return;
	}

	var shirtCollection *mongo.Collection = database.MongoKatuseDb.DB.Collection("shirts");
	var pantsCollection *mongo.Collection = database.MongoKatuseDb.DB.Collection("pants");

	for key,val := range (*currSalePtr).Items {
		if key[:5] == "Shirt" {
			updateDb(shirtCollection, val, key);
		}else{
			updateDb(pantsCollection, val, key);
		}
	}
	
}
func updateDb(col *mongo.Collection, amount int, title string){
	var currItems Shirt = Shirt{};
	var ptrCurrItems *Shirt = &currItems;
	filter := bson.D{{Key: "title", Value: title}}
	var err error = col.FindOne(context.TODO(), filter).Decode(ptrCurrItems);
	if err != nil {
		fmt.Println(err);
	}
	var nwQuan int = currItems.Quantity - amount;
	update := bson.D{ { "$set", bson.D{{"quantity", nwQuan}} } };
	_, upErr := col.UpdateOne(context.TODO(), filter, update);
	if upErr != nil {
		fmt.Println(upErr);
	}
	
}

func SalesRoute(g *gin.RouterGroup){
	g.POST("/addSale", addSaleToDb);
	g.PATCH("/udpateInventory", updateInvetory);
}