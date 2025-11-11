package availability

import (
	structur "AutoShop/Structur"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Availab map[string]structur.AvailabilitySt

func (a Availab) LoadFromTXT(filename, category string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла", filename, ":", err)
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
			continue
		}

		name := strings.Join(parts[:len(parts)-1], " ")
		price, _ := strconv.ParseFloat(parts[len(parts)-1], 64)
		a[name] = structur.AvailabilitySt{Category: category, Name: name, Price: price}
	}
	fmt.Printf("Загружено %d позиций (%s)\n", len(a), category)
}
