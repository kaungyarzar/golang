package main

import "fmt"

type Drive interface {
	drive() string
}

type CarFactory interface {
	offroad() Drive
	city() Drive
}

func GetCarFactory(vendor string) CarFactory {
	if vendor == "toyota" {
		return &Toyota{}
	}
	if vendor == "honda" {
		return &Honda{}
	}
	return nil
}

type ToyotaOffroad struct {
}

func (v *ToyotaOffroad) drive() string {
	return "Toyota offroad is driving.."
}

type ToyotaCity struct {
}

func (v *ToyotaCity) drive() string {
	return "Toyota city is driving.."
}

type Toyota struct {
}

func (v *Toyota) offroad() Drive {
	return &ToyotaOffroad{}
}
func (v *Toyota) city() Drive {
	return &ToyotaCity{}
}

type HondaOffroad struct {
}

func (v *HondaOffroad) drive() string {
	return "Honda offroad is driving.."
}

type HondaCity struct {
}

func (v *HondaCity) drive() string {
	return "Honda city is driving.."
}

type Honda struct {
}

func (v *Honda) offroad() Drive {
	return &HondaOffroad{}
}
func (v *Honda) city() Drive {
	return &HondaCity{}
}

func main() {
	car := GetCarFactory("honda")
	fmt.Println(car.city().drive())
	fmt.Println(car.offroad().drive())
}
