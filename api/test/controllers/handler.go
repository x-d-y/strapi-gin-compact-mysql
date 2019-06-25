package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
	getRoutes "github.com/xdy/gin/utils/getRoutes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Routers struct {
}

var cPath = "test"

func bindHandler() map[string]reflect.Value {
	var ruTest Routers
	type ControllerMapsType map[string]reflect.Value
	crMap := make(ControllerMapsType, 0)
	vf := reflect.ValueOf(&ruTest)
	mNum := vf.NumMethod()
	vft := vf.Type()
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		crMap[mName] = vf.Method(i)
	}
	return crMap
}

var collection *mongo.Collection

func LoadHandler(router *gin.Engine, routerHandeler map[string][]getRoutes.RouteInfo, client_ *mongo.Client) {
	collection_ := client_.Database("test").Collection("trainers")
	collection = collection_
	crMap := bindHandler()
	//fmt.Println(crMap)
	for k, v := range routerHandeler {
		if k == cPath {
			groupRoute := router.Group(k)
			for _, u := range v {
				u_ := u
				Handler := []rune(u_.Handler)
				if Handler[0] >= 97 && Handler[0] <= 122 {
					Handler[0] -= 32
				}
				if u_.Method == "GET" {
					groupRoute.GET(u_.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u_.Method == "POST" {
					groupRoute.POST(u_.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u_.Method == "PUT" {
					groupRoute.PUT(u_.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u_.Method == "DELETE" {
					groupRoute.DELETE(u_.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u_.Method == "PATCH" {
					groupRoute.PATCH(u_.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else {
					fmt.Println("unmatch any method")
					return
				}
			}
		}
	}
}

/**************************************  custom  ******************************************/
type Model struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
	City string `form:"city"`
}

func (r *Routers) InsertOne(c *gin.Context) {
	var model Model
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, &model); err != nil {
		log.Printf("transfer to json err")
	}
	insertResult, err := collection.InsertOne(context.TODO(), model)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
func (r *Routers) UpdateOne(c *gin.Context) {
	var model Model
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, &model); err != nil {
		log.Printf("transfer to json err")
	}
	paramas := c.Param("_id")
	_ = paramas

	objctId, _ := primitive.ObjectIDFromHex("5d10a4897150fd13f7d8f3e1")
	updateResult, err := collection.UpdateOne(context.TODO(), bson.D{{"_id", objctId}}, bson.D{{"$set", model}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	fmt.Println("this is test2Handler")
}
func (r *Routers) FindOne(c *gin.Context) {
	paramas := c.Param("_id")
	_ = paramas
	var result Model
	objctId, _ := primitive.ObjectIDFromHex(paramas)

	err := collection.FindOne(context.TODO(), bson.D{{"_id", objctId}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

}
func (r *Routers) Find(c *gin.Context) {
	var model Model
	if c.ShouldBind(&model) == nil {
	}
	findOptions := options.Find()
	var results []*Model
	cur, err := collection.Find(context.TODO(), model, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem Model
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}
	fmt.Println(len(results))

}
func (r *Routers) DeleteOne(c *gin.Context) {
	paramas := c.Param("_id")
	_ = paramas
	objctId, _ := primitive.ObjectIDFromHex("5d10a4897150fd13f7d8f3e1")
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{"_id", objctId}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}
