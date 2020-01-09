package parser

import (
	"errors"
	"fmt"
)

const (
	KEY_DEFAULT = 0
	KEY_CALCULATED = 1
	KEY_GIVEN = 2
)

type Key struct {
	Name 	string
	Value  	bool
	State 	int
}

func (k Key) Eval(key string) (bool, error) {
	if k.Name == key {
		e := errors.New(fmt.Sprintf("Error: %s is self-referring.\n", key))
		return k.Value, e
	}
	return k.Value, nil
}

func (key Key) Init(val bool) (e error) {
	key.Value = true
	if key.State == KEY_DEFAULT {
		key.State = KEY_GIVEN
	} else {
		e = errors.New(fmt.Sprintf("Warning: %s defined true multiple times.\n", key.Name))
	}
	return
}
func (key Key) Set(val bool) (e error) {
	if key.State == KEY_DEFAULT {
		key.State = KEY_CALCULATED
		key.Value = val
	} else if key.State == KEY_GIVEN {
		e = errors.New(fmt.Sprintf("Error: %s violates your statements.\n", key.Name))
	} else if key.State == KEY_CALCULATED {
		e = errors.New(fmt.Sprintf("Error: %s was already calculated to be %t.\n", key.Name, key.Value))
	}
	return
}

