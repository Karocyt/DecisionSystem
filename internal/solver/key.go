package solver

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

func (k *Key) Eval(key string) (mybool bool, e error) { // Never evaluate subtree
	fmt.Println("\tKey Eval", k.Name, key)				//  Might need an array of strings to check everything
	//defer fmt.Println("\tEnd Key Eval", k.Name, key)
	if k.Name == key {
		fmt.Println("BUG ?!")
		return k.Value, errors.New(fmt.Sprintf("Error: %s is self-referring.\n", key))
	}
	if len(k.rules) > 0 && k.State == KEY_DEFAULT { ///////////////////////// To change when one big op
		mybool, e = k.rules[0].Eval(key)
		if e != nil {
			return
		}
		e = k.Set(mybool)
		return
	}
	return k.Value, e
}

func (key *Key) Set(val bool) (e error) {
	fmt.Println("\tKey Set", key.Name, val)
	// defer fmt.Println("\tEnd Key Set", key.Name, val)
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
	return fmt.Sprintf("{%s:%t:%d} => %t", key.Name, key.Value, key.State, val)
}

