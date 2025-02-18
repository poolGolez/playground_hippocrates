package main

import "fmt"

func main() {
	fmt.Println("Section 6: Lists and Slices")

	zodiacs := []string{"Aquarius", "Libra", "Pisces", "Leo", "Virgo", "Aries", "Taurus", "Scorpio", "Sagitarrius", "Gemini", "Cancer", "Capricorn"}

	strongZodiacs := zodiacs[:5]

	fmt.Printf("Zodiacs        > Length: %v, Capacity: %v\n", len(zodiacs), cap(zodiacs))
	fmt.Println(zodiacs)
	fmt.Printf("Strong Zodiacs > Length: %v, Capacity: %v\n", len(strongZodiacs), cap(strongZodiacs))
	fmt.Println(strongZodiacs)

	allZodiacs := strongZodiacs[:cap(strongZodiacs)]
	fmt.Printf("&zodiacs:%p   &allZodiacs: %p\n", &zodiacs, &allZodiacs)
}
