package main

type user struct {
	Name  string
	Age   int
	Phone string
}

var Users = map[string]*user{
	"1": {
		Name:  "jingpc",
		Age:   28,
		Phone: "13800138000",
	},
	"2": {
		Name:  "lizy",
		Age:   26,
		Phone: "139666666",
	},
}
