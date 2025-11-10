package main

import (
	client "AutoShop/Client"
	availability "AutoShop/availability"
	"fmt"
)

func main() {

	orders := make(client.Ordr)
	availabilitys := make(availability.Availab)

	fmt.Println("Добро пожаловать в магазин автозапчастей!")
	fmt.Println("Выберите категорию товара:", "1)Headlights", "2)Motors", "3)Transmission")

	client.EnterOrder(orders, availabilitys)
}
