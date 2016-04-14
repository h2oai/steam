package web

//go:generate piping -idl service.pipe -go service.go -ts ../../gui/src/proxy.ts

import (
// "fmt"
// "github.com/h2oai/steam/lib/uniq"
//"time"
)

// func NewID() (ID, error) {
// 	u, err := uniq.NewID()
// 	return ID(u), err
// }

// func NewTicket() (*Ticket, error) {
// 	id, err := NewID()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Ticket{id}, nil
// }

// func ToTimestamp(t time.Time) Timestamp {
// 	return Timestamp(uniq.Timestamp(t))
// }

// func Now() Timestamp {
// 	return Timestamp(uniq.Now())
// }

// func NewEvent(id ID, eventType EventType, origin, message string) *Event {
// 	return &Event{
// 		id,
// 		eventType,
// 		origin,
// 		message,
// 		Now(),
// 	}
// }

// func NewQuitEvent(id ID, origin string) *Event {
// 	return &Event{
// 		id,
// 		EventQuit,
// 		origin,
// 		"",
// 		Now(),
// 	}
// }

// func (e *Event) String() string {
// 	return fmt.Sprintf("Event job=%s, type=%s, origin=%s: %s", e.ID, e.Type, e.Origin, e.Message)
// }
