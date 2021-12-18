package main

import (
	"errors"
	"fmt"
	"log"
)

type ErrTop struct {
	msg string
	mid error
}

func (et ErrTop) Error() string {
	return et.msg
}

func (et ErrTop) Unwrap() error {
	return et.mid
}

type ErrMiddle struct {
	msg    string
	bottom error
}

func (em ErrMiddle) Error() string {
	return em.msg
}

func (em ErrMiddle) Unwrap() error {
	return em.bottom
}

type ErrBottom struct {
	msg string
}

func (eb ErrBottom) Error() string {
	return eb.msg
}

func main() {
	bottom := ErrBottom{"I'm the bottom error"}
	mid := ErrMiddle{"I'm the middle error", bottom}
	top := ErrTop{"I'm the top error", mid}

	if errors.Is(top, mid) { // top contain mid
		log.Println("We have a middle error")
	}

	if !errors.Is(mid, top) { //mid not contain top
		log.Println("This is not a top level error")
	}

	var newBot ErrBottom
	if errors.As(top, &newBot) {
		log.Println("Found bottom level error", newBot)
	}
	log.Printf("My error chain: \n%w\n", top)

	newBot = ErrBottom{"This is a new bottom error"}
	newErr := fmt.Errorf("This is another error: %w", newBot)

	if errors.Is(newErr, newBot) {
		log.Println("This has a bottom error in it")
	}
}
