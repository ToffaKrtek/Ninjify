package server

import (
	"fmt"
	"net"
	"testing"
)

func TestGetPublicAddress(t *testing.T) {
	publicAddr, err := getPublicAddress()
	if err != nil {
		t.Fatalf("Ошибка при получении публичного адреса: %v", err)
	}

	if publicAddr == "" {
		t.Fatal("Публичный адрес пустой")
	}

	if _, err := net.ResolveUDPAddr("udp", publicAddr); err != nil {
		t.Fatalf("Неверный формат публичного адреса: %v", err)
	}

	fmt.Println("Полученный публичный адрес:", publicAddr)
}
