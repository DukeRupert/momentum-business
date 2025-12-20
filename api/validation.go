package main

import (
	"regexp"
	"strings"
)

// ContactForm represents the contact form submission
type ContactForm struct {
	FirstName         string   `json:"first-name"`
	LastName          string   `json:"last-name"`
	Email             string   `json:"email"`
	PhoneNumber       string   `json:"phone-number"`
	AnnualRevenue     string   `json:"annual-revenue"`
	Services          []string `json:"services"`
	Message           string   `json:"message"`
	Website           string   `json:"website"`              // Honeypot field
	TurnstileResponse string   `json:"cf-turnstile-response"` // Cloudflare Turnstile token
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationResult represents the result of validation
type ValidationResult struct {
	Valid  bool              `json:"valid"`
	Errors []ValidationError `json:"errors"`
}

// Valid annual revenue values
var validRevenueRanges = map[string]bool{
	"under-100k": true,
	"100k-500k":  true,
	"500k-1m":    true,
	"1m-5m":      true,
	"over-5m":    true,
}

// Valid service values
var validServices = map[string]bool{
	"essentials":       true,
	"growth-strategy":  true,
	"complete-support": true,
	"consulting":       true,
	"cleanup":          true,
}

// Regex patterns
var (
	namePattern  = regexp.MustCompile(`^[a-zA-Z\s'\-]+$`)
	emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	phonePattern = regexp.MustCompile(`^[\+]?[1-9]?[\d\s\-\(\)\.]{10,15}$`)
)

// Validate validates the contact form
func (f *ContactForm) Validate() ValidationResult {
	result := ValidationResult{Valid: true, Errors: []ValidationError{}}

	// Validate first name
	firstName := strings.TrimSpace(f.FirstName)
	if firstName == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "first-name",
			Message: "First name is required",
		})
	} else if len(firstName) < 2 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "first-name",
			Message: "First name must be at least 2 characters",
		})
	} else if len(firstName) > 50 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "first-name",
			Message: "First name must be less than 50 characters",
		})
	} else if !namePattern.MatchString(firstName) {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "first-name",
			Message: "First name can only contain letters, spaces, hyphens, and apostrophes",
		})
	}

	// Validate last name
	lastName := strings.TrimSpace(f.LastName)
	if lastName == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "last-name",
			Message: "Last name is required",
		})
	} else if len(lastName) < 2 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "last-name",
			Message: "Last name must be at least 2 characters",
		})
	} else if len(lastName) > 50 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "last-name",
			Message: "Last name must be less than 50 characters",
		})
	} else if !namePattern.MatchString(lastName) {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "last-name",
			Message: "Last name can only contain letters, spaces, hyphens, and apostrophes",
		})
	}

	// Validate email
	email := strings.TrimSpace(f.Email)
	if email == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "email",
			Message: "Email is required",
		})
	} else if len(email) > 254 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "email",
			Message: "Email must be less than 254 characters",
		})
	} else if !emailPattern.MatchString(email) {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "email",
			Message: "Please enter a valid email address",
		})
	}

	// Validate phone number
	phone := strings.TrimSpace(f.PhoneNumber)
	if phone == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "phone-number",
			Message: "Phone number is required",
		})
	} else if !phonePattern.MatchString(phone) {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "phone-number",
			Message: "Please enter a valid phone number",
		})
	}

	// Validate annual revenue
	revenue := strings.TrimSpace(f.AnnualRevenue)
	if revenue == "" {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "annual-revenue",
			Message: "Please select your annual revenue range",
		})
	} else if !validRevenueRanges[revenue] {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "annual-revenue",
			Message: "Please select a valid revenue range",
		})
	}

	// Validate services
	if len(f.Services) == 0 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "services",
			Message: "Please select at least one service you're interested in",
		})
	} else {
		for _, service := range f.Services {
			if !validServices[service] {
				result.Errors = append(result.Errors, ValidationError{
					Field:   "services",
					Message: "Invalid service selected: " + service,
				})
				break
			}
		}
	}

	// Validate message (optional but has max length)
	if len(f.Message) > 2000 {
		result.Errors = append(result.Errors, ValidationError{
			Field:   "message",
			Message: "Message must be less than 2000 characters",
		})
	}

	result.Valid = len(result.Errors) == 0
	return result
}
