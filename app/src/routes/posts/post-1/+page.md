---
layout: '../BlogLayout.svelte'
title: 'Getting Started with mdsvex'
date: '2025-01-15'
author: 'Jane Developer'
tags: ['svelte', 'mdx', 'markdown', 'tutorial']
description: 'Learn how to combine the power of Svelte components with Markdown using mdsvex'
readingTime: '5'
---

<script>
  import { onMount } from 'svelte';
  
  // You can use Svelte logic in your markdown!
  let clickCount = 0;
  let mounted = false;
  
  onMount(() => {
    mounted = true;
  });
  
  function handleClick() {
    clickCount += 1;
  }
</script>

Welcome to this **mdsvex** tutorial! This markdown file demonstrates how you can seamlessly blend Markdown content with Svelte components.

## What is mdsvex?

mdsvex is a preprocessor for Svelte that allows you to write Svelte components inside Markdown files. Think of it as MDX for the Svelte ecosystem.

## Interactive Components

Here's a simple interactive component right inside our markdown:

<div class="interactive-demo">
  <p>You've clicked the button <strong>{clickCount}</strong> times!</p>
  <button on:click={handleClick} class="demo-button">
    Click me!
  </button>
  
  {#if mounted}
    <p class="mounted-message">✅ Component is mounted and reactive!</p>
  {/if}
</div>

## Code Examples

You can still use regular markdown code blocks:

```javascript

But you can also create interactive code examples with Svelte:

<div class="code-demo">
  <h4>Live Code Output:</h4>
  <p>2 + 2 = {2 + 2}</p>
  <p>Current time: {new Date().toLocaleTimeString()}</p>
</div>

## Conditional Content

{#if clickCount > 0}
  🎉 **Congratulations!** You've discovered the interactive nature of this document.
{:else}
  👆 Try clicking the button above to see conditional content!
{/if}

## Lists Work Great Too

Here are some benefits of using mdsvex:

- ✨ **Rich interactivity**: Add Svelte components anywhere
- 📝 **Familiar syntax**: Standard Markdown with Svelte enhancements  
- 🎨 **Custom layouts**: Use layout components like the one wrapping this content
- 🔧 **Full Svelte power**: Access to stores, lifecycle functions, and more

## Mathematical Expressions

You can even include dynamic calculations:

The area of a circle with radius 5 is **{Math.PI * Math.pow(5, 2).toFixed(2)}** square units.

---

*This entire post is written in Markdown but powered by Svelte components!*

<style>
  .interactive-demo {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    padding: 2rem;
    border-radius: 12px;
    margin: 2rem 0;
    text-align: center;
  }
  
  .demo-button {
    background: white;
    color: #667eea;
    border: none;
    padding: 0.75rem 1.5rem;
    border-radius: 8px;
    font-weight: 600;
    cursor: pointer;
    transition: transform 0.2s ease;
  }
  
  .demo-button:hover {
    transform: translateY(-2px);
  }
  
  .mounted-message {
    margin-top: 1rem;
    font-size: 0.9rem;
    opacity: 0.9;
  }
  
  .code-demo {
    background: #f8f9fa;
    border: 2px solid #e9ecef;
    border-radius: 8px;
    padding: 1.5rem;
    margin: 1.5rem 0;
  }
  
  .code-demo h4 {
    margin-top: 0;
    color: #495057;
  }
</style>
