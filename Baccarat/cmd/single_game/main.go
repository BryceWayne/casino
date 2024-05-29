package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card struct represents a single playing card
type Card struct {
	Value int
	Suit  string
}

// Deck struct represents a deck of playing cards
type Deck struct {
	Cards []Card
}

// Hand struct represents a hand of cards
type Hand struct {
	Cards []Card
}

// Initialize a new shoe of cards with multiple decks
func NewShoe(numDecks int) *Deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Ace is 1

	deck := Deck{}

	for i := 0; i < numDecks; i++ {
		for _, suit := range suits {
			for _, value := range values {
				deck.Cards = append(deck.Cards, Card{Value: value, Suit: suit})
			}
		}
	}

	return &deck
}

// Shuffle the deck of cards
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// Draw a card from the deck
func (d *Deck) Draw() Card {
	if len(d.Cards) == 0 {
		panic("no cards left in the deck")
	}
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

// Calculate the value of a hand in Baccarat
func (h *Hand) Value() int {
	total := 0
	for _, card := range h.Cards {
		if card.Value >= 10 {
			total += 0
		} else {
			total += card.Value
		}
	}
	return total % 10
}

// Deal initial hands to the player and banker
func dealInitialHands(deck *Deck) (Hand, Hand) {
	playerHand := Hand{Cards: []Card{deck.Draw(), deck.Draw()}}
	bankerHand := Hand{Cards: []Card{deck.Draw(), deck.Draw()}}
	return playerHand, bankerHand
}

// Determine if a hand should draw a third card based on Baccarat rules
func shouldDrawThirdCard(handValue int, otherHandValue int, isPlayer bool) bool {
	if handValue <= 5 && isPlayer {
		return true
	}
	if !isPlayer {
		if handValue <= 2 {
			return true
		}
		switch handValue {
		case 3:
			return otherHandValue != 8
		case 4:
			return otherHandValue >= 2 && otherHandValue <= 7
		case 5:
			return otherHandValue >= 4 && otherHandValue <= 7
		case 6:
			return otherHandValue == 6 || otherHandValue == 7
		}
	}
	return false
}

// Deal third card based on Baccarat rules
func dealThirdCard(deck *Deck, playerHand, bankerHand *Hand) {
	playerValue := playerHand.Value()
	bankerValue := bankerHand.Value()

	if shouldDrawThirdCard(playerValue, bankerValue, true) {
		playerHand.Cards = append(playerHand.Cards, deck.Draw())
		playerValue = playerHand.Value()
	}

	if shouldDrawThirdCard(bankerValue, playerValue, false) {
		bankerHand.Cards = append(bankerHand.Cards, deck.Draw())
	}
}

// Determine the winner
func determineWinner(playerHand, bankerHand Hand) string {
	playerValue := playerHand.Value()
	bankerValue := bankerHand.Value()

	if playerValue > bankerValue {
		return "Player"
	} else if bankerValue > playerValue {
		return "Banker"
	} else {
		return "Tie"
	}
}

func main() {
	// Initialize and shuffle the deck
	deck := NewShoe(8)
	deck.Shuffle()

	// Deal initial hands
	playerHand, bankerHand := dealInitialHands(deck)

	// Deal third cards based on Baccarat rules
	dealThirdCard(deck, &playerHand, &bankerHand)

	// Determine the winner
	winner := determineWinner(playerHand, bankerHand)

	// Print results
	fmt.Printf("Player's hand: %+v (value: %d)\n", playerHand.Cards, playerHand.Value())
	fmt.Printf("Banker's hand: %+v (value: %d)\n", bankerHand.Cards, bankerHand.Value())
	fmt.Printf("Winner: %s\n", winner)
}
