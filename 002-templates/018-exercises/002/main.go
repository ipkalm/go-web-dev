package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	ZIP     uint32
}

type city struct {
	City   string
	Hotels []hotel
}

type region struct {
	Region string
	Cities []city
}

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	b := []region{
		{
			Region: "Northen CA",
			Cities: []city{
				{
					City: "San Jose",
					Hotels: []hotel{
						{
							Name:    "Marriot",
							Address: "301 S Market St, San Jose, CA",
							ZIP:     95113,
						},
						{
							Name:    "Holiday Inn",
							Address: "1350 N 1st St, San Jose, CA",
							ZIP:     95112,
						},
					},
				},
				{
					City: "San Francisco",
					Hotels: []hotel{
						{
							Name:    "Hyatt Centric Fisherman's Wharf",
							Address: "555 North Point St, San Francisco, CA",
							ZIP:     94133,
						},
						{
							Name:    "Beck's Motor Lodge",
							Address: "2222 Market St, San Francisco, CA",
							ZIP:     94114,
						},
						{
							Name:    "Geary Parkway Motel",
							Address: "4750 Geary Blvd, San Francisco, CA ",
							ZIP:     94118,
						},
					},
				},
			},
		},
		{
			Region: "Central CA",
			Cities: []city{
				{
					City: "Fresno",
					Hotels: []hotel{
						{
							Name:    "Hyatt Place Fresno",
							Address: "7333 N Fresno St, Fresno, CA",
							ZIP:     93720,
						},
						{
							Name:    "Courtyard by Marriott Fresno",
							Address: "140 E Shaw Ave, Fresno, CA",
							ZIP:     93710,
						},
					},
				},
				{
					City: "Visalia",
					Hotels: []hotel{
						{
							Name:    "Visalia Marriott at the Convention Center",
							Address: "300 S Court St, Visalia, CA",
							ZIP:     93291,
						},
						{
							Name:    "Residence Inn by Marriott Visalia",
							Address: "205 N Plaza Dr, Visalia, CA",
							ZIP:     93291,
						},
					},
				},
			},
		},
		{
			Region: "Southern CA",
			Cities: []city{
				{
					City: "San Diego",
					Hotels: []hotel{
						{
							Name:    "The El Cordova Hotel",
							Address: "1351 Orange Ave, Coronado, CA",
							ZIP:     92118,
						},
						{
							Name:    "Embassy Suites by Hilton San Diego Bay Downtown",
							Address: "601 Pacific Hwy, San Diego, CA ",
							ZIP:     92101,
						},
					},
				},
			},
		},
	}

	err := t.ExecuteTemplate(os.Stdout, "index.gohtml", b)
	if err != nil {
		log.Fatalln(err)
	}
}
