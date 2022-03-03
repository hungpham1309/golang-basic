package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// create a new type of 'deck'
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFile(filename string) (deck, error) {
	byteData, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	data := string(byteData)
	s := strings.Split(data, ",")
	return s, nil
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for i := range d {
		position := r.Intn(len(d) - 1)

		d[i], d[position] = d[position], d[i]
	}
}
