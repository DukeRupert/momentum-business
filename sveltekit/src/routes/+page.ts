import type { PageLoad } from './$types';

export const load: PageLoad = () => {
    // SEO Configuration - Centralized for easy management
    const seoData = {
        title: "Momentum Business Solutions - QuickBooks Experts & Financial Consulting",
        description: "Expert bookkeeping, payroll processing, and financial consulting for growing businesses. QuickBooks ProAdvisor certified. Serving businesses with $500K-$5M revenue. Get clarity, save time.",
        canonical: "https://momentumbusiness.org",
        ogImage: "https://momentumbusiness.org/og-image.jpg",
        keywords: [
            "QuickBooks ProAdvisor",
            "bookkeeping services", 
            "payroll processing",
            "financial consulting",
            "small business accounting",
            "Richland WA",
            "Tri-Cities Washington",
            "CFO advisory",
            "business planning",
            "monthly bookkeeping",
            "QuickBooks Online",
            "financial planning",
            "business solutions"
        ],
        businessInfo: {
            name: "Momentum Business Solutions, LLC",
            founder: "Cade Maldonado",
            phone: "(509) 554-8022",
            email: "cade@momentumbusiness.org",
            linkedIn: "https://www.linkedin.com/company/momentum-business-solutions-llc",
            address: {
                locality: "Richland",
                region: "WA", 
                country: "US",
                coordinates: {
                    lat: "46.2857",
                    lng: "-119.2844"
                }
            }
        },
        services: [
            {
                name: "QuickBooks Setup & Management",
                description: "Professional QuickBooks Online setup, maintenance, and ongoing support",
                price: "Starting at $750/month"
            },
            {
                name: "Bookkeeping Services", 
                description: "Monthly bookkeeping, account reconciliation, and financial reporting",
                price: "Included in packages"
            },
            {
                name: "Payroll Processing",
                description: "Complete payroll management for businesses with up to 10 employees", 
                price: "Up to 10 employees included"
            },
            {
                name: "Financial Consulting",
                description: "Strategic CFO advisory, budgeting, forecasting, and business planning",
                price: "$150/hour or monthly packages"
            }
        ]
    };

    // Generate structured data
    const structuredData = {
        "@context": "https://schema.org",
        "@type": ["Organization", "ProfessionalService"],
        "name": seoData.businessInfo.name,
        "description": seoData.description,
        "url": seoData.canonical,
        "logo": `${seoData.canonical}/logo.png`,
        "image": seoData.ogImage,
        "founder": {
            "@type": "Person", 
            "name": seoData.businessInfo.founder,
            "jobTitle": "CEO/Owner",
            "email": seoData.businessInfo.email
        },
        "contactPoint": {
            "@type": "ContactPoint",
            "telephone": seoData.businessInfo.phone,
            "email": seoData.businessInfo.email,
            "contactType": "customer service",
            "availableLanguage": "English"
        },
        "address": {
            "@type": "PostalAddress",
            "addressLocality": seoData.businessInfo.address.locality,
            "addressRegion": seoData.businessInfo.address.region,
            "addressCountry": seoData.businessInfo.address.country
        },
        "geo": {
            "@type": "GeoCoordinates",
            "latitude": seoData.businessInfo.address.coordinates.lat,
            "longitude": seoData.businessInfo.address.coordinates.lng
        },
        "serviceArea": {
            "@type": "Country",
            "name": "United States"
        },
        "hasOfferCatalog": {
            "@type": "OfferCatalog",
            "name": "Financial Services",
            "itemListElement": seoData.services.map(service => ({
                "@type": "Offer",
                "itemOffered": {
                    "@type": "Service",
                    "name": service.name,
                    "description": service.description
                }
            }))
        },
        "aggregateRating": {
            "@type": "AggregateRating", 
            "ratingValue": "5.0",
            "ratingCount": "15",
            "bestRating": "5"
        },
        "sameAs": [
            seoData.businessInfo.linkedIn
        ]
    };

    return {
        seo: seoData,
        structuredData
    };
};