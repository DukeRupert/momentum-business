export interface ContactFormData {
    'first-name': string;
    'last-name': string;
    email: string;
    'phone-number': string;
    'annual-revenue': string;
    services: string[];
    message: string;
}

export interface ThankYouEmailOptions {
    formData: ContactFormData;
    postmarkToken: string;
    fromEmail?: string;
}

// Available services
export type Service =
    | 'essentials'
    | 'payroll'
    | 'planning'
    | 'tax-prep'
    | 'bookkeeping';

export async function sendContactFormEmail(formData: ContactFormData, postmarkToken: string, to: string, from: string) {
    // Helper function to format service names
    function formatService(service: Service) {
        const serviceMap = {
            'essentials': 'QuickBooks Essentials',
            'payroll': 'Payroll Management',
            'planning': 'Business Planning',
            'tax-prep': 'Tax Preparation',
            'bookkeeping': 'Full Bookkeeping'
        };
        return serviceMap[service] || service;
    }

    // Helper function to format annual revenue
    function formatRevenue(revenue: string) {
        const revenueMap = {
            'under-100k': 'Under $100,000',
            '100k-500k': '$100,000 - $500,000',
            '500k-1m': '$500,000 - $1,000,000',
            '1m-plus': '$1,000,000+'
        };
        return revenueMap[revenue] || revenue;
    }

    // Helper function to get service class for styling
    function getServiceClass(service: Service) {
        const classMap = {
            'essentials': 'bookkeeping',
            'payroll': 'payroll',
            'planning': 'consulting',
            'tax-prep': 'cleanup',
            'bookkeeping': 'bookkeeping'
        };
        return classMap[service] || 'bookkeeping';
    }

    // Get current timestamp in PST
    const timestamp = new Date().toLocaleString('en-US', {
        weekday: 'long',
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: 'numeric',
        minute: '2-digit',
        timeZone: 'America/Los_Angeles',
        timeZoneName: 'short'
    });

    // Create the HTML email body
    const htmlBody = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>New Lead: Contact Form Submission - ${formData['first-name']} ${formData['last-name']}</title>
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
        @media (max-width: 600px) {
            .info-grid {
                grid-template-columns: 1fr;
            }
            body {
                padding: 10px;
            }
            .email-container {
                padding: 24px;
            }
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
            <strong>Submitted:</strong> ${timestamp} | <strong>Source:</strong> Website Contact Form | <strong>Form Version:</strong> v2.1
        </div>

        <div class="section">
            <h2>Contact Information</h2>
            <div class="info-grid">
                <div class="info-item">
                    <div class="info-label">Full Name</div>
                    <div class="info-value">${formData['first-name']} ${formData['last-name']}</div>
                </div>
                <div class="info-item">
                    <div class="info-label">Email Address</div>
                    <div class="info-value">${formData.email}</div>
                </div>
                <div class="info-item">
                    <div class="info-label">Phone Number</div>
                    <div class="info-value">${formData['phone-number']}</div>
                </div>
                <div class="info-item revenue-highlight">
                    <div class="info-label">Annual Revenue</div>
                    <div class="info-value">${formatRevenue(formData['annual-revenue'])}</div>
                </div>
            </div>
        </div>

        <div class="section">
            <h2>Services of Interest</h2>
            <div class="services-list">
                <div class="info-label" style="margin-bottom: 12px;">Client selected the following services:</div>
                ${formData.services.map(service => `<span class="service-tag ${getServiceClass(service)}">${formatService(service)}</span>`).join('')}
            </div>
        </div>

        ${formData.message ? `
        <div class="section">
            <h2>Client Message</h2>
            <div class="message-box">
                "${formData.message}"
            </div>
        </div>
        ` : ''}

        <div class="footer">
            <p><strong>Momentum Business Solutions</strong></p>
            <p>QuickBooks Online | Payroll Processing | Financial Consulting | Strategic Planning</p>
            <p>Email: cade@momentumbusiness.org | Phone: (509) 554-8022</p>
        </div>
    </div>
</body>
</html>`;

    // Create plain text version
    const textBody = `
NEW QUALIFIED LEAD - Momentum Business Solutions
===============================================

SUBMISSION DETAILS:
Submitted: ${timestamp}
Source: Website Contact Form
Form Version: v2.1

CONTACT INFORMATION:
-------------------
Name: ${formData['first-name']} ${formData['last-name']}
Email: ${formData.email}
Phone: ${formData['phone-number']}
Annual Revenue: ${formatRevenue(formData['annual-revenue'])}

SERVICES OF INTEREST:
--------------------
Client selected the following services:
${formData.services.map(service => `• ${formatService(service)}`).join('\n')}

${formData.message ? `CLIENT MESSAGE:
---------------
"${formData.message}"

` : ''}CONTACT INFORMATION:
-------------------
Momentum Business Solutions
Where Strategy Meets Execution

QuickBooks Online | Payroll Processing | Financial Consulting | Strategic Planning
Email: cade@momentumbusiness.org
Phone: (509) 554-8022

---
This email was generated from your website contact form.
`;

    // Send email via Postmark
    try {
        console.log(`Api token: ${postmarkToken}`)
        const response = await fetch('https://api.postmarkapp.com/email', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'X-Postmark-Server-Token': postmarkToken
            },
            body: JSON.stringify({
                From: from,
                To: to,
                Subject: `New Lead: Contact Form Submission - ${formData['first-name']} ${formData['last-name']}`,
                TextBody: textBody,
                HtmlBody: htmlBody,
                MessageStream: 'outbound'
            })
        });

        if (!response.ok) {
            console.log(response)
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const result = await response.json();
        return { success: true, data: result };
    } catch (error: any) {
        console.error('Error sending email:', error);
        return { success: false, error: error.message };
    }
}

export async function sendThankYouEmail({
    formData,
    postmarkToken,
    fromEmail = 'cade@momentumbusiness.org'
}: ThankYouEmailOptions) {

    // Helper function to format service names for customer-facing text
    function formatServiceForCustomer(service: Service): string {
        const serviceMap: Record<Service, string> = {
            'essentials': 'QuickBooks Setup & Management',
            'payroll': 'Payroll Processing',
            'planning': 'Strategic Business Planning',
            'tax-prep': 'Tax Preparation Services',
            'bookkeeping': 'Full-Service Bookkeeping'
        };
        return serviceMap[service];
    }

    // Create HTML email body
    const htmlBody = `
<!DOCTYPE html>
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
        .services-summary {
            background: #f0fdf4;
            border: 2px solid #22c55e;
            border-radius: 8px;
            padding: 20px;
            margin: 24px 0;
        }
        .services-summary h3 {
            color: #15803d;
            margin: 0 0 12px 0;
            font-size: 16px;
            font-weight: 600;
        }
        .service-item {
            color: #166534;
            margin: 4px 0;
            font-weight: 500;
        }
        .service-item::before {
            content: "✓ ";
            color: #22c55e;
            font-weight: bold;
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
        .cta-button {
            display: inline-block;
            background: #53945c;
            color: white;
            padding: 12px 24px;
            border-radius: 6px;
            text-decoration: none;
            font-weight: 600;
            margin: 16px 0;
        }
        @media (max-width: 600px) {
            body {
                padding: 10px;
            }
            .email-container {
                padding: 24px;
            }
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
            Thank you, ${formData['first-name']}!
        </div>

        <div class="main-content">
            <p>We sincerely appreciate you taking the time to reach out to Momentum Business Solutions. Your inquiry about our financial management services has been received and is very important to us.</p>

            <p>We understand that managing your business finances can be complex, and we're here to help streamline those processes so you can focus on what you do best – growing your business.</p>
        </div>

        <div class="services-summary">
            <h3>Services You're Interested In:</h3>
            ${formData.services.map(service =>
        `<div class="service-item">${formatServiceForCustomer(service)}</div>`
    ).join('')}
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
</html>`;

    // Create plain text version
    const textBody = `
Thank you, ${formData['first-name']}!

We sincerely appreciate you taking the time to reach out to Momentum Business Solutions. Your inquiry about our financial management services has been received and is very important to us.

We understand that managing your business finances can be complex, and we're here to help streamline those processes so you can focus on what you do best – growing your business.

SERVICES YOU'RE INTERESTED IN:
${formData.services.map(service => `• ${formatServiceForCustomer(service)}`).join('\n')}

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
`;

    // Send email via Postmark
    try {
        const response = await fetch('https://api.postmarkapp.com/email', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'X-Postmark-Server-Token': postmarkToken
            },
            body: JSON.stringify({
                From: fromEmail,
                To: formData.email,
                Subject: `Thank you for your interest in Momentum Business Solutions`,
                TextBody: textBody,
                HtmlBody: htmlBody,
                MessageStream: 'outbound'
            })
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const result = await response.json();
        return { success: true, data: result };
    } catch (error) {
        console.error('Error sending thank you email:', error);
        return { success: false, error: (error as Error).message };
    }
}