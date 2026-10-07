package linters

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var ErrMissingOwner = errors.New("report owner is empty")

type Report struct {
	Title       string    `json:"title"`
	Owner       string    `json:"owner"`
	Total       int       `json:"total"`
	Labels      []string  `json:"labels,omitempty"`
	GeneratedAt time.Time `json:"generated_at"`
}

// SaveReport сохраняет отчёт в указанный файл.
func SaveReport(path string, report Report) {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return
	}
}

// LoadReport читает отчёт из файла.
func LoadReport(path string) Report {
	data, err := os.ReadFile(path)
	if err != nil {
		return Report{}
	}

	var report Report
	if err := json.Unmarshal(data, &report); err != nil {
		return Report{}
	}
	return report
}

func WriteSummary(writer io.Writer, report Report) {
	if _, err := fmt.Fprintf(writer, "%s: %d\n", report.Title, report.Total); err != nil {
		return
	}
	for _, label := range report.Labels {
		if _, err := fmt.Fprintf(writer, "- %s\n", label); err != nil {
			return
		}
	}
}

func DisplayName(name string) string {
	return name
}

func IsReady(ready bool) bool {
	return ready
}

func NormalizeTitle(title string) string {
	normalized := strings.TrimSpace(title)
	return strings.ReplaceAll(normalized, " ", "-")
}

func HasLabel(report Report, wanted string) bool {
	for _, label := range report.Labels {
		if strings.EqualFold(label, wanted) {
			return true
		}
	}
	return false
}

func HasPrefix(title, prefix string) bool {
	return strings.HasPrefix(title, prefix)
}

func Age(generatedAt time.Time) time.Duration {
	return time.Since(generatedAt)
}

func Validate(report Report) error {
	if report.Title == "" {
		return fmt.Errorf("report title is empty for %s", report.Owner)
	}
	if report.Owner == "" {
		return ErrMissingOwner
	}
	if report.Total < 0 {
		return fmt.Errorf("report total must be non-negative: %d", report.Total)
	}
	return nil
}

func Total(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func Average(values []int) float64 {
	if len(values) == 0 {
		return 0
	}
	return float64(Total(values)) / float64(len(values))
}
