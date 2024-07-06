package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct{
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    OrderID   primitive.ObjectID `json:"order_id" bson:"order_id"`
    Amount    float64            `json:"amount"`
    Status    string             `json:"status"`
    CreatedAt time.Time          `json:"created_at"`
}

func GetPayments(c *fiber.Ctx) error{

}

func CreatePayment(c *fiber.Ctx) error{

}

func UpdatePayment(c *fiber.Ctx) error{

}

func DeletePayment(c *fiber.Ctx) error{
    
}
 