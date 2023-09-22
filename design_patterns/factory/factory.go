package main

import "fmt"

type Payment interface {
	GetCardID() string
	SetID(cardid string)
}

type VistaPayment struct {
	CardID string
}

type MasterPayment struct {
	CardID string
}

func (v *VistaPayment) GetCardID() string {
	return v.CardID
}

func (v *MasterPayment) GetCardID() string {
	return v.CardID
}

func (v *VistaPayment) SetID(cardid string) {
	v.CardID = cardid
}

func (v *MasterPayment) SetID(cardid string) {
	v.CardID = cardid
}

func GetPayment(paytype string) Payment {
	switch paytype {
	case "vista":
		return &VistaPayment{}
	case "master":
		return &MasterPayment{}
	default:
		panic("Not support")
	}
}

func main() {
	payment := GetPayment("master")
	payment.SetID("thisismaster")
	fmt.Println(payment.GetCardID())
}
