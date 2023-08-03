// responsavel por receber e injetar as dependencias relacionadas à alertas
package vehicles

import "fmt"

type VehicleModule struct {
}

func (a *VehicleModule) Initialize() error {

	fmt.Println("VehicleModule initialized!")
	return nil
}

func CreateModule() IVehicleModule {
	return &VehicleModule{}
}
