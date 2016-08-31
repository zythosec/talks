package main

import "fmt"

// ONE OMIT
type MicroPanda struct {
	Name string
}

func (m MicroPanda) Speak() string {
	return fmt.Sprintf("Woof! I'm %s", m.Name)
}

type SecuriCat struct {
	Name string
}

func (s SecuriCat) Speak() string {
	return fmt.Sprintf("Meow! I'm %s", s.Name)
}

type Speaker interface {
	Speak() string
}

func MakeSpeak(s Speaker) string {
	return s.Speak()
}

// ONE OMIT

// START OMIT
func main() {
	lola := &MicroPanda{Name: "Lola"}
	cinnamon := &SecuriCat{Name: "Cinnamon"}
	nutmeg := &SecuriCat{Name: "Nutmeg"}

	friends := []Speaker{lola, cinnamon, nutmeg}
	for _, friend := range friends {
		fmt.Println(MakeSpeak(friend))
	}
}

// END OMIT
