package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Parking struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SlotNo int                `json:"slot_no,omitempty" bson:"slot_no,omitempty"`
	Status string             `json:"status,omitempty" bson:"status,omitempty"`
	Car    []Car              `json:"car,omitempty" bson:"car,omitempty"`
}
type ParkingCheck struct {
	SlotNo int    `json:"slot_no,omitempty" bson:"slot_no,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty"`
	Car    *Car   `json:"car,omitempty" bson:"car,omitempty"`
}

type Car struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RegistrationNo int                `json:"registration_no,omitempty" bson:"registration_no,omitempty"`
	Color          string             `json:"color,omitempty" bson:"color,omitempty"`
}

type Parked struct {
	SlotNo int `json:"slot_no,omitempty" bson:"slot_no,omitempty"`
}
