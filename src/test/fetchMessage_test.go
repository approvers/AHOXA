package test

import (
	"change-status-go/src"
	"log"
	"testing"
)

func pingMessageTest(t *testing.T) {
	result, Err := src.FetchMessage("ping")
	if Err != nil {
		t.Fatalf("failed at Test fetchMessage_ping: %#v", Err)
	}
	if result != "Pong!" {
		t.Fatalf("failed test")
	}
}

func helpMessageTest(t *testing.T) {
	result, Err := src.FetchMessage("help")
	if Err != nil {
		t.Fatalf("failed at Test fetchMessage_ping: %#v", Err)
	}
	if result != "Pong!" {
		t.Fatalf("failed test")
	}
}

func fetchMessageFailedTest(t *testing.T) {
	result, Err := src.FetchMessage("%hogehoge")
	if Err == nil {
		log.Fatal("failed to test,please check your code.")
	}
	if result != src.DefaultMessage {
		log.Fatal("not returned true message.")
	}
}
