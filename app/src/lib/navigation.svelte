<script>
  import { fade, fly } from 'svelte/transition';
  
  const navigation = [
    { name: 'Services', href: '/services' },
    { name: 'About', href: '/about' },
    { name: 'Contact', href: '/contact' },
  ];

  let mobileMenuOpen = $state(false);

  function openMobileMenu() {
    mobileMenuOpen = true;
  }

  function closeMobileMenu() {
    mobileMenuOpen = false;
  }

  // Close menu when clicking outside
  /**
	 * @param {{ target: any; currentTarget: any; }} event
	 */
  function handleBackdropClick(event) {
    if (event.target === event.currentTarget) {
      closeMobileMenu();
    }
  }

  // Handle escape key
  /**
	 * @param {{ key: string; }} event
	 */
  function handleKeydown(event) {
    if (event.key === 'Escape' && mobileMenuOpen) {
      closeMobileMenu();
    }
  }
</script>

<svelte:window onkeydown={handleKeydown} />

<header class="bg-white">
  <nav aria-label="Global" class="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8">
    <a href="/" class="-m-1.5 p-1.5">
      <span class="sr-only">Your Company</span>
      <img
        alt=""
        src="https://tailwindcss.com/plus-assets/img/logos/mark.svg?color=indigo&shade=600"
        class="h-8 w-auto"
      />
    </a>
    
    <div class="flex lg:hidden">
      <button
        type="button"
        onclick={openMobileMenu}
        class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700"
      >
        <span class="sr-only">Open main menu</span>
        <!-- Bars3Icon -->
        <svg aria-hidden="true" class="size-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
        </svg>
      </button>
    </div>
    
    <div class="hidden lg:flex lg:gap-x-12">
      {#each navigation as item}
        <a href={item.href} class="text-sm/6 font-semibold text-gray-900">
          {item.name}
        </a>
      {/each}
      <a href="#" class="text-sm/6 font-semibold text-gray-900">
        Log in <span aria-hidden="true">&rarr;</span>
      </a>
    </div>
  </nav>

  <!-- Mobile menu overlay and panel -->
  {#if mobileMenuOpen}
    <!-- Backdrop -->
    <div 
      class="fixed inset-0 z-50 lg:hidden" 
      onclick={handleBackdropClick}
      role="button"
      tabindex="0"
      onkeydown={(e) => e.key === 'Enter' && closeMobileMenu()}
      transition:fade={{ duration: 200 }}
    ></div>
    
    <!-- Mobile menu panel -->
    <div 
      class="fixed inset-y-0 right-0 z-50 w-full overflow-y-auto bg-white p-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10 lg:hidden"
      transition:fly={{ x: 300, duration: 300 }}
    >
      <div class="flex items-center justify-between">
        <a href="#" class="-m-1.5 p-1.5">
          <span class="sr-only">Your Company</span>
          <img
            alt=""
            src="https://tailwindcss.com/plus-assets/img/logos/mark.svg?color=primary&shade=600"
            class="h-8 w-auto"
          />
        </a>
        
        <button
          type="button"
          onclick={closeMobileMenu}
          class="-m-2.5 rounded-md p-2.5 text-gray-700"
        >
          <span class="sr-only">Close menu</span>
          <!-- XMarkIcon -->
          <svg aria-hidden="true" class="size-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
      
      <div class="mt-6 flow-root">
        <div class="-my-6 divide-y divide-gray-500/10">
          <div class="space-y-2 py-6">
            {#each navigation as item}
              <a
                href={item.href}
                class="-mx-3 block rounded-lg px-3 py-2 text-base/7 font-semibold text-gray-900 hover:bg-gray-50"
              >
                {item.name}
              </a>
            {/each}
          </div>
          
          <div class="py-6">
            <a
              href="#"
              class="-mx-3 block rounded-lg px-3 py-2.5 text-base/7 font-semibold text-gray-900 hover:bg-gray-50"
            >
              Log in
            </a>
          </div>
        </div>
      </div>
    </div>
  {/if}
</header>