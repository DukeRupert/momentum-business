<script lang="ts">
  import { onMount } from 'svelte';
	import { quadOut, quintIn, quintOut } from 'svelte/easing';
  import { fly, fade } from 'svelte/transition';

  const testimonials = [
    {
      companyLogo: "https://tailwindcss.com/plus-assets/img/logos/tuple-logo-gray-900.svg",
      companyName: "Tuple",
      quote: "Amet amet eget scelerisque tellus sit neque faucibus non eleifend. Integer eu praesent at a. Ornare arcu gravida natoque erat et cursus tortor consequat at. Vulputate gravida sociis enim nullam ultricies habitant malesuada lorem ac. Tincidunt urna dui pellentesque sagittis.",
      authorImage: "https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
      authorName: "Judith Black",
      authorTitle: "CEO of Tuple"
    },
    {
      companyLogo: "https://tailwindcss.com/plus-assets/img/logos/reform-logo-gray-900.svg",
      companyName: "Reform",
      quote: "Excepteur veniam labore ullamco eiusmod. Pariatur consequat proident duis dolore nulla veniam reprehenderit nisi officia voluptate incididunt exercitation exercitation elit. Nostrud veniam sint dolor nisi ullamco.",
      authorImage: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
      authorName: "Joseph Rodriguez",
      authorTitle: "CEO of Reform"
    }
  ];

  let visible = [false, false];
  let testimonialsSection: HTMLElement;

  onMount(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            // Animate first testimonial immediately
            visible[0] = true;
            
            // Animate second testimonial with a delay
            setTimeout(() => {
              visible[1] = true;
            }, 400); // 200ms delay
            
            // Stop observing once animations are triggered
            observer.unobserve(entry.target);
          }
        });
      },
      {
        threshold: 0.5, // Trigger when 20% of the section is visible
        rootMargin: '-50px' // Start animation 50px before entering viewport
      }
    );

    if (testimonialsSection) {
      observer.observe(testimonialsSection);
    }

    // Cleanup
    return () => {
      if (testimonialsSection) {
        observer.unobserve(testimonialsSection);
      }
    };
  });
</script>

<section 
  bind:this={testimonialsSection}
  id="testimonials" 
  class="bg-white py-24 sm:py-32"
>
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <div class="mx-auto grid max-w-2xl grid-cols-1 lg:mx-0 lg:max-w-none lg:grid-cols-2">
      {#each testimonials as testimonial, index}
        {#if visible[index]}
          <div 
            class="flex flex-col pb-10 sm:pb-16 {index === 0 ? 'lg:pr-8 lg:pb-0 xl:pr-20' : 'border-t border-gray-900/10 pt-10 sm:pt-16 lg:border-t-0 lg:border-l lg:pt-0 lg:pl-8 xl:pl-20'}"
            in:fly={{ y: 50, duration: 600, delay: 200, opacity: 100, easing: quadOut }}
          >
            <img 
              src={testimonial.companyLogo} 
              alt={testimonial.companyName} 
              class="h-12 self-start" 
            />
            <figure class="mt-10 flex flex-auto flex-col justify-between">
              <blockquote class="text-body-lg text-gray-900">
                <p>"{testimonial.quote}"</p>
              </blockquote>
              <figcaption class="mt-10 flex items-center gap-x-6">
                <img 
                  src={testimonial.authorImage} 
                  alt={testimonial.authorName} 
                  class="size-14 rounded-full bg-gray-50" 
                />
                <div class="text-base">
                  <div class="text-kicker font-primary-semibold text-primary-600 uppercase tracking-wide">{testimonial.authorName}</div>
                  <div class="mt-1 text-caption text-gray-600">{testimonial.authorTitle}</div>
                </div>
              </figcaption>
            </figure>
          </div>
        {/if}
      {/each}
    </div>
  </div>
</section>