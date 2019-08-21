package roles

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Role define un role
type Role struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Role        string             `json:"email" bson:"email"`
	Description string             `json:"username" bson:"username"`
}
