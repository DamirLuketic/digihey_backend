package handlers

import "github.com/DamirLuketic/digihey_backend/db"

func parseCreateVehicleEntry(make, model, oid string, year int64) db.Vehicle {
	return db.Vehicle{
		Make:  make,
		Model: model,
		OID:   oid,
		Year:  year,
	}
}
