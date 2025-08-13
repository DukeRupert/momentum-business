<script>
	import Seo from '$lib/seo.svelte';
	import SectionHeader from '$lib/header.svelte';
	import Cleanup from '$lib/cleanup.svelte';
	import Premium from '$lib/premium.svelte';
	import Consulting from '$lib/consulting.svelte';

	// Services data with new Momentum Business Solutions packages
	const servicesData = [
		{
			id: 'essentials',
			name: 'Essentials Package',
			tier: 'Tier 1',
			price: '$750',
			period: '/month',
			timeCommitment: '2–3 hours/month',
			description: 'For businesses that need clean, accurate financials and basic insights.',
			features: [
				'Monthly bookkeeping and account reconciliation (up to 3 bank/credit accounts)',
				'Basic financial reports: Profit & Loss, Balance Sheet, Cash Flow',
				'Month-end close and error checks',
				'QuickBooks Online maintenance and categorization',
				'Email support (48–72 hr response time)'
			],
			bestFor: 'Solo operators or small teams looking for reliable, organized books.',
			popular: false,
			ctaText: 'Get Started',
			ctaHref: '/contact?service=essentials'
		},
		{
			id: 'growth-strategy',
			name: 'Growth Strategy Package',
			tier: 'Tier 2',
			price: '$1,250',
			period: '/month',
			timeCommitment: '4–6 hours/month',
			description: 'For growing businesses that want financial strategy and light HR support.',
			features: [
				'Everything in Tier 1, plus:',
				'Monthly 60-minute strategic CFO advisory call',
				'Forecasting and budgeting templates',
				'Labor cost tracking and analysis',
				'Payroll support for up to 10 employees',
				'Employee onboarding/offboarding guidance',
				'Priority email support (24–48 hr response time)'
			],
			bestFor: 'Businesses ready to scale with a proactive financial partner.',
			popular: true,
			ctaText: 'Most Popular',
			ctaHref: '/contact?service=growthStrategy'
		}
	];

	// Process steps
	const processSteps = [
		{
			step: '1',
			title: 'Discovery Call',
			description:
				"We'll discuss your current financial situation, pain points, and goals to determine the best fit."
		},
		{
			step: '2',
			title: 'Onboarding',
			description:
				"We'll set up secure access to your systems, review your books, and establish processes — typically within 1-2 weeks."
		},
		{
			step: '3',
			title: 'Ongoing Support',
			description:
				'Monthly financial reports, regular check-ins, and ongoing support to keep your finances organized and strategic.'
		}
	];
</script>

<!-- Simple SEO tags -->
<Seo
	title="Services & Pricing"
	description="Monthly bookkeeping packages starting at $750/month. QuickBooks setup, payroll processing, and financial consulting for growing businesses. Based in Richland, WA."
	canonical="/services"
/>

<!-- Page Header -->
<SectionHeader
	title="Monthly Service Packages"
	description="Comprehensive financial solutions designed to give you clarity, confidence, and peace of mind so you can focus on growing your business."
/>

<!-- Services Grid -->
<section class="bg-white py-12 sm:py-24">
	<div class="mx-auto max-w-7xl px-6 lg:px-8">
		<div class="mx-auto grid max-w-2xl grid-cols-1 gap-8 lg:mx-0 lg:max-w-none lg:grid-cols-2">
			{#each servicesData as service}
				<div
					id={service.id}
					class="relative rounded-2xl border border-gray-200 p-8 shadow-sm transition-shadow hover:shadow-md {service.popular
						? 'ring-primary-600 ring-2'
						: ''}"
				>
					{#if service.popular}
						<div class="absolute -top-4 left-1/2 -translate-x-1/2">
							<span
								class="bg-primary-600 text-kicker font-primary-semibold inline-flex items-center rounded-full px-4 py-1 text-white"
							>
								Most Popular
							</span>
						</div>
					{/if}

					<div class="mb-4">
						<span
							class="bg-primary-50 text-primary-700 inline-flex items-center rounded-md px-2 py-1 text-xs font-medium"
						>
							{service.tier}
						</span>
					</div>

					<div class="mb-4 flex items-baseline justify-between">
						<h3 class="text-subhead font-display-semibold text-gray-900">
							{service.name}
						</h3>
					</div>

					<div class="mb-4">
						<div class="flex items-baseline">
							<span class="text-headline font-display-bold text-gray-900">
								{service.price}
							</span>
							<span class="text-body ml-1 text-gray-600">
								{service.period}
							</span>
						</div>
						<p class="text-caption text-primary-600 mt-1 font-medium">
							Estimated Time Commitment: {service.timeCommitment}
						</p>
					</div>

					<p class="text-body mb-6 text-gray-600">
						{service.description}
					</p>

					<div class="mb-6">
						<h4 class="mb-3 text-sm font-semibold text-gray-900">Includes:</h4>
						<ul class="space-y-3">
							{#each service.features as feature}
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
									<span class="text-body text-gray-700">{feature}</span>
								</li>
							{/each}
						</ul>
					</div>

					<div class="mb-8 rounded-lg bg-gray-50 p-4">
						<h4 class="mb-2 text-sm font-semibold text-gray-900">Best for:</h4>
						<p class="text-caption text-gray-600">
							{service.bestFor}
						</p>
					</div>

					<div class="mt-auto">
						<a
							href={service.ctaHref}
							class="text-body font-primary-semibold block w-full rounded-md px-4 py-2.5 text-center transition-colors {service.popular
								? 'bg-primary-600 hover:bg-primary-500 text-white'
								: 'text-primary-600 ring-primary-600 hover:bg-primary-50 bg-white ring-1'}"
						>
							{service.ctaText}
						</a>
					</div>
				</div>
			{/each}
		</div>
		<Consulting />
	</div>
</section>

<Premium />

<Cleanup />

<!-- How It Works Process -->
<section class="bg-gray-50 py-24 sm:py-32">
	<div class="mx-auto max-w-7xl px-6 lg:px-8">
		<div class="mx-auto max-w-2xl lg:text-center">
			<h2 class="text-headline font-display-semibold tracking-tight text-gray-900">How it works</h2>
			<p class="text-body-lg mt-4 text-gray-600">
				Getting started is simple. We handle the complexity so you can focus on your business.
			</p>
		</div>

		<div class="mx-auto mt-16 max-w-2xl lg:max-w-none">
			<div class="grid max-w-xl grid-cols-1 gap-8 lg:max-w-none lg:grid-cols-3">
				{#each processSteps as step}
					<div class="text-center">
						<div
							class="bg-primary-600 mx-auto flex size-16 items-center justify-center rounded-full"
						>
							<span class="text-subhead font-display-bold text-white">
								{step.step}
							</span>
						</div>
						<h3 class="text-subhead font-display-semibold mt-6 text-gray-900">
							{step.title}
						</h3>
						<p class="text-body mt-4 text-gray-600">
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
				Interested in partnering with Momentum Business Solutions?
			</h2>
			<p class="text-body-lg mt-4 text-gray-600">
				Let's schedule a discovery call to explore which package is right for your business.
			</p>
			<div class="mt-8 flex items-center justify-center gap-6">
				<a
					href="/contact"
					class="bg-primary-600 text-body font-primary-semibold shadow-xs hover:bg-primary-500 focus-visible:outline-primary-600 rounded-md px-6 py-3 text-white focus-visible:outline-2 focus-visible:outline-offset-2"
				>
					Schedule Discovery Call
				</a>
				<a href="/about" class="text-body font-primary-semibold text-gray-900">
					Learn more about us <span aria-hidden="true">→</span>
				</a>
			</div>

			<div class="mt-8 border-t border-gray-200 pt-8">
				<div class="text-center">
					<p class="text-body font-primary-semibold mb-2 text-gray-900">Cade Maldonado</p>
					<p class="text-caption text-gray-600">CEO/Owner | QuickBooks ProAdvisor Certified</p>
					<p class="text-caption mt-1 text-gray-600">
						<a
							href="mailto:cade@momentumbusiness.org"
							class="text-primary-600 hover:text-primary-500"
						>
							cade@momentumbusiness.org
						</a>
					</p>
				</div>
			</div>
		</div>
	</div>
</section>
