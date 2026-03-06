package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.yaml.in/yaml/v2"
)

type Notifier interface {
	Notify(ctx context.Context, title, body string) error
}

type MultiNotifier []Notifier

func (m MultiNotifier) Notify(ctx context.Context, title, body string) error {
	var errs []error
	for _, n := range m {
		if err := n.Notify(ctx, title, body); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("some notifiers failed: %v", errs)
	}
	return nil
}

// Mattermost notifier

type MattermostNotifier struct {
	WebhookURL string
	Channel    string
}

func (m MattermostNotifier) Notify(ctx context.Context, title, body string) error {
	payload := map[string]string{
		"text": fmt.Sprintf("**%s**\n%s", title, body),
	}

	if m.Channel != "" {
		payload["channel"] = m.Channel
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, m.WebhookURL, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("mattermost webhook returned %s", resp.Status)
	}

	return nil
}

// configs

type AlertRule struct {
	TermMonths     int           `yaml:"term_months"`
	AlertBeforeRaw string        `yaml:"alert_before"`
	AlertBefore    time.Duration `yaml:"-"`
}

type AlertConfig struct {
	Rules []AlertRule `yaml:"rules"`
}

func (c *AlertConfig) Prepare() error {
	for i := range c.Rules {
		d, err := time.ParseDuration(c.Rules[i].AlertBeforeRaw)
		if err != nil {
			return fmt.Errorf("rule for term %d: %w", c.Rules[i].TermMonths, err)
		}
		c.Rules[i].AlertBefore = d
	}
	return nil
}

func (c *AlertConfig) FindRule(termMonths int) *AlertRule {
	for i := range c.Rules {
		if c.Rules[i].TermMonths == termMonths {
			return &c.Rules[i]
		}
	}
	return nil
}

func LoadAlertConfig(path string) (*AlertConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg AlertConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if err := cfg.Prepare(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type ContractPeriod struct {
	PeriodStart time.Time
	PeriodEnd   time.Time
	TimeLeft    time.Duration
}

func computeContractPeriod(now, start time.Time, termMonths int) ContractPeriod {
	periodStart := start
	for periodStart.AddDate(0, termMonths, 0).Before(now) {
		periodStart = periodStart.AddDate(0, termMonths, 0)
	}

	periodEnd := periodStart.AddDate(0, termMonths, 0)
	return ContractPeriod{
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
		TimeLeft:    periodEnd.Sub(now),
	}
}

func shouldAlert(termMonths int, period ContractPeriod, cfg *AlertConfig) bool {
	rule := cfg.FindRule(termMonths)
	if rule == nil {
		return false
	}
	return period.TimeLeft <= rule.AlertBefore
}

func formatTimeLeft(p ContractPeriod) string {
	now := time.Now()
	end := p.PeriodEnd

	months := 0
	t := now

	for t.AddDate(0, 1, 0).Before(end) || t.AddDate(0, 1, 0).Equal(end) {
		t = t.AddDate(0, 1, 0)
		months++
	}

	days := int(end.Sub(t).Hours() / 24)

	switch {
	case months > 0 && days > 0:
		return fmt.Sprintf("%d months %d days", months, days)
	case months > 0:
		return fmt.Sprintf("%d months", months)
	case days > 0:
		return fmt.Sprintf("%d days", days)
	default:
		return "less than a day"
	}
}
