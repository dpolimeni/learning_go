// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dpolimeni/fiber_app/ent/events"
	"github.com/dpolimeni/fiber_app/ent/reservations"
	"github.com/dpolimeni/fiber_app/ent/user"
)

// Reservations is the model entity for the Reservations schema.
type Reservations struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReservationsQuery when eager-loading is set.
	Edges               ReservationsEdges `json:"edges"`
	events_reservations *int32
	user_reservations   *int
	selectValues        sql.SelectValues
}

// ReservationsEdges holds the relations/edges for other nodes in the graph.
type ReservationsEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Event holds the value of the event edge.
	Event *Events `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReservationsEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReservationsEdges) EventOrErr() (*Events, error) {
	if e.loadedTypes[1] {
		if e.Event == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: events.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reservations) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reservations.FieldID:
			values[i] = new(sql.NullString)
		case reservations.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case reservations.ForeignKeys[0]: // events_reservations
			values[i] = new(sql.NullInt64)
		case reservations.ForeignKeys[1]: // user_reservations
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reservations fields.
func (r *Reservations) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reservations.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
		case reservations.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case reservations.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field events_reservations", value)
			} else if value.Valid {
				r.events_reservations = new(int32)
				*r.events_reservations = int32(value.Int64)
			}
		case reservations.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_reservations", value)
			} else if value.Valid {
				r.user_reservations = new(int)
				*r.user_reservations = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Reservations.
// This includes values selected through modifiers, order, etc.
func (r *Reservations) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Reservations entity.
func (r *Reservations) QueryUser() *UserQuery {
	return NewReservationsClient(r.config).QueryUser(r)
}

// QueryEvent queries the "event" edge of the Reservations entity.
func (r *Reservations) QueryEvent() *EventsQuery {
	return NewReservationsClient(r.config).QueryEvent(r)
}

// Update returns a builder for updating this Reservations.
// Note that you need to call Reservations.Unwrap() before calling this method if this Reservations
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reservations) Update() *ReservationsUpdateOne {
	return NewReservationsClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Reservations entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reservations) Unwrap() *Reservations {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reservations is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reservations) String() string {
	var builder strings.Builder
	builder.WriteString("Reservations(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ReservationsSlice is a parsable slice of Reservations.
type ReservationsSlice []*Reservations