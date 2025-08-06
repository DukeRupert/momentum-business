<script>
	import SectionHeader from '$lib/header.svelte';
	import MissionStatement from '$lib/mission.svelte';
	import TeamSection from '$lib/team.svelte';
	import FeaturesGrid from '$lib/featuresGrid.svelte';
	import Cta from '$lib/cta.svelte';

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
