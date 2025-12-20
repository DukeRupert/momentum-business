<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import Navigation from '$lib/navigation.svelte';
	import Footer from '$lib/footer.svelte';
	import Banner from '$lib/banner.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="bg-white">
	<Banner />
	<Navigation />
	{@render children?.()}
	<Footer />
</div>
