package db

type DataStore interface {
	CreateVehicle(vehicle Vehicle) (Vehicle, error)
	DeleteVehicle(vehicleID string) error
	GetVehicles() ([]Vehicle, error)
}
