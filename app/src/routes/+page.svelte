<script lang="ts">
	import Hero from '$lib/hero.svelte';
	import Features from '$lib/features.svelte';
	import Testimonials from '$lib/testimonials.svelte';
	import Cleanup from '$lib/cleanup.svelte';
	import Faq from '$lib/faq.svelte';
	import Blog from '$lib/blog.svelte';
	import CallToAction from '$lib/cta.svelte';

	// Get SEO data from load function
    let { data }: PageProps = $props();
    const { seo, structuredData } = data;
</script>

<svelte:head>
    <!-- Primary Meta Tags -->
    <title>{seo.title}</title>
    <meta name="title" content={seo.title} />
    <meta name="description" content={seo.description} />
    <meta name="keywords" content={seo.keywords.join(', ')} />
    <meta name="author" content={seo.businessInfo.name} />
    <meta name="robots" content="index, follow" />
    <link rel="canonical" href={seo.canonical} />

    <!-- Open Graph / Facebook -->
    <meta property="og:type" content="website" />
    <meta property="og:url" content={seo.canonical} />
    <meta property="og:title" content={seo.title} />
    <meta property="og:description" content={seo.description} />
    <meta property="og:image" content={seo.ogImage} />
    <meta property="og:image:width" content="1200" />
    <meta property="og:image:height" content="630" />
    <meta property="og:site_name" content={seo.businessInfo.name} />
    <meta property="og:locale" content="en_US" />

    <!-- Twitter/X -->
    <meta property="twitter:card" content="summary_large_image" />
    <meta property="twitter:url" content={seo.canonical} />
    <meta property="twitter:title" content={seo.title} />
    <meta property="twitter:description" content={seo.description} />
    <meta property="twitter:image" content={seo.ogImage} />

    <!-- Business-specific Meta Tags -->
    <meta name="geo.region" content="US-{seo.businessInfo.address.region}" />
    <meta name="geo.placename" content={seo.businessInfo.address.locality} />
    <meta name="geo.position" content="{seo.businessInfo.address.coordinates.lat};{seo.businessInfo.address.coordinates.lng}" />
    <meta name="ICBM" content="{seo.businessInfo.address.coordinates.lat}, {seo.businessInfo.address.coordinates.lng}" />

    <!-- Structured Data -->
    {@html `<script type="application/ld+json">${JSON.stringify(structuredData)}</script>`}
    
    <!-- Preconnect to external domains for performance -->
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
</svelte:head>

<Hero />
<!-- <Testimonials /> -->
<Features />
<Cleanup />
<Faq />
<!-- <Blog /> -->
<CallToAction />
