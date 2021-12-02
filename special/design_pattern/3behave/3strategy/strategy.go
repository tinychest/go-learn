package _strategy

import "fmt"

type IAttack interface {
	Attack()
}

type swordAttack struct{}
type gunAttack struct{}
type stickAttack struct{}

func (a *swordAttack) Attack() {
	fmt.Println("slash")
}

func (a *gunAttack) Attack() {
	fmt.Println("shot")
}

func (a *stickAttack) Attack() {
	fmt.Println("stab")
}

var strategies = map[string]IAttack{
	"sword": new(swordAttack),
	"gun":   new(gunAttack),
	"stick": new(stickAttack),
}

func NewAttack(method string) (IAttack, error) {
	v, ok := strategies[method]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", method)
	}
	return v, nil
}
