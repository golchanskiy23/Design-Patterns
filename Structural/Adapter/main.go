package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type Round interface {
	getRadius() float64
}

// round interface
type SquarePeg struct {
	Width float64
}

// round interface

type RoundHole struct {
	Radius float64
}

// round interface

type RoundPeg struct {
	radius float64
}

func (p RoundPeg) getRadius() float64 {
	return p.radius
}

type SquarePegAdapter struct {
	squarePeg SquarePeg
}

func (p SquarePegAdapter) getRadius() float64 {
	return p.squarePeg.Width * math.Sqrt(2) / 2
}

func (r RoundHole) fits(peg Round) error {
	if r.Radius < peg.getRadius() {
		return errors.New("too small roundhole")
	}
	return nil
}

func main() {
	var w, r float64
	fmt.Scanf("%f %f\n", &w, &r)
	roundHole := RoundHole{Radius: r}
	squarePeg := SquarePeg{Width: w}
	peg := SquarePegAdapter{squarePeg}
	if err := roundHole.fits(peg); err != nil {
		log.Fatalf("error during pushing: %v", err)
	}
	log.Print("Connected!!!")
}
