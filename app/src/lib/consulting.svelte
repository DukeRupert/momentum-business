<script>
	let {
		serviceType = 'Financial Consulting & FP&A',
		description = 'Need expert financial guidance? Momentum offers one-on-one consulting for business owners looking to improve financial planning, cash flow forecasting, and business strategy.',
		ctaText = 'Book a Session Today',
		ctaHref = '/contact?service=consulting',
		price = '$150/hour',
		showPrice = true,
		showServices = true,
		services = [
			{
				name: 'Financial Planning & Analysis (FP&A)',
				description: 'Strategic insights to drive profitability'
			},
			{
				name: 'Budgeting & Forecasting',
				description: 'Helping you plan for sustainable growth'
			},
			{
				name: 'Cash Flow Management',
				description: 'Optimizing income, expenses, and working capital'
			},
			{
				name: 'Business Strategy & Performance Review',
				description: 'Identifying opportunities to scale efficiently'
			}
		],
		closingMessage = "Book a session today to gain clarity on your business finances and take control of your company's financial future!",
		variant = 'primary' // "primary" or "secondary"
	} = $props();

	// Determine styling based on variant
	let cardClasses = $derived(
		variant === 'primary'
			? 'text-primary-600 ring-primary-200 hover:ring-primary-300 focus-visible:outline-primary-600'
			: 'text-gray-600 ring-gray-200 hover:ring-gray-300 focus-visible:outline-gray-600'
	);
</script>

<div class="rounded-3xl mt-8 p-8 ring-1 ring-gray-900/10 sm:p-10">
	<div class="lg:flex lg:items-start lg:gap-x-8">
		<div class="lg:flex-auto">
			<!-- Header -->
			<div class="mb-4 flex items-baseline gap-x-3">
				<h3 class="text-subhead font-display-semibold text-primary-600">
					{serviceType}
				</h3>
				{#if showPrice}
					<span class="text-body font-primary-medium text-gray-500">
						{price}
					</span>
				{/if}
			</div>

			<!-- Description -->
			<p class="text-body mb-6 text-gray-600">
				{description}
			</p>

			<!-- Services List -->
			{#if showServices && services.length > 0}
				<div class="mb-6">
					<p class="text-body font-primary-semibold mb-4 text-gray-900">
						For {price}, we help with:
					</p>
					<ul class="space-y-3">
						{#each services as service}
							<li class="flex items-start gap-3">
								<svg
									class="text-primary-600 mt-0.5 size-5 flex-none"
									fill="currentColor"
									viewBox="0 0 20 20"
								>
									<path
										fill-rule="evenodd"
										d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z"
										clip-rule="evenodd"
									/>
								</svg>
								<div>
									<span class="text-body font-primary-semibold text-gray-900">
										{service.name}
									</span>
									<span class="text-body text-gray-600"> – {service.description}</span>
								</div>
							</li>
						{/each}
					</ul>
				</div>
			{/if}

			<!-- Closing Message -->
			{#if closingMessage}
				<p class="text-body font-primary-medium mb-8 text-gray-700">
					{closingMessage}
				</p>
			{/if}
		</div>

		<!-- CTA Button -->
		<div class="lg:mt-0 lg:flex-none">
			<a
				href={ctaHref}
				class="bg-primary-600 text-body font-primary-semibold hover:bg-primary-500 focus-visible:outline-primary-600 inline-flex items-center rounded-md px-4 py-2.5 text-white shadow-sm transition-colors focus-visible:outline-2 focus-visible:outline-offset-2"
			>
				{ctaText} <span aria-hidden="true" class="ml-2">→</span>
			</a>
		</div>
	</div>
</div>
