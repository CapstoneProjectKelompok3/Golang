package helper

type UnitCount struct {
	UnitPolisi    int
	UnitAmbulance int
	UnitDamkar    int
	UnitDishub    int
	UnitSAR       int
}

var (
	UnitPolisiInput    = 0
	UnitAmbulanceInput = 0
	UnitDamkarInput    = 0
	UnitDishubInput    = 0
	UnitSARInput       = 0
)

func InputUnit(input UnitCount) UnitCount {
	UnitPolisiInput = input.UnitPolisi
	UnitAmbulanceInput = input.UnitAmbulance
	UnitDamkarInput = input.UnitDamkar
	UnitDishubInput = input.UnitDishub
	UnitSARInput = input.UnitSAR
	return UnitCount{
		UnitPolisi:    UnitPolisiInput,
		UnitAmbulance: UnitAmbulanceInput,
		UnitDamkar:    UnitDamkarInput,
		UnitDishub:    UnitDishubInput,
		UnitSAR:       UnitSARInput,
	}
}

func UnitDriver() UnitCount {
	return UnitCount{
		UnitPolisi:    UnitPolisiInput,
		UnitAmbulance: UnitAmbulanceInput,
		UnitDamkar:    UnitDamkarInput,
		UnitDishub:    UnitDishubInput,
		UnitSAR:       UnitSARInput,
	}
}
