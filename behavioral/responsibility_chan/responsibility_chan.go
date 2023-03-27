package main

import "fmt"

type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

type PatientHandler interface {
	Execute(*Patient) error
	SetNext(PatientHandler) PatientHandler
	Do(*Patient) error
}

// RegistrationHandler is the handler of registration abstruct
type Next struct {
	nextHandler PatientHandler
}

func (n *Next) SetNext(h PatientHandler) PatientHandler {
	n.nextHandler = h
	return h
}

func (n *Next) Execute(p *Patient) (err error) {
	if n.nextHandler != nil {
		if err = n.nextHandler.Do(p); err != nil {
			return
		}

		return n.nextHandler.Execute(p)
	}

	return
}

// RegistrationHandler is the handler of registration 接待
type Reception struct {
	Next
}

func (c *Reception) Do(p *Patient) (err error) {
	if p.RegistrationDone {
		fmt.Println("Registration Done")
		return
	}

	fmt.Println("Register registration")
	p.RegistrationDone = true
	return
}

// DoctorCheckUpHandler is the handler of doctor check up 诊所
type Clinic struct {
	Next
}

func (c *Clinic) Do(p *Patient) (err error) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor check up Done")
		return
	}

	fmt.Println("Doctor check up")
	p.DoctorCheckUpDone = true
	return
}

// PaymentHandler is the handler of payment 收费
type Cashier struct {
	Next
}

func (c *Cashier) Do(p *Patient) (err error) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
		return
	}

	fmt.Println("Payment")
	p.PaymentDone = true
	return
}

// PharmacyHandler is the handler of pharmacy 药房
type Pharmacy struct {
	Next
}

func (c *Pharmacy) Do(p *Patient) (err error) {
	if p.MedicineDone {
		fmt.Println("Medicine Done")
		return
	}

	fmt.Println("Medicine")
	p.MedicineDone = true
	return
}

// StartHandler is the handler of start 开始
type StartHandler struct {
	Next
}

func (c *StartHandler) Do(p *Patient) (err error) {
	fmt.Println("Start")
	return
}

func main() {
	patient := &Patient{
		Name:              "John",
		RegistrationDone:  false,
		DoctorCheckUpDone: false,
		MedicineDone:      false,
		PaymentDone:       false,
	}

	patientHealthHandler := &StartHandler{}
	reception := &Reception{}
	clinic := &Clinic{}
	pharmacy := &Pharmacy{}
	cashier := &Cashier{}
	patientHealthHandler.SetNext(reception).SetNext(clinic).SetNext(pharmacy).SetNext(cashier)

	if err := patientHealthHandler.Execute(patient); err != nil {
		fmt.Println("Fail Error", err.Error())
		return
	}

	fmt.Println("Success")
}
