package availability

import (
	structur "AutoShop/Structur"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Availab map[string]structur.AvailabilitySt

func (a Availab) PrintMotors() {
	data, err := os.ReadFile("TXT/motors.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла motors.txt:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			fmt.Println("Пропущена цена или название")
			continue
		}

		name := strings.Join(parts[:len(parts)-1], " ")
		priceStr := parts[len(parts)-1]
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			fmt.Printf("Ошибка в строке", line, ":", err)
			continue
		}

		a[name] = structur.AvailabilitySt{Category: "Motors", Name: name, Price: price}
	}

	fmt.Println("Загружено %d позиций моторов\n", len(a))

}
