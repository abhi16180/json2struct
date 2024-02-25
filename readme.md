# JSON to Struct Generator

This tool converts JSON data into Go struct definitions.

## Overview

The JSON to Struct Generator is a simple tool written in Go that generates Go struct definitions based on JSON input. It's useful for quickly creating Go struct types that match the structure of JSON data.

## Features

- Converts JSON data to Go struct definitions
- Handles nested structures, arrays, and complex JSON data
- Copies generated struct definitions to the clipboard

## Usage

Clone the repository and run the tool with the following command:

```bash
go run main.go <json-file path>
```

sample.json
```json
{
  "name": "John",
  "age": 30,
  "cars": {
    "car1": "Ford",
    "car2": "BMW",
    "car3": "Fiat",
    "custom": {
      "name": "custom",
      "age": 30,
      "is_student": true,
      "grades": [
        1,
        2,
        3
      ],
      "address": {
        "city": "New York",
        "zip_code": "10001",
        "coordinates": {
          "latitude": 40.7128,
          "longitude": 74.0060
        }
      },
      "friends": [
        {
          "city": "New York",
          "zip_code": "10001"
        }
      ]
    }
  }
}
```

sample output
```go
type Base struct {
	Cars Cars    `json:"cars"`
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}
type Cars struct {
	Car2   string `json:"car2"`
	Car3   string `json:"car3"`
	Custom Custom `json:"custom"`
	Car1   string `json:"car1"`
}
type Custom struct {
	Age       float64   `json:"age"`
	IsStudent bool      `json:"is_student"`
	Grades    []float64 `json:"grades"`
	Address   Address   `json:"address"`
	Friends   []Address `json:"friends"`
	Name      string    `json:"name"`
}
type Address struct {
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	Coordinates Friends `json:"coordinates"`
}
type Friends struct {
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
```

