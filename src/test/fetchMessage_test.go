package test

import (
	"change-status-go/src"
	"testing"
)

func TestPingMessage(t *testing.T) {
	result, Err := src.FetchMessage("ping")
	if Err != nil {
		t.Fatalf("failed at Test fetchMessage_ping: %#v", Err)
	}
	if result != "Pong!" {
		t.Fatalf("failed test")
	}
}

func TestHelpMessage(t *testing.T) {
	result, Err := src.FetchMessage("help")
	if Err != nil {
		t.Fatalf("failed at Test fetchMessage_ping: %#v", Err)
	}
	if result != src.HelpMessage {
		t.Fatalf("failed test")
	}
}

func TestFetchMessageFailed(t *testing.T) {
	result, Err := src.FetchMessage("%unsupportedCommandTest")
	if Err == nil {
		t.Fatal("failed to test,please check your code.")
	}
	if result != src.DefaultMessage {
		t.Fatal("not returned true message.")
	}
}
