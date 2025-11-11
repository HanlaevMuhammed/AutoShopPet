package main

import (
	availability "AutoShop/Availability"
	client "AutoShop/Client"
	"fmt"
)

func main() {
	orders := make(client.Ordr)
	availabilitys := make(availability.Availab)

	fmt.Println("Добро пожаловать в магазин автозапчастей!")

	availabilitys.LoadFromTXT("TXT/headlights.txt", "Headlights")
	availabilitys.LoadFromTXT("TXT/motors.txt", "Motors")
	availabilitys.LoadFromTXT("TXT/transmission.txt", "Transmission")

	client.EnterOrder(orders, availabilitys)
	orders.PrintByCategory()
}
