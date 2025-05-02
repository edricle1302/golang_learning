package builder

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}
