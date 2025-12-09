package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Vec3 struct {
	x, y, z int
}

func main() {
	clusters, err := readInput("input.txt")
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}

	fmt.Println("Part1:", part1(clusters, 1000))
	fmt.Println("Part2:", part2(clusters))
}

func part1(clusters [][]Vec3, n int) int {
	var dists []int
	for n > 0 {
		min, minA, minB, minAP, minBP := minClusterDist(clusters)

		var newDist []int
		for _, d := range dists {
			if d <= min {
				n--
				continue
			}
			newDist = append(newDist, d)
		}
		if n <= 0 || minA == -1 || minB == -1 || minAP == -1 || minBP == -1 {
			break
		}
		dists = append(newDist, clustersPDist(clusters[minA], clusters[minB], minAP, minBP)...)
		clusters[minA] = append(clusters[minA], clusters[minB]...)
		clusters = append(clusters[:minB], clusters[minB+1:]...)
		n--
	}

	sort.Slice(clusters, func(i, j int) bool {
		return len(clusters[i]) > len(clusters[j])
	})

	product := 1
	limit := min(len(clusters), 3)
	for i := range limit {
		product *= len(clusters[i])
	}

	return product
}

func part2(clusters [][]Vec3) int {
	for {
		_, minA, minB, minAP, minBP := minClusterDist(clusters)
		if len(clusters) == 2 {
			return clusters[minA][minAP].x * clusters[minB][minBP].x
		}
		clusters[minA] = append(clusters[minA], clusters[minB]...)
		clusters = append(clusters[:minB], clusters[minB+1:]...)
	}
}

func minClusterDist(clusters [][]Vec3) (int, int, int, int, int) {
	minA, minB, minAP, minBP := -1, -1, -1, -1
	min := math.MaxInt
	for a := range len(clusters) {
		for b := a + 1; b < len(clusters); b++ {
			dist, ap, bp := clusterDist(clusters[a], clusters[b])
			if dist < min {
				minA, minB = a, b
				minAP, minBP = ap, bp
				min = dist
			}
		}
	}
	return min, minA, minB, minAP, minBP
}

func clustersPDist(c1, c2 []Vec3, a, b int) []int {
	var dist []int
	for i, p1 := range c1 {
		for j, p2 := range c2 {
			if a == i && b == j {
				continue
			}
			dist = append(dist, pointDist(p1, p2))
		}
	}
	return dist
}

func clusterDist(a, b []Vec3) (int, int, int) {
	min := math.MaxInt
	minA := 0
	minB := 0
	for a, ap := range a {
		for b, bp := range b {
			dist := pointDist(ap, bp)
			if dist < min {
				min = dist
				minA = a
				minB = b
			}
		}
	}
	return min, minA, minB
}

func pointDist(a, b Vec3) int {
	xs := a.x - b.x
	ys := a.y - b.y
	zs := a.z - b.z
	return xs*xs + ys*ys + zs*zs
}

func readInput(filename string) ([][]Vec3, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points [][]Vec3
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("expects 3D point coordinates")
		}
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		num3, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			return nil, fmt.Errorf("invalid point format")
		}

		var cluster []Vec3
		cluster = append(cluster, Vec3{num1, num2, num3})
		points = append(points, cluster)
	}
	return points, nil
}
