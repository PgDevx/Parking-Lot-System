package validation

type ValidateCreateParkingLot struct {
	NoOfSlot int `json:"no_of_slot" validate:"required,gte=10"`
}

type ValidateParkCar struct {
	Code           string `json:"code,omitempty" bson:"code,omitempty"`
	RegistrationNo int    `json:"registration_no" validate:"required"`
	Color          string `json:"color" validate:"required"`
}

type ValidateLeaveParking struct {
	RegistrationNo int `json:"registration_no" validate:"required"`
}

type ValidateReturnSameColorVehicles struct {
	Color string `json:"color" validate:"required"`
}

type ValidateParkedSlot struct {
	RegistrationNo int `json:"registration_no" validate:"required"`
}
