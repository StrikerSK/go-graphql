package observer

type ServiceObserver interface {
	GetName() string
	FindAll() (interface{}, error)
	FindByID(id interface{}) (interface{}, error)
	CreateData(input interface{}) error
	UpdateData(input interface{}) error
	DeleteData(id interface{}) error
}
