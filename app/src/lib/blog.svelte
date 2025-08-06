<script>
  const blogPosts = [
    {
      date: "2020-03-16",
      dateDisplay: "Mar 16, 2020",
      category: "Marketing",
      categoryHref: "#",
      title: "Boost your conversion rate",
      titleHref: "#",
      excerpt: "Illo sint voluptas. Error voluptates culpa eligendi. Hic vel totam vitae illo. Non aliquid explicabo necessitatibus unde. Sed exercitationem placeat consectetur nulla deserunt vel. Iusto corrupti dicta.",
      authorImage: "https://images.unsplash.com/photo-1519244703995-f4e0f30006d5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
      authorName: "Michael Foster",
      authorTitle: "Co-Founder / CTO",
      authorHref: "#"
    },
    {
      date: "2020-03-10",
      dateDisplay: "Mar 10, 2020",
      category: "Sales",
      categoryHref: "#",
      title: "How to use search engine optimization to drive sales",
      titleHref: "#",
      excerpt: "Optio cum necessitatibus dolor voluptatum provident commodi et. Qui aperiam fugiat nemo cumque.",
      authorImage: "https://images.unsplash.com/photo-1517841905240-472988babdf9?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
      authorName: "Lindsay Walton",
      authorTitle: "Front-end Developer",
      authorHref: "#"
    },
    {
      date: "2020-02-12",
      dateDisplay: "Feb 12, 2020",
      category: "Business",
      categoryHref: "#",
      title: "Improve your customer experience",
      titleHref: "#",
      excerpt: "Cupiditate maiores ullam eveniet adipisci in doloribus nulla minus. Voluptas iusto libero adipisci rem et corporis. Nostrud sint anim sunt aliqua. Nulla eu labore irure incididunt velit cillum quis magna dolore.",
      authorImage: "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
      authorName: "Tom Cook",
      authorTitle: "Director of Product",
      authorHref: "#"
    }
  ];
</script>

<div class="bg-white py-24 sm:py-32">
  <div id="blog" class="mx-auto max-w-7xl px-6 lg:px-8">
    <div class="mx-auto max-w-2xl lg:mx-0">
      <h2 class="text-4xl font-semibold tracking-tight text-pretty text-gray-900 sm:text-5xl">From the blog</h2>
      <p class="mt-2 text-lg/8 text-gray-600">Learn how to grow your business with our expert advice.</p>
    </div>
    <div id="blog-posts" class="mx-auto mt-10 grid max-w-2xl grid-cols-1 gap-x-8 gap-y-16 border-t border-gray-200 pt-10 sm:mt-16 sm:pt-16 lg:mx-0 lg:max-w-none lg:grid-cols-3">
      {#each blogPosts as post}
        <article class="flex max-w-xl flex-col items-start justify-between">
          <div class="flex items-center gap-x-4 text-xs">
            <time datetime={post.date} class="text-gray-500">{post.dateDisplay}</time>
            <a href={post.categoryHref} class="relative z-10 rounded-full bg-gray-50 px-3 py-1.5 font-medium text-gray-600 hover:bg-gray-100">{post.category}</a>
          </div>
          <div class="group relative grow">
            <h3 class="mt-3 text-lg/6 font-semibold text-gray-900 group-hover:text-gray-600">
              <a href={post.titleHref}>
                <span class="absolute inset-0"></span>
                {post.title}
              </a>
            </h3>
            <p class="mt-5 line-clamp-3 text-sm/6 text-gray-600">{post.excerpt}</p>
          </div>
          <div class="relative mt-8 flex items-center gap-x-4 justify-self-end">
            <img src={post.authorImage} alt={post.authorName} class="size-10 rounded-full bg-gray-50" />
            <div class="text-sm/6">
              <p class="font-semibold text-gray-900">
                <a href={post.authorHref}>
                  <span class="absolute inset-0"></span>
                  {post.authorName}
                </a>
              </p>
              <p class="text-gray-600">{post.authorTitle}</p>
            </div>
          </div>
        </article>
      {/each}
    </div>
  </div>
</div>