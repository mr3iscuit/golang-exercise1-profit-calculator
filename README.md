# Profit Calculator CLI

## Description
This CLI application is written using the Cobra library to provide a more engaging and professional command-line experience. As someone experienced in Go and many other programming languages, I find basic input/output with `fmt.Scan` and `fmt.Println` to be repetitive and not interesting. To make this project more enjoyable and educational, I have used Cobra for command parsing and help messages, and used some design patterns such as Singleton for state management.

A command-line tool to calculate profit, Earnings Before Tax (EBT), and profit ratios based on your revenue, tax rate, and expenses. Built with Go and Cobra.

## Features
- Calculate net profit after tax and expenses
- Calculate Earnings Before Tax (EBT)
- Compute profit ratios
- Easy-to-use command-line interface

## Installation & Build

Clone the repository and build the binary:

```sh
git clone https://github.com/mr3iscuit/golang-exercise1-profit-calculator.git
cd golang-exercise1-profit-calculator
go build -o profit-calc
```

## Usage

Run the CLI with the desired command and flags:

```sh
./profit-calc [command] [flags]
```

### Global Flags
- `--revenue, -r`   Total revenue (income) amount (required)
- `--tax-rate, -t`  Tax rate as a percentage (e.g., 15 for 15%) (required)
- `--expenses, -e`  Total expenses amount (required)

### Commands

#### Calculate Net Profit
```sh
./profit-calc profit --revenue 10000 --tax-rate 15 --expenses 3000
```

#### Calculate Earnings Before Tax (EBT)
```sh
./profit-calc ebt --revenue 10000 --tax-rate 15 --expenses 3000
```

#### Calculate Profit Ratio
```sh
./profit-calc ratio --revenue 10000 --tax-rate 15 --expenses 3000
```

## Example Output
```
Earnings Before Tax: 5950.000000
```

## Help
For more information on commands and flags:
```sh
./profit-calc --help
./profit-calc profit --help
./profit-calc ebt --help
./profit-calc ratio --help
```

## License
MIT 
