# Profit Calculator CLI

A command-line tool to calculate profit, Earnings Before Tax (EBT), and profit ratios based on your revenue, tax rate, and expenses. Built with Go and Cobra.

## Features
- Calculate net profit after tax and expenses
- Calculate Earnings Before Tax (EBT)
- Compute profit ratios
- Easy-to-use command-line interface

## Installation & Build

Clone the repository and build the binary:

```sh
git clone https://github.com/yourusername/profit-calc.git
cd profit-calc
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