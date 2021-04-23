package app

import (
	"context"
	"my/v1/errors"
	"my/v1/model"
	"my/v1/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateParkingLot(n int) ([]interface{}, error) {
	var parking []interface{}

	for i := 1; i <= n; i++ {
		parking = append(parking, model.Parking{
			SlotNo: i,
			Status: "empty",
		})
	}
	res, err := mongo.NewMongoStorage().Database.Collection(model.ParkingColl).InsertMany(context.TODO(), parking)
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func getEmptyParkingSlots() (int, error) {
	var empty model.Parking
	err := mongo.NewMongoStorage().Database.Collection(model.ParkingColl).FindOne(context.TODO(), bson.M{"status": "empty"}).Decode(&empty)
	if err != nil {
		return 0, errors.DBError.Wrapf(err, "Failed to querry Database")
	}
	slot := empty.SlotNo
	return slot, nil
}

func ParkCar(regNo int, color string) (*model.ParkingCheck, error) {

	car := &model.Car{
		RegistrationNo: regNo,
		Color:          color,
	}

	res, err := mongo.NewMongoStorage().Database.Collection(model.CarColl).InsertOne(context.TODO(), car)
	if err != nil {
		return nil, err
	}
	carID := res.InsertedID.(primitive.ObjectID)
	car.ID = carID
	slot, err := getEmptyParkingSlots()
	if err != nil {
		return nil, errors.NotFound.Wrapf(err, "Parking slot Full")
	}

	filter := bson.M{"slot_no": slot}
	update := bson.M{"$push": bson.M{"car": car}, "$set": bson.M{"status": "filled"}}
	ress, err := mongo.NewMongoStorage().Database.Collection(model.ParkingColl).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, errors.DBError.Wrap(err, "Failed to Park car at the parking slot")
	}
	if ress.MatchedCount == 0 {
		return nil, errors.NotFound.New("Parking slot not found")
	}
	if ress.ModifiedCount == 0 {
		return nil, errors.DBError.New("Failed to update parking lot")
	}
	park := &model.ParkingCheck{
		SlotNo: slot,
		Status: "filled",
		Car:    car,
	}

	return park, nil
}

func GetStatusOfParkingLot() ([]model.Parking, error) {
	var parking []model.Parking
	cur, err := mongo.NewMongoStorage().Database.Collection(model.ParkingColl).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, errors.DBError.Wrapf(err, "Failed to querry database")
	}
	if err := cur.All(context.TODO(), &parking); err != nil {
		return nil, err
	}

	return parking, nil
}

func RemoveCar(regNO int) (bool, error) {

	res, err := mongo.NewMongoStorage().Database.Collection(model.CarColl).DeleteOne(context.TODO(), bson.M{"registration_no": regNO})
	if err != nil {
		return false, err
	}
	if res.DeletedCount == 0 {
		return false, errors.NotFound.New("Car not found in Parking lot")
	}

	filter := bson.M{"car.registration_no": regNO}
	update := bson.M{"$unset": bson.M{"car": ""}, "$set": bson.M{"status": "empty"}}
	ress, err := mongo.NewMongoStorage().Database.Collection(model.ParkingColl).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, errors.DBError.Wrapf(err, "Failed to querry database")
	}
	if ress.MatchedCount == 0 {
		return false, errors.NotFound.New("Parking slot not found")
	}
	if ress.ModifiedCount == 0 {
		return false, errors.DBError.New("Failed to update parking lot")
	}

	return true, nil
}

func GetSameColorCar(color string) ([]model.Car, error) {
	var car []model.Car
	cur, err := mongo.NewMongoStorage().Database.Collection(model.CarColl).Find(context.TODO(), bson.M{"color": color})
	if err != nil {
		return nil, errors.DBError.Wrapf(err, "Failed to querry database")
	}
	if err := cur.All(context.TODO(), &car); err != nil {
		return nil, err
	}
	if car == nil {
		return nil, errors.NotFound.Newf("Car of color %s not present", color)
	}
	return car, nil
}

func GetParkedSlot(regNo int) (*model.Parked, error) {
	var p *model.Parked

	filter := bson.M{"car.registration_no": regNo}
	err := mongo.NewMongoStorage().Database.Collection(model.ParkingColl).FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		return nil, errors.DBError.Wrapf(err, "Failed to querry database")
	}
	return p, nil
}
