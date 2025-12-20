<script lang="ts">
    // Simple SEO component for consistent meta tags across pages
    let {
        title = '',
        description = '',
        canonical = '',
        noindex = false,
        ogImage = 'https://momentumbusiness.org/og-image.jpg'
    } = $props();

    // Default business info
    const businessName = 'Momentum Business Solutions';
    const baseUrl = 'https://momentumbusiness.org';
    
    // Generate full title
    const fullTitle = title ? `${title} | ${businessName}` : `${businessName} - QuickBooks Experts & Financial Consulting`;
    
    // Generate canonical URL
    const canonicalUrl = canonical ? `${baseUrl}${canonical}` : baseUrl;
    
    // Default description if none provided
    const metaDescription = description || 'Expert bookkeeping, payroll processing, and financial consulting for growing businesses. QuickBooks ProAdvisor certified in Richland, WA.';
</script>

<svelte:head>
    <!-- Primary Meta Tags -->
    <title>{fullTitle}</title>
    <meta name="description" content={metaDescription} />
    <meta name="robots" content={noindex ? 'noindex, nofollow' : 'index, follow'} />
    <link rel="canonical" href={canonicalUrl} />

    <!-- Open Graph -->
    <meta property="og:title" content={fullTitle} />
    <meta property="og:description" content={metaDescription} />
    <meta property="og:url" content={canonicalUrl} />
    <meta property="og:image" content={ogImage} />
    <meta property="og:type" content="website" />
    <meta property="og:site_name" content={businessName} />

    <!-- Twitter -->
    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:title" content={fullTitle} />
    <meta name="twitter:description" content={metaDescription} />
    <meta name="twitter:image" content={ogImage} />
</svelte:head>