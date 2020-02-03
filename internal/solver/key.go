package solver

import (
	"errors"
	"fmt"
    "github.com/fatih/color"
)

type Key struct {
	Name 	string
	Child	Node
}

var boldBlack *color.Color = color.New(color.Bold, color.FgBlack)
var boldRed *color.Color = color.New(color.Bold, color.FgRed)

func (k *Key) Eval(keys []string) (mybool bool, e error) {
	for _, item := range keys {
        if item == k.Name {
            return false, errors.New(
                fmt.Sprintf("%s Variable %s is self-referring while solving for %s.",
                boldRed.Sprint("SolvingError:"),
                boldBlack.Sprintf("'%s'", item),
                boldBlack.Sprintf("'%s'", keys[0])))
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

