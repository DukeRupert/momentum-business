package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

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
