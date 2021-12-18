package main

import (
	"log"
	"os"
	"testing"
)

func Test_doIt(t *testing.T) {
	inputFile, err := os.Open("input_test")
	if err != nil {
		log.Fatal(err)
	}

	result := doIt(inputFile)
	if result != 315 {
		t.Errorf("doIt = %d; want 315", result)
	}
}

func Test_packet1(t *testing.T) {
	packet := readPacket(hexToBinary("D2FE28"))

	if packet.version != 6 {
		t.Errorf("doIt = %d; want 6", packet.version)
	}
	if packet.typeID != 4 {
		t.Errorf("doIt = %d; want 4", packet.typeID)
	}
	if packet.value != 2021 {
		t.Errorf("doIt = %d; want 2021", packet.value)
	}
}

func Test_packet2(t *testing.T) {
	packet := readPacket(hexToBinary("38006F45291200"))

	if packet.version != 1 {
		t.Errorf("doIt = %d; want 1", packet.version)
	}
	if packet.typeID != 6 {
		t.Errorf("doIt = %d; want 6", packet.typeID)
	}
	if packet.subPacketsTotalLen != 27 {
		t.Errorf("doIt = %d; want 27", packet.subPacketsTotalLen)
	}
}

func Test_packet3(t *testing.T) {
	packet := readPacket(hexToBinary("EE00D40C823060"))

	if packet.version != 7 {
		t.Errorf("doIt = %d; want 7", packet.version)
	}
	if packet.typeID != 3 {
		t.Errorf("doIt = %d; want 3", packet.typeID)
	}
	if packet.numSubPackets != 3 {
		t.Errorf("doIt = %d; want 3", packet.totalLength)
	}
}

func Test_packet4(t *testing.T) {
	packet := readPacket(hexToBinary("8A004A801A8002F478"))
	if versionSum(packet) != 16 {
		t.Errorf("doIt = %d; want 16", packet.version)
	}

	packet = readPacket(hexToBinary("620080001611562C8802118E34"))
	if versionSum(packet) != 12 {
		t.Errorf("doIt = %d; want 12", packet.version)
	}

	packet = readPacket(hexToBinary("C0015000016115A2E0802F182340"))
	if versionSum(packet) != 23 {
		t.Errorf("doIt = %d; want 23", packet.version)
	}

	packet = readPacket(hexToBinary("A0016C880162017C3686B18A3D4780"))
	if versionSum(packet) != 31 {
		t.Errorf("doIt = %d; want 31", packet.version)
	}
}
