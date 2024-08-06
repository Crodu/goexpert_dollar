package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Crodu/goexpert_dollar/server"
)

// GetExchange gets the exchange rate.
func GetExchange() error {
	url := "http://localhost:8080/cotacao"

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}

	var exchange server.Exchange

	err = json.Unmarshal(body, &exchange)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return err
	}

	fmt.Println("Exchange rate:", exchange.Usdbrl.Bid)
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", exchange.Usdbrl.Bid))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("Bid saved to cotacao.txt")

	return nil

}
