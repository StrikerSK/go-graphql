package observer

type CustomObserver struct {
	observers   []ServiceObserver
	mainService string
}

// Main Service name will be needed to use in case only one service is needed
func (co *CustomObserver) SetMainService(newMain string) {
	co.mainService = newMain
}

func (co *CustomObserver) Register(newObserver ServiceObserver) {
	co.observers = append(co.observers, newObserver)
}

func (co *CustomObserver) Unregister(observerName string) {
	var filteredObservers []ServiceObserver

	for _, item := range co.observers {
		if item.GetName() != observerName {
			filteredObservers = append(filteredObservers, item)
		}
	}

	co.observers = filteredObservers
}

func (co *CustomObserver) CreateData(input interface{}) error {
	for _, item := range co.observers {
		if err := item.CreateData(input); err != nil {
			return err
		}
	}

	return nil
}

func (co *CustomObserver) UpdateData(input interface{}) error {
	for _, item := range co.observers {
		if err := item.UpdateData(input); err != nil {
			return err
		}
	}

	return nil
}

func (co *CustomObserver) DeleteData(id interface{}) error {
	for _, item := range co.observers {
		if err := item.DeleteData(id); err != nil {
			return err
		}
	}

	return nil
}

func (co *CustomObserver) FindAll() (result interface{}, err error) {
	for _, item := range co.observers {
		if item.GetName() == co.mainService {
			result, err = item.FindAll()
			break
		}
	}

	return
}

func (co *CustomObserver) FindByID(id interface{}) (result interface{}, err error) {
	for _, item := range co.observers {
		if item.GetName() == co.mainService {
			result, err = item.FindByID(id)
			break
		}
	}

	return
}
