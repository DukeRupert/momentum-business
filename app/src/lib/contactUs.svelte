<script>
  let { 
    title = "Contact Us",
    subtitle = "Let's start the conversation.",
    description = "Ready to take control of your finances? Get in touch with us today to discuss how we can help streamline your bookkeeping, payroll, and financial planning.",
    contactInfo = {
      email: "cade@momentumbusiness.org",
      phone: "(509) 554-8022",
      location: "Serving clients remotely nationwide"
    }
  } = $props();
</script>

<div class="relative isolate bg-white">
  <div class="mx-auto grid max-w-7xl grid-cols-1 lg:grid-cols-2">
    <!-- Contact Information Side -->
    <div class="relative px-6 pt-24 pb-20 sm:pt-32 lg:static lg:px-8 lg:py-48">
      <div class="mx-auto max-w-xl lg:mx-0 lg:max-w-lg">
        <!-- Background Pattern -->
        <div class="absolute inset-y-0 left-0 -z-10 w-full overflow-hidden bg-gray-100 ring-1 ring-gray-900/10 lg:w-1/2">
          <svg aria-hidden="true" class="absolute inset-0 size-full mask-[radial-gradient(100%_100%_at_top_right,white,transparent)] stroke-gray-200">
            <defs>
              <pattern id="contact-pattern" width="200" height="200" x="100%" y="-1" patternUnits="userSpaceOnUse">
                <path d="M130 200V.5M.5 .5H200" fill="none" />
              </pattern>
            </defs>
            <rect width="100%" height="100%" stroke-width="0" class="fill-white" />
            <svg x="100%" y="-1" class="overflow-visible fill-gray-50">
              <path d="M-470.5 0h201v201h-201Z" stroke-width="0" />
            </svg>
            <rect width="100%" height="100%" fill="url(#contact-pattern)" stroke-width="0" />
          </svg>
          <!-- Gradient Blur -->
          <div aria-hidden="true" class="absolute top-[calc(100%-13rem)] -left-56 hidden transform-gpu blur-3xl lg:top-[calc(50%-7rem)] lg:left-[max(-14rem,calc(100%-59rem))]">
            <div style="clip-path: polygon(74.1% 56.1%, 100% 38.6%, 97.5% 73.3%, 85.5% 100%, 80.7% 98.2%, 72.5% 67.7%, 60.2% 37.8%, 52.4% 32.2%, 47.5% 41.9%, 45.2% 65.8%, 27.5% 23.5%, 0.1% 35.4%, 17.9% 0.1%, 27.6% 23.5%, 76.1% 2.6%, 74.1% 56.1%)" class="aspect-[1155/678] w-[72.1875rem] bg-gradient-to-tr from-[#80caff] to-[#4f46e5] opacity-10"></div>
          </div>
        </div>
        
        <!-- Contact Content -->
        <h2 class="text-headline-lg font-display-semibold tracking-tight text-pretty text-gray-900 sm:text-display-sm">
          {title}
        </h2>
        <p class="mt-2 text-subhead font-primary-medium text-primary-600">
          {subtitle}
        </p>
        <p class="mt-6 text-body-lg text-gray-600">
          {description}
        </p>
        
        <!-- Contact Details -->
        <dl class="mt-10 space-y-4 text-body text-gray-600">
          <div class="flex gap-x-4">
            <dt class="flex-none">
              <span class="sr-only">Email</span>
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" aria-hidden="true" class="h-7 w-6 text-gray-400">
                <path d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
            </dt>
            <dd>
              <a href="mailto:{contactInfo.email}" class="hover:text-gray-900">
                {contactInfo.email}
              </a>
            </dd>
          </div>
          
          <div class="flex gap-x-4">
            <dt class="flex-none">
              <span class="sr-only">Telephone</span>
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" aria-hidden="true" class="h-7 w-6 text-gray-400">
                <path d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 0 0 2.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 0 1-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 0 0-1.091-.852H4.5A2.25 2.25 0 0 0 2.25 4.5v2.25Z" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
            </dt>
            <dd>
              <a href="tel:{contactInfo.phone.replace(/\D/g, '')}" class="hover:text-gray-900">
                {contactInfo.phone}
              </a>
            </dd>
          </div>
          
          <div class="flex gap-x-4">
            <dt class="flex-none">
              <span class="sr-only">Location</span>
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" aria-hidden="true" class="h-7 w-6 text-gray-400">
                <path d="M12 21a9.004 9.004 0 0 0 8.716-6.747M12 21a9.004 9.004 0 0 1-8.716-6.747M12 21c2.485 0 4.5-4.03 4.5-9S14.485 3 12 3s-4.5 4.03-4.5 9 2.015 9 4.5 9Z" stroke-linecap="round" stroke-linejoin="round" />
                <path d="M2 12h20" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
            </dt>
            <dd>{contactInfo.location}</dd>
          </div>
        </dl>
      </div>
    </div>
    
    <!-- Contact Form Side -->
    <form action="#" method="POST" class="px-6 pt-20 pb-24 sm:pb-32 lg:px-8 lg:py-48">
      <div class="mx-auto max-w-xl lg:mr-0 lg:max-w-lg">
        <div class="grid grid-cols-1 gap-x-8 gap-y-6 sm:grid-cols-2">
          <div>
            <label for="first-name" class="block text-caption font-primary-semibold text-gray-900">First name</label>
            <div class="mt-2.5">
              <input 
                id="first-name" 
                type="text" 
                name="first-name" 
                autocomplete="given-name" 
                class="block w-full rounded-md bg-white px-3.5 py-2 text-body text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-primary-600" 
              />
            </div>
          </div>
          
          <div>
            <label for="last-name" class="block text-caption font-primary-semibold text-gray-900">Last name</label>
            <div class="mt-2.5">
              <input 
                id="last-name" 
                type="text" 
                name="last-name" 
                autocomplete="family-name" 
                class="block w-full rounded-md bg-white px-3.5 py-2 text-body text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-primary-600" 
              />
            </div>
          </div>
          
          <div class="sm:col-span-2">
            <label for="email" class="block text-caption font-primary-semibold text-gray-900">Email</label>
            <div class="mt-2.5">
              <input 
                id="email" 
                type="email" 
                name="email" 
                autocomplete="email" 
                class="block w-full rounded-md bg-white px-3.5 py-2 text-body text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-primary-600" 
              />
            </div>
          </div>
          
          <div class="sm:col-span-2">
            <label for="phone-number" class="block text-caption font-primary-semibold text-gray-900">Phone number</label>
            <div class="mt-2.5">
              <input 
                id="phone-number" 
                type="tel" 
                name="phone-number" 
                autocomplete="tel" 
                class="block w-full rounded-md bg-white px-3.5 py-2 text-body text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-primary-600" 
              />
            </div>
          </div>
          
          <div class="sm:col-span-2">
            <label for="annual-revenue" class="block text-caption font-primary-semibold text-gray-900">Annual revenue</label>
            <div class="mt-2.5">
              <select 
                id="annual-revenue" 
                name="annual-revenue" 
                class="block w-full rounded-md bg-white px-3.5 py-2 text-body text-gray-900 outline-1 -outline-offset-1 outline-gray-300 focus:outline-2 focus:-outline-offset-2 focus:outline-primary-600"
              >
                <option value="">Select revenue range</option>
                <option value="under-100k">Under $100K</option>
                <option value="100k-500k">$100K - $500K</option>
                <option value="500k-1m">$500K - $1M</option>
                <option value="1m-5m">$1M - $5M</option>
                <option value="over-5m">Over $5M</option>
              </select>
            </div>
          </div>
          
          <div class="sm:col-span-2">
            <fieldset>
              <legend class="block text-caption font-primary-semibold text-gray-900 mb-4">Services you're interested in</legend>
              <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
                <div class="flex items-start gap-3">
                  <div class="flex h-6 items-center">
                    <input 
                      id="service-bookkeeping" 
                      name="services" 
                      value="bookkeeping" 
                      type="checkbox" 
                      class="size-4 rounded border-gray-300 text-primary-600 focus:ring-primary-600 focus:ring-offset-0"
                    />
                  </div>
                  <label for="service-bookkeeping" class="text-body text-gray-700">
                    <span class="font-primary-medium">Basic Bookkeeping</span>
                    <span class="block text-caption text-gray-500">Starting at $750/month</span>
                  </label>
                </div>
                
                <div class="flex items-start gap-3">
                  <div class="flex h-6 items-center">
                    <input 
                      id="service-bookkeeping-payroll" 
                      name="services" 
                      value="bookkeeping-payroll" 
                      type="checkbox" 
                      class="size-4 rounded border-gray-300 text-primary-600 focus:ring-primary-600 focus:ring-offset-0"
                    />
                  </div>
                  <label for="service-bookkeeping-payroll" class="text-body text-gray-700">
                    <span class="font-primary-medium">Bookkeeping + Payroll</span>
                    <span class="block text-caption text-gray-500">Starting at $1,250/month</span>
                  </label>
                </div>
                
                <div class="flex items-start gap-3">
                  <div class="flex h-6 items-center">
                    <input 
                      id="service-consulting" 
                      name="services" 
                      value="consulting" 
                      type="checkbox" 
                      class="size-4 rounded border-gray-300 text-primary-600 focus:ring-primary-600 focus:ring-offset-0"
                    />
                  </div>
                  <label for="service-consulting" class="text-body text-gray-700">
                    <span class="font-primary-medium">Financial Consulting</span>
                    <span class="block text-caption text-gray-500">$150/hour</span>
                  </label>
                </div>
                
                <div class="flex items-start gap-3">
                  <div class="flex h-6 items-center">
                    <input 
                      id="service-cleanup" 
                      name="services" 
                      value="cleanup" 
                      type="checkbox" 
                      class="size-4 rounded border-gray-300 text-primary-600 focus:ring-primary-600 focus:ring-offset-0"
                    />
                  </div>
                  <label for="service-cleanup" class="text-body text-gray-700">
                    <span class="font-primary-medium">Bookkeeping Cleanup</span>
                    <span class="block text-caption text-gray-500">Custom quote</span>
                  </label>
                </div>
              </div>
            </fieldset>
          </div>
          
          <div class="sm:col-span-2">
            <label for="message" class="block text-caption font-primary-semibold text-gray-900">Message</label>
            <div class="mt-2.5">
              <textarea 
                id="message" 
                name="message" 
                rows="4" 
                class="block w-full rounded-md bg-white px-3.5 py-2 text-body text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-primary-600"
              ></textarea>
            </div>
          </div>
        </div>
        
        <div class="mt-8 flex justify-end">
          <button 
            type="submit" 
            class="rounded-md bg-primary-600 px-3.5 py-2.5 text-center text-caption font-primary-semibold text-white shadow-xs hover:bg-primary-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
          >
            Send message
          </button>
        </div>
      </div>
    </form>
  </div>
</div>