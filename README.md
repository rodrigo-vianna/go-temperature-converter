# Temperature Converter

This Go program converts temperatures between Kelvin and Celsius.

## Description

The program takes a temperature value as a command-line argument and converts it between Kelvin and Celsius. By default, it converts from Kelvin to Celsius. Use the `-k` flag to convert from Celsius to Kelvin.

## Formulas

- **Celsius to Kelvin**: \( K = C + 273.15 \)
- **Kelvin to Celsius**: \( C = K - 273.15 \)

## Usage

1. **Basic Usage**:
    - To convert from Kelvin to Celsius (default):
      ```sh
      go run main.go <temperature>
      ```
    - Example:
      ```sh
      go run main.go 300
      ```
      Output:
      ```
      Temperature: 26.85 C
      ```

2. **Convert from Celsius to Kelvin**:
    - Use the `-k` flag:
      ```sh
      go run main.go <temperature> -k
      ```
    - Example:
      ```sh
      go run main.go 26.85 -k
      ```
      Output:
      ```
      Temperature: 300.00 K
      ```

## Installation

To build the program, you need to have Go installed. Clone the repository and build the binary:

```sh
git clone <repository_url>
cd <repository_directory>
go build -o temp_converter main.go
```

## Running the Program

Run the built binary with the required arguments:

```sh
./temp_converter <temperature> [-k]
```

## Examples

- Convert 300 Kelvin to Celsius:
  ```sh
  ./temp_converter 300
  ```
  Output:
  ```
  Temperature: 26.85 C
  ```

- Convert 26.85 Celsius to Kelvin:
  ```sh
  ./temp_converter 26.85 -k
  ```
  Output:
  ```
  Temperature: 300.00 K
  ```

## Code Overview

The program consists of two main functions for temperature conversion and a `main` function to handle command-line arguments and perform the conversion.

- `convertKelvinToCelsius(k float64) float64`: Converts Kelvin to Celsius.
- `convertCelsiusToKelvin(c float64) float64`: Converts Celsius to Kelvin.
- `main()`: Handles command-line arguments, performs conversions, and prints the result.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
