package register

type Driver interface {
	Create()
	Update(a string, b int) error
}

var Drivers = make(map[string]Driver)

func Register(name string, driver Driver) {
	Drivers[name] = driver
}
