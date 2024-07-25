package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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
func solveVRP(loads []Load) map[int][]string {
	depot := Point{0, 0}
	maxDriveTime := 720.0 // 12 hours in minutes
	drivers := make(map[int][]string)
	driverID := 0
	sort.Slice(loads, func(i, j int) bool {
		return depot.distanceTo(loads[i].pickup)+loads[i].pickup.distanceTo(loads[i].dropoff) < depot.distanceTo(loads[j].pickup)+loads[j].pickup.distanceTo(loads[j].dropoff)
	})
	currentDriverTime := 0.0
	var currentDriverLoads []string
	for _, load := range loads {
		tripTime := depot.distanceTo(load.pickup) + load.pickup.distanceTo(load.dropoff) + load.dropoff.distanceTo(depot)
		if currentDriverTime+tripTime <= maxDriveTime {
			currentDriverLoads = append(currentDriverLoads, load.id)
			currentDriverTime += tripTime
		} else {
			drivers[driverID] = currentDriverLoads
			driverID++
			currentDriverLoads = []string{load.id}
			currentDriverTime = tripTime
		}
	}
	if len(currentDriverLoads) > 0 {
		drivers[driverID] = currentDriverLoads
	}
	return drivers
}
