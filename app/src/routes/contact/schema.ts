import { z } from 'zod/v4';

export const schema = z.object({
    'first-name': z
        .string()
        .min(1, 'First name is required')
        .min(2, 'First name must be at least 2 characters')
        .max(50, 'First name must be less than 50 characters')
        .regex(/^[a-zA-Z\s'-]+$/, 'First name can only contain letters, spaces, hyphens, and apostrophes'),

    'last-name': z
        .string()
        .min(1, 'Last name is required')
        .min(2, 'Last name must be at least 2 characters')
        .max(50, 'Last name must be less than 50 characters')
        .regex(/^[a-zA-Z\s'-]+$/, 'Last name can only contain letters, spaces, hyphens, and apostrophes'),

    email: z
        .email('Please enter a valid email address')
        .min(1, 'Email is required')
        .max(254, 'Email must be less than 254 characters'),

    'phone-number': z
        .string()
        .min(1, 'Phone number is required')
        .regex(
            /^[\+]?[1-9]?[\d\s\-\(\)\.]{10,15}$/,
            'Please enter a valid phone number'
        ),

    'annual-revenue': z
        .string()
        .min(1, 'Please select your annual revenue range'),

    services: z
        .array(z.string())
        .min(1, 'Please select at least one service you\'re interested in')
        .refine(
            (services) => services.every(service =>
                ['essentials', 'growth-strategy', 'executive-operations', 'complete-support', 'cleanup'].includes(service)
            ),
            'Invalid service selection'
        ),

    message: z
        .string()
        .max(2000, 'Message must be less than 2000 characters')
})

export type Schema = typeof schema