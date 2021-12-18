package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	sumPacketID     = 0
	productPacketID = 1
	minPacketID     = 2
	maxPacketID     = 3
	literalPacketID = 4
	gtPacketID      = 5
	ltPacketID      = 6
	eqPacketID      = 7
)

const literalBlockSize = 5

type header struct {
	version, typeID, totalLength, subPacketsTotalLen, numSubPackets uint64
}

type packet struct {
	header
	value   uint64
	inner   []packet
	bitsize int
}

func hexToBinary(input string) (result string) {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range decoded {
		result += fmt.Sprintf("%08b", b)
	}

	return
}

func binToDecimal(input string) (dec uint64) {
	dec, err := strconv.ParseUint(input, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func readPacket(input string) packet {
	pheader := header{
		version: binToDecimal(input[0:3]),
		typeID:  binToDecimal(input[3:6]),
	}

	if pheader.typeID == literalPacketID {
		// literal value
		var valueData string
		curIndex := 6

		for {
			block := input[curIndex : curIndex+literalBlockSize]
			valueData += block[1:]
			if block[0] == '0' {
				break
			}
			curIndex += literalBlockSize
		}

		return packet{
			header:  pheader,
			value:   binToDecimal(valueData),
			bitsize: curIndex + literalBlockSize,
		}
	}

	pheader.totalLength = binToDecimal(input[6:7])

	var bitsize int
	var inner []packet
	if pheader.totalLength == 0 {
		// the next 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet
		pheader.subPacketsTotalLen = binToDecimal(input[7 : 7+15])
		subpackets := input[7+15 : 7+15+pheader.subPacketsTotalLen]
		bitsize = int(7 + 15 + pheader.subPacketsTotalLen)

		for len(subpackets) > 0 {
			newSubPacket := readPacket(subpackets)
			subpackets = subpackets[newSubPacket.bitsize:]
			inner = append(inner, newSubPacket)

		}
	} else {
		// the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet
		pheader.numSubPackets = binToDecimal(input[7 : 7+11])

		curIndex := 7 + 11
		for i := 0; i < int(pheader.numSubPackets); i++ {
			newSubPacket := readPacket(input[curIndex:])
			curIndex += newSubPacket.bitsize
			inner = append(inner, newSubPacket)
		}

		bitsize = curIndex
	}

	return packet{
		header:  pheader,
		bitsize: bitsize,
		inner:   inner,
	}
}

func evaluate(p packet) int {

	switch p.typeID {
	case sumPacketID:
		result := 0
		for _, sp := range p.inner {
			result += evaluate(sp)
		}
		return result
	case productPacketID:
		result := 1
		for _, sp := range p.inner {
			result *= evaluate(sp)
		}
		return result
	case minPacketID:
		result := -1
		for _, sp := range p.inner {
			value := evaluate(sp)
			if result == -1 || value < result {
				result = value
			}
		}
		return result
	case maxPacketID:
		result := -1
		for _, sp := range p.inner {
			value := evaluate(sp)
			if result == -1 || value > result {
				result = value
			}
		}
		return result
	case literalPacketID:
		return int(p.value)
	case gtPacketID:
		if evaluate(p.inner[0]) > evaluate(p.inner[1]) {
			return 1
		}
		return 0
	case ltPacketID:
		if evaluate(p.inner[0]) < evaluate(p.inner[1]) {
			return 1
		}
		return 0
	case eqPacketID:
		if evaluate(p.inner[0]) == evaluate(p.inner[1]) {
			return 1
		}
		return 0
	}

	panic(fmt.Sprintf("Unknown type id %v", p.typeID))
}

func doIt(input io.Reader) int {
	packetData, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}
	packet := readPacket(hexToBinary(strings.TrimRight(string(packetData), "\r\n")))

	return evaluate(packet)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
