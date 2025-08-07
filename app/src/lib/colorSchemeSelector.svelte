<script>
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';

	let selectedScheme = $state('green');
	let isExpanded = $state(false);

	// Color palettes
	const colorSchemes = {
		green: {
			name: 'Tom Thumb',
			colors: {
				50: '#f4f9f5',
				100: '#e6f2e7',
				200: '#cde5d1',
				300: '#a6cfac',
				400: '#77b17f',
				500: '#54935d',
				600: '#417848',
				700: '#355e3b',
				800: '#2e4d33',
				900: '#27402b',
				950: '#122115'
			}
		},
		blue: {
			name: 'Cloud Burst',
			colors: {
				50: '#f1f5fd',
				100: '#dfe9fa',
				200: '#c6d9f7',
				300: '#9ec1f2',
				400: '#709fea',
				500: '#4f7ee2',
				600: '#3a61d6',
				700: '#314ec4',
				800: '#2d41a0',
				900: '#293a7f',
				950: '#1e2654'
			}
		},
		purple: {
			name: 'Deep Purple',
			colors: {
				50: '#f4f3ff',
				100: '#ebe9fe',
				200: '#d9d6fe',
				300: '#bfb8fd',
				400: '#a090fa',
				500: '#8b5cf6',
				600: '#7c3aed',
				700: '#6d28d9',
				800: '#5b21b6',
				900: '#4c1d95',
				950: '#2e1065'
			}
		},
		charcoal: {
			name: 'Steel Gray',
			colors: {
				50: '#f8f9fa',
				100: '#f1f3f4',
				200: '#e8eaed',
				300: '#dadce0',
				400: '#bdc1c6',
				500: '#9aa0a6',
				600: '#80868b',
				700: '#5f6368',
				800: '#3c4043',
				900: '#202124',
				950: '#171717'
			}
		},
		teal: {
			name: 'Ocean Depths',
			colors: {
				50: '#f0fdfa',
				100: '#ccfbf1',
				200: '#99f6e4',
				300: '#5eead4',
				400: '#2dd4bf',
				500: '#14b8a6',
				600: '#0d9488',
				700: '#0f766e',
				800: '#115e59',
				900: '#134e4a',
				950: '#042f2e'
			}
		}
	};

	// Apply color scheme to CSS custom properties
	function applyColorScheme(scheme) {
		if (!browser) return;

		const colors = colorSchemes[scheme].colors;
		const root = document.documentElement;

		Object.entries(colors).forEach(([shade, color]) => {
			root.style.setProperty(`--primary-${shade}`, color);
		});

		// Store preference in localStorage
		localStorage.setItem('colorScheme', scheme);
	}

	// Handle scheme change
	function handleSchemeChange(event) {
		selectedScheme = event.target.value;
		applyColorScheme(selectedScheme);
	}

	// Toggle expanded state
	function toggleExpanded() {
		isExpanded = !isExpanded;
	}

	// Load saved preference on mount
	onMount(() => {
		if (browser) {
			const savedScheme = localStorage.getItem('colorScheme') || 'green';
			selectedScheme = savedScheme;
			applyColorScheme(savedScheme);
		}
	});

	// Reactive statement to apply colors when selectedScheme changes
	$effect(() => {
		applyColorScheme(selectedScheme);
	});
</script>

<div
	class="fixed right-0 top-4 z-50 transition-transform duration-300 ease-in-out"
	class:translate-x-56={!isExpanded}
>
	<!-- Collapsible Tab -->
	<button
		on:click={toggleExpanded}
		class="absolute -left-10 top-4 rounded-l-lg border border-gray-200 bg-white p-2 shadow-md transition-colors hover:bg-gray-50"
		class:shadow-lg={!isExpanded}
	>
		<div class="flex flex-col items-center gap-1">
			<!-- Color indicator dot -->
			<div
				class="h-3 w-3 rounded-full border border-gray-200"
				style="background-color: {colorSchemes[selectedScheme].colors[600]}"
			></div>

			<!-- Expand/Collapse icon -->
			<svg
				class="h-4 w-4 text-gray-600 transition-transform duration-200"
				class:rotate-180={isExpanded}
				fill="none"
				stroke="currentColor"
				viewBox="0 0 24 24"
			>
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
			</svg>
		</div>
	</button>

	<!-- Main Panel -->
	<div class="mr-4 min-w-[200px] rounded-l-lg border border-gray-200 bg-white p-4 shadow-lg">
		<div class="mb-3 flex items-center justify-between">
			<div class="flex items-center gap-2">
				<!-- Palette Icon -->
				<svg class="h-4 w-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z"
					/>
				</svg>
				<span class="text-sm font-medium text-gray-900">Color Scheme</span>
			</div>

			<!-- Close button -->
			<button on:click={toggleExpanded} class="text-gray-400 transition-colors hover:text-gray-600">
				<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M6 18L18 6M6 6l12 12"
					/>
				</svg>
			</button>
		</div>

		<div class="space-y-2">
			{#each Object.entries(colorSchemes) as [key, scheme]}
				<label
					class="flex cursor-pointer items-center gap-3 rounded-md p-2 transition-colors hover:bg-gray-50"
				>
					<input
						type="radio"
						name="colorScheme"
						value={key}
						bind:group={selectedScheme}
						on:change={handleSchemeChange}
						class="h-4 w-4 border-gray-300 bg-gray-100 text-blue-600 focus:ring-blue-500"
					/>

					<div class="flex flex-1 items-center gap-2">
						<div
							class="h-4 w-4 rounded-full border border-gray-200"
							style="background-color: {scheme.colors[600]}"
						></div>
						<span class="text-sm text-gray-700">{scheme.name}</span>
					</div>
				</label>
			{/each}
		</div>

		<!-- Color Preview -->
		<div class="mt-4 border-t border-gray-100 pt-3">
			<div class="mb-2 text-xs text-gray-500">Preview:</div>
			<div class="flex gap-1">
				{#each [500, 600, 700, 800] as shade}
					<div
						class="h-6 w-6 rounded border border-gray-200"
						style="background-color: {colorSchemes[selectedScheme].colors[shade]}"
						title="{shade}: {colorSchemes[selectedScheme].colors[shade]}"
					></div>
				{/each}
			</div>
		</div>

		<!-- Demo Elements -->
		<div class="mt-4 space-y-2 border-t border-gray-100 pt-3">
			<div class="mb-2 text-xs text-gray-500">Live Demo:</div>

			<button
				class="w-full rounded-md px-3 py-2 text-sm font-medium text-white transition-colors hover:opacity-90"
				style="background-color: {colorSchemes[selectedScheme].colors[600]}"
			>
				Primary Button
			</button>

			<div
				class="w-full rounded-md border-2 px-3 py-2 text-center text-sm transition-colors"
				style="border-color: {colorSchemes[selectedScheme].colors[600]}; color: {colorSchemes[
					selectedScheme
				].colors[600]}; background-color: {colorSchemes[selectedScheme].colors[50]}"
			>
				Outlined Element
			</div>
		</div>
	</div>
</div>
