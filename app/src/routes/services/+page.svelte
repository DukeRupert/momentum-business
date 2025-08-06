<script>
  import SectionHeader from '$lib/header.svelte';

  // Services data with pricing and features
  const servicesData = [
    {
      id: 'bookkeeping',
      name: "Bookkeeping",
      price: "$750",
      period: "/month",
      description: "For businesses that want clean, up-to-date books and peace of mind.",
      features: [
        "Monthly transaction categorization & reconciliation",
        "Monthly financial reports (Profit & Loss, Balance Sheet, Cash Flow)",
        "Monthly business review meeting",
        "Ongoing email support for financial questions"
      ],
      popular: false,
      ctaText: "Get Started",
      ctaHref: "/contact?service=bookkeeping"
    },
    {
      id: 'payroll',
      name: "Bookkeeping + Payroll",
      price: "$1,250",
      period: "/month",
      description: "Everything in the Bookkeeping package, plus:",
      features: [
        "All Bookkeeping services included",
        "Payroll processing for your employees",
        "Payroll tax filings and year-end forms (W-2s, 1099s)",
        "Employee direct deposit setup",
        "Seamless integration with QuickBooks Online Payroll"
      ],
      popular: true,
      ctaText: "Most Popular",
      ctaHref: "/contact?service=payroll"
    },
    {
      id: 'consulting',
      name: "Financial Consulting",
      price: "$150",
      period: "/hour",
      description: "For business owners seeking deeper insight and strategic support.",
      features: [
        "Budgeting and forecasting",
        "Cash flow planning",
        "Profitability analysis",
        "Financial planning & analysis (FP&A)"
      ],
      popular: false,
      ctaText: "Book Session",
      ctaHref: "/contact?service=consulting"
    },
    {
      id: 'fractional-cfo',
      name: "Fractional CFO Services",
      price: "Custom",
      period: "pricing",
      description: "For growing businesses ready for executive-level financial strategy without the full-time cost.",
      features: [
        "Monthly or quarterly financial leadership",
        "Strategic financial planning",
        "Business model evaluation",
        "Custom KPI development and tracking",
        "Forecasting and long-range planning",
        "Support for investor or lender conversations"
      ],
      popular: false,
      ctaText: "Learn More",
      ctaHref: "/contact?service=cfo",
      note: "Ideal for companies preparing for growth, seeking funding, or needing senior-level insight to scale efficiently."
    }
  ];

  // Process steps
  const processSteps = [
    {
      step: "1",
      title: "Discovery Call",
      description: "We'll discuss your current financial situation, pain points, and goals to determine the best fit."
    },
    {
      step: "2",
      title: "Onboarding",
      description: "We'll set up secure access to your systems, review your books, and establish processes — typically within 1-2 weeks."
    },
    {
      step: "3",
      title: "Ongoing Support",
      description: "Monthly financial reports, regular check-ins, and ongoing support to keep your finances organized and strategic."
    }
  ];
</script>

<!-- Page Header -->
<SectionHeader 
  title="Services" 
  description="Comprehensive financial solutions designed to give you clarity, confidence, and peace of mind so you can focus on growing your business."
/>

<!-- Services Grid -->
<section class="bg-white py-24 sm:py-32">
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <div class="mx-auto grid max-w-2xl grid-cols-1 gap-8 lg:mx-0 lg:max-w-none lg:grid-cols-2">
      {#each servicesData as service}
        <div id={service.id} class="relative rounded-2xl border border-gray-200 p-8 shadow-sm hover:shadow-md transition-shadow {service.popular ? 'ring-2 ring-primary-600' : ''}">
          {#if service.popular}
            <div class="absolute -top-4 left-1/2 -translate-x-1/2">
              <span class="inline-flex items-center rounded-full bg-primary-600 px-4 py-1 text-kicker font-primary-semibold text-white">
                Most Popular
              </span>
            </div>
          {/if}
          
          <div class="flex items-baseline justify-between">
            <h3 class="text-subhead font-display-semibold text-gray-900">
              {service.name}
            </h3>
            <div class="text-right">
              <span class="text-headline font-display-bold text-gray-900">
                {service.price}
              </span>
              <span class="text-body text-gray-600">
                {service.period}
              </span>
            </div>
          </div>
          
          <p class="mt-4 text-body text-gray-600">
            {service.description}
          </p>
          
          <ul class="mt-8 space-y-3">
            {#each service.features as feature}
              <li class="flex items-start gap-3">
                <svg class="size-5 flex-none text-primary-600 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z" clip-rule="evenodd" />
                </svg>
                <span class="text-body text-gray-700">{feature}</span>
              </li>
            {/each}
          </ul>
          
          {#if service.note}
            <p class="mt-6 text-caption text-gray-500 italic">
              {service.note}
            </p>
          {/if}
          
          <div class="mt-8">
            <a 
              href={service.ctaHref}
              class="block w-full rounded-md text-center px-4 py-2.5 text-body font-primary-semibold transition-colors {service.popular ? 'bg-primary-600 text-white hover:bg-primary-500' : 'bg-white text-primary-600 ring-1 ring-primary-600 hover:bg-primary-50'}"
            >
              {service.ctaText}
            </a>
          </div>
        </div>
      {/each}
    </div>
  </div>
</section>

<!-- How It Works Process -->
<section class="bg-gray-50 py-24 sm:py-32">
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <div class="mx-auto max-w-2xl lg:text-center">
      <h2 class="text-headline font-display-semibold tracking-tight text-gray-900">
        How it works
      </h2>
      <p class="mt-4 text-body-lg text-gray-600">
        Getting started is simple. We handle the complexity so you can focus on your business.
      </p>
    </div>
    
    <div class="mx-auto mt-16 max-w-2xl lg:max-w-none">
      <div class="grid max-w-xl grid-cols-1 gap-8 lg:max-w-none lg:grid-cols-3">
        {#each processSteps as step}
          <div class="text-center">
            <div class="mx-auto flex size-16 items-center justify-center rounded-full bg-primary-600">
              <span class="text-subhead font-display-bold text-white">
                {step.step}
              </span>
            </div>
            <h3 class="mt-6 text-subhead font-display-semibold text-gray-900">
              {step.title}
            </h3>
            <p class="mt-4 text-body text-gray-600">
              {step.description}
            </p>
          </div>
        {/each}
      </div>
    </div>
  </div>
</section>

<!-- Call to Action -->
<section class="bg-white py-24 sm:py-32">
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <div class="mx-auto max-w-2xl text-center">
      <h2 class="text-headline font-display-semibold tracking-tight text-gray-900">
        Ready to get started?
      </h2>
      <p class="mt-4 text-body-lg text-gray-600">
        Let's discuss which service is the best fit for your business. Schedule a free consultation today.
      </p>
      <div class="mt-8 flex items-center justify-center gap-6">
        <a 
          href="/contact" 
          class="rounded-md bg-primary-600 px-6 py-3 text-body font-primary-semibold text-white shadow-xs hover:bg-primary-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
        >
          Schedule Free Consultation
        </a>
        <a href="/about" class="text-body font-primary-semibold text-gray-900">
          Learn more about us <span aria-hidden="true">→</span>
        </a>
      </div>
    </div>
  </div>
</section>