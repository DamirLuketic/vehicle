package handlers

import "github.com/DamirLuketic/vehicle/db"

func parseCreateVehicleEntry(make, model, oid string, year int64) db.Vehicle {
	return db.Vehicle{
		Make:  make,
		Model: model,
		OID:   oid,
		Year:  year,
	}
}

func parseVehicleDataToJSON(vehicle db.Vehicle) VehicleJSON {
	return VehicleJSON{
		Make:  vehicle.Make,
		Model: vehicle.Model,
		Year:  vehicle.Year,
		Oid:   vehicle.OID,
	}
}

func parseVehiclesDataToJSON(vehicles []db.Vehicle) (vehiclesJSON []VehicleJSON) {
	for _, vehicle := range vehicles {
		vehiclesJSON = append(vehiclesJSON, parseVehicleDataToJSON(vehicle))
	}
	return vehiclesJSON
}
