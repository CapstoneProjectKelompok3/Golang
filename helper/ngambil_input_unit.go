package helper

type UnitCount struct {
	UnitPolisi    int
	UnitAmbulance int
	UnitDamkar    int
	UnitDishub    int
	UnitSAR       int
}

func InputUnit(input UnitCount) UnitCount {
	return UnitCount{
		UnitPolisi:    input.UnitPolisi,
		UnitAmbulance: input.UnitAmbulance,
		UnitDamkar:    input.UnitDamkar,
		UnitDishub:    input.UnitDishub,
		UnitSAR:       input.UnitSAR,
	}
}