<script>
  const footerData = {
    navigation: [
      { name: "About Us", href: "/about-us" },
      { name: "Services", href: "#" },
      { name: "Pricing", href: "#" },
      { name: "Contact", href: "#" },
      { name: "Privacy", href: "#" },
      { name: "Terms", href: "#" }
    ],
    socialLinks: [
      {
        name: "LinkedIn",
        href: "https://www.linkedin.com/company/momentum-business-solutions-llc",
        icon: `<svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true" class="size-6">
          <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
        </svg>`
      },
      {
        name: "Facebook",
        href: "#",
        icon: `<svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true" class="size-6">
          <path d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12c0 4.991 3.657 9.128 8.438 9.878v-6.987h-2.54V12h2.54V9.797c0-2.506 1.492-3.89 3.777-3.89 1.094 0 2.238.195 2.238.195v2.46h-1.26c-1.243 0-1.63.771-1.63 1.562V12h2.773l-.443 2.89h-2.33v6.988C18.343 21.128 22 16.991 22 12z" clip-rule="evenodd" fill-rule="evenodd" />
        </svg>`
      },
      {
        name: "X",
        href: "#",
        icon: `<svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true" class="size-6">
          <path d="M13.6823 10.6218L20.2391 3H18.6854L12.9921 9.61788L8.44486 3H3.2002L10.0765 13.0074L3.2002 21H4.75404L10.7663 14.0113L15.5685 21H20.8131L13.6819 10.6218H13.6823ZM11.5541 13.0956L10.8574 12.0991L5.31391 4.16971H7.70053L12.1742 10.5689L12.8709 11.5655L18.6861 19.8835H16.2995L11.5541 13.096V13.0956Z" />
        </svg>`
      },
      {
        name: "YouTube",
        href: "#",
        icon: `<svg viewBox="0 0 24 24" fill="currentColor" aria-hidden="true" class="size-6">
          <path d="M19.812 5.418c.861.23 1.538.907 1.768 1.768C21.998 8.746 22 12 22 12s0 3.255-.418 4.814a2.504 2.504 0 0 1-1.768 1.768c-1.56.419-7.814.419-7.814.419s-6.255 0-7.814-.419a2.505 2.505 0 0 1-1.768-1.768C2 15.255 2 12 2 12s0-3.255.417-4.814a2.507 2.507 0 0 1 1.768-1.768C5.744 5 11.998 5 11.998 5s6.255 0 7.814.418ZM15.194 12 10 15V9l5.194 3Z" clip-rule="evenodd" fill-rule="evenodd" />
        </svg>`
      }
    ],
    copyright: {
      year: "2024",
      company: "Momentum Financial Solutions",
      text: "All rights reserved."
    }
  };
</script>

<footer class="bg-white">
  <div class="mx-auto max-w-7xl overflow-hidden px-6 py-20 sm:py-24 lg:px-8">
    <nav aria-label="Footer" class="-mb-6 flex flex-wrap justify-center gap-x-12 gap-y-3 text-body">
      {#each footerData.navigation as link}
        <a href={link.href} class="text-gray-600 hover:text-gray-900 font-primary-medium">
          {link.name}
        </a>
      {/each}
    </nav>
    
    <div class="mt-16 flex justify-center gap-x-10">
      {#each footerData.socialLinks as social}
        <a href={social.href} class="text-gray-600 hover:text-gray-800">
          <span class="sr-only">{social.name}</span>
          {@html social.icon}
        </a>
      {/each}
    </div>
    
    <p class="mt-10 text-center text-caption text-gray-600 font-primary-normal">
      &copy; {footerData.copyright.year} {footerData.copyright.company}. {footerData.copyright.text}
    </p>
  </div>
</footer>