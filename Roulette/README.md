# Roulette Simulation

This project simulates a Roulette game using a modified Martingale betting strategy. The simulation runs a specified number of times and reports the win and loss rates based on the given strategy.

## Game Rules

Roulette is a casino game where players bet on the outcome of a spinning wheel. The wheel has 38 slots, numbered 0-36 and 37 (representing 00). Players can place bets on various outcomes, such as specific numbers, ranges of numbers, or colors (red or black).

## Betting Strategy

This simulation uses a modified Martingale strategy for two different bets:

1. **Second 12 (numbers 19-36)**
2. **Third 12 (numbers 25-36)**

### Martingale Steps

For both bets, the following steps are followed:
- Step 0: Bet $25
- Step 1: Bet $50
- Step 2: Bet $150
- Step 3: Bet $450
- Step 4: Bet $850
- Step 5: Bet $1,350
- Step 6: Reset to step 0

On a win, the bet amount resets to Step 0. On a loss, the bet amount progresses to the next step. If the bet amount exceeds the available balance, the game quits.

## Program Usage

To run the simulation, use the following command:

```shell
go run main.go --balance <initial_balance> --simulations <number_of_simulations>
```

### Command-Line Arguments

- `--balance`: Initial balance (default: 10000)
- `--simulations`: Number of simulations to run (default: 1000000)

## Example

Here's an example of running the simulation with an initial balance of $5000 and 10,000,000 simulations:

```shell
go run main.go --balance 5000 --simulations 10000000
10000000 / 10000000 [------------------------------------------------------------------------------------------] 100.00% 399633 p/s
After 10000000 simulations:
Win rate: 92.63%
Lose rate: 7.37%
```

### Example Results

| Simulation Count | Initial Balance | Win Rate | Lose Rate |
|------------------|-----------------|----------|-----------|
| 10,000,000       | $10,000         | 92.63%   | 7.37%     |

## Conclusion

This simulation provides insights into the effectiveness of a modified Martingale betting strategy in Roulette. By running multiple simulations, you can observe the win and loss rates and make informed decisions based on the outcomes.