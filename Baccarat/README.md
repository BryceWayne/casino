# Baccarat Simulation Game

This Go program simulates a series of Baccarat games, recording the results and calculating the win rate after a specified number of simulations. The program also implements a specific betting strategy to manage bets and determine when to hit or stand based on the rules of Baccarat.

## How to Play Baccarat

### Objective
The goal of Baccarat is to bet on the hand that you believe will have a total closest to 9. You can bet on the Player's hand, the Banker's hand, or a tie.

### Card Values
- Cards 2-9: Face value
- 10s and face cards (King, Queen, Jack): 0 points
- Aces: 1 point

### Gameplay
1. Two hands are dealt: the Player's hand and the Banker's hand.
2. Each hand starts with two cards.
3. If either hand totals 8 or 9, it is called a "natural," and no more cards are drawn.
4. If neither hand has a natural, additional cards may be drawn based on specific rules.

### Drawing Rules

In Baccarat, the rules for drawing a third card differ for the Player's and Banker's hands. Here's a detailed explanation:

#### Player's Hand
- The Player always draws a third card if the total of the first two cards is 5 or less.
- The Player stands (does not draw a third card) if the total is 6 or 7.

#### Banker's Hand
The rules for the Banker's third card are more complex and depend on both the Banker's current total and whether the Player has drawn a third card:

- If the Player stands (i.e., the Player's total is 6 or 7), the Banker follows the same rules as the Player:
  - The Banker draws a third card if the total is 5 or less.
  - The Banker stands if the total is 6 or 7.

- If the Player draws a third card, the Banker's action depends on the value of the Banker's first two cards and the value of the Player's third card:

  - **Banker total is 0, 1, or 2**: Banker always draws a third card.
  - **Banker total is 3**: Banker draws unless the Player's third card is an 8.
  - **Banker total is 4**: Banker draws if the Player's third card is 2, 3, 4, 5, 6, or 7.
  - **Banker total is 5**: Banker draws if the Player's third card is 4, 5, 6, or 7.
  - **Banker total is 6**: Banker draws if the Player's third card is 6 or 7.
  - **Banker total is 7**: Banker always stands.

These rules ensure that the drawing process is predetermined and does not involve any decision-making by the players during the game.

## Betting Strategy
**Source Video:** [Baccarat Strategy: How to Win at Baccarat with 99.7% Winrate](https://www.youtube.com/watch?v=g1JpoE2UyF8)

The program follows a specific betting strategy to manage bets and determine when to hit or stand:

1. **Initial Bet**: The initial bet value is set by the user.
2. **Bet Type**: The initial bet is placed on the Player's hand.
3. **Step**: A step counter is used to follow the betting strategy.

### Betting Sequence
1. Bet on Player.
2. If the bet is lost, double the bet and switch to Banker.
3. If the bet is lost again, double the bet and switch back to Player.
4. Continue this sequence, doubling the bet and following the specific sequence of Player and Banker bets until either a win occurs or the table limit is reached.

### Detailed Sequence
- **Step 1**: Bet on Player.
- **Step 2**: If the bet is lost, double the bet and switch to Banker.
- **Step 3**: If the bet is lost, double the bet and switch back to Player.
- **Step 4**: If the bet is lost, double the bet and stay on Player.
- **Step 5**: If the bet is lost, double the bet and switch to Banker.
- **Step 6**: If the bet is lost, double the bet and stay on Banker.
- **Step 7**: If the bet is lost, reset the bet to the initial bet value, switch back to Player, and reset the step counter to 1.

### Resetting
- If the bet is won at any step, the bet value is reset to the initial bet, the bet type is set to Player, and the step counter is reset to 1.
- The balance is updated based on the bet and the result. A 5% commission is deducted from Banker wins.

This strategy is designed to recover losses and achieve a profit equal to the initial bet after a series of losses, while also managing the risk by resetting the bet after a certain number of steps.

## Program Structure

### Card Struct
Represents a single playing card with a value and suit.

### Deck Struct
Represents a deck of playing cards. Supports shuffling and drawing cards.

### Hand Struct
Represents a hand of cards and calculates the hand value according to Baccarat rules.

### GameHistory Struct
Records the details of a single game, including player name, bet value, bet type, player and banker values, winner, and balance.

### Functions
- `NewShoe`: Initializes a new shoe of cards with multiple decks.
- `Shuffle`: Shuffles the deck of cards.
- `Draw`: Draws a card from the deck.
- `Value`: Calculates the value of a hand in Baccarat.
- `dealInitialHands`: Deals initial hands to the player and banker.
- `shouldDrawThirdCard`: Determines if a hand should draw a third card based on Baccarat rules.
- `dealThirdCard`: Deals third cards based on Baccarat rules.
- `determineWinner`: Determines the winner between the player and banker.
- `loadGameHistory`: Loads game history from a JSON file.
- `saveGameHistory`: Saves game history to a JSON file.
- `playGame`: Plays a single game and returns the result.
- `runSimulation`: Runs a single simulation.
- `main`: Runs simulations concurrently and reports the win rate.

## Running the Program
To run the program, use the following command-line arguments:

- `-name`: Player's name (default: "Player")
- `-bet`: Initial bet value (default: 100)
- `-balance`: Player's balance (default: 1000)
- `-simulations`: Number of simulations to run (default: 10000)
- `-tablelimit`: Table limit for betting (default: 1000)
- `-decks`: Number of decks in the shoe (default: 8)

---

Example:
```sh
go run main.go -name "Alice" -bet 50 -balance 2500 -simulations 500000 -tablelimit 2000 -decks 6
```

Output:
```
500000 / 500000 [-------------------------------------------------------------] 100.00% 962 p/s
Game history saved successfully.
Win rate after 500000 simulations: 64.28%
```