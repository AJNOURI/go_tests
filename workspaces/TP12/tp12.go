package main

import (
	"fmt"
)

type gender string

const (
	m gender = "male"
	f gender = "female"
)

type work string

const (
	developer work = "developer"
	sysadmin  work = "sysadmin"
	manager   work = "manager"
)

type currency string

const (
	eur currency = "EUR"
	usd currency = "USD"
)

type address struct {
	city    string
	country string
}

type company struct {
	name         string
	headquarters address
}

var znk = company{"Zenika", address{"Paris", "France"}}
var msf = company{"Microsoft", address{"Redmond", "United States"}}
var apl = company{"Apple", address{"Cupertino", "United States"}}
var fbk = company{"Facebook", address{"Menlo Park", "United States"}}
var amz = company{"Amazon", address{"Seattle", "United States"}}

type salary struct {
	amount   float64
	currency currency
}

type person struct {
	name   string
	sex    gender
	office address
}

type certification string

const (
	csm     certification = "Certified Scrum Master"
	spring  certification = "Spring Professional"
	mongodb certification = "MongoDB Professional"
	java8   certification = "Java SE 8 Programmer OCE"
	pmp     certification = "PMI Professional Project Manager"
	cissp   certification = "Certified Information Systems Security Professional"
)

type worker struct {
	person
	company        company
	role           work
	monthlyPay     salary
	certifications []certification
}

var workers = []worker{
	{
		person{"Marcel", m, address{"France", "Nantes"}},
		znk,
		developer,
		salary{424242.42, eur},
		[]certification{java8, csm},
	},
	{
		person{"Mark", m, address{"Menlo Park", "United States"}},
		fbk,
		developer,
		salary{424242.42, usd},
		[]certification{pmp, java8, cissp},
	},
}

// To complete : create certificationsNum() method

// To complete : create isOfficeCompanyHeadquarters() method

func main() {
	for _, _ = range workers {
		var name string    // To complete
		var certsNum int   // To complete
		var worksAtHq bool // To complete
		var company string // To complete
		fmt.Printf("%s has %d certifications", name, certsNum)
		if worksAtHq {
			fmt.Printf(" and his/her office is at %s headquarters\n", company)
		} else {
			fmt.Printf(" and his/her office is far from %s headquarters\n", company)
		}
	}
}
