package validator

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
)

type Validator interface {
	Validate(string) ValidationResult
}

type ValidationResult struct {
	IsValid    bool
	Errors     []string
	Warnings   []string
	Normalized string ///?
}

type StringArr []string

type URLValidator struct {
	allowedSchemes map[string]bool
	allowedTLDs    map[string]bool
	blockedTLDs    map[string]bool
	maxLength      int
	minLength      int
}

func NewUrlValidator() *URLValidator {
	return &URLValidator{
		allowedSchemes: map[string]bool{
			"http":  true,
			"https": true,
			"ftp":   true,
		},
		blockedTLDs: map[string]bool{
			"onion": true, // Tor сети
			"i2p":   true, // Анонимные сети
			"local": true, // Локальные адреса??
		},
		allowedTLDs: map[string]bool{
			"com": true, "org": true, "net": true, "io": true,
			"ru": true, "dev": true, "app": true,
		},
		maxLength: 2048,
		minLength: 10,
	}
}

func (sArr StringArr) String() string {
	return strings.Join([]string(sArr), ",")
}

func (v *URLValidator) Validate(rawUrl string) ValidationResult {
	result := ValidationResult{
		IsValid:  false,
		Errors:   []string{},
		Warnings: []string{},
	}
	if len(rawUrl) > v.maxLength {
		result.Errors = append(result.Errors, fmt.Sprintf("URL too long: %d", len(rawUrl)))
		return result
	}

	if len(rawUrl) < v.minLength {
		result.Errors = append(result.Errors, fmt.Sprintf("URL too short: %d", len(rawUrl)))
	}

	normalized, err := v.normalizeURL(rawUrl)
	if err != nil {
		result.Errors = append(result.Errors, err.Error())
		return result
	}

	result.Normalized = normalized

	parsed, err := url.Parse(normalized)
	if err != nil {
		result.Errors = append(result.Errors, err.Error())
		return result
	}

	if err := v.validateScheme(parsed); err != nil {
		result.Errors = append(result.Errors, err.Error())
		return result
	}

	if err := v.validateHost(parsed); err != nil {
		result.Errors = append(result.Errors, err.Error())
		return result
	}

	if ok := v.isPhishingLike(parsed.Host); ok {
		result.Errors = append(result.Errors, fmt.Errorf("phishing-like url").Error())
		return result
	}

	return result
}

func (v *URLValidator) normalizeURL(rawUrl string) (string, error) {
	if !regexp.MustCompile(`^[a-zA-Z]+://`).MatchString(rawUrl) {
		rawUrl = "https://" + rawUrl
	}

	decoded, err := url.QueryUnescape(rawUrl)
	if err != nil {
		return "", fmt.Errorf("URL decoding failed")
	}

	decoded = strings.TrimSpace(decoded)

	parsed, err := url.Parse(decoded)
	if err != nil {
		return "", fmt.Errorf("URL parsing failed")
	}

	parsed.Host = strings.ToLower(parsed.Host)

	return parsed.String(), nil
}

func (v *URLValidator) validateScheme(parsed *url.URL) error {
	if parsed.Scheme == "" {
		return fmt.Errorf("URL scheme is required")
	}

	if !v.allowedSchemes[strings.ToLower(parsed.Scheme)] {
		return fmt.Errorf("URL Scheme: %s is not allowed", parsed.Scheme)
	}

	restrictedSchemes := map[string]bool{
		"javascript": true,
		"data":       true,
		"file":       true,
		"vbscript":   true,
		"telnet":     true,
	}

	if restrictedSchemes[strings.ToLower(parsed.Scheme)] {
		return fmt.Errorf("URL Scheme: %s is restricted", parsed.Scheme)
	}

	return nil

}

func (v *URLValidator) isPhishingLike(host string) bool {
	phishingPatterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)paypal`),
		regexp.MustCompile(`(?i)bank`),
		regexp.MustCompile(`(?i)login`),
		regexp.MustCompile(`(?i)account`),
		regexp.MustCompile(`-security-`),
	}

	for _, pattern := range phishingPatterns {
		if pattern.MatchString(host) {
			return true
		}
	}
	return false
}

func (v *URLValidator) validateTLD(host string) error {
	parts := strings.Split(host, ".")
	if len(parts) < 2 {
		return fmt.Errorf("invalid domain format")
	}

	tld := parts[len(parts)-1]

	// Проверка заблокированных TLD
	if v.blockedTLDs[strings.ToLower(tld)] {
		return fmt.Errorf("TLD '%s' is blocked", tld)
	}

	// Если есть белый список TLD
	if len(v.allowedTLDs) > 0 && !v.allowedTLDs[strings.ToLower(tld)] {
		return fmt.Errorf("TLD '%s' is not allowed", tld)
	}

	return nil
}

func (v *URLValidator) validateHost(parsed *url.URL) error {
	if parsed.Host == "" {
		return fmt.Errorf("host is required")
	}

	// Проверка IP адресов
	if ip := net.ParseIP(parsed.Host); ip != nil {
		if v.isPrivateIP(ip) {
			return fmt.Errorf("private IP addresses are not allowed")
		}
		if v.isReservedIP(ip) {
			return fmt.Errorf("reserved IP addresses are not allowed")
		}
	}

	// Проверка TLD
	if err := v.validateTLD(parsed.Host); err != nil {
		return err
	}

	return nil
}

func (v *URLValidator) isPrivateIP(ip net.IP) bool {
	privateBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
		"::1/128",
		"fc00::/7",
	}

	for _, block := range privateBlocks {
		_, network, _ := net.ParseCIDR(block)
		if network.Contains(ip) {
			return true
		}
	}
	return false
}

func (v *URLValidator) isReservedIP(ip net.IP) bool {
	reservedBlocks := []string{
		"0.0.0.0/8",
		"169.254.0.0/16",
		"224.0.0.0/4",
	}

	for _, block := range reservedBlocks {
		_, network, _ := net.ParseCIDR(block)
		if network.Contains(ip) {
			return true
		}
	}
	return false
}
