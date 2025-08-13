<script lang="ts">
    import SEO from '$lib/seo.svelte';
    import Logo from '$lib/assets/logo.png';

    // Form state
    let email: string = $state('');
    let password: string = $state('');
    let isLoading: boolean = $state(false);

    // Form validation
    let emailError: string = $state('');
    let passwordError: string = $state('');

    function validateEmail(email: string): boolean {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    }

    function handleSubmit(event: Event): void {
        event.preventDefault();
        
        // Reset errors
        emailError = '';
        passwordError = '';

        // Validate form
        let hasErrors = false;

        if (!email) {
            emailError = 'Email is required';
            hasErrors = true;
        } else if (!validateEmail(email)) {
            emailError = 'Please enter a valid email address';
            hasErrors = true;
        }

        if (!password) {
            passwordError = 'Password is required';
            hasErrors = true;
        } else if (password.length < 8) {
            passwordError = 'Password must be at least 8 characters';
            hasErrors = true;
        }

        if (hasErrors) return;

        // Simulate loading state
        isLoading = true;
        
        // For demo purposes, show alert after delay
        setTimeout(() => {
            isLoading = false;
            alert('Login functionality coming soon! This is just a preview of the interface.');
        }, 1500);
    }
</script>

<SEO 
    title="Client Login"
    description="Secure client portal login for Momentum Business Solutions. Access your financial reports, documents, and account information."
    canonical="/login"
    noindex={true}
/>

<div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
        <img 
            src={Logo} 
            alt="Momentum Business Solutions" 
            class="mx-auto h-24 sm:h-36 lg:h-48 w-auto rounded-sm"
        />
        <h2 class="mt-10 text-center text-2xl/9 font-bold tracking-tight text-gray-900">
            Sign in to your account
        </h2>
    </div>

    <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
        <form class="space-y-6" onsubmit={handleSubmit}>
            <div>
                <label for="email" class="block text-sm/6 font-medium text-gray-900">
                    Email address
                </label>
                <div class="mt-2">
                    <input 
                        id="email" 
                        type="email" 
                        name="email" 
                        required 
                        autocomplete="email"
                        bind:value={email}
                        class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-primary-600 sm:text-sm/6 {emailError ? 'outline-red-300 focus:outline-red-600' : ''}"
                        placeholder="Enter your email"
                    />
                    {#if emailError}
                        <p class="mt-2 text-sm text-red-600">{emailError}</p>
                    {/if}
                </div>
            </div>

            <div>
                <div class="flex items-center justify-between">
                    <label for="password" class="block text-sm/6 font-medium text-gray-900">
                        Password
                    </label>
                    <div class="text-sm">
                        <a href="#" class="font-semibold text-primary-600 hover:text-primary-500">
                            Forgot password?
                        </a>
                    </div>
                </div>
                <div class="mt-2">
                    <input 
                        id="password" 
                        type="password" 
                        name="password" 
                        required 
                        autocomplete="current-password"
                        bind:value={password}
                        class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-primary-600 sm:text-sm/6 {passwordError ? 'outline-red-300 focus:outline-red-600' : ''}"
                        placeholder="Enter your password"
                    />
                    {#if passwordError}
                        <p class="mt-2 text-sm text-red-600">{passwordError}</p>
                    {/if}
                </div>
            </div>

            <div>
                <button 
                    type="submit"
                    disabled={isLoading}
                    class="flex w-full justify-center rounded-md bg-primary-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-primary-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600 disabled:opacity-50 disabled:cursor-not-allowed"
                >
                    {#if isLoading}
                        <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                        Signing in...
                    {:else}
                        Sign in
                    {/if}
                </button>
            </div>
        </form>

        <p class="mt-10 text-center text-sm/6 text-gray-500">
            Need access to your financial dashboard?
            <a href="/contact" class="font-semibold text-primary-600 hover:text-primary-500">
                Contact us
            </a>
        </p>
    </div>
</div>