# VRP Solution

## Overview

The VRP Solution project provides a tool for solving Vehicle Routing Problems (VRP). The program reads input files defining VRP scenarios and generates schedules for drivers, optimizing the total cost based on defined criteria.

## Features

- Parses VRP problem files to extract load and location data.
- Generates schedules for drivers that minimize the total cost.
- Can be run from a precompiled executable or directly from Go source code.

## Prerequisites

- **Go**: Ensure Go is installed on your system to compile and run the Go code.

## How to Run

### Compile the Program

If you don't have the precompiled executable, you can compile the program using Go:

#### Build the Executable:

```
go build -o vrp_solution.exe main.go
```
#### Run the Executable:

```
.\vrp_solution.exe .\problem\problem1.txt
```
#### Run the Go Program Directly:

```
go run main.go ./problem/problem1.txt
```
