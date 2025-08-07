<script>
  import { slide } from 'svelte/transition';
  import { quintOut } from 'svelte/easing';

  let { 
    message = {
      highlight: "QuickBooks Cleanup Special",
      text: "Get your books organized before year-end. Professional cleanup service for just $750 – most completed within 7-10 business days",
      linkText: "Get Started",
      linkHref: "/services#quickbooks-cleanup"
    },
    backgroundColor = "bg-primary-600",
    onDismiss = null // Optional callback when banner is dismissed
  } = $props();

  let showBanner = $state(true)

  function dismissBanner() {
    showBanner = false;
    if (onDismiss) {
      onDismiss();
    }
  }
</script>

{#if showBanner}
  <div class="flex items-center gap-x-6 {backgroundColor} px-6 py-2.5 sm:px-3.5 sm:before:flex-1" transition:slide={{ duration: 400, easing: quintOut }}>
    <p class="text-caption text-white">
      <a href={message.linkHref} class="hover:text-white/90 transition-colors">
        <strong class="font-primary-semibold">{message.highlight}</strong>
        <svg viewBox="0 0 2 2" aria-hidden="true" class="mx-2 inline size-0.5 fill-current">
          <circle r="1" cx="1" cy="1" />
        </svg>
        {message.text}&nbsp;
        <span aria-hidden="true">→</span>
      </a>
    </p>
    
    <div class="flex flex-1 justify-end">
      <button 
        type="button" 
        onclick={dismissBanner}
        class="-m-3 p-3 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white hover:text-white/80 transition-colors"
        aria-label="Dismiss announcement"
      >
        <span class="sr-only">Dismiss</span>
        <svg viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" class="size-5 text-white">
          <path d="M6.28 5.22a.75.75 0 0 0-1.06 1.06L8.94 10l-3.72 3.72a.75.75 0 1 0 1.06 1.06L10 11.06l3.72 3.72a.75.75 0 1 0 1.06-1.06L11.06 10l3.72-3.72a.75.75 0 0 0-1.06-1.06L10 8.94 6.28 5.22Z" />
        </svg>
      </button>
    </div>
  </div>
{/if}