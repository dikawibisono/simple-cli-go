package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	id int
	name, address, job string
}

//create dummy data
var classmate = []Biodata{
	{id: 1, name: "Dika", address: "Karawang", job: "Backend Developer"},
	{id: 2, name: "Wibi", address: "Bandung", job: "Admin"},
	{id: 3, name: "Sono", address: "Jakarta", job: "Accountant"},
	{id: 4, name: "Diki", address: "Magelang", job: "Backend Developer"},
	{id: 5, name: "Wiba", address: "Purwakarta", job: "Design Grafis"},
	{id: 6, name: "Soni", address: "Bekasi", job: "UI/UX"},
	{id: 7, name: "Daka", address: "Bandung", job: "Software Engineer"},
	{id: 8, name: "Wobi", address: "Karawang", job: "Mobile Developer"},
	{id: 9, name: "Sina", address: "Jakarta", job: "Frontend Developer"},
}

func main(){
	noId, err := strconv.Atoi(os.Args[1])

	index := noId - 1

	if err != nil {
		fmt.Println("Warning: Please input integer")
	} else if index < len(classmate){
		fmt.Println("No Absen: ", classmate[index].id)
		fmt.Println("Nama: ", classmate[index].name)
		fmt.Println("Alamat: ", classmate[index].address)
		fmt.Println("Job: ", classmate[index].job)
	} else {
		fmt.Println("Data not found")
	}
}