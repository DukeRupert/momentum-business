package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// TurnstileResponse represents Cloudflare's siteverify response
type TurnstileResponse struct {
	Success     bool     `json:"success"`
	ErrorCodes  []string `json:"error-codes,omitempty"`
	ChallengeTS string   `json:"challenge_ts,omitempty"`
	Hostname    string   `json:"hostname,omitempty"`
}

// verifyTurnstile verifies the Turnstile token with Cloudflare
func verifyTurnstile(token, secretKey, remoteIP string) (bool, error) {
	if secretKey == "" {
		log.Println("TURNSTILE_SECRET_KEY not set, skipping verification")
		return true, nil
	}

	payload := map[string]string{
		"secret":   secretKey,
		"response": token,
	}
	if remoteIP != "" {
		payload["remoteip"] = remoteIP
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(
		"https://challenges.cloudflare.com/turnstile/v0/siteverify",
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result TurnstileResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return false, err
	}

	if !result.Success {
		log.Printf("Turnstile verification failed: %v", result.ErrorCodes)
	}

	return result.Success, nil
}

// ContactResponse represents the response from the contact endpoint
type ContactResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
	Error   string            `json:"error,omitempty"`
	Errors  []ValidationError `json:"errors,omitempty"`
	Data    *ContactData      `json:"data,omitempty"`
}

// ContactData contains data to pass back to the client
type ContactData struct {
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var form ContactForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Printf("Failed to decode request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ContactResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	// Check honeypot field - if filled, it's a bot
	// Return fake success to not alert the bot
	if strings.TrimSpace(form.Website) != "" {
		log.Printf("Honeypot triggered - likely bot submission from IP: %s", r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ContactResponse{
			Success: true,
			Message: "Message sent successfully",
		})
		return
	}

	// Verify Turnstile token
	turnstileSecret := os.Getenv("TURNSTILE_SECRET_KEY")
	if form.TurnstileResponse == "" && turnstileSecret != "" {
		log.Printf("Missing Turnstile token from IP: %s", r.RemoteAddr)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ContactResponse{
			Success: false,
			Error:   "Please complete the security check",
		})
		return
	}

	if turnstileSecret != "" {
		verified, err := verifyTurnstile(form.TurnstileResponse, turnstileSecret, r.RemoteAddr)
		if err != nil {
			log.Printf("Turnstile verification error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ContactResponse{
				Success: false,
				Error:   "Security verification failed. Please try again.",
			})
			return
		}
		if !verified {
			log.Printf("Turnstile verification failed for IP: %s", r.RemoteAddr)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ContactResponse{
				Success: false,
				Error:   "Security check failed. Please try again.",
			})
			return
		}
	}

	// Trim whitespace from all string fields
	form.FirstName = strings.TrimSpace(form.FirstName)
	form.LastName = strings.TrimSpace(form.LastName)
	form.Email = strings.TrimSpace(form.Email)
	form.PhoneNumber = strings.TrimSpace(form.PhoneNumber)
	form.AnnualRevenue = strings.TrimSpace(form.AnnualRevenue)
	form.Message = strings.TrimSpace(form.Message)

	// Validate form
	validationResult := form.Validate()
	if !validationResult.Valid {
		log.Printf("Validation failed: %+v", validationResult.Errors)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ContactResponse{
			Success: false,
			Error:   "Validation failed",
			Errors:  validationResult.Errors,
		})
		return
	}

	// Get environment variables for email
	postmarkToken := os.Getenv("POSTMARK_TOKEN")
	postmarkTo := os.Getenv("POSTMARK_TO")
	postmarkFrom := os.Getenv("POSTMARK_FROM")

	if postmarkToken == "" || postmarkTo == "" || postmarkFrom == "" {
		log.Printf("Missing email configuration: token=%v, to=%v, from=%v",
			postmarkToken != "", postmarkTo != "", postmarkFrom != "")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ContactResponse{
			Success: false,
			Error:   "Server configuration error",
		})
		return
	}

	// Send notification email to business
	if err := SendContactFormEmail(&form, postmarkToken, postmarkTo, postmarkFrom); err != nil {
		log.Printf("Failed to send contact form email: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ContactResponse{
			Success: false,
			Error:   "Failed to send message. Please try again.",
		})
		return
	}

	// Send thank you email to customer
	if err := SendThankYouEmail(&form, postmarkToken, postmarkFrom); err != nil {
		// Log the error but don't fail the request
		log.Printf("Failed to send thank you email: %v", err)
	}

	log.Printf("Contact form submitted successfully: %s %s <%s>",
		form.FirstName, form.LastName, form.Email)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ContactResponse{
		Success: true,
		Message: "Message sent successfully",
		Data: &ContactData{
			FirstName: form.FirstName,
			Email:     form.Email,
		},
	})
}
