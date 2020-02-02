package solver

import (
	"errors"
	"fmt"
)

type Key struct {
	Name 	string
	Child	Node
}

func (k *Key) Eval(keys []string) (mybool bool, e error) {
	for _, item := range keys {
        if item == k.Name {
            return false, errors.New(fmt.Sprintf("Error: %s is self-referring.\n", item))
        }
    }
    if k.Child == nil {
    	return false, e
    }
    mybool, e = k.Child.Eval(append(keys, k.Name))
    if (mybool) {
    	var op True
    	k.Child = &op
    }
    return
}

func (key Key) String() string {
	return fmt.Sprintf("{%s:%T}", key.Name, key.Child)
}

