package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Item      string             `json:"item"`
    Quantity  int                `json:"quantity"`
    Price     int            `json:"price"`
    CreatedAt time.Time          `json:"created_at"`
}

var orderCollection *mongo.Collection

func init(){
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://srivastavasamarth94:1fJxNQp9n7WZNuRG@cluster0.8kt0yvj.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	orderCollection = client.Database("api-gateway").Collection("oders")
}

func GetOrders(c *fiber.Ctx) error {
    var orders Order
    cursor,err := orderCollection.Find(context.Background(),bson.M{})
    if err!=nil{
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Error in getting the orders"})
    }
    defer cursor.Close(context.Background())

    if err := cursor.All(context.Background(),&orders); err!=nil{
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"eeror":"Error in decoding the orders"})
    }
    return c.JSON(orders);
}

func CreateOrder(c *fiber.Ctx) error {
    var order Order
    if err := c.BodyParser(&order); err!=nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Request"})
    }
    order.ID = primitive.NewObjectID()
    order.CreatedAt = time.Now()

    _,err := orderCollection.InsertOne(context.Background(),order)
    if err!=nil{
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Could not create order"})
    }
    return c.Status(fiber.StatusCreated).JSON(order)
}

func UpdateOrder(c *fiber.Ctx) error{
   id := c.Params("id")
   objID,err := primitive.ObjectIDFromHex(id)
   if err!=nil{
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid order Id"})
   }
   var order Order
   if err := c.BodyParser(&order); err!=nil{
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Request"})
   }
   update := bson.M{
    "$set":order,
   }

   _,err = orderCollection.UpdateOne(context.Background(),bson.M{"_id":objID},update)
   if err!=nil{
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Could not update the order"})
   }
   return c.JSON(order);
}

func DeleteOrder(c *fiber.Ctx) error{
   id := c.Params("id")
   objID,err := primitive.ObjectIDFromHex(id)
   if err!=nil{
    c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Id"})
   }
   _,err = orderCollection.DeleteOne(context.Background(),bson.M{"_id":objID})
   if err!=nil{
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Could not delete order"})
   }
   return c.JSON(fiber.Map{"message":"Order deleted successfully"})
}