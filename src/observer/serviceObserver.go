package observer

type ServiceObserver interface {
	GetName() string
	FindAll() interface{}
	FindByID(id interface{}) (interface{}, error)
	CreateData(input interface{}) error
	UpdateData(id interface{}, input interface{}) error
	DeleteData(id interface{}) error
}
