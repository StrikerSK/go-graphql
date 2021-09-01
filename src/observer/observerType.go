package observer

import "github.com/strikersk/go-graphql/src"

type CustomObserver struct {
	Observers []ServiceObserver
}

func (co *CustomObserver) Register(newObserver ServiceObserver) {
	co.Observers = append(co.Observers, newObserver)
}

func (co *CustomObserver) Unregister(observerName string) {
	var filteredObservers []ServiceObserver

	for _, item := range co.Observers {
		if item.GetName() != observerName {
			filteredObservers = append(filteredObservers, item)
		}
	}

	co.Observers = filteredObservers
}

func (co *CustomObserver) CreateData(input interface{}) error {
	for _, item := range co.Observers {
		if err := item.CreateData(input); err != nil {
			return err
		}
	}

	return nil
}

func (co *CustomObserver) UpdateData(id interface{}, input interface{}) error {
	for _, item := range co.Observers {
		if err := item.UpdateData(id, input); err != nil {
			return err
		}
	}

	return nil
}

func (co *CustomObserver) DeleteData(id interface{}) error {
	for _, item := range co.Observers {
		if err := item.DeleteData(id); err != nil {
			return err
		}
	}

	return nil
}

func (co *CustomObserver) FindAll() (result interface{}) {
	for _, item := range co.Observers {
		if item.GetName() == src.ServiceName {
			result = item.FindAll()
			break
		}
	}

	return
}

func (co *CustomObserver) FindByID(id interface{}) (result interface{}, err error) {
	for _, item := range co.Observers {
		if item.GetName() == src.ServiceName {
			result, err = item.FindByID(id)
			break
		}
	}

	return
}
