// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/dpolimeni/fiber_app/ent/events"
	"github.com/dpolimeni/fiber_app/ent/schema"
	"github.com/dpolimeni/fiber_app/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventsFields := schema.Events{}.Fields()
	_ = eventsFields
	// eventsDescName is the schema descriptor for name field.
	eventsDescName := eventsFields[1].Descriptor()
	// events.DefaultName holds the default value on creation for the name field.
	events.DefaultName = eventsDescName.Default.(string)
	// eventsDescCapacity is the schema descriptor for capacity field.
	eventsDescCapacity := eventsFields[2].Descriptor()
	// events.CapacityValidator is a validator for the "capacity" field. It is called by the builders before save.
	events.CapacityValidator = func() func(int16) error {
		validators := eventsDescCapacity.Validators
		fns := [...]func(int16) error{
			validators[0].(func(int16) error),
			validators[1].(func(int16) error),
		}
		return func(capacity int16) error {
			for _, fn := range fns {
				if err := fn(capacity); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// eventsDescDescription is the schema descriptor for description field.
	eventsDescDescription := eventsFields[3].Descriptor()
	// events.DefaultDescription holds the default value on creation for the description field.
	events.DefaultDescription = eventsDescDescription.Default.(string)
	// eventsDescID is the schema descriptor for id field.
	eventsDescID := eventsFields[0].Descriptor()
	// events.IDValidator is a validator for the "id" field. It is called by the builders before save.
	events.IDValidator = eventsDescID.Validators[0].(func(int32) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[0].Descriptor()
	// user.DefaultAge holds the default value on creation for the age field.
	user.DefaultAge = userDescAge.Default.(int)
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[5].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
}
