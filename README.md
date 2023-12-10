
# CSV File Sorter

This script reads a file in `.csv` format and allows you to choose the criteria for sorting the data based on the inputs from the first line.

## How to Use

Run the script using the following command:

```bash
go run main.go ./source-file.csv ./destination-file.csv
```

Replace `./source-file.csv` with the name and path of the file to be read, and `./destination-file.csv` with the name and path for the newly generated sorted file based on the chosen criteria.

## Requirements

- Go (Golang) installed

## Usage Instructions

1. Ensure Go is properly installed on your system.
2. Run go mod tidy in your project directory to tidy up and ensure all dependencies are in order..
3. Run the script with the provided command format.
4. Choose the criteria for sorting based on the inputs from the first line of the CSV file.

## Example

```bash
go run main.go ./data.csv ./sorted_data.csv
```

## License

This project is licensed under the [MIT License](LICENSE).

