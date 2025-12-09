package cmd

import (
	"testing"
	"time"
)

func TestComputeContractPeriod_OneYearTerm(t *testing.T) {
	// Contract starts on 2022-01-01 with a 12-month term
	start := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	termMonths := 12

	// "now" is in the middle of 2024
	now := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)

	period := computeContractPeriod(now, start, termMonths)

	// Expected current contract period:
	//   2024-01-01 → 2025-01-01
	wantStart := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	wantEnd := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	if !period.PeriodStart.Equal(wantStart) {
		t.Fatalf("PeriodStart = %s, want %s", period.PeriodStart, wantStart)
	}
	if !period.PeriodEnd.Equal(wantEnd) {
		t.Fatalf("PeriodEnd = %s, want %s", period.PeriodEnd, wantEnd)
	}

	if period.TimeLeft <= 0 {
		t.Fatalf("TimeLeft should be > 0, got %s", period.TimeLeft)
	}
}

func TestShouldAlert_WithMatchingRule(t *testing.T) {
	cfg := &AlertConfig{
		Rules: []AlertRule{
			{TermMonths: 12, AlertBeforeRaw: "2160h"}, // ~90 days
		},
	}
	if err := cfg.Prepare(); err != nil {
		t.Fatalf("Prepare() error = %v", err)
	}

	// 80 days left — should alert
	period := ContractPeriod{
		TimeLeft: 80 * 24 * time.Hour,
	}

	if !shouldAlert(12, period, cfg) {
		t.Fatalf("shouldAlert() = false, want true for 80 days left and 90-day rule")
	}

	// 120 days left — should NOT alert
	period = ContractPeriod{
		TimeLeft: 120 * 24 * time.Hour,
	}

	if shouldAlert(12, period, cfg) {
		t.Fatalf("shouldAlert() = true, want false for 120 days left and 90-day rule")
	}
}

func TestShouldAlert_NoRule(t *testing.T) {
	cfg := &AlertConfig{
		Rules: []AlertRule{
			{TermMonths: 6, AlertBeforeRaw: "2160h"},
		},
	}
	if err := cfg.Prepare(); err != nil {
		t.Fatalf("Prepare() error = %v", err)
	}

	// No rule for 12-month contracts — should NOT alert
	period := ContractPeriod{
		TimeLeft: 10 * 24 * time.Hour,
	}

	if shouldAlert(12, period, cfg) {
		t.Fatalf("shouldAlert() = true, want false when no rule exists")
	}
}

func TestFormatTimeLeft_MonthsAndDays(t *testing.T) {
	// Use current time to compute end date
	now := time.Now().UTC()

	// End date in 2 months and 10 days
	end := now.AddDate(0, 2, 10)

	period := ContractPeriod{
		PeriodEnd: end,
	}

	got := formatTimeLeft(period)

	// Expected format is: "2 months 10 days"
	if got != "2 months 10 days" {
		t.Fatalf("formatTimeLeft() = %q, want %q", got, "2 months 10 days")
	}
}

func TestFormatTimeLeft_OnlyDays(t *testing.T) {
	now := time.Now().UTC()
	end := now.Add(5 * 24 * time.Hour)

	period := ContractPeriod{
		PeriodEnd: end,
	}

	got := formatTimeLeft(period)

	if got != "5 days" {
		t.Fatalf("formatTimeLeft() = %q, want %q", got, "5 days")
	}
}
