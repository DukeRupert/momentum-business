<script>
  let { 
    title = "Section Title", 
    description = "Section description goes here.",
    kicker = null 
  } = $props();
</script>

<div class="bg-white py-24 sm:py-32">
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <div class="mx-auto max-w-2xl lg:mx-0">
      {#if kicker}
        <h2 class="text-kicker font-primary-semibold text-primary-600 uppercase tracking-wide">
          {kicker}
        </h2>
      {/if}
      
      <h1 class="text-display-sm font-display-bold tracking-tight text-gray-900 sm:text-display {kicker ? 'mt-2' : ''}">
        {title}
      </h1>
      
      <p class="mt-8 text-body-lg font-primary-medium text-pretty text-gray-500 sm:text-subhead">
        {description}
      </p>
    </div>
  </div>
</div>