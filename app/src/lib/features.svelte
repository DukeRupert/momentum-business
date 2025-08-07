<script lang="ts">
	const features = [
		{
			name: 'Complete bookkeeping service.',
			description:
				'Monthly transaction recording, bank reconciliation, and financial reports in QuickBooks Online. Always accurate, always on time, always tax-ready.',
			icon: 'shield-check'
		},
		{
			name: 'Full-service payroll processing.',
			description:
				'We handle payroll runs, tax filings, direct deposits, and compliance for up to 10 employees. Everything automated so you never miss a deadline.',
			icon: 'users'
		},
		{
			name: 'Strategic financial consulting.',
			description:
				'Get expert guidance on budgeting, forecasting, and cash flow optimization. Make data-driven decisions with our $150/hour FP&A consulting.',
			icon: 'chart-bar'
		}
	];

	function getIconSvg(iconName: string) {
		const icons: Record<string, string> = {
			'shield-check': `
        <svg aria-hidden="true" class="absolute top-1 left-1 size-5 text-primary-600" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M9.661 2.237a.531.531 0 0 1 .678 0 11.947 11.947 0 0 0 7.078 2.749.5.5 0 0 1 .479.425c.069.52.104 1.05.104 1.589 0 5.162-3.26 9.563-7.834 11.256a.48.48 0 0 1-.332 0C5.26 16.564 2 12.163 2 7c0-.538.035-1.069.104-1.589a.5.5 0 0 1 .48-.425 11.947 11.947 0 0 0 7.077-2.75Zm4.196 5.954a.75.75 0 0 0-1.214-.882l-3.236 4.53L7.73 9.77a.75.75 0 0 0-1.06 1.061l2.25 2.25a.75.75 0 0 0 1.137-.089l3.75-5.25Z" clip-rule="evenodd" />
        </svg>
      `,
			'users': `
        <svg aria-hidden="true" class="absolute top-1 left-1 size-5 text-primary-600" fill="currentColor" viewBox="0 0 20 20">
          <path d="M7 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6ZM14.5 9a2.5 2.5 0 1 0 0-5 2.5 2.5 0 0 0 0 5ZM1.615 16.428a1.224 1.224 0 0 1-.569-1.175 6.002 6.002 0 0 1 11.908 0c.058.467-.172.92-.57 1.174A9.953 9.953 0 0 1 7 18a9.953 9.953 0 0 1-5.385-1.572ZM14.5 16h-.106c.07-.297.088-.611.048-.933a7.47 7.47 0 0 0-1.588-3.755 4.502 4.502 0 0 1 5.874 2.636.818.818 0 0 1-.36.98A7.465 7.465 0 0 1 14.5 16Z" />
        </svg>
      `,
			'chart-bar': `
        <svg aria-hidden="true" class="absolute top-1 left-1 size-5 text-primary-600" fill="currentColor" viewBox="0 0 20 20">
          <path d="M15.5 2A1.5 1.5 0 0 1 17 3.5v13a1.5 1.5 0 0 1-1.5 1.5h-6A1.5 1.5 0 0 1 8 16.5v-6A1.5 1.5 0 0 1 9.5 9h6ZM9.5 16.5v-7h6v7h-6Z" />
          <path d="M7 6.5A1.5 1.5 0 0 1 5.5 8H2a1.5 1.5 0 0 1-1.5-1.5v-3A1.5 1.5 0 0 1 2 2h3.5A1.5 1.5 0 0 1 7 3.5v3Z" />
        </svg>
      `
		};
		return icons[iconName] || '';
	}
</script>

<div class="overflow-hidden bg-white py-24 sm:py-32">
	<div class="mx-auto max-w-7xl px-6 lg:px-8">
		<div
			class="mx-auto grid max-w-2xl grid-cols-1 gap-x-8 gap-y-16 sm:gap-y-20 lg:mx-0 lg:max-w-none lg:grid-cols-2"
		>
			<div class="lg:ml-auto lg:pl-4 lg:pt-4">
				<div id="features" class="lg:max-w-lg">
					<h2 class="text-primary-600 text-base/7 font-semibold">Our services</h2>
					<p
						class="mt-2 text-pretty text-4xl font-semibold tracking-tight text-gray-900 sm:text-5xl"
					>
						Everything you need to stay organized
					</p>
					<p class="mt-6 text-lg/8 text-gray-600">
						From basic bookkeeping to payroll processing and strategic consulting, we provide
						comprehensive financial services tailored to growing businesses like yours.
					</p>
					<dl class="mt-10 max-w-xl space-y-8 text-base/7 text-gray-600 lg:max-w-none">
						{#each features as feature}
							<div class="relative pl-9">
								<dt class="inline font-semibold text-gray-900">
									{@html getIconSvg(feature.icon)}
									{feature.name}
								</dt>
								<dd class="inline">{feature.description}</dd>
							</div>
						{/each}
					</dl>
				</div>
			</div>
			<div class="flex items-start justify-end lg:order-first">
				<img
					alt="Product screenshot"
					src="https://as2.ftcdn.net/v2/jpg/10/58/87/61/1000_F_1058876100_ZN9yjJLPq3jODb1YCRFLPTcTs48siq3V.jpg"
					width={2432}
					height={1442}
					class="w-3xl sm:w-228 max-w-none rounded-xl shadow-xl ring-1 ring-gray-400/10"
				/>
			</div>
		</div>
	</div>
</div>
