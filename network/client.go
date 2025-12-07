package client

import (
	"fmt"
	"log"
	"net"
	"slices"
)

type Client struct {
	ipAddress string
	port      int
}

func NewClient(address string, port int) *Client {
	return &Client{
		ipAddress: address,
		port:      port,
	}
}

func (c *Client) sendQuery(query []byte) []byte {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", c.ipAddress, c.port))
	if err != nil {
		log.Fatal("Connection error", err)
	}
	defer conn.Close()

	if _, err := conn.Write(query); err != nil {
		log.Fatal("Could not send query", err)
	}

	response := make([]byte, 1024)
	lengthOfResponse, err := conn.Read(response)
	if err != nil {
		log.Fatal("Could not read response", err)
	}

	if !hasTheSameId(query, response) {
		log.Fatalf("Response doesn't have the same ID as the message. q:%v, r:%v\n", query, response)
	}

	return response[:lengthOfResponse]
}

func hasTheSameId(query, response []byte) bool {
	return slices.Equal(query[:2], response[:2])
}
