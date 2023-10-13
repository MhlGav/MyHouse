package main

import "fmt"

type Room struct {
	Name      string
	Area      float64
	Height    float64
	Furniture []string
}

type Apartment struct {
	Area   float64
	Height float64
	Rooms  []Room
}

type FamilyMember struct {
	Name string
}

type Family struct {
	Members []FamilyMember
	Pets    []string
}

func main() {
	bedroom := Room{
		Name:      "Спальня",
		Area:      15.0,
		Height:    3.0,
		Furniture: []string{"шкаф-купе", "комод"},
	}

	livingRoom := Room{
		Name:      "Гостиная",
		Area:      20.0,
		Height:    3.0,
		Furniture: []string{"диван", "телевизор"},
	}

	kitchen := Room{
		Name:      "Кухня",
		Area:      10.0,
		Height:    3.0,
		Furniture: []string{"кухонный гарнитур", "микроволновая печь", "газовая плита"},
	}

	apartment := Apartment{
		Area:   46.0,
		Height: 3.0,
		Rooms:  []Room{bedroom, livingRoom, kitchen},
	}

	familyMembers := []FamilyMember{
		{Name: "Папа"},
		{Name: "Мама"},
		{Name: "Я"},
		{Name: "Сестра"},
	}

	pets := []string{"Йоркширский терьер"}

	family := Family{
		Members: familyMembers,
		Pets:    pets,
	}

	describeApartment(apartment, family)
}

func describeApartment(apartment Apartment, family Family) {
	fmt.Println("Описание квартиры:")
	fmt.Printf("- Площадь: %.1f квадратных метров\n", apartment.Area)
	fmt.Printf("- Высота потолков: %.1f метров\n", apartment.Height)
	fmt.Printf("- Количество комнат: %d\n", len(apartment.Rooms))

	fmt.Println("\nМебель в комнатах:")
	for _, room := range apartment.Rooms {
		fmt.Printf("- %s:\n", room.Name)
		for _, furniture := range room.Furniture {
			fmt.Printf("  - %s\n", furniture)
		}
	}

	fmt.Println("\nСемья:")
	for _, member := range family.Members {
		fmt.Println("- ", member.Name)
	}

	fmt.Println("\nДомашний питомец:")
	for _, pet := range family.Pets {
		fmt.Println("- ", pet)
	}
}
