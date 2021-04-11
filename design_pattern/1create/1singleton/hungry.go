package singleton

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = new(Singleton)
}

func GetInstance() *Singleton {
	return singleton
}
