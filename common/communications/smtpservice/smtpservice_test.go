package smtpservice

import (
	"testing"

	"github.com/aaabigfish/goex/common/communications/base"
	"github.com/aaabigfish/goex/config"
)

var s SMTPservice

func TestSetup(t *testing.T) {
	t.Parallel()
	cfg := &config.Config{Communications: base.CommunicationsConfig{}}
	commsCfg := cfg.GetCommunicationsConfig()
	s.Setup(&commsCfg)
}

func TestConnect(t *testing.T) {
	if err := s.Connect(); err != nil {
		t.Error("smtpservice Connect() error", err)
	}
}

func TestPushEvent(t *testing.T) {
	err := s.PushEvent(base.Event{})
	if err == nil {
		t.Error("smtpservice PushEvent() error cannot be nil")
	}
}

func TestSend(t *testing.T) {
	err := s.Send("", "")
	if err == nil {
		t.Error("smtpservice Send() error cannot be nil")
	}
	err = s.Send("subject", "alertmessage")
	if err == nil {
		t.Error("smtpservice Send() error cannot be nil")
	}
}
