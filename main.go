package main

import (
	"lcdTest/display"
	"log"
)

func main() {

	channel := make(chan []string, 5)
	// Create a new lcd struct
	lcd, err := display.NewHD44780(channel, "0x27", 1)
	if err != nil {
		log.Println(err)
		return
	}

	// Show test messages to the LCD
	lcd.DisplayMessage("Hello", "World", "Test", "Sentence")
	lcd.DisplayMessage("Hello", "World", "Test", "Sentence", "Odd sentence")
	// Clear the display
	lcd.ClearDisplay()
}
