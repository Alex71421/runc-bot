package main

import (
	"fmt"

	"github.com/Alex71421/runc-bot/internal/event"
)

func main() {
	speedRace, err := event.New("speedrace", "Соревнования скорость", "https://speed")
	if err != nil {
		fmt.Println(err)
		return
	}
	speedRace.DateText = "21-22 ноября"
	text1 := "Подробнее Регистрация"
	speedRace.Status = event.ParseStatus(text1)
	fmt.Println(speedRace.Status)
	fmt.Println(speedRace)

	moscowHalf, err := event.New("moscowhalf", "Московский полумарафон", "https://half")
	if err != nil {
		fmt.Println(err)
		return
	}
	moscowHalf.DateText = "25-26 апреля"
	text2 := "  РЕГИСТРАЦИЯ ПОЗДНЕЕ  "
	moscowHalf.Status = event.ParseStatus(text2)
	fmt.Println(moscowHalf.Status)
}
