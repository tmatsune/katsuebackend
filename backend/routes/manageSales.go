package routes

import (
	"context"
	"fmt"
	//"fmt"
	"katsuebackend/backend/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"katsuebackend/backend/models"
)
/*
var allClothes[] models.Item = []models.Item {
	{Title:"Shirt0", Sold: 0}, {Title:"Shirt1", Sold: 0}, {Title:"Shirt2", Sold: 0}, {Title:"Shirt3", Sold: 0},
	{Title:"Shirt4", Sold: 0},{Title:"Shirt5", Sold: 0}, {Title:"Pants0", Sold: 0}, {Title:"Pants1", Sold: 0},
	{Title:"Pants2", Sold: 0},{Title:"Pants3", Sold: 0},{Title:"Pants4", Sold: 0},
};
*/
func geNewMap()([]models.Item){
	var allClothes[] models.Item = []models.Item {
	{Title:"Shirt0", Sold: 0}, {Title:"Shirt1", Sold: 0}, {Title:"Shirt2", Sold: 0}, {Title:"Shirt3", Sold: 0},
	{Title:"Shirt4", Sold: 0},{Title:"Shirt5", Sold: 0}, {Title:"Pants0", Sold: 0}, {Title:"Pants1", Sold: 0},
	{Title:"Pants2", Sold: 0},{Title:"Pants3", Sold: 0},{Title:"Pants4", Sold: 0},
	};
	return allClothes;
}

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
func getAllSales(c *gin.Context){
	var allSales[] Sale;
	//var salesColelction *mongo.Collection = database.MongoKatuseDb.DB.Collection("sales");
	var err error;
	retData(&allSales, &err);
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg":"problem with database"});
		fmt.Print(err)
		return;
	}
	c.IndentedJSON(200, allSales)
}

func ProccessData(c *gin.Context){
	var allSales[] Sale;
	var err error;
	retData(&allSales, &err);
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg":"problem with database"});
		fmt.Print(err)
		return;
	}
	var allClothes[] models.Item = geNewMap()
	//------------make map-------//
	var m  = map[string]*models.Item{}
	for i := 0; i < len(allClothes); i++ {
		m[allClothes[i].Title] = &allClothes[i]
	}
	//-------------create models in 
	var nodes[] models.Item = []models.Item {};
	for _,sale := range allSales {
		for k,v := range sale.Items{
			var no models.Item = models.Item{ Title:k, Sold:v };
			nodes = append(nodes, no);
		}
	}
	
	for _,val := range nodes {
		var curNum int = m[val.Title].Sold + val.Sold
		//fmt.Println(m[val.title])
		chagneVal(m[val.Title], curNum)
	}
	//for _,v := range m {
	//	fmt.Println(*v)
	//}

	c.IndentedJSON(200, m);

}
func retData(sals *[]Sale , er *error){
	var salesColelction *mongo.Collection = database.MongoKatuseDb.DB.Collection("sales");
	var err error;
	var cursor *mongo.Cursor;
	cursor, err = salesColelction.Find(context.TODO(), bson.D{});
	if err != nil {
		(*er) = err;
	}
	err = cursor.All(context.TODO(), sals);
	if err != nil {
		(*er) = err;
	}
}

func chagneVal(i *models.Item, num int){
	i.Sold = num
}

func SalesRoute(g *gin.RouterGroup){
	g.POST("/addSale", addSaleToDb);
	g.PUT("/udpateInventory", updateInvetory);
	g.GET("/getAllSales", getAllSales);
	g.GET("/procData", ProccessData)
}