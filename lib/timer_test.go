package lib

import (
	"testing"
	"time"
)

const tps = 60

func TestNewTimer(t *testing.T) {
	tmr := NewTimer(time.Second, tps)
	expected := int(time.Second.Milliseconds()) * tps / 1000
	if tmr.targetTicks != expected {
		t.Errorf("expected targetTicks %d, got %d", expected, tmr.targetTicks)
	}
}

func TestTimerNotReadyInitially(t *testing.T) {
	tmr := NewTimer(time.Second, tps)
	if tmr.IsReady() {
		t.Error("expected timer not to be ready initially")
	}
}

func TestTimerBecomesReady(t *testing.T) {
	tmr := NewTimer(time.Second, tps)
	for range tmr.targetTicks {
		tmr.Update()
	}
	if !tmr.IsReady() {
		t.Error("expected timer to be ready after target ticks")
	}
}

func TestTimerNotReadyBeforeTarget(t *testing.T) {
	tmr := NewTimer(time.Second, tps)
	for range tmr.targetTicks - 1 {
		tmr.Update()
	}
	if tmr.IsReady() {
		t.Error("expected timer not to be ready before target ticks")
	}
}

func TestTimerUpdateStopsAtTarget(t *testing.T) {
	tmr := NewTimer(time.Second, tps)
	for range tmr.targetTicks + 100 {
		tmr.Update()
	}
	if tmr.currentTicks != tmr.targetTicks {
		t.Errorf("expected currentTicks to stop at %d, got %d", tmr.targetTicks, tmr.currentTicks)
	}
}

func TestTimerReset(t *testing.T) {
	tmr := NewTimer(time.Second, tps)
	for range tmr.targetTicks {
		tmr.Update()
	}
	tmr.Reset()
	if tmr.IsReady() {
		t.Error("expected timer not to be ready after reset")
	}
	if tmr.currentTicks != 0 {
		t.Errorf("expected currentTicks 0 after reset, got %d", tmr.currentTicks)
	}
}

func TestTimerResetAllowsReuse(t *testing.T) {
	tmr := NewTimer(time.Second, tps)
	for range tmr.targetTicks {
		tmr.Update()
	}
	tmr.Reset()
	for range tmr.targetTicks {
		tmr.Update()
	}
	if !tmr.IsReady() {
		t.Error("expected timer to be ready again after reset and re-arming")
	}
}

func TestTimerZeroDuration(t *testing.T) {
	tmr := NewTimer(0, tps)
	if !tmr.IsReady() {
		t.Error("expected zero-duration timer to be ready immediately")
	}
}

func TestTimerMinimumDuration(t *testing.T) {
	d := time.Duration(1000/tps+1) * time.Millisecond
	tmr := NewTimer(d, tps)
	if tmr.targetTicks < 1 {
		t.Error("expected at least 1 tick for duration >= ceil(1000/TPS)")
	}
}
