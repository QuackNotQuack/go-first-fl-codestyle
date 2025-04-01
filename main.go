package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func attack(charName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(3, 5))
	}

	if charClass == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(5, 10))
	}

	if charClass == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func defence(char_name, char_class string) string {
	var blocked int
	switch char_class {
	case "warrior":
		blocked = 10 + randint(5, 10)
	case "mage":
		blocked = 10 + randint(-2, 2)
	case "healer":
		blocked = 10 + randint(2, 5)
	default:
		return "неизвестный класс персонажа"
	}
	return fmt.Sprintf("%s блокировал %d урона.", char_name, blocked)
}

// обратите внимание на "if else" и на "else"
func special(charName, charClass string) string {
	specials := map[string]func(string) string{
		"warrior": func(name string) string {
			return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", name, 80+25)
		},
		"mage": func(name string) string {
			return fmt.Sprintf("%s применил специальное умение `Атака %d`", name, 5+40)
		},
		"healer": func(name string) string {
			return fmt.Sprintf("%s применил специальное умение `Защита %d`", name, 10+30)
		},
	}

	if f, ok := specials[charClass]; ok {
		return f(charName)
	}
	return "неизвестный класс персонажа"
}

// здесь обратите внимание на имена параметров
func startTraining(char_name, char_class string) string {
	if char_class == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", char_name)
	}

	if char_class == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", char_name)
	}

	if char_class == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", char_name)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		if cmd == "attack" {
			fmt.Println(attack(char_name, char_class))
		}

		if cmd == "defence" {
			fmt.Println(defence(char_name, char_class))
		}

		if cmd == "special" {
			fmt.Println(special(char_name, char_class))
		}
	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func choiseCharClass() string {
	var approveChoice string
	var charClass string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)
		if charClass == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if charClass == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if charClass == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return charClass
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var char_name string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &char_name)

	fmt.Printf("Здравствуй, %s\n", char_name)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	char_class := choiseCharClass()

	fmt.Println(startTraining(char_name, char_class))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
