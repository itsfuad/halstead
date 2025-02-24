# Halstead Metrics

This project provides a tool to analyze source code and calculate Halstead metrics. Halstead metrics are a set of software metrics introduced by Maurice Halstead to measure the complexity of a program.

## Features

- Supports multiple programming languages including C#, Java, JavaScript, Python, and Ruby.
- Calculates various Halstead metrics such as:
    - Number of distinct operators (N1)
    - Number of distinct operands (N2)
    - Total number of operators (n1)
    - Total number of operands (n2)
    - Program vocabulary (N)
    - Program length (n)
    - Calculated program length (Np)
    - Calculated program volume (V)
    - Calculated program difficulty (D)
    - Calculated program effort (E)
    - Calculated program time (T)
    - Calculated program bugs (B)

## Usage

To use the tool, run the `main.go` file with the `-filepath` flag pointing to the source code file you want to analyze. For example:

```sh
go run main.go -filepath path/to/sourcecode.ext
```

## Example

Here is an example of how to use the tool to analyze a Python file:

```sh
go run main.go -filepath Dataset/PalindromeChecker.py
```

## Requirements

- Go 1.23.2 or higher

## Installation

1. Clone the repository:
     ```sh
     git clone https://github.com/itsfuad/halstead.git
     ```
2. Navigate to the project directory:
     ```sh
     cd halstead-metrics
     ```
3. Run the tool:
     ```sh
     go run main.go -filepath path/to/sourcecode.ext
     ```

## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Acknowledgements

- Maurice Halstead for introducing the Halstead metrics.
- The open-source community for providing various tools and libraries.
