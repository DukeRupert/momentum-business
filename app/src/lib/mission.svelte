<script>
  // Props for the mission section
  let { 
    title = "Our Mission",
    content = {
      mainText: "At Momentum Business Solutions, we believe small businesses deserve the same financial clarity and strategic insight as big corporations. Our mission is to help you move forward with confidence — through precise bookkeeping, reliable payroll, and smart financial planning.",
      supportingText: "We understand that managing finances shouldn't keep you awake at night or pull you away from growing your business. That's why we handle the complexity so you can focus on what you do best — serving your customers and building your vision."
    },
    stats = [
      {
        value: "10+ years",
        label: "Combined financial expertise"
      },
      {
        value: "$500K-$5M",
        label: "Annual revenue clients we serve"
      },
      {
        value: "1-2 weeks",
        label: "Typical onboarding time"
      }
    ]
  } = $props();
</script>

<section class="bg-gray-50 py-24 sm:py-32">
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <div class="mx-auto max-w-2xl lg:mx-0 lg:max-w-none">
      <h2 class="text-headline-lg font-display-semibold tracking-tight text-pretty text-gray-900 sm:text-display-sm">
        {title}
      </h2>
      
      <div class="mt-8 flex flex-col gap-x-8 gap-y-16 lg:flex-row lg:gap-y-20">
        <!-- Mission Text -->
        <div class="lg:w-full lg:max-w-2xl lg:flex-auto">
          <p class="text-subhead text-gray-600 font-primary-medium">
            {content.mainText}
          </p>
          <p class="mt-8 text-body text-gray-700">
            {content.supportingText}
          </p>
        </div>
        
        <!-- Stats/Highlights -->
        <div class="lg:flex lg:flex-auto lg:justify-center">
          <dl class="w-full max-w-xs space-y-8 lg:w-64 xl:w-80">
            {#each stats as stat}
              <div class="flex flex-col-reverse gap-y-2">
                <dt class="text-body text-gray-600">
                  {stat.label}
                </dt>
                <dd class="text-headline-lg font-display-bold tracking-tight text-primary-600">
                  {stat.value}
                </dd>
              </div>
            {/each}
          </dl>
        </div>
      </div>
    </div>
  </div>
</section>