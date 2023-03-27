package main

import "fmt"

type PS5 interface {
	StartGPUEngine()
	GetPrice() int64
}

type PS5WithCD struct {
}

func (p *PS5WithCD) StartGPUEngine() {
	fmt.Println("start engine")
}

func (p *PS5WithCD) GetPrice() int64 {
	return 10000
}

type PS5WithDigital struct {
}

func (p *PS5WithDigital) StartGPUEngine() {
	fmt.Println("start normal gpu engine")
}

func (p *PS5WithDigital) GetPrice() int64 {
	return 20000
}

type PS5WithPlus struct {
	pS5Machine PS5
}

func (p *PS5WithPlus) StartGPUEngine() {
	p.pS5Machine.StartGPUEngine()
	fmt.Println("start plus gpu engine")
}

func (p *PS5WithPlus) GetPrice() int64 {
	return p.pS5Machine.GetPrice() + 1000
}

type PS5WithTopicColor struct {
	pS5Machine PS5
}

func (p *PS5WithTopicColor) StartGPUEngine() {
	p.pS5Machine.StartGPUEngine()
	fmt.Println("start topic color gpu engine")
}

func (p *PS5WithTopicColor) GetPrice() int64 {
	return p.pS5Machine.GetPrice() + 5000
}

func main() {
	ps5 := &PS5WithCD{}
	ps5.StartGPUEngine()
	fmt.Println(ps5.GetPrice())

	ps5WithDigital := &PS5WithDigital{}
	ps5WithDigital.StartGPUEngine()
	fmt.Println(ps5WithDigital.GetPrice())

	ps5WithPlus := &PS5WithPlus{pS5Machine: ps5}
	ps5WithPlus.StartGPUEngine()
	fmt.Println(ps5WithPlus.GetPrice())

	ps5WithTopicColor := &PS5WithTopicColor{pS5Machine: ps5}
	ps5WithTopicColor.StartGPUEngine()
	fmt.Println(ps5WithTopicColor.GetPrice())
}
