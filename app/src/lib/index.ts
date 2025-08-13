/**
 * Configuration options for slugify function
 */
export interface SlugifyOptions {
    /** Maximum length of the slug (default: 60) */
    maxLength?: number;
    /** Whether to use strict mode for more aggressive character removal (default: false) */
    strict?: boolean;
}

/**
 * Convert a string to a URL-friendly slug
 * @param text - The text to slugify
 * @param options - Configuration options
 * @returns The slugified text
 */
export function slugify(text: string, options: SlugifyOptions = {}): string {
    const { maxLength = 60, strict = false } = options;
    
    if (!text || typeof text !== 'string') {
        return '';
    }

    let slug = text
        // Convert to lowercase
        .toLowerCase()
        // Remove HTML tags if any
        .replace(/<[^>]*>/g, '')
        // Replace spaces, underscores, and multiple hyphens with single hyphen
        .replace(/[\s_-]+/g, '-')
        // Remove special characters (keep alphanumeric, hyphens)
        .replace(/[^a-z0-9-]/g, '')
        // Remove leading/trailing hyphens
        .replace(/^-+|-+$/g, '');

    if (strict) {
        // In strict mode, be more aggressive about removing characters
        slug = slug.replace(/[^a-z0-9-]/g, '');
    }

    // Ensure no double hyphens
    slug = slug.replace(/-+/g, '-');

    // Truncate to maxLength if specified
    if (maxLength && slug.length > maxLength) {
        slug = slug.substring(0, maxLength);
        // Don't end with a hyphen after truncation
        slug = slug.replace(/-+$/, '');
    }

    return slug;
}

/**
 * Generate a unique slug by appending a number if needed
 * @param baseSlug - The base slug
 * @param existingSlugs - Array of existing slugs to check against
 * @returns A unique slug
 */
export function createUniqueSlug(baseSlug: string, existingSlugs: string[] = []): string {
    let slug = baseSlug;
    let counter = 1;

    while (existingSlugs.includes(slug)) {
        slug = `${baseSlug}-${counter}`;
        counter++;
    }

    return slug;
}

/**
 * Slugify specifically for blog posts with sensible defaults
 * @param title - Blog post title
 * @param existingSlugs - Optional array of existing slugs
 * @returns Blog post slug
 */
export function slugifyBlogPost(title: string, existingSlugs: string[] = []): string {
    const baseSlug = slugify(title, { maxLength: 60 });
    
    if (existingSlugs.length > 0) {
        return createUniqueSlug(baseSlug, existingSlugs);
    }
    
    return baseSlug;
}

// Export as default for convenience
export default slugify;