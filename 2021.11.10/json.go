package main

//json编码解码
import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"titile`
	Year   int      `json:"year"`
	Price  float32  `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	godfather := Movie{
		"godfather",
		1975,
		12.3,
		[]string{"Al pacino", "Malon Brando"},
	}

	jsonstr, err := json.Marshal(godfather)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("jsonStr is : %s\n", jsonstr)

	mGodFather := Movie{}
	err = json.Unmarshal(jsonstr, &mGodFather)
	if err != nil {
		return
	}
	fmt.Printf("struct is : %v\n", mGodFather)
}
