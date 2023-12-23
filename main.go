package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type RegionCity struct {
	Region string
	City   string
}

func main() {
	templates := template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		region := getRegion()
		city := getCity(region)

		data := RegionCity{
			Region: region,
			City:   city,
		}

		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/destination", func(w http.ResponseWriter, r *http.Request) {
		region := getRegion()
		city := getCity(region)

		data := RegionCity{
			Region: region,
			City:   city,
		}

		if err := templates.ExecuteTemplate(w, "region-city.html", data); err != nil {
			panic(err)
		}
	})

	log.Fatal(http.ListenAndServe(":10101", nil))
}

var cities = map[string]map[string]map[int]string{
	"NorthEast": {
		"odd": {
			2:  "New York",
			3:  "New York",
			4:  "New York",
			5:  "Albany",
			6:  "Boston",
			7:  "Buffalo",
			8:  "Boston",
			9:  "Portland ME",
			10: "New York",
			11: "New York",
			12: "New York",
		},
		"even": {
			2:  "New York",
			3:  "Washington DC",
			4:  "Pittsburgh",
			5:  "Pittsburgh",
			6:  "Philadelphia",
			7:  "Washington DC",
			8:  "Philadelphia",
			9:  "Baltimore",
			10: "Baltimore",
			11: "Baltimore",
			12: "New York",
		},
	},
	"SouthEast": {
		"odd": {
			2:  "Charlotte",
			3:  "Charlotte",
			4:  "Chattanooga",
			5:  "Atlanta",
			6:  "Atlanta",
			7:  "Atlanta",
			8:  "Richmond",
			9:  "Knoxville",
			10: "Mobile",
			11: "Knoxville",
			12: "Mobile",
		},
		"even": {
			2:  "Norfolk",
			3:  "Norfolk",
			4:  "Norfolk",
			5:  "Charleston",
			6:  "Miami",
			7:  "Jacksonville",
			8:  "Miami",
			9:  "Tampa",
			10: "Tampa",
			11: "Mobile",
			12: "Norfolk",
		},
	},
	"NorthCentral": {
		"odd": {
			2:  "Cleveland",
			3:  "Cleveland",
			4:  "Cleveland",
			5:  "Cleveland",
			6:  "Detroit",
			7:  "Detroit",
			8:  "Indianapolis",
			9:  "Milwaukee",
			10: "Milwaukee",
			11: "Chicago",
			12: "Milwaukee",
		},
		"even": {
			2:  "Cincinnati",
			3:  "Chicago",
			4:  "Cincinnati",
			5:  "Cincinnati",
			6:  "Columbus",
			7:  "Chicago",
			8:  "Chicago",
			9:  "St. Louis",
			10: "St. Louis",
			11: "St. Louis",
			12: "Chicago",
		},
	},
	"SouthCentral": {
		"odd": {
			2:  "Memphis",
			3:  "Memphis",
			4:  "Memphis",
			5:  "Little Rock",
			6:  "New Orleans",
			7:  "Birmingham",
			8:  "Louisville",
			9:  "Nashville",
			10: "Nashville",
			11: "Louisville",
			12: "Memphis",
		},
		"even": {
			2:  "Shreveport",
			3:  "Shreveport",
			4:  "Dallas",
			5:  "New Orleans",
			6:  "Dallas",
			7:  "San Antonio",
			8:  "Houston",
			9:  "Houston",
			10: "Fort Worth",
			11: "Fort Worth",
			12: "Fort Worth",
		},
	},
	"Plains": {
		"odd": {
			2:  "Kansas City",
			3:  "Kansas City",
			4:  "Denver",
			5:  "Denver",
			6:  "Denver",
			7:  "Kansas City",
			8:  "Kansas City",
			9:  "Kansas City",
			10: "Pueblo",
			11: "Pueblo",
			12: "Oklahoma City",
		},
		"even": {
			2:  "Oklahoma City",
			3:  "St. Paul",
			4:  "Minneapolis",
			5:  "St. Paul",
			6:  "Minneapolis",
			7:  "Oklahoma City",
			8:  "Des Moines",
			9:  "Omaha",
			10: "Omaha",
			11: "Fargo",
			12: "Fargo",
		},
	},
	"NorthWest": {
		"odd": {
			2:  "Spokane",
			3:  "Spokane",
			4:  "Seattle",
			5:  "Seattle",
			6:  "Seattle",
			7:  "Seattle",
			8:  "Rapid City",
			9:  "Casper",
			10: "Billings",
			11: "Billings",
			12: "Spokane",
		},
		"even": {
			2:  "Spokane",
			3:  "Salt Lake City",
			4:  "Salt Lake City",
			5:  "Salt Lake City",
			6:  "Portland OR",
			7:  "Portland OR",
			8:  "Portland OR",
			9:  "Pocatello",
			10: "Butte",
			11: "Butte",
			12: "Portland OR",
		},
	},
	"SouthWest": {
		"odd": {
			2:  "San Diego",
			3:  "San Diego",
			4:  "Reno",
			5:  "San Diego",
			6:  "Sacramento",
			7:  "Las Vegas",
			8:  "Phoenix",
			9:  "El Paso",
			10: "Tucumcari",
			11: "Phoenix",
			12: "Phoenix",
		},
		"even": {
			2:  "Los Angeles",
			3:  "Oakland",
			4:  "Oakland",
			5:  "Oakland",
			6:  "Los Angeles",
			7:  "Los Angeles",
			8:  "Los Angeles",
			9:  "San Francisco",
			10: "San Francisco",
			11: "San Francisco",
			12: "San Francisco",
		},
	},
}

func getCity(region string) string {

	oddOrEven := rollRedDie()
	roll := rollDice(2)
	city := cities[region][oddOrEven][roll]

	return city
}

var regions = map[string]map[int]string{
	"odd": {
		2:  "Plains",
		3:  "SouthEast",
		4:  "SouthEast",
		5:  "SouthEast",
		6:  "NorthCentral",
		7:  "NorthCentral",
		8:  "NorthEast",
		9:  "NorthEast",
		10: "NorthEast",
		11: "NorthEast",
		12: "NorthEast",
	},
	"even": {
		2:  "SouthWest",
		3:  "SouthCentral",
		4:  "SouthCentral",
		5:  "SouthCentral",
		6:  "SouthWest",
		7:  "SouthWest",
		8:  "Plains",
		9:  "NorthWest",
		10: "NorthWest",
		11: "Plains",
		12: "NorthWest",
	},
}

func getRegion() string {

	oddOrEven := rollRedDie()
	roll := rollDice(2)
	region := regions[oddOrEven][roll]

	return region
}

func rollRedDie() string {
	roll := rollDice(1)
	if roll%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func rollDice(quantity int) int {
	total := 0

	for i := 0; i < quantity; i++ {
		rand.Seed(time.Now().UnixNano())
		total += rand.Intn(6) + 1
	}

	return total
}
