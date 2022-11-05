package factory_method

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{Gun: Gun{power: 1, name: "musket"}}
}
