<script>
	const faqData = {
		title: 'Frequently asked questions',
		questions: [
			{
				question: 'What industries do you serve?',
				answer:
					'We work with service-based and product-based businesses — including real estate, retail, creatives, contractors, and more.'
			},
			{
				question: 'Do I need to be located in Washington State?',
				answer:
					'Nope! We work remotely with clients across the U.S. using secure cloud-based tools.'
			},
			{
				question: 'What accounting software do you use?',
				answer:
					'We use QuickBooks Online along with integrated tools for payroll, reporting, and document sharing.'
			},
			{
				question: 'How does onboarding work?',
				answer:
					"We'll review your books, set up secure access, outline expectations, and launch your monthly services — typically within 1–2 weeks."
			},
			{
				question: 'Can you help clean up old books?',
				answer:
					"Yes! We offer bookkeeping cleanup as a separate service. We'll assess your books and give you a custom quote."
			}
		]
	};
</script>

<div class="bg-white">
	<div class="mx-auto max-w-7xl px-6 py-24 sm:py-32 lg:px-8 lg:py-40">
		<div id="faq" class="mx-auto max-w-4xl">
			<h2
				class="text-headline-lg font-display-semibold sm:text-display-sm tracking-tight text-gray-900"
			>
				{faqData.title}
			</h2>

			<dl class="mt-16 divide-y divide-gray-900/10">
				{#each faqData.questions as faq, index}
					<div class="py-6 first:pt-0 last:pb-0">
						<details class="group">
							<summary
								class="flex w-full cursor-pointer list-none items-start justify-between text-left text-gray-900"
							>
								<span class="text-body font-primary-semibold pr-6">
									{faq.question}
								</span>
								<span class="ml-6 flex h-7 flex-shrink-0 items-center">
									<!-- Plus icon (shown when closed) -->
									<svg
										viewBox="0 0 24 24"
										fill="none"
										stroke="currentColor"
										stroke-width="1.5"
										aria-hidden="true"
										class="size-6 group-open:hidden"
									>
										<path d="M12 6v12m6-6H6" stroke-linecap="round" stroke-linejoin="round" />
									</svg>
									<!-- Minus icon (shown when open) -->
									<svg
										viewBox="0 0 24 24"
										fill="none"
										stroke="currentColor"
										stroke-width="1.5"
										aria-hidden="true"
										class="hidden size-6 group-open:block"
									>
										<path d="M18 12H6" stroke-linecap="round" stroke-linejoin="round" />
									</svg>
								</span>
							</summary>

							<dd class="mt-2 pr-12">
								<p class="text-body text-gray-600">
									{faq.answer}
								</p>
							</dd>
						</details>
					</div>
				{/each}
			</dl>
		</div>
	</div>
</div>

<style>
	/* Remove default details styling and add custom styling */
	details > summary {
		display: flex;
		align-items: flex-start;
		justify-content: space-between;
	}

	details > summary::-webkit-details-marker {
		display: none;
	}

	details > summary::marker {
		content: '';
	}

	/* Smooth transition for content */
	details[open] dd {
		animation: slideDown 0.2s ease-out;
	}

	@keyframes slideDown {
		from {
			opacity: 0;
			transform: translateY(-10px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	/* Hover effect for better UX */
	summary:hover {
		color: #374151; /* gray-700 */
	}
</style>
