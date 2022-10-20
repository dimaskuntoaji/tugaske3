package controller

import (
	"fmt"
	"math/rand"
	"encoding/json"
	"text/template"
	"net/http"
	"tugaske3/structs"
)

type Wind struct {
	Int int
}

type Water struct {
	Int int
}

func (c *Wind) StatusWind() string {
	var status string

	switch {
	case c.Int < 6:
		status = "aman"
	case c.Int >= 7 && c.Int <= 15:
		status = "siaga"
	case c.Int > 15:
		status = "bahaya"
	}

	return status
}

func (c *Water) StatusWater() string {
	var status string

	switch {
	case c.Int < 5:
		status = "aman"
	case c.Int >= 6 && c.Int <= 8:
		status = "siaga"
	case c.Int > 8:
		status = "bahaya"
	}

	return status
}

func UpdateWater(c chan int) {
	wa := rand.Intn(99) + 1
	fmt.Println("Water", wa)
	c <- wa
}

func UpdateWind(c chan int) {
	wi := rand.Intn(99) + 1
	fmt.Println("Wind", wi)
	c <- wi
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	Result := structs.Value{}
	var waterInt int
	var windInt int
	c := make(chan int)

	go UpdateWater(c)
	waterInt = <-c

	go UpdateWind(c)
	windInt = <-c

	dataWater := Water{Int: waterInt}
	dataWind := Wind{Int: windInt}

	statWater := dataWater.StatusWater()
	statWind := dataWind.StatusWind()

	Result = structs.Value{
		WaterValue:  waterInt,
		WindValue:   windInt,
		WaterStatus: statWater,
		WindStatus:  statWind,
	}

	_, err := json.Marshal(Result)
	if err != nil {
		panic(err)
	}

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("./template/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, Result)
		return
	}
	http.Error(w, "invalid method", http.StatusBadRequest)
}