package observer

import (
	"github.com/strikersk/go-graphql/src"
	"log"
	"sync"
)

var cacheLock = &sync.Mutex{}
var observerInstance *CustomObserver

func GetObserverInstance() *CustomObserver {
	//To prevent expensive lock operations
	//This means that the observerInstance field is already populated
	if observerInstance == nil {
		cacheLock.Lock()
		defer cacheLock.Unlock()

		//Only one goroutine can create the singleton instance.
		if observerInstance == nil {
			log.Println("Creating Observer instance")

			tempInstance := CustomObserver{}

			//Set TodoService also as main service for observer
			tempInstance.Register(&src.TodoService{})
			tempInstance.SetMainService(src.ServiceName)

			tempInstance.Register(&src.LoggingTodoService{})

			log.Println("Observer initialization: created")
			observerInstance = &tempInstance
		} else {
			log.Println("Application Observer instance already created!")
		}
	} else {
		//log.Println("Application Observer instance already created!")
	}

	return observerInstance
}
