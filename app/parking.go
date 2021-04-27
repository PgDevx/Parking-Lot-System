package app

import (
	"context"
	"fmt"
	"my/v1/errors"
	"my/v1/model"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (a *App) CreateParkingLot(n int) ([]interface{}, error) {
	var parking []interface{}
	var pid []model.Parking
	code, err := a.generateParkingCode()
	if err != nil {
		return nil, err
	}

	lots := &model.ParkingLots{
		Code: code,
	}
	ress, err := a.MongoDB.Database.Collection(model.ParkingLotsColl).InsertOne(context.TODO(), lots)
	if err != nil {
		return nil, errors.DBError.Wrapf(err, "Failed to query database")
	}
	lots.ID = ress.InsertedID.(primitive.ObjectID)

	for i := 1; i <= n; i++ {
		parking = append(parking, model.Parking{
			Code:   code,
			SlotNo: i,
			Status: "empty",
		})
	}
	res, err := a.MongoDB.Database.Collection(model.ParkingColl).InsertMany(context.TODO(), parking)
	if err != nil {
		return nil, errors.DBError.Wrapf(err, "Failed to query database")
	}

	opts := options.Find().SetProjection(bson.M{"_id": 1})
	cur, err := a.MongoDB.Database.Collection(model.ParkingColl).Find(context.TODO(), bson.M{"code": code}, opts)
	if err != nil {
		return nil, errors.DBError.Wrapf(err, "Failed to query database")
	}
	if err := cur.All(context.TODO(), &pid); err != nil {
		return nil, err
	}
	var id []primitive.ObjectID
	for _, v := range pid {
		id = append(id, v.ID)
	}

	filter := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"parking": id}}
	pk, err := a.MongoDB.Database.Collection(model.ParkingLotsColl).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, errors.DBError.Wrap(err, "Failed to Park car at the parking slot")
	}
	if pk.MatchedCount == 0 {
		return nil, errors.NotFound.New("Parking slot not found")
	}
	if pk.ModifiedCount == 0 {
		return nil, errors.DBError.New("Failed to update parking lot")
	}

	return res.InsertedIDs, nil
}

func (a *App) generateParkingCode() (string, error) {
	var parking *model.Parking
	opts := options.FindOne().SetProjection(bson.M{"_id": 0, "code": 1}).SetSort(bson.M{"_id": -1})
	err := a.MongoDB.Database.Collection(model.ParkingLotsColl).FindOne(context.TODO(), bson.M{}, opts).Decode(&parking)
	var prevCodeInt int
	if err != nil {
		if err == mongo.ErrNoDocuments {
			prevCodeInt = 0
		} else {
			return "", errors.DBError.Wrap(err, "Failed to set code")
		}
	} else {
		prevCode := strings.SplitN(parking.Code, "-", -1)
		prevCodeInt, _ = strconv.Atoi(prevCode[len(prevCode)-1])
	}
	parkingCode := fmt.Sprintf("P-%s", strconv.Itoa(prevCodeInt+1))

	return parkingCode, nil
}

func (a *App) getEmptyParkingSlots(code string) (int, primitive.ObjectID, error) {
	var empty model.Parking
	filter := bson.M{"code": code, "status": "empty"}
	err := a.MongoDB.Database.Collection(model.ParkingColl).FindOne(context.TODO(), filter).Decode(&empty)
	if err != nil {
		return 0, primitive.NilObjectID, errors.DBError.Wrapf(err, "Failed to querry Database")
	}
	slot := empty.SlotNo
	ID := empty.ID
	return slot, ID, nil
}

func (a *App) ParkCar(code string, regNo int, color string) (*model.Car, error) {

	slot, id, err := a.getEmptyParkingSlots(code)
	if err != nil {
		return nil, errors.NotFound.Wrapf(err, "Parking slot Full")
	}
	car := &model.Car{
		ParkingID:      id,
		RegistrationNo: regNo,
		Color:          color,
	}

	res, err := a.MongoDB.Database.Collection(model.CarColl).InsertOne(context.TODO(), car)
	if err != nil {
		return nil, errors.DBError.Wrap(err, "Failed to Park car at the parking slot")
	}
	car.ID = res.InsertedID.(primitive.ObjectID)

	filter := bson.M{"code": code, "slot_no": slot}
	update := bson.M{"$set": bson.M{"status": "filled"}}
	ress, err := a.MongoDB.Database.Collection(model.ParkingColl).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, errors.DBError.Wrap(err, "Failed to Park car at the parking slot")
	}
	if ress.MatchedCount == 0 {
		return nil, errors.NotFound.New("Parking slot not found")
	}
	if ress.ModifiedCount == 0 {
		return nil, errors.DBError.New("Failed to update parking lot")
	}

	return car, nil
}

func (a *App) GetStatusOfParkingLot() ([]model.ParkingCheck, error) {
	var parking []model.ParkingCheck

	lookupStage := bson.D{
		{
			Key: "$lookup", Value: bson.M{
				"from":         "car",
				"localField":   "_id",
				"foreignField": "parking_id",
				"as":           "car",
			},
		},
	}

	cur, err := a.MongoDB.Database.Collection(model.ParkingColl).Aggregate(context.TODO(), mongo.Pipeline{lookupStage})
	if err != nil {
		return nil, errors.DBError.Wrapf(err, "Failed to querry database")
	}
	if err := cur.All(context.TODO(), &parking); err != nil {
		return nil, err
	}

	return parking, nil
}

func (a *App) RemoveCar(regNO int) (bool, error) {
	var car model.Car
	err := a.MongoDB.Database.Collection(model.CarColl).FindOne(context.TODO(), bson.M{"registration_no": regNO}).Decode(&car)
	if err != nil {
		return false, errors.DBError.Wrapf(err, "Failed to querry Database")
	}

	res, err := a.MongoDB.Database.Collection(model.CarColl).DeleteOne(context.TODO(), bson.M{"registration_no": regNO})
	if err != nil {
		return false, err
	}
	if res.DeletedCount == 0 {
		return false, errors.NotFound.New("Car not found in Parking lot")
	}

	filter := bson.M{"_id": car.ParkingID}
	update := bson.M{"$set": bson.M{"status": "empty"}}
	ress, err := a.MongoDB.Database.Collection(model.ParkingColl).UpdateOne(context.TODO(), filter, update)
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

func (a *App) GetSameColorCar(color string) ([]model.Car, error) {
	var car []model.Car
	cur, err := a.MongoDB.Database.Collection(model.CarColl).Find(context.TODO(), bson.M{"color": color})
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

func (a *App) GetParkedSlot(regNo int) (*model.Parked, error) {
	var p *model.Parked

	var car model.Car
	err := a.MongoDB.Database.Collection(model.CarColl).FindOne(context.TODO(), bson.M{"registration_no": regNo}).Decode(&car)
	if err != nil {
		return nil, errors.DBError.Wrapf(err, "Failed to querry Database")
	}

	filter := bson.M{"_id": car.ParkingID}
	err = a.MongoDB.Database.Collection(model.ParkingColl).FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		return nil, errors.DBError.Wrapf(err, "Failed to querry database")
	}
	return p, nil
}
