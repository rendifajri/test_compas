package structs

type User struct {
	Id           string             `json:"id" db:"id"`
	Name         string             `json:"name" db:"name"`
	Point        UserPoint          `json:"point"`
	PointHistory []UserPointHistory `json:"point_history"`
}

type UserPoint struct {
	Point int `json:"point" db:"point"`
}

type UserPointHistory struct {
	Point      int `json:"point" db:"point"`
	PointFinal int `json:"point_final" db:"point_final"`
}
