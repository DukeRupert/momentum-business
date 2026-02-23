package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// PostmarkEmail represents an email to send via Postmark
type PostmarkEmail struct {
	From          string `json:"From"`
	To            string `json:"To"`
	Subject       string `json:"Subject"`
	TextBody      string `json:"TextBody"`
	HtmlBody      string `json:"HtmlBody"`
	MessageStream string `json:"MessageStream"`
}

// PostmarkResponse represents the response from Postmark API
type PostmarkResponse struct {
	To          string `json:"To"`
	SubmittedAt string `json:"SubmittedAt"`
	MessageID   string `json:"MessageID"`
	ErrorCode   int    `json:"ErrorCode"`
	Message     string `json:"Message"`
}

// sendEmail sends an email via Postmark
func sendEmail(token string, email PostmarkEmail) error {
	body, err := json.Marshal(email)
	if err != nil {
		return fmt.Errorf("failed to marshal email: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.postmarkapp.com/email", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Postmark-Server-Token", token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var pmResp PostmarkResponse
		json.Unmarshal(respBody, &pmResp)
		return fmt.Errorf("postmark error %d: %s", pmResp.ErrorCode, pmResp.Message)
	}

	return nil
}

// formatRevenue converts revenue code to human-readable string
func formatRevenue(revenue string) string {
	revenueMap := map[string]string{
		"under-100k": "Under $100,000",
		"100k-500k":  "$100,000 - $500,000",
		"500k-1m":    "$500,000 - $1,000,000",
		"1m-5m":      "$1,000,000 - $5,000,000",
		"over-5m":    "Over $5,000,000",
	}
	if formatted, ok := revenueMap[revenue]; ok {
		return formatted
	}
	return revenue
}

// formatService converts service code to human-readable string
func formatService(service string) string {
	serviceMap := map[string]string{
		"essentials":       "Essentials Package",
		"growth-strategy":  "Growth Strategy Package",
		"complete-support": "Complete Business Support",
		"consulting":       "Financial Consulting",
		"cleanup":          "QuickBooks Cleanup",
	}
	if formatted, ok := serviceMap[service]; ok {
		return formatted
	}
	return service
}

// getServiceClass returns the CSS class for a service
func getServiceClass(service string) string {
	classMap := map[string]string{
		"essentials":       "bookkeeping",
		"growth-strategy":  "payroll",
		"complete-support": "consulting",
		"consulting":       "consulting",
		"cleanup":          "cleanup",
	}
	if class, ok := classMap[service]; ok {
		return class
	}
	return "bookkeeping"
}

// SendContactFormEmail sends the notification email to the business
func SendContactFormEmail(form *ContactForm, token, to, from string) error {
	// Get timestamp in PST
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		loc = time.UTC
	}
	timestamp := time.Now().In(loc).Format("Monday, January 2, 2006 at 3:04 PM MST")

	// Build services tags HTML
	var serviceTagsHTML strings.Builder
	for _, service := range form.Services {
		serviceTagsHTML.WriteString(fmt.Sprintf(
			`<span class="service-tag %s">%s</span>`,
			getServiceClass(service),
			formatService(service),
		))
	}

	// Build services list for text email
	var servicesList strings.Builder
	for _, service := range form.Services {
		servicesList.WriteString(fmt.Sprintf("* %s\n", formatService(service)))
	}

	// Build message section HTML
	messageHTML := ""
	if form.Message != "" {
		messageHTML = fmt.Sprintf(`
        <div class="section">
            <h2>Client Message</h2>
            <div class="message-box">
                "%s"
            </div>
        </div>
        `, form.Message)
	}

	// Build message section text
	messageText := ""
	if form.Message != "" {
		messageText = fmt.Sprintf(`CLIENT MESSAGE:
---------------
"%s"

`, form.Message)
	}

	htmlBody := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>New Lead: Contact Form Submission - %s %s</title>
    <style>
        body {
            font-family: 'Manrope', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            line-height: 1.6;
            color: #374151;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f8fafc;
        }
        .email-container {
            background: white;
            border-radius: 12px;
            padding: 32px;
            box-shadow: 0 4px 6px rgba(0,0,0,0.05);
            border: 1px solid #e2e8f0;
        }
        .header {
            border-bottom: 3px solid #53945c;
            padding-bottom: 24px;
            margin-bottom: 32px;
        }
        .company-name {
            color: #53945c;
            font-size: 28px;
            font-weight: 700;
            margin: 0;
            font-family: 'Outfit', sans-serif;
        }
        .tagline {
            color: #64748b;
            font-size: 15px;
            margin: 6px 0 0 0;
            font-weight: 500;
        }
        .lead-priority {
            display: inline-block;
            background: #53945c;
            color: white;
            padding: 6px 16px;
            border-radius: 20px;
            font-size: 13px;
            font-weight: 600;
            margin-top: 12px;
        }
        .section {
            margin-bottom: 28px;
        }
        .section h2 {
            color: #1f2937;
            font-size: 18px;
            font-weight: 600;
            margin-bottom: 16px;
            border-bottom: 2px solid #e5e7eb;
            padding-bottom: 8px;
            font-family: 'Outfit', sans-serif;
        }
        .info-grid {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 16px;
            margin-bottom: 20px;
        }
        .info-item {
            background: #f4f9f5;
            padding: 16px;
            border-radius: 8px;
            border-left: 4px solid #53945c;
        }
        .info-label {
            font-weight: 600;
            color: #374151;
            font-size: 14px;
            margin-bottom: 6px;
            text-transform: uppercase;
            letter-spacing: 0.5px;
        }
        .info-value {
            color: #1f2937;
            font-size: 15px;
            font-weight: 500;
        }
        .revenue-highlight {
            background: #dfe9fa;
            border-left-color: #4f7ee2;
        }
        .services-list {
            background: #dfe9fa;
            padding: 20px;
            border-radius: 8px;
            border-left: 4px solid #4f7ee2;
        }
        .service-tag {
            display: inline-block;
            background: #53945c;
            color: white;
            padding: 6px 14px;
            border-radius: 18px;
            font-size: 13px;
            font-weight: 500;
            margin-right: 10px;
            margin-bottom: 6px;
            text-transform: capitalize;
        }
        .service-tag.bookkeeping { background: #53945c; }
        .service-tag.payroll { background: #4f7ee2; }
        .service-tag.consulting { background: #417848; }
        .service-tag.cleanup { background: #709fea; }
        .message-box {
            background: #f8fafc;
            border: 2px solid #e2e8f0;
            border-radius: 8px;
            padding: 20px;
            font-style: italic;
            color: #475569;
            line-height: 1.7;
        }
        .footer {
            margin-top: 32px;
            padding-top: 24px;
            border-top: 2px solid #e5e7eb;
            text-align: center;
            color: #64748b;
            font-size: 13px;
        }
        .submission-meta {
            background: #f1f5f9;
            padding: 12px 16px;
            border-radius: 6px;
            font-size: 12px;
            color: #64748b;
            margin-bottom: 20px;
        }
    </style>
</head>
<body>
    <div class="email-container">
        <div class="header">
            <h1 class="company-name">Momentum Business Solutions</h1>
            <p class="tagline">Where Strategy Meets Execution</p>
            <span class="lead-priority">New Qualified Lead</span>
        </div>

        <div class="submission-meta">
            <strong>Submitted:</strong> %s | <strong>Source:</strong> Website Contact Form
        </div>

        <div class="section">
            <h2>Contact Information</h2>
            <div class="info-grid">
                <div class="info-item">
                    <div class="info-label">Full Name</div>
                    <div class="info-value">%s %s</div>
                </div>
                <div class="info-item">
                    <div class="info-label">Email Address</div>
                    <div class="info-value">%s</div>
                </div>
                <div class="info-item">
                    <div class="info-label">Phone Number</div>
                    <div class="info-value">%s</div>
                </div>
                <div class="info-item revenue-highlight">
                    <div class="info-label">Annual Revenue</div>
                    <div class="info-value">%s</div>
                </div>
            </div>
        </div>

        <div class="section">
            <h2>Services of Interest</h2>
            <div class="services-list">
                <div class="info-label" style="margin-bottom: 12px;">Client selected the following services:</div>
                %s
            </div>
        </div>

        %s

        <div class="footer">
            <p><strong>Momentum Business Solutions</strong></p>
            <p>QuickBooks Online | Payroll Processing | Financial Consulting | Strategic Planning</p>
            <p>Email: cade@momentumbusiness.org | Phone: (509) 554-8022</p>
        </div>
    </div>
</body>
</html>`,
		form.FirstName, form.LastName,
		timestamp,
		form.FirstName, form.LastName,
		form.Email,
		form.PhoneNumber,
		formatRevenue(form.AnnualRevenue),
		serviceTagsHTML.String(),
		messageHTML,
	)

	textBody := fmt.Sprintf(`NEW QUALIFIED LEAD - Momentum Business Solutions
===============================================

SUBMISSION DETAILS:
Submitted: %s
Source: Website Contact Form

CONTACT INFORMATION:
-------------------
Name: %s %s
Email: %s
Phone: %s
Annual Revenue: %s

SERVICES OF INTEREST:
--------------------
Client selected the following services:
%s
%sCONTACT INFORMATION:
-------------------
Momentum Business Solutions
Where Strategy Meets Execution

QuickBooks Online | Payroll Processing | Financial Consulting | Strategic Planning
Email: cade@momentumbusiness.org
Phone: (509) 554-8022

---
This email was generated from your website contact form.
`,
		timestamp,
		form.FirstName, form.LastName,
		form.Email,
		form.PhoneNumber,
		formatRevenue(form.AnnualRevenue),
		servicesList.String(),
		messageText,
	)

	email := PostmarkEmail{
		From:          from,
		To:            to,
		Subject:       fmt.Sprintf("New Lead: Contact Form Submission - %s %s", form.FirstName, form.LastName),
		TextBody:      textBody,
		HtmlBody:      htmlBody,
		MessageStream: "outbound",
	}

	return sendEmail(token, email)
}

// SendThankYouEmail sends a thank you email to the customer
func SendThankYouEmail(form *ContactForm, token, from string) error {
	htmlBody := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Thank You for Your Interest - Momentum Business Solutions</title>
    <style>
        body {
            font-family: 'Manrope', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            line-height: 1.6;
            color: #374151;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f8fafc;
        }
        .email-container {
            background: white;
            border-radius: 12px;
            padding: 32px;
            box-shadow: 0 4px 6px rgba(0,0,0,0.05);
            border: 1px solid #e2e8f0;
        }
        .header {
            text-align: center;
            border-bottom: 3px solid #53945c;
            padding-bottom: 24px;
            margin-bottom: 32px;
        }
        .company-name {
            color: #53945c;
            font-size: 28px;
            font-weight: 700;
            margin: 0;
            font-family: 'Outfit', sans-serif;
        }
        .tagline {
            color: #64748b;
            font-size: 15px;
            margin: 6px 0 0 0;
            font-weight: 500;
        }
        .greeting {
            font-size: 24px;
            color: #1f2937;
            font-weight: 600;
            margin-bottom: 20px;
            text-align: center;
        }
        .main-content {
            font-size: 16px;
            line-height: 1.7;
            color: #374151;
            margin-bottom: 32px;
        }
        .timeline-box {
            background: #dfe9fa;
            border-left: 4px solid #4f7ee2;
            padding: 20px;
            border-radius: 8px;
            margin: 24px 0;
        }
        .timeline-box h3 {
            color: #1e40af;
            margin: 0 0 12px 0;
            font-size: 18px;
            font-weight: 600;
        }
        .timeline-box p {
            margin: 0;
            color: #1e3a8a;
            font-weight: 500;
        }
        .contact-info {
            background: #f8fafc;
            border: 1px solid #e2e8f0;
            border-radius: 8px;
            padding: 20px;
            margin: 24px 0;
        }
        .contact-info h3 {
            color: #374151;
            margin: 0 0 12px 0;
            font-size: 16px;
            font-weight: 600;
        }
        .contact-detail {
            margin: 8px 0;
            color: #4b5563;
        }
        .contact-detail strong {
            color: #374151;
        }
        .footer {
            margin-top: 32px;
            padding-top: 24px;
            border-top: 2px solid #e5e7eb;
            text-align: center;
            color: #64748b;
            font-size: 13px;
        }
    </style>
</head>
<body>
    <div class="email-container">
        <div class="header">
            <h1 class="company-name">Momentum Business Solutions</h1>
            <p class="tagline">Where Strategy Meets Execution</p>
        </div>

        <div class="greeting">
            Thank you, %s!
        </div>

        <div class="main-content">
            <p>We sincerely appreciate you taking the time to reach out to Momentum Business Solutions. Your inquiry about our financial management services has been received and is very important to us.</p>

            <p>We understand that managing your business finances can be complex, and we're here to handle the bookkeeping, payroll, and reporting so you can focus on what you do best - growing your business.</p>
        </div>

        <div class="timeline-box">
            <h3>What Happens Next?</h3>
            <p><strong>Within 24 hours:</strong> Cade from our team will personally review your submission and reach out to discuss your specific needs and how we can best support your business goals.</p>
        </div>

        <div class="contact-info">
            <h3>In the Meantime</h3>
            <p>If you have any urgent questions or would like to speak with us immediately, please don't hesitate to reach out:</p>
            <div class="contact-detail"><strong>Email:</strong> cade@momentumbusiness.org</div>
            <div class="contact-detail"><strong>Phone:</strong> (509) 554-8022</div>
        </div>

        <div class="main-content">
            <p>We look forward to the opportunity to partner with you and help your business achieve its financial goals.</p>

            <p>Best regards,<br>
            <strong>The Momentum Business Solutions Team</strong></p>
        </div>

        <div class="footer">
            <p><strong>Momentum Business Solutions</strong></p>
            <p>QuickBooks Online | Payroll Processing | Financial Consulting | Strategic Planning</p>
            <p>Email: cade@momentumbusiness.org | Phone: (509) 554-8022</p>
        </div>
    </div>
</body>
</html>`, form.FirstName)

	textBody := fmt.Sprintf(`Thank you, %s!

We sincerely appreciate you taking the time to reach out to Momentum Business Solutions. Your inquiry about our financial management services has been received and is very important to us.

We understand that managing your business finances can be complex, and we're here to handle the bookkeeping, payroll, and reporting so you can focus on what you do best - growing your business.

WHAT HAPPENS NEXT?
Within 24 hours: Cade from our team will personally review your submission and reach out to discuss your specific needs and how we can best support your business goals.

IN THE MEANTIME:
If you have any urgent questions or would like to speak with us immediately, please don't hesitate to reach out:

Email: cade@momentumbusiness.org
Phone: (509) 554-8022

We look forward to the opportunity to partner with you and help your business achieve its financial goals.

Best regards,
The Momentum Business Solutions Team

---
Momentum Business Solutions
Where Strategy Meets Execution

QuickBooks Online | Payroll Processing | Financial Consulting | Strategic Planning
Email: cade@momentumbusiness.org | Phone: (509) 554-8022
`, form.FirstName)

	email := PostmarkEmail{
		From:          from,
		To:            form.Email,
		Subject:       "Thank you for your interest in Momentum Business Solutions",
		TextBody:      textBody,
		HtmlBody:      htmlBody,
		MessageStream: "outbound",
	}

	return sendEmail(token, email)
}
