package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
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

// GameHistory struct represents a record of a single game
type GameHistory struct {
	PlayerName  string `json:"player_name"`
	BetValue    int    `json:"bet_value"`
	BetType     string `json:"bet_type"`
	PlayerValue int    `json:"player_value"`
	BankerValue int    `json:"banker_value"`
	Winner      string `json:"winner"`
	Balance     int    `json:"balance"`
}

// Initialize a new shoe of cards with multiple decks
func NewShoe(numDecks int) *Deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

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

// Load game history from JSON file
func loadGameHistory(filePath string) ([]GameHistory, error) {
	var history []GameHistory
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return history, err
	}
	err = json.Unmarshal(data, &history)
	return history, err
}

// Save game history to JSON file
func saveGameHistory(filePath string, history []GameHistory) error {
	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

// Play a single game and return the result
func playGame(playerName string, betValue int, betType string, balance int, deck *Deck, numDecks int, houseEdge float64) (string, int, int, GameHistory) {
	// Shuffle the deck if it's close to being empty
	if len(deck.Cards) < 6 {
		deck = NewShoe(numDecks)
		deck.Shuffle()
	}

	// Deal initial hands
	playerHand, bankerHand := dealInitialHands(deck)

	// Deal third cards based on Baccarat rules
	dealThirdCard(deck, &playerHand, &bankerHand)

	// Determine the winner
	winner := determineWinner(playerHand, bankerHand)

	// Update balance based on the bet and the result
	if betType == winner {
		if betType == "Banker" {
			// Deduct 5% commission on Banker wins
			balance += int(float64(betValue) * houseEdge)
		} else {
			balance += betValue
		}
	} else {
		balance -= betValue
	}

	// Record the game result in history
	game := GameHistory{
		PlayerName:  playerName,
		BetValue:    betValue,
		BetType:     betType,
		PlayerValue: playerHand.Value(),
		BankerValue: bankerHand.Value(),
		Winner:      winner,
		Balance:     balance,
	}

	return winner, betValue, balance, game
}

// Run a single simulation
func runSimulation(playerName string, initialBet int, initialBalance int, tableLimit int, numDecks int, houseEdge float64, resultChan chan<- bool, historyChan chan<- []GameHistory, wg *sync.WaitGroup) {
	defer wg.Done()
	// Initialize variables
	balance := initialBalance
	betValue := initialBet
	betType := "Player"
	step := 1
	var gameHistories []GameHistory

	// Create and shuffle the initial shoe
	deck := NewShoe(numDecks)
	deck.Shuffle()

	// Play the game until we win $1,000 or lose all our money
	for balance > 0 && balance < initialBalance+1000 {
		// Check if betValue exceeds table limit
		if betValue > tableLimit {
			betValue = tableLimit
		}

		winner, _, newBalance, gameHistory := playGame(playerName, betValue, betType, balance, deck, numDecks, houseEdge)
		gameHistories = append(gameHistories, gameHistory)

		if winner == betType {
			// Won the bet, reset to step 1
			betValue = initialBet
			betType = "Player"
			step = 1
		} else {
			// Lost the bet, follow the strategy
			betValue *= 2
			switch step {
			case 1:
				betType = "Banker"
			case 2:
				betType = "Player"
			case 3:
				betType = "Player"
			case 4:
				betType = "Banker"
			case 5:
				betType = "Banker"
			case 6:
				betValue = initialBet
				betType = "Player"
				step = 0
			}
			step++
		}
		balance = newBalance
	}

	resultChan <- (balance >= initialBalance+1000)
	historyChan <- gameHistories
}

// Main function to run simulations and report win rate
func main() {
	// Define command-line arguments
	playerName := flag.String("name", "Player", "Player's name")
	initialBet := flag.Int("bet", 100, "Initial bet value")
	initialBalance := flag.Int("balance", 5000, "Player's balance (optional)")
	numSimulations := flag.Int("simulations", 10000, "Number of simulations to run")
	tableLimit := flag.Int("tablelimit", 1000, "Table limit for betting")
	numDecks := flag.Int("decks", 8, "Number of decks in the shoe")
	houseEdge := flag.Float64("houseEdge", 0.95, "Number of decks in the shoe")

	flag.Parse()

	// Run simulations concurrently
	resultChan := make(chan bool, *numSimulations)
	historyChan := make(chan []GameHistory, *numSimulations)
	var wg sync.WaitGroup
	bar := pb.StartNew(*numSimulations)

	for i := 0; i < *numSimulations; i++ {
		wg.Add(1)
		go func() {
			defer bar.Increment()
			runSimulation(*playerName, *initialBet, *initialBalance, *tableLimit, *numDecks, *houseEdge, resultChan, historyChan, &wg)
		}()
	}

	wg.Wait()
	close(resultChan)
	close(historyChan)
	bar.Finish()

	// Collect results
	winCount := 0
	var allGameHistories []GameHistory
	for result := range resultChan {
		if result {
			winCount++
		}
	}

	for histories := range historyChan {
		allGameHistories = append(allGameHistories, histories...)
	}

	// Save the complete game history to JSON file
	err := saveGameHistory("game_history.json", allGameHistories)
	if err != nil {
		fmt.Println("Error saving game history:", err)
	} else {
		fmt.Println("Game history saved successfully.")
	}

	// Report win rate
	winRate := float64(winCount) / float64(*numSimulations) * 100
	fmt.Printf("Win rate after %d simulations: %.2f%%\n", *numSimulations, winRate)
}
