### Description:

- Dockerized golang API with MySQL DB.
- On API start MySQL DB is initialized, with proper vehicle table.
- OID is used as a unique identifier.
- On DB initialization data from "Vehicles.json" are read and set in DB.
- API contains three endpoints:
  - /api/vehicles [GET]
     - return all vehicles data
  - /api/create_vehicle [POST]
     - create new vehicle entry
     - body: [CreateVehicleRequest](api/handlers/types.go)
  - /api/delete_vehicle [DELETE]
     - remove vehicle entry
     - Query Param: "oid"
- Basic Auth used for API:
  - Development username: "digihey"
  - Development password: "digihey"
- The port on which API is exposed is "8080"

## Local run:
- docker-compose build
- docker-compose -f .\docker-compose.yaml up
