package observer

import (
	"encoding/json"
	"log"

	"github.com/gofiber/websocket/v2"
)


type Observer interface {
	Update(data interface{})
}

type Subject struct {
	observers []Observer
}

type WebsocketObserver struct {
	Conn *websocket.Conn
}

var SubjectInstance = &Subject{}



func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Detach(o Observer) {
	var newObservers []Observer
	for _, observer := range s.observers {
		if observer != o {
			newObservers = append(newObservers, observer)
		}
	}
	s.observers = newObservers
}

func (s *Subject) Notify(data interface{}) {
	for _, observer := range s.observers {
		observer.Update(data)
	}
}

func (wso *WebsocketObserver) Update(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling data: %s", err)
		return
	}

	err = wso.Conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Printf("Error sending message: %s", err)
		return
	}

}

