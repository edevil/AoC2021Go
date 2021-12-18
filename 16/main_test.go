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
	packet := readPacket(hexToBinary("C200B40A82"))
	if evaluate(packet) != 3 {
		t.Errorf("doIt = %d; want 3", packet.version)
	}

	packet = readPacket(hexToBinary("04005AC33890"))
	if evaluate(packet) != 54 {
		t.Errorf("doIt = %d; want 54", packet.version)
	}

	packet = readPacket(hexToBinary("880086C3E88112"))
	if evaluate(packet) != 7 {
		t.Errorf("doIt = %d; want 7", packet.version)
	}

	packet = readPacket(hexToBinary("CE00C43D881120"))
	if evaluate(packet) != 9 {
		t.Errorf("doIt = %d; want 9", packet.version)
	}

	packet = readPacket(hexToBinary("D8005AC2A8F0"))
	if evaluate(packet) != 1 {
		t.Errorf("doIt = %d; want 1", packet.version)
	}

	packet = readPacket(hexToBinary("F600BC2D8F"))
	if evaluate(packet) != 0 {
		t.Errorf("doIt = %d; want 0", packet.version)
	}

	packet = readPacket(hexToBinary("9C005AC2F8F0"))
	if evaluate(packet) != 0 {
		t.Errorf("doIt = %d; want 0", packet.version)
	}

	packet = readPacket(hexToBinary("9C0141080250320F1802104A08"))
	if evaluate(packet) != 1 {
		t.Errorf("doIt = %d; want 1", packet.version)
	}
}
