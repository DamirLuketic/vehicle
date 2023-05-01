package handlers

type VehicleJSON struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int64  `json:"year"`
	Oid   string `json:"oid"`
}
