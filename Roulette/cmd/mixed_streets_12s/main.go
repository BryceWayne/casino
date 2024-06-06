package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
)

// BetType represents the type of bet in Roulette
type BetType string

const (
	Second12     BetType = "Second12"
	Third12      BetType = "Third12"
	ThirdStreet  BetType = "ThirdStreet"
	FourthStreet BetType = "FourthStreet"
)

// Bet represents a single bet in the game
type Bet struct {
	Type      BetType
	BetAmount int
}

// Result represents the result of a single spin
type Result struct {
	Number    int
	Outcome   string
	Balance   int
	SpinCount int
}

// Simulate a single spin of the Roulette wheel
func spinWheel(european bool) int {
	if european {
		return rand.Intn(37) // 0-36
	} else {
		return rand.Intn(38) // 0-37 (0-36, 37 represents 00)
	}
}

// Determine the outcome of a bet based on the spin result
func determineOutcome(bet Bet, number int) (string, int) {
	outcome := "Lose"
	payout := 0

	switch bet.Type {
	case Second12:
		if number >= 19 && number <= 36 {
			outcome = "Win"
			payout = bet.BetAmount * 2
		}
	case Third12:
		if number >= 25 && number <= 36 {
			outcome = "Win"
			payout = bet.BetAmount * 2
		}
	case ThirdStreet:
		if number >= 7 && number <= 9 {
			outcome = "Win"
			payout = bet.BetAmount * 11
		}
	case FourthStreet:
		if number >= 10 && number <= 12 {
			outcome = "Win"
			payout = bet.BetAmount * 11
		}
	}

	return outcome, payout
}

// Run a single simulation and return the result
func runSimulation(european bool, initialBalance int, profitGoal int, wg *sync.WaitGroup, resultChan chan<- Result) {
	defer wg.Done()

	bet1 := Bet{Type: Second12, BetAmount: 100}
	bet2 := Bet{Type: Third12, BetAmount: 100}
	bet3 := Bet{Type: ThirdStreet, BetAmount: 25}
	bet4 := Bet{Type: FourthStreet, BetAmount: 25}

	balance := initialBalance
	spinCount := 0

	for balance > 0 && balance < initialBalance+profitGoal {

		// Quit the game if the bet amount exceeds the available balance
		if bet1.BetAmount > balance || bet2.BetAmount > balance || bet3.BetAmount > balance || bet4.BetAmount > balance {
			break
		}

		number := spinWheel(european)
		spinCount++

		outcome1, payout1 := determineOutcome(bet1, number)
		outcome2, payout2 := determineOutcome(bet2, number)
		outcome3, payout3 := determineOutcome(bet3, number)
		outcome4, payout4 := determineOutcome(bet4, number)

		if outcome1 == "Win" {
			balance += payout1
		} else {
			balance -= bet1.BetAmount
		}

		if outcome2 == "Win" {
			balance += payout2
		} else {
			balance -= bet2.BetAmount
		}

		if outcome3 == "Win" {
			balance += payout3
		} else {
			balance -= bet3.BetAmount
		}

		if outcome4 == "Win" {
			balance += payout4
		} else {
			balance -= bet4.BetAmount
		}
	}

	result := Result{
		Number:    -1, // Spin number is not relevant for the final result
		Outcome:   fmt.Sprintf("Final balance: %d", balance),
		Balance:   balance,
		SpinCount: spinCount,
	}

	resultChan <- result
}

// Calculate the standard deviation for a slice of integers
func calculateStandardDeviation(data []int, mean float64) float64 {
	var sumOfSquares float64
	for _, value := range data {
		sumOfSquares += math.Pow(float64(value)-mean, 2)
	}
	variance := sumOfSquares / float64(len(data))
	return math.Sqrt(variance)
}

// Main function to run simulations and report outcomes
func main() {
	// Define command-line arguments
	initialBalance := flag.Int("balance", 10_000, "Initial balance")
	numSimulations := flag.Int("simulations", 1_000_000, "Number of simulations to run")
	european := flag.Bool("european", false, "Use European wheel (single 0)")
	profitGoal := flag.Int("profit", 1_000, "Profit goal")

	flag.Parse()

	// Run simulations concurrently
	resultChan := make(chan Result, *numSimulations)
	var wg sync.WaitGroup
	bar := pb.StartNew(*numSimulations)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < *numSimulations; i++ {
		wg.Add(1)
		go func() {
			defer bar.Increment()
			runSimulation(*european, *initialBalance, *profitGoal, &wg, resultChan)
		}()
	}

	wg.Wait()
	close(resultChan)
	bar.Finish()

	// Collect and report results
	winCount := 0
	loseCount := 0
	totalSpins := 0
	spinCounts := make([]int, 0, *numSimulations)
	winResults := make([]int, 0, *numSimulations)

	for result := range resultChan {
		if result.Balance >= *initialBalance+1000 {
			winCount++
			winResults = append(winResults, 1)
		} else {
			loseCount++
			winResults = append(winResults, 0)
		}
		totalSpins += result.SpinCount
		spinCounts = append(spinCounts, result.SpinCount)
	}

	winRate := float64(winCount) / float64(*numSimulations) * 100
	loseRate := float64(loseCount) / float64(*numSimulations) * 100
	averageSpins := float64(totalSpins) / float64(*numSimulations)
	stdDevWinRate := calculateStandardDeviation(winResults, float64(winCount)/float64(*numSimulations)) * 100

	fmt.Printf("After %d simulations:\n", *numSimulations)
	fmt.Printf("Win rate: %.2f%% (± %.2f%%)\n", winRate, stdDevWinRate)
	fmt.Printf("Lose rate: %.2f%%\n", loseRate)
	fmt.Printf("Average number of spins: %.2f\n", averageSpins)
}
