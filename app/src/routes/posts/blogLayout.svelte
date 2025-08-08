<script>
  export let title = '';
  export let date = '';
  export let author = '';
  export let tags = [];
  export let description = '';
  export let readingTime = '';
</script>

<svelte:head>
  <title>{title}</title>
  <meta name="description" content={description} />
</svelte:head>

<article class="blog-post">
  <header class="post-header">
    <h1 class="post-title">{title}</h1>
    
    <div class="post-meta">
      {#if author}
        <span class="author">By {author}</span>
      {/if}
      
      {#if date}
        <time class="date" datetime={date}>
          {new Date(date).toLocaleDateString('en-US', { 
            year: 'numeric', 
            month: 'long', 
            day: 'numeric' 
          })}
        </time>
      {/if}
      
      {#if readingTime}
        <span class="reading-time">{readingTime} min read</span>
      {/if}
    </div>
    
    {#if tags && tags.length > 0}
      <div class="tags">
        {#each tags as tag}
          <span class="tag">#{tag}</span>
        {/each}
      </div>
    {/if}
  </header>

  <main class="post-content">
    <!-- This is where the markdown content will be rendered -->
    <slot />
  </main>

  <footer class="post-footer">
    <p>Thanks for reading! Feel free to share your thoughts.</p>
  </footer>
</article>

<style>
  .blog-post {
    max-width: 800px;
    margin: 0 auto;
    padding: 2rem;
    font-family: system-ui, -apple-system, sans-serif;
    line-height: 1.6;
    color: #333;
  }

  .post-header {
    margin-bottom: 3rem;
    border-bottom: 2px solid #eee;
    padding-bottom: 2rem;
  }

  .post-title {
    font-size: 2.5rem;
    font-weight: 700;
    margin: 0 0 1rem 0;
    color: #2c3e50;
  }

  .post-meta {
    display: flex;
    gap: 1rem;
    align-items: center;
    font-size: 0.9rem;
    color: #666;
    margin-bottom: 1rem;
    flex-wrap: wrap;
  }

  .author {
    font-weight: 500;
  }

  .date {
    font-style: italic;
  }

  .reading-time {
    background: #f8f9fa;
    padding: 0.25rem 0.5rem;
    border-radius: 4px;
    font-size: 0.8rem;
  }

  .tags {
    display: flex;
    gap: 0.5rem;
    flex-wrap: wrap;
  }

  .tag {
    background: #e3f2fd;
    color: #1976d2;
    padding: 0.25rem 0.75rem;
    border-radius: 20px;
    font-size: 0.8rem;
    font-weight: 500;
  }

  .post-content {
    margin-bottom: 3rem;
  }

  /* Styling for markdown content */
  .post-content :global(h2) {
    font-size: 1.8rem;
    margin: 2rem 0 1rem 0;
    color: #2c3e50;
  }

  .post-content :global(h3) {
    font-size: 1.4rem;
    margin: 1.5rem 0 0.75rem 0;
    color: #34495e;
  }

  .post-content :global(p) {
    margin-bottom: 1.25rem;
  }

  .post-content :global(blockquote) {
    border-left: 4px solid #3498db;
    padding-left: 1.5rem;
    margin: 1.5rem 0;
    font-style: italic;
    color: #555;
  }

  .post-content :global(code) {
    background: #f8f9fa;
    padding: 0.2rem 0.4rem;
    border-radius: 4px;
    font-family: 'Fira Code', monospace;
    font-size: 0.9rem;
  }

  .post-content :global(pre) {
    background: #282c34;
    color: #abb2bf;
    padding: 1.5rem;
    border-radius: 8px;
    overflow-x: auto;
    margin: 1.5rem 0;
  }

  .post-content :global(pre code) {
    background: none;
    padding: 0;
    color: inherit;
  }

  .post-footer {
    border-top: 1px solid #eee;
    padding-top: 2rem;
    text-align: center;
    color: #666;
    font-style: italic;
  }

  /* Responsive design */
  @media (max-width: 768px) {
    .blog-post {
      padding: 1rem;
    }
    
    .post-title {
      font-size: 2rem;
    }
    
    .post-meta {
      flex-direction: column;
      align-items: flex-start;
      gap: 0.5rem;
    }
  }
</style>
