package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Parking struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Code   string             `json:"code,omitempty" bson:"code,omitempty"`
	SlotNo int                `json:"slot_no,omitempty" bson:"slot_no,omitempty"`
	Status string             `json:"status,omitempty" bson:"status,omitempty"`
}
type ParkingLots struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Code      string               `json:"code,omitempty" bson:"code,omitempty"`
	ParkingID []primitive.ObjectID `json:"parking,omitempty" bson:"parking,omitempty"`
}
type ParkingCheck struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Code   string             `json:"code,omitempty" bson:"code,omitempty"`
	SlotNo int                `json:"slot_no,omitempty" bson:"slot_no,omitempty"`
	Status string             `json:"status,omitempty" bson:"status,omitempty"`
	Car    []Car              `json:"car,omitempty" bson:"car,omitempty"`
}

type Car struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ParkingID      primitive.ObjectID `json:"parking_id,omitempty" bson:"parking_id,omitempty"`
	RegistrationNo int                `json:"registration_no,omitempty" bson:"registration_no,omitempty"`
	Color          string             `json:"color,omitempty" bson:"color,omitempty"`
}

type Parked struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Code   string             `json:"code,omitempty" bson:"code,omitempty"`
	SlotNo int                `json:"slot_no,omitempty" bson:"slot_no,omitempty"`
}
