package core

// Item is the struct that represents the item in the database
type Item struct {
	Site        string  `bson:"site"`
	ID          string  `bson:"id"`
	Price       float64 `bson:"price"`
	StartTime   string  `bson:"start_time"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Nickname    string  `bson:"nickname"`
}
