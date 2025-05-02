package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	// main.go
	manufacturingVehicle := ManufacturingDirector{}
	bicycleBuilder := &BicycleBuilder{}

	manufacturingVehicle.SetBuilder(bicycleBuilder)
	manufacturingVehicle.Construct()

	bicycle := bicycleBuilder.GetVehicle()
	fmt.Printf("Vehicle is %s has %d wheels, %d seats.", bicycle.Structure, bicycle.Wheels, bicycle.Seats)

}
