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
	Child	Node /// might need a pointer ?
}

func (k *Key) Eval(keys []string) (mybool bool, e error) { // Never evaluate subtree
	//fmt.Println("\tKey Eval", k.Name, keys)				//  Might need an array of strings to check everything
	//defer fmt.Println("\tEnd Key Eval", k.Name, key)
	for _, item := range keys {
        if item == k.Name {
            return false, errors.New(fmt.Sprintf("Error: %s is self-referring.\n", item))
        }
    }
    mybool, e = k.Child.Eval(append(keys, k.Name))
    if e != nil {
    	return
    }
    e = k.Set(mybool)
    if (mybool) {
    	var op True
    	k.Child = &op
    } else {
    	var op False
    	k.Child = &op
    }
    return
}

func (key *Key) Set(val bool) (e error) {
	//fmt.Println("\tKey Set", key.Name, val)
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
	val := key.Value
	return fmt.Sprintf("{%s:%t:%d} => %t", key.Name, key.Value, key.State, val)
}

