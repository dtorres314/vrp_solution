package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x float64
	y float64
}

func (p Point) distanceTo(other Point) float64 {
	return math.Sqrt((p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y))
}

type Load struct {
	id      string
	pickup  Point
	dropoff Point
}

func parsePoint(coord string) Point {
	coords := strings.Split(coord, ",")
	x, _ := strconv.ParseFloat(coords[0], 64)
	y, _ := strconv.ParseFloat(coords[1], 64)
	return Point{x, y}
}

func parseInput(filePath string) ([]Load, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var loads []Load
	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip header
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		id := parts[0]
		pickupCoords := strings.Trim(parts[1], "()")
		dropoffCoords := strings.Trim(parts[2], "()")
		pickup := parsePoint(pickupCoords)
		dropoff := parsePoint(dropoffCoords)
		loads = append(loads, Load{id, pickup, dropoff})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return loads, nil
}

func solveVRP(loads []Load) (map[int][]string, float64) {
	depot := Point{0, 0}
	maxDriveTime := 720.0 // 12 hours in minutes
	drivers := make(map[int][]string)
	driverID := 0
	sort.SliceStable(loads, func(i, j int) bool {
		distanceI := depot.distanceTo(loads[i].pickup) + loads[i].pickup.distanceTo(loads[i].dropoff)
		distanceJ := depot.distanceTo(loads[j].pickup) + loads[j].pickup.distanceTo(loads[j].dropoff)
		if distanceI == distanceJ {
			return loads[i].id < loads[j].id
		}
		return distanceI < distanceJ
	})
	currentDriverTime := 0.0
	totalDrivenMinutes := 0.0
	var currentDriverLoads []string
	for _, load := range loads {
		tripTime := depot.distanceTo(load.pickup) + load.pickup.distanceTo(load.dropoff) + load.dropoff.distanceTo(depot)
		if currentDriverTime+tripTime <= maxDriveTime {
			currentDriverLoads = append(currentDriverLoads, load.id)
			currentDriverTime += tripTime
		} else {
			drivers[driverID] = currentDriverLoads
			totalDrivenMinutes += currentDriverTime
			driverID++
			currentDriverLoads = []string{load.id}
			currentDriverTime = tripTime
		}
	}
	if len(currentDriverLoads) > 0 {
		drivers[driverID] = currentDriverLoads
		totalDrivenMinutes += currentDriverTime
	}
	totalCost := 500*float64(len(drivers)) + totalDrivenMinutes
	return drivers, totalCost
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go {path_to_problem_or_directory}")
		return
	}

	inputPath := os.Args[1]

	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		fmt.Printf("Error checking input path: %v\n", err)
		return
	}

	if fileInfo.IsDir() {
		// Handle directory case
		files, err := os.ReadDir(inputPath)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			return
		}

		var costs []float64
		var sumRunTime float64

		for _, entry := range files {
			if entry.IsDir() {
				continue
			}
			fmt.Println(entry.Name())
			fmt.Println("\trunning...")
			filePath := filepath.Join(inputPath, entry.Name())

			start := time.Now()
			loads, err := parseInput(filePath)
			if err != nil {
				fmt.Printf("Error reading input file: %v\n", err)
				continue
			}
			solution, totalCost := solveVRP(loads)
			elapsed := time.Since(start)

			fmt.Printf("\trun time: %s\n", elapsed)
			if elapsed.Seconds() > 30 {
				fmt.Println("\t\tRun time constraint of 30s exceeded! Please reduce program runtime!")
			}
			sumRunTime += elapsed.Seconds()

			fmt.Println("\tevaluating solution...")
			for _, driverLoads := range solution {
				fmt.Printf("[%s]\n", strings.Join(driverLoads, ","))
			}
			fmt.Printf("Total cost: %.2f\n", totalCost)

			costs = append(costs, totalCost)
		}

		if len(costs) > 0 {
			var meanCost float64
			for _, cost := range costs {
				meanCost += cost
			}
			meanCost /= float64(len(costs))
			fmt.Printf("mean cost: %.2f\n", meanCost)
			fmt.Printf("mean run time: %.2f ms\n", (sumRunTime*1000)/float64(len(costs)))
		}
	} else {
		// Handle file case
		fmt.Println(inputPath)
		fmt.Println("\trunning...")
		start := time.Now()
		loads, err := parseInput(inputPath)
		if err != nil {
			fmt.Printf("Error reading input file: %v\n", err)
			return
		}
		solution, totalCost := solveVRP(loads)
		elapsed := time.Since(start)

		fmt.Printf("\trun time: %s\n", elapsed)
		if elapsed.Seconds() > 30 {
			fmt.Println("\t\tRun time constraint of 30s exceeded! Please reduce program runtime!")
		}

		fmt.Println("\tevaluating solution...")
		for _, driverLoads := range solution {
			fmt.Printf("[%s]\n", strings.Join(driverLoads, ","))
		}
		fmt.Printf("Total cost: %.2f\n", totalCost)
	}
}
