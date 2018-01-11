package pubgopgg

import (
	"fmt"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	fmt.Println("===== TESTING PLAYER =====")

	c, err := New()
	player, err := c.GetPlayer("json_","na","fpp","2018-01")
	if err != nil  {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", player)

	player, err = c.GetPlayer("Napora","na","fpp","2018-01")
	if err != nil  {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", player)

	player, err = c.GetPlayer("--------","na","fpp","2018-01")
	if err != nil  {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", player)

	player, err = c.GetPlayer("JSon_","EU","fpp","2018-01")
	if err != nil  {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", player)
}