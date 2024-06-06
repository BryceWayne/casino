# Roulette Simulation

This project simulates a Roulette game using a modified Martingale betting strategy. The simulation runs a specified number of times and reports the win and loss rates based on the given strategy.

## Game Rules

Roulette is a casino game where players bet on the outcome of a spinning wheel. The wheel has 38 slots, numbered 0-36 and 37 (representing 00). Players can place bets on various outcomes, such as specific numbers, ranges of numbers, or colors (red or black).

## Betting Strategy

This simulation uses a modified Martingale strategy for two different bets:

1. **Second 18 (numbers 19-36)**
2. **Third 18 (numbers 25-36)**

### Martingale Steps

For both bets, the following steps are followed:
- Step 0: Bet $25
- Step 1: Bet $50
- Step 2: Bet $150
- Step 3: Bet $450
- Step 4: Bet $850
- Step 5: Reset to step 0

On a win, the bet amount resets to Step 0. On a loss, the bet amount progresses to the next step. If the bet amount exceeds the available balance, the game quits.

## Program Usage

To run the simulation, use the following command:

```shell
go run main.go --balance <initial_balance> --profit <profit_goal> --simulations <number_of_simulations>
```

### Command-Line Arguments

- `--balance`: Initial balance (default: 10000)
- `--profit`: Profit goal to end the game (default: 1000)
- `--simulations`: Number of simulations to run (default: 1000000)

## Example

Here's an example of running the simulation with different profit goals and 100,000 simulations each:

```shell
go run main.go --profit 500 --balance 10000 --simulations 100000
100000 / 100000 [----------------------------------------------------------------------------------------------] 100.00% 984775 p/s
After 100000 simulations:
Win rate: 97.05% (± 16.91%)
Lose rate: 2.95%
Average number of spins: 19.48

go run main.go --profit 2500 --balance 10000 --simulations 100000
100000 / 100000 [----------------------------------------------------------------------------------------------] 100.00% 322533 p/s
After 100000 simulations:
Win rate: 93.16% (± 25.24%)
Lose rate: 6.84%
Average number of spins: 56.48

go run main.go --profit 5000 --balance 10000 --simulations 100000
100000 / 100000 [----------------------------------------------------------------------------------------------] 100.00% 156546 p/s
After 100000 simulations:
Win rate: 90.00% (± 30.01%)
Lose rate: 10.00%
Average number of spins: 105.74
```

### Example Results

| Simulation Count | Initial Balance | Profit Goal | Win Rate (± 1 Std Dev) | Lose Rate | Average Spins |
|------------------|-----------------|-------------|------------------------|-----------|---------------|
| 100,000          | $10,000         | $500        | 97.05% (± 16.91%)      | 2.95%     | 19.48         |
| 100,000          | $10,000         | $2,500      | 93.16% (± 25.24%)      | 6.84%     | 56.48         |
| 100,000          | $10,000         | $5,000      | 90.00% (± 30.01%)      | 10.00%    | 105.74        |

## Conclusion

This simulation provides insights into the effectiveness of a modified Martingale betting strategy in Roulette. By running multiple simulations, you can observe the win and loss rates, as well as the average number of spins needed to meet the goal, along with the standard deviation, allowing you to make informed decisions based on the outcomes.
```

This README includes the game rules, betting strategy, program usage instructions, and example results formatted into a table. The table clearly presents the simulation count, initial balance, profit goal, win rate with standard deviation, lose rate, and average spins for easy reference.