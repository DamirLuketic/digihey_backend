package handlers

type CreateVehicleRequest struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int64  `json:"year"`
	Oid   string `json:"oid"`
}
