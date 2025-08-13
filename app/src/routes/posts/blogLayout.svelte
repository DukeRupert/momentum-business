<script lang="ts">
	import { slugifyBlogPost } from '$lib';
	import Cta from '$lib/cta.svelte';
	import Seo from '$lib/seo.svelte';
	export let category = '';
	export let title = '';
	export let subtitle = '';
	export let date = '';
	export let author = '';
	export let authorRole = '';
	export let authorImage = '';
	export let tags: string[] = [];
	export let description = '';
	export let readingTime = '';
	export let featuredImage = '';
	export let featuredImageCaption = '';
	export let highlights: { title: string; description: string }[] = [];

	// Auto-generate slug from title
	$: slug = slugifyBlogPost(title);
</script>

<Seo {title} {description} canonical={`/posts/${slug}`} />

<div class="bg-white px-6 py-32 lg:px-8">
	<div class="mx-auto max-w-3xl text-base/7 text-gray-700">
		<!-- Header Section -->
		<header class="post-header mb-10 border-b-2 border-gray-100 pb-8">
			{#if category}
				<p class="text-primary-600 text-base/7 font-semibold">{category}</p>
			{/if}

			<h1 class="mt-2 text-pretty text-4xl font-semibold tracking-tight text-gray-900 sm:text-5xl">
				{title}
			</h1>

			{#if subtitle}
				<p class="mt-6 text-xl/8 text-gray-600">{subtitle}</p>
			{/if}

			<!-- Post Meta Information -->
			<div class="mt-6 flex flex-wrap items-center gap-4 text-sm text-gray-500">
				{#if author}
					<div class="flex items-center gap-2">
						{#if authorImage}
							<img src={authorImage} alt={author} class="size-6 rounded-full bg-gray-50" />
						{/if}
						<span class="font-semibold text-gray-900">{author}</span>
						{#if authorRole}
							<span>– {authorRole}</span>
						{/if}
					</div>
				{/if}

				{#if date}
					<time class="font-medium" datetime={date}>
						{new Date(date).toLocaleDateString('en-US', {
							year: 'numeric',
							month: 'long',
							day: 'numeric'
						})}
					</time>
				{/if}

				{#if readingTime}
					<span class="rounded-md bg-gray-100 px-2 py-1 text-xs font-medium">
						{readingTime} min read
					</span>
				{/if}
			</div>

			{#if tags && tags.length > 0}
				<div class="mt-4 flex flex-wrap gap-2">
					{#each tags as tag}
						<span class="bg-primary-50 text-primary-700 rounded-full px-3 py-1 text-sm font-medium">
							#{tag}
						</span>
					{/each}
				</div>
			{/if}
		</header>

		<!-- Main Content Area -->
		<div class="max-w-2xl text-gray-600">
			<slot name="intro" />

			<!-- Highlights/Features List -->
			{#if highlights && highlights.length > 0}
				<ul role="list" class="mt-8 max-w-xl space-y-8 text-gray-600">
					{#each highlights as highlight}
						<li class="flex gap-x-3">
							<svg
								viewBox="0 0 20 20"
								fill="currentColor"
								data-slot="icon"
								aria-hidden="true"
								class="text-primary-600 mt-1 size-5 flex-none"
							>
								<path
									d="M10 18a8 8 0 1 0 0-16 8 8 0 0 0 0 16Zm3.857-9.809a.75.75 0 0 0-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 1 0-1.06 1.061l2.5 2.5a.75.75 0 0 0 1.137-.089l4-5.5Z"
									clip-rule="evenodd"
									fill-rule="evenodd"
								/>
							</svg>
							<span>
								<strong class="font-semibold text-gray-900">{highlight.title}</strong>
								{highlight.description}
							</span>
						</li>
					{/each}
				</ul>
			{/if}

			<slot name="content" />
		</div>

		<!-- Featured Image -->
		{#if featuredImage}
			<figure class="mt-16">
				<img
					src={featuredImage}
					alt={featuredImageCaption || 'Featured image'}
					class="aspect-video rounded-xl bg-gray-50 object-cover"
				/>
				{#if featuredImageCaption}
					<figcaption class="mt-4 flex gap-x-2 text-sm/6 text-gray-500">
						<svg
							viewBox="0 0 20 20"
							fill="currentColor"
							data-slot="icon"
							aria-hidden="true"
							class="mt-0.5 size-5 flex-none text-gray-300"
						>
							<path
								d="M18 10a8 8 0 1 1-16 0 8 8 0 0 1 16 0Zm-7-4a1 1 0 1 1-2 0 1 1 0 0 1 2 0ZM9 9a.75.75 0 0 0 0 1.5h.253a.25.25 0 0 1 .244.304l-.459 2.066A1.75 1.75 0 0 0 10.747 15H11a.75.75 0 0 0 0-1.5h-.253a.25.25 0 0 1-.244-.304l.459-2.066A1.75 1.75 0 0 0 9.253 9H9Z"
								clip-rule="evenodd"
								fill-rule="evenodd"
							/>
						</svg>
						{featuredImageCaption}
					</figcaption>
				{/if}
			</figure>
		{/if}

		<!-- Additional Content -->
		<div class="mt-16 max-w-2xl text-gray-600">
			<slot name="additional" />
			<slot />
		</div>
	</div>
</div>
<Cta />

<!-- <style>
  /* Global styles for markdown content */
  :global(.prose h2) {
    @apply text-3xl font-semibold tracking-tight text-pretty text-gray-900 mt-16 mb-6;
  }

  :global(.prose h3) {
    @apply text-2xl font-semibold tracking-tight text-gray-900 mt-12 mb-4;
  }

  :global(.prose p) {
    @apply mt-6 text-base/7 text-gray-600;
  }

  :global(.prose blockquote) {
    @apply mt-10 border-l border-primary-600 pl-9;
  }

  :global(.prose blockquote p) {
    @apply font-semibold text-gray-900 text-lg;
  }

  :global(.prose code) {
    @apply bg-gray-100 text-gray-900 px-2 py-1 rounded text-sm font-mono;
  }

  :global(.prose pre) {
    @apply bg-gray-900 text-gray-100 p-6 rounded-xl overflow-x-auto mt-8;
  }

  :global(.prose pre code) {
    @apply bg-transparent text-inherit p-0;
  }

  :global(.prose ul) {
    @apply mt-6 space-y-2;
  }

  :global(.prose li) {
    @apply text-gray-600;
  }

  :global(.prose strong) {
    @apply font-semibold text-gray-900;
  }

  :global(.prose a) {
    @apply text-primary-600 hover:text-primary-500 underline decoration-primary-200 hover:decoration-primary-300;
  }
</style> -->
