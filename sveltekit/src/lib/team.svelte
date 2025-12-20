<script lang="ts">
  // TypeScript interface definitions
  interface SocialLink {
    platform: string;
    url: string;
    icon: string;
  }

  interface TeamMember {
    name: string;
    role: string;
    image: string;
    bio: string;
    socialLinks: SocialLink[];
  }

  interface TeamData {
    title: string;
    description: string;
    members: TeamMember[];
  }

  // Svelte 5 props with TypeScript and default values
  let { 
    teamData = {
      title: "Meet the Momentum team",
      description: "We're a dedicated group of financial professionals who are passionate about helping businesses achieve financial clarity and peace of mind.",
      members: [
        {
          name: "Sarah Johnson",
          role: "Founder & Lead Bookkeeper",
          image: "https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=8&w=1024&h=1024&q=80",
          bio: "With over 12 years of experience in financial management, Sarah founded Momentum to help business owners focus on what they do best. She specializes in cash flow optimization and strategic financial planning.",
          socialLinks: [
            {
              platform: "LinkedIn",
              url: "#",
              icon: `<svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" class="size-5">
                <path d="M16.338 16.338H13.67V12.16c0-.995-.017-2.277-1.387-2.277-1.39 0-1.601 1.086-1.601 2.207v4.248H8.014v-8.59h2.559v1.174h.037c.356-.675 1.227-1.387 2.526-1.387 2.703 0 3.203 1.778 3.203 4.092v4.711zM5.005 6.575a1.548 1.548 0 11-.003-3.096 1.548 1.548 0 01.003 3.096zm-1.337 9.763H6.34v-8.59H3.667v8.59zM17.668 1H2.328C1.595 1 1 1.581 1 2.298v15.403C1 18.418 1.595 19 2.328 19h15.34c.734 0 1.332-.582 1.332-1.299V2.298C19 1.581 18.402 1 17.668 1z" clip-rule="evenodd" fill-rule="evenodd" />
              </svg>`
            }
          ]
        },
        {
          name: "Michael Chen",
          role: "Senior Bookkeeper & Payroll Specialist",
          image: "https://images.unsplash.com/photo-1519244703995-f4e0f30006d5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=8&w=1024&h=1024&q=80",
          bio: "Michael brings 8 years of payroll and compliance expertise to our team. He ensures our clients never worry about tax deadlines or employee payment processing, handling everything with precision and care.",
          socialLinks: [
            {
              platform: "LinkedIn",
              url: "#",
              icon: `<svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" class="size-5">
                <path d="M16.338 16.338H13.67V12.16c0-.995-.017-2.277-1.387-2.277-1.39 0-1.601 1.086-1.601 2.207v4.248H8.014v-8.59h2.559v1.174h.037c.356-.675 1.227-1.387 2.526-1.387 2.703 0 3.203 1.778 3.203 4.092v4.711zM5.005 6.575a1.548 1.548 0 11-.003-3.096 1.548 1.548 0 01.003 3.096zm-1.337 9.763H6.34v-8.59H3.667v8.59zM17.668 1H2.328C1.595 1 1 1.581 1 2.298v15.403C1 18.418 1.595 19 2.328 19h15.34c.734 0 1.332-.582 1.332-1.299V2.298C19 1.581 18.402 1 17.668 1z" clip-rule="evenodd" fill-rule="evenodd" />
              </svg>`
            }
          ]
        },
        {
          name: "Jennifer Martinez",
          role: "Financial Consultant & Client Success",
          image: "https://images.unsplash.com/photo-1506794778202-cad84cf45f1d?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=8&w=1024&h=1024&q=80",
          bio: "Jennifer specializes in strategic financial planning and helps our clients make data-driven decisions. With her background in business consulting, she translates complex financial data into actionable insights for growth.",
          socialLinks: [
            {
              platform: "LinkedIn", 
              url: "#",
              icon: `<svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" class="size-5">
                <path d="M16.338 16.338H13.67V12.16c0-.995-.017-2.277-1.387-2.277-1.39 0-1.601 1.086-1.601 2.207v4.248H8.014v-8.59h2.559v1.174h.037c.356-.675 1.227-1.387 2.526-1.387 2.703 0 3.203 1.778 3.203 4.092v4.711zM5.005 6.575a1.548 1.548 0 11-.003-3.096 1.548 1.548 0 01.003 3.096zm-1.337 9.763H6.34v-8.59H3.667v8.59zM17.668 1H2.328C1.595 1 1 1.581 1 2.298v15.403C1 18.418 1.595 19 2.328 19h15.34c.734 0 1.332-.582 1.332-1.299V2.298C19 1.581 18.402 1 17.668 1z" clip-rule="evenodd" fill-rule="evenodd" />
              </svg>`
            }
          ]
        }
      ]
    } as TeamData
  } = $props();
</script>

<div class="bg-white py-24 md:py-32">
  <div class="mx-auto grid max-w-7xl grid-cols-1 lg:grid-cols-5 gap-20 px-6 lg:px-8">
    <div class="max-w-2xl lg:col-span-2">
      <h2 class="text-headline-lg font-display-semibold tracking-tight text-pretty text-gray-900 sm:text-display-sm">
        {teamData.title}
      </h2>
      <p class="mt-6 text-body-lg text-gray-600">
        {teamData.description}
      </p>
    </div>
    
    <ul role="list" class="divide-y divide-gray-200 lg:col-span-3">
      {#each teamData.members as member}
        <li class="flex flex-col gap-10 py-12 first:pt-0 last:pb-0 sm:flex-row">
          <img 
            src={member.image} 
            alt={member.name} 
            class="aspect-[4/5] w-52 flex-none rounded-2xl object-cover outline-1 -outline-offset-1 outline-black/5" 
          />
          <div class="max-w-xl flex-auto">
            <h3 class="text-subhead font-display-semibold tracking-tight text-gray-900">
              {member.name}
            </h3>
            <p class="text-body text-gray-600 font-primary-medium">
              {member.role}
            </p>
            <p class="mt-6 text-body text-gray-600">
              {member.bio}
            </p>
            <ul role="list" class="mt-6 flex gap-x-6">
              {#each member.socialLinks as social}
                <li>
                  <a href={social.url} class="text-gray-400 hover:text-gray-500">
                    <span class="sr-only">{social.platform}</span>
                    {@html social.icon}
                  </a>
                </li>
              {/each}
            </ul>
          </div>
        </li>
      {/each}
    </ul>
  </div>
</div>