package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

type PageData struct {
	Regions       interface{}
	Region        string
	Cities        interface{}
	City          string
	RegionRedRoll int
	RegionRoll    int
	CityRedRoll   int
	CityRoll      int
}

func main() {
	templates := template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		region, regionRedRoll, regionRoll := getRegion()
		city, cityRedRoll, cityRoll := getCity(region)

		data := PageData{
			Regions:       Regions,
			Region:        region,
			Cities:        Cities,
			City:          city,
			RegionRedRoll: regionRedRoll,
			RegionRoll:    regionRoll,
			CityRedRoll:   cityRedRoll,
			CityRoll:      cityRoll,
		}

		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/destination", func(w http.ResponseWriter, r *http.Request) {
		region := r.URL.Query().Get("region")
		var regionRedRoll int
		var regionRoll int
		if region == "" {
			region, regionRedRoll, regionRoll = getRegion()
		}
		city, cityRedRoll, cityRoll := getCity(region)

		data := PageData{
			Region:        region,
			City:          city,
			RegionRedRoll: regionRedRoll,
			RegionRoll:    regionRoll,
			CityRedRoll:   cityRedRoll,
			CityRoll:      cityRoll,
		}

		if err := templates.ExecuteTemplate(w, "region-city.html", data); err != nil {
			panic(err)
		}
	})

	log.Fatal(http.ListenAndServe(":10101", nil))
}

func getCity(region string) (city string, redRoll int, roll int) {

	oddOrEven, redRoll := rollRedDie()
	roll = rollDice(2)
	city = Cities[region][oddOrEven][roll]

	return
}

func getRegion() (region string, redRoll int, roll int) {

	oddOrEven, redRoll := rollRedDie()
	roll = rollDice(2)
	region = Regions[oddOrEven][roll]

	return
}

func rollRedDie() (string, int) {
	roll := rollDice(1)
	if roll%2 == 0 {
		return "Even", roll
	} else {
		return "Odd", roll
	}
}

func rollDice(quantity int) (total int) {

	for i := 0; i < quantity; i++ {
		total += rand.Intn(6) + 1
	}

	return
}
