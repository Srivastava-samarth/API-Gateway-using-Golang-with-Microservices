package handlers

import (
	"context"
	"log"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Payment struct{
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    OrderID   primitive.ObjectID `json:"order_id" bson:"order_id"`
    Amount    float64            `json:"amount"`
    Status    string             `json:"status"`
    CreatedAt time.Time          `json:"created_at"`
}

var paymentsCollection *mongo.Collection

func init(){
    mongoUri := os.Getenv("MONGODB_URI")
    client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	paymentsCollection = client.Database("api-gateway").Collection("payments")
}

func GetPayments(c *fiber.Ctx) error{
    var payments []Payment;
    cursor,err := paymentsCollection.Find(context.Background(),bson.M{});
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch payments"})
    }
    defer cursor.Close(context.Background());

    for cursor.Next(context.Background()){
        var payment Payment
        if err:= cursor.Decode(&payment); err!=nil{
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Error decoding payments"})
        }
        payments = append(payments,payment)
    }
    return c.JSON(payments);
}

func CreatePayment(c *fiber.Ctx) error{
    var payment Payment;
    if err:=c.BodyParser(&payment);err!=nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Request"})
    }
    payment.ID = primitive.NewObjectID();
    payment.CreatedAt = time.Now();
    payment.Status = "Pending"

    _,err := paymentsCollection.InsertOne(context.Background(),payment)
    if err!=nil{
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Could not process payment"})
    }
    return c.Status(fiber.StatusCreated).JSON(payment);
}

func UpdatePayment(c *fiber.Ctx) error{
    id := c.Params("id")
    objID,err := primitive.ObjectIDFromHex(id);
    if err!=nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid payment ID"})
    }
    var payment Payment;
    if err:=c.BodyParser(&payment);err!=nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid request"})
    }
    update := bson.M{
        "$set":payment,
    }
    _,err = paymentsCollection.UpdateOne(context.Background(),bson.M{"_id":objID},update);
    if err!=nil{
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Error in updating the payment"})
    }
    return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message":"Payment updated successfully"})
}

func DeletePayment(c *fiber.Ctx) error{
    id := c.Params("id");
    objID,err := primitive.ObjectIDFromHex(id);
    if err!=nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid Pyament ID"})
    }
    _,err = paymentsCollection.DeleteOne(context.Background(),bson.M{"_id":objID})
    if err!=nil{
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Could not delete teh payment"})
    }
    return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message":"Payment deleted successfully"})
}
 