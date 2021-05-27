package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"testing"
)

func TestConnect(t *testing.T) {
	cfg := elasticsearch.Config{
		Addresses:             []string{"http://127.0.0.1:9200"},
		Username:              "",
		Password:              "",
		CloudID:               "",
		APIKey:                "",
		ServiceToken:          "",
		Header:                nil,
		CACert:                nil,
		RetryOnStatus:         nil,
		DisableRetry:          false,
		EnableRetryOnTimeout:  false,
		MaxRetries:            0,
		DiscoverNodesOnStart:  false,
		DiscoverNodesInterval: 0,
		EnableMetrics:         false,
		EnableDebugLogger:     false,
		DisableMetaHeader:     false,
		RetryBackoff:          nil,
		Transport:             nil,
		Logger:                nil,
		Selector:              nil,
		ConnectionPoolFunc:    nil,
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	resp, err := client.Info()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()

	fmt.Printf("%s\n", resp)
}
