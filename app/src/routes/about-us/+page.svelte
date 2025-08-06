<script>
	import SectionHeader from '$lib/header.svelte';
	import MissionStatement from '$lib/mission.svelte';
	import TeamSection from '$lib/team.svelte';
	import FeaturesGrid from '$lib/featuresGrid.svelte';
	import Cta from '$lib/cta.svelte';

	// Mission section data
	const missionData = {
		title: 'Our Mission',
		description:
			'At Momentum Business Solutions, we believe small businesses deserve the same financial clarity and strategic insight as big corporations. Our mission is to help you move forward with confidence — through precise bookkeeping, reliable payroll, and smart financial planning.'
	};

	// Team data with Cade as the only member
	const teamData = {
		title: 'Meet the Team',
		description:
			'Get to know the financial experts who are dedicated to helping your business succeed.',
		members: [
			{
				name: 'Cade Maldonado',
				role: 'Founder & CEO',
				image:
					'https://media.licdn.com/dms/image/v2/D4E03AQEwiyCOi13jIA/profile-displayphoto-shrink_200_200/profile-displayphoto-shrink_200_200/0/1715788195232?e=1757548800&v=beta&t=T7diNCm8jF622h4bHiOKfmB67z0jkpSyGIB6fZuDED4',
				bio: 'Cade brings over a decade of financial management experience across corporate and small business settings. He is QuickBooks ProAdvisor Certified and specializes in helping business owners gain clarity and confidence through clean books, streamlined processes, and actionable insights.',
				socialLinks: [
					{
						platform: 'LinkedIn',
						url: 'https://www.linkedin.com/in/cade-maldonado-a71801104',
						icon: `<svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" class="size-5">
              <path d="M16.338 16.338H13.67V12.16c0-.995-.017-2.277-1.387-2.277-1.39 0-1.601 1.086-1.601 2.207v4.248H8.014v-8.59h2.559v1.174h.037c.356-.675 1.227-1.387 2.526-1.387 2.703 0 3.203 1.778 3.203 4.092v4.711zM5.005 6.575a1.548 1.548 0 11-.003-3.096 1.548 1.548 0 01.003 3.096zm-1.337 9.763H6.34v-8.59H3.667v8.59zM17.668 1H2.328C1.595 1 1 1.581 1 2.298v15.403C1 18.418 1.595 19 2.328 19h15.34c.734 0 1.332-.582 1.332-1.299V2.298C19 1.581 18.402 1 17.668 1z" clip-rule="evenodd" fill-rule="evenodd" />
            </svg>`
					}
				]
			}
		]
	};

	// Why Momentum features
	const whyMomentumFeatures = [
		{
			title: 'Tailored for small businesses',
			description:
				'We understand the unique challenges of growing businesses and provide solutions that scale with you.',
			icon: `<svg aria-hidden="true" class="size-6 text-primary-600" fill="currentColor" viewBox="0 0 20 20">
        <path fill-rule="evenodd" d="M10 2a4 4 0 0 0-4 4v1H5a1 1 0 0 0-.994.89l-1 9A1 1 0 0 0 4 18h12a1 1 0 0 0 .994-1.11l-1-9A1 1 0 0 0 15 7h-1V6a4 4 0 0 0-4-4ZM8 6V5a2 2 0 1 1 4 0v1H8Zm-1 3a1 1 0 1 1 2 0 1 1 0 0 1-2 0Zm6 0a1 1 0 1 1 2 0 1 1 0 0 1-2 0Z" clip-rule="evenodd" />
      </svg>`
		},
		{
			title: 'Transparent communication',
			description:
				'Clear reporting, regular updates, and straightforward explanations — no financial jargon or hidden surprises.',
			icon: `<svg aria-hidden="true" class="size-6 text-primary-600" fill="currentColor" viewBox="0 0 20 20">
        <path fill-rule="evenodd" d="M18 10a8 8 0 1 1-16 0 8 8 0 0 1 16 0Zm-7-4a1 1 0 1 1-2 0 1 1 0 0 1 2 0ZM9 9a.75.75 0 0 0 0 1.5h.253a.25.25 0 0 1 .244.304l-.459 2.066A1.75 1.75 0 0 0 10.747 15H11a.75.75 0 0 0 0-1.5h-.253a.25.25 0 0 1-.244-.304l.459-2.066A1.75 1.75 0 0 0 9.253 9H9Z" clip-rule="evenodd" />
      </svg>`
		},
		{
			title: 'Financial insights that go beyond the numbers',
			description:
				"We don't just track your finances — we help you understand what they mean and how to use them for strategic growth.",
			icon: `<svg aria-hidden="true" class="size-6 text-primary-600" fill="currentColor" viewBox="0 0 20 20">
        <path fill-rule="evenodd" d="M1 4.75C1 3.784 1.784 3 2.75 3h14.5c.966 0 1.75.784 1.75 1.75v10.515a1.75 1.75 0 0 1-1.75 1.75h-1.5c-.078 0-.155-.005-.23-.015H4.48c-.075.01-.152.015-.23.015h-1.5A1.75 1.75 0 0 1 1 15.265V4.75Zm16.5 7.385V11.01a.25.25 0 0 0-.25-.25h-1.5a.25.25 0 0 0-.25.25v1.125c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25Zm0 2.005a.25.25 0 0 0-.25-.25h-1.5a.25.25 0 0 0-.25.25v1.125c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25v-1.125Z" clip-rule="evenodd" />
      </svg>`
		}
	];
</script>

<!-- Page Header -->
<SectionHeader
	title="About Us"
	description="Learn about our mission to help small businesses achieve financial clarity and confidence."
/>
<MissionStatement />
<TeamSection {teamData} />
<FeaturesGrid />
<Cta />
