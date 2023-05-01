package data

type VehicleJSON struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int64  `json:"year"`
	Id    struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
}
