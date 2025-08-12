<script lang="ts">
	interface ZodErrors {
		formErrors: any[];
		fieldErrors: FieldErrors;
	}

	interface FieldErrors {
		'first-name': string[];
		'last-name': string[];
		email: string[];
		'phone-number': string[];
		'annual-revenue': string[];
		message: string[];
	}

	let { errors = null }: { errors: ZodErrors | null } = $props();

	// Helper function to format field names for display
	function formatFieldName(fieldName: string) {
		return fieldName
			.split('-')
			.map((word) => word.charAt(0).toUpperCase() + word.slice(1))
			.join(' ');
	}

	// Check if there are any errors to display
	let hasErrors = $derived(
		errors &&
			((errors.formErrors && errors.formErrors.length > 0) ||
				(errors.fieldErrors && Object.keys(errors.fieldErrors).length > 0))
	);

	// Get all field errors as a flat array for easier iteration
	let fieldErrorsList = $derived(
		errors?.fieldErrors
			? Object.entries(errors.fieldErrors).map(([field, fieldErrors]) => ({
					field,
					displayName: formatFieldName(field),
					errors: fieldErrors
				}))
			: []
	);
</script>

{#if hasErrors}
	<div class="mt-6 rounded-md border border-red-200 bg-red-50 p-4">
		<div class="flex">
			<div class="flex-shrink-0">
				<!-- Error icon -->
				<svg
					class="h-5 w-5 text-red-400"
					viewBox="0 0 20 20"
					fill="currentColor"
					aria-hidden="true"
				>
					<path
						fill-rule="evenodd"
						d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.28 7.22a.75.75 0 00-1.06 1.06L8.94 10l-1.72 1.72a.75.75 0 101.06 1.06L10 11.06l1.72 1.72a.75.75 0 101.06-1.06L11.06 10l1.72-1.72a.75.75 0 00-1.06-1.06L10 8.94 8.28 7.22z"
						clip-rule="evenodd"
					/>
				</svg>
			</div>
			<div class="ml-3">
				<h3 class="text-sm font-medium text-red-800">Please correct the following errors:</h3>
				<div class="mt-2">
					<ul class="space-y-1 text-sm text-red-700">
						<!-- Form-level errors -->
						{#if errors && errors.formErrors && errors.formErrors.length > 0}
							{#each errors.formErrors as formError}
								<li class="flex items-start">
									<span class="mr-2">•</span>
									<span>{formError}</span>
								</li>
							{/each}
						{/if}

						<!-- Field-level errors -->
						{#each fieldErrorsList as { field, displayName, errors: fieldErrors }}
							{#each fieldErrors as error}
								<li class="flex items-start">
									<span class="mr-2">•</span>
									<span><strong>{displayName}:</strong> {error}</span>
								</li>
							{/each}
						{/each}
					</ul>
				</div>
			</div>
		</div>
	</div>
{/if}
