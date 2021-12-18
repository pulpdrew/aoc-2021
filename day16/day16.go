package day16

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var hexToBinary = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func Run() {
	hex, err := os.ReadFile("./day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert hex string to binary
	builder := strings.Builder{}
	for _, r := range string(hex) {
		builder.WriteString(hexToBinary[r])
	}
	binary := builder.String()

	// Part 1
	p, _ := parsePacket([]rune(binary))
	versionSum := sumOfVersions(p)
	fmt.Printf("Day 16 - Part 1: The sum of versions is %v\n", versionSum)

	// Part 1
	fmt.Printf("Day 16 - Part 1: The value of the packet is %v\n", valueOf(p))
}

type packet struct {
	version      int
	typeId       int
	literal      int
	lengthTypeId int
	children     []packet
}

func parsePacket(data []rune) (p packet, remaining []rune) {
	remaining = data

	version, _ := strconv.ParseInt(string(remaining[0:3]), 2, 0)
	remaining = remaining[3:]
	p.version = int(version)

	typeId, _ := strconv.ParseInt(string(remaining[0:3]), 2, 0)
	remaining = remaining[3:]
	p.typeId = int(typeId)

	if typeId == 4 {
		groupPrefix := '1'
		literalStringValue := strings.Builder{}
		for groupPrefix == '1' {
			groupPrefix = remaining[0]
			literalStringValue.WriteString(string(remaining[1:5]))
			remaining = remaining[5:]
		}

		literalIntValue, _ := strconv.ParseInt(literalStringValue.String(), 2, 0)
		p.literal = int(literalIntValue)

		return p, remaining
	}

	lengthTypeId, _ := strconv.ParseInt(string(remaining[0:1]), 2, 0)
	remaining = remaining[1:]
	p.lengthTypeId = int(lengthTypeId)

	if lengthTypeId == 0 {

		subPacketLength, _ := strconv.ParseInt(string(remaining[0:15]), 2, 0)
		remaining = remaining[15:]

		originalRemainingLength := len(remaining)
		for len(remaining) > originalRemainingLength-int(subPacketLength) {
			child, r := parsePacket(remaining)
			remaining = r
			p.children = append(p.children, child)
		}

	} else {

		subPacketCount, _ := strconv.ParseInt(string(remaining[0:11]), 2, 0)
		remaining = remaining[11:]

		for i := 0; i < int(subPacketCount); i++ {
			child, r := parsePacket(remaining)
			remaining = r
			p.children = append(p.children, child)
		}
	}

	return p, remaining
}

func sumOfVersions(p packet) (sum int) {
	sum += p.version
	for _, child := range p.children {
		sum += sumOfVersions(child)
	}

	return sum
}

func valueOf(p packet) (value int) {

	switch p.typeId {
	case 0: // Sum
		for _, child := range p.children {
			value += valueOf(child)
		}
	case 1: // Product
		value = 1
		for _, child := range p.children {
			value *= valueOf(child)
		}
	case 2: // Min
		value = math.MaxInt
		for _, child := range p.children {
			childValue := valueOf(child)
			if childValue < value {
				value = childValue
			}
		}
	case 3: // Max
		value = math.MinInt
		for _, child := range p.children {
			childValue := valueOf(child)
			if childValue > value {
				value = childValue
			}
		}
	case 4: // literal
		value = p.literal
	case 5: // Greater
		if valueOf(p.children[0]) > valueOf(p.children[1]) {
			value = 1
		} else {
			value = 0
		}
	case 6: // Less
		if valueOf(p.children[0]) < valueOf(p.children[1]) {
			value = 1
		} else {
			value = 0
		}
	case 7: // Equal
		if valueOf(p.children[0]) == valueOf(p.children[1]) {
			value = 1
		} else {
			value = 0
		}
	}

	return value
}
