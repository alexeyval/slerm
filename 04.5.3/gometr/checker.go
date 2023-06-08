package gometr

import (
	"fmt"
	"strings"
)

type Checker struct {
	items []Checkable
}

func (ch *Checker) Add(items []Checkable) {
	fmt.Println(&Checker{
		items: items,
	})
	ch.items = append(ch.items, items...)
}

func (ch *Checker) Check() {
	for _, v := range ch.items {
		if !v.Health() {
			fmt.Println(v.GetID(), "не работает")
		}
	}
}

func (ch *Checker) String() string {
	var IDs []string
	for _, v := range ch.items {
		IDs = append(IDs, v.GetID())
	}
	return strings.Join(IDs, " ")
}

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}
