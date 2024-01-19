package models

// EstimatedFloors : ""
type EstimatedFloors struct {
	PropertyFloor `bson:",inline"`
}

// RefEstimatedFloors : ""
type RefEstimatedFloors struct {
	RefPropertyFloor `bson:",inline"`
}
