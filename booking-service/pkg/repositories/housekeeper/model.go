package housekeeper

type HouseKeeper struct {
	HouseKeeperID string `json:"house_keeper_id" bson:"house_keeper_id"`
	Name          string `json:"name" bson:"name"`
	PhoneNumber   string `json:"phone_number" bson:"phone_number"`
}
