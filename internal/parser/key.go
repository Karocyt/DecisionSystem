package parser

import (
	"errors"
	"fmt"
)

const (
	KEY_DEFAULT = 0
	KEY_COMPUTED = 1
	KEY_GIVEN = 2
)

type Key struct {
	Name 	string
	Value  	bool
	State 	int
	rules	[]Defines
}

func (k Key) Eval(key string) (mybool bool, e error) {
	if k.Name == key {
		return k.Value, errors.New(fmt.Sprintf("Error: %s is self-referring.\n", key))
	}
	// val := k.Value
	// for i, rule := range k.rules {
	// 	fmt.Println("rule ", i, ": ", rule)
	// 	e = rule.Apply()
	// 	if e != nil {
	// 		return k.Value, e
	// 	}
	// }
	// if k.State != KEY_DEFAULT && val != k.Value {
	// 	e = errors.New(fmt.Sprintf("Error: %s was already supposed to be %t.\n", k.Name, k.Value))
	// }
	return k.Value, e
}

func (key Key) Set(val bool) (e error) {
	if key.State == KEY_DEFAULT {
		key.State = KEY_COMPUTED
		key.Value = val
	} else if key.State == KEY_GIVEN && !val {
		e = errors.New(fmt.Sprintf("Error: %s violates your statements.\n", key.Name))
	} else if key.State == KEY_COMPUTED && val != key.Value {
		e = errors.New(fmt.Sprintf("Error: %s was already calculated to be %t.\n", key.Name, key.Value))
	}
	return
}

func (key Key) String() string {
	val, _ := key.Eval(key.Name) 
	return fmt.Sprintf("{%s:%t}", key.Name, val)
}

