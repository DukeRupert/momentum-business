<script>
  let {
    sectionTitle = "Ready for comprehensive business support?",
    sectionDescription = "Move beyond basic bookkeeping to proactive leadership with our complete business partnership package.",
    serviceTitle = "Complete Business Support",
    serviceDescription = "Running a small business comes with no shortage of challenges—balancing finances, managing a team, and planning for growth. We help you move from reactive problem-solving to proactive leadership.",
    pricingLabel = "Monthly partnership investment",
    price = "$2,500",
    currency = "USD",
    ctaText = "Schedule Discovery Call",
    ctaHref = "/contact?service=completeSupport",
    disclaimer = "Custom packages available for businesses with unique needs",
    features = [
      "Monthly bookkeeping & bank reconciliation",
      "Financial reporting and dashboards", 
      "Monthly CFO-style strategy calls",
      "Budgeting, forecasting & cash flow planning",
      "QuickBooks optimization and cleanup",
      "Pricing & profitability analysis",
      "Employee onboarding/offboarding guidance",
      "Payroll support (up to 10 employees)",
      "Labor cost analysis & org chart review",
      "HR process consulting & compliance",
      "Goal setting & KPI tracking",
      "Tax season readiness support"
    ]
  } = $props();
</script>

<section id="complete-support" class="bg-white py-24 sm:py-32">
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <div class="mx-auto max-w-4xl sm:text-center">
      <h2 class="text-display-sm font-display-bold tracking-tight text-pretty text-gray-900 sm:text-display sm:text-balance">
        {sectionTitle}
      </h2>
      <p class="mx-auto mt-6 max-w-2xl text-body-lg font-primary-medium text-pretty text-gray-500 sm:text-subhead">
        {sectionDescription}
      </p>
    </div>
    
    <div class="mx-auto mt-16 max-w-2xl rounded-3xl ring-1 ring-gray-200 sm:mt-20 lg:mx-0 lg:flex lg:max-w-none">
      <!-- Service Details -->
      <div class="p-8 sm:p-10 lg:flex-auto">
        <h3 class="text-headline font-display-semibold tracking-tight text-gray-900">
          {serviceTitle}
        </h3>
        <p class="mt-6 text-body text-gray-600">
          {serviceDescription}
        </p>
        
        <div class="mt-10 flex items-center gap-x-4">
          <h4 class="flex-none text-caption font-primary-semibold text-primary-600">
            What's included each month
          </h4>
          <div class="h-px flex-auto bg-gray-100"></div>
        </div>
        
        <ul role="list" class="mt-8 grid grid-cols-1 gap-4 text-body text-gray-600 sm:grid-cols-2 sm:gap-6">
          {#each features as feature}
            <li class="flex gap-x-3">
              <svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" class="h-6 w-5 flex-none text-primary-600">
                <path d="M16.704 4.153a.75.75 0 0 1 .143 1.052l-8 10.5a.75.75 0 0 1-1.127.075l-4.5-4.5a.75.75 0 0 1 1.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 0 1 1.05-.143Z" clip-rule="evenodd" fill-rule="evenodd" />
              </svg>
              {feature}
            </li>
          {/each}
        </ul>
      </div>
      
      <!-- Pricing Card -->
      <div class="-mt-2 p-2 lg:mt-0 lg:w-full lg:max-w-md lg:shrink-0">
        <div class="rounded-2xl bg-gray-50 py-10 text-center ring-1 ring-inset ring-gray-900/5 lg:flex lg:flex-col lg:justify-center lg:py-16">
          <div class="mx-auto max-w-xs px-8">
            <p class="text-body font-primary-semibold text-gray-600">
              {pricingLabel}
            </p>
            <p class="mt-6 flex items-baseline justify-center gap-x-2">
              <span class="text-display font-display-bold tracking-tight text-gray-900">
                {price}
              </span>
              <span class="text-caption font-primary-semibold tracking-wide text-gray-600">
                {currency}
              </span>
            </p>
            <a 
              href={ctaHref}
              class="mt-10 block w-full rounded-md bg-primary-600 px-3 py-2 text-center text-body font-primary-semibold text-white shadow-xs hover:bg-primary-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
            >
              {ctaText}
            </a>
            <p class="mt-6 text-caption text-gray-600">
              {disclaimer}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>