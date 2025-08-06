<script>
  const ctaData = {
    headline: "Ready to take control of your finances?",
    description: "Stop losing sleep over QuickBooks and cash flow management. Join hundreds of business owners who've found peace of mind with our expert financial services.",
    primaryButton: {
      text: "Get started today",
      href: "#"
    },
    secondaryButton: {
      text: "View our packages",
      href: "#"
    }
  };
</script>

<div class="bg-primary-700">
  <div class="px-6 py-24 sm:px-6 sm:py-32 lg:px-8">
    <div class="mx-auto max-w-2xl text-center">
      <h2 class="text-headline-lg font-display-bold tracking-tight text-balance text-white sm:text-display-sm">
        {ctaData.headline}
      </h2>
      <p class="mx-auto mt-6 max-w-xl text-body-lg text-pretty text-primary-200">
        {ctaData.description}
      </p>
      <div class="mt-10 flex items-center justify-center gap-x-6">
        <a 
          href={ctaData.primaryButton.href} 
          class="rounded-md bg-white px-3.5 py-2.5 text-body font-primary-semibold text-primary-600 shadow-xs hover:bg-primary-50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
        >
          {ctaData.primaryButton.text}
        </a>
        <a href={ctaData.secondaryButton.href} class="text-body font-primary-semibold text-white">
          {ctaData.secondaryButton.text}
          <span aria-hidden="true">→</span>
        </a>
      </div>
    </div>
  </div>
</div>