<script>
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
</script>

<section id="testimonials" class="bg-white py-24 sm:py-32">
  <div class="mx-auto max-w-7xl px-6 lg:px-8">
    <div class="mx-auto grid max-w-2xl grid-cols-1 lg:mx-0 lg:max-w-none lg:grid-cols-2">
      {#each testimonials as testimonial, index}
        <div class="flex flex-col pb-10 sm:pb-16 {index === 0 ? 'lg:pr-8 lg:pb-0 xl:pr-20' : 'border-t border-gray-900/10 pt-10 sm:pt-16 lg:border-t-0 lg:border-l lg:pt-0 lg:pl-8 xl:pl-20'}">
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
      {/each}
    </div>
  </div>
</section>