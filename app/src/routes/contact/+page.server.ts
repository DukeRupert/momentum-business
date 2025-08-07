import { fail, type ServerLoad } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import { schema } from './schema';
import { z } from 'zod/v4'

export const load: ServerLoad = () => {
    let e = {}
    return { errors: e}
}

export const actions: Actions = {
    default: async ({ request }) => {
        console.log("Default form action")
        const formData = await request.formData();

        // Extract form data
        const rawData = {
            'first-name': formData.get('first-name') as string,
            'last-name': formData.get('last-name') as string,
            email: formData.get('email') as string,
            'phone-number': formData.get('phone-number') as string,
            'annual-revenue': formData.get('annual-revenue') as string,
            services: formData.getAll('services') as string[],
            message: formData.get('message') as string
        };

        console.log("raw data:")
        console.log(rawData)
        // Validate the form data
        const result = schema.safeParse(rawData);
        console.log("result:")
        console.log(result)
        if (!result.success) {
            let e = z.flattenError(result.error)
            console.log('Flatten Errors: ')
            console.log(e)
            return fail(400, {
                errors: e,
                data: rawData
            });
        }

        // If validation passes, you can process the form data here
        const validatedData = result.data;

        try {
            // TODO: Add your form processing logic here
            // Examples:
            // - Send email notification
            // - Save to database
            // - Integrate with CRM
            // - Send to third-party service

            console.log('Form submitted successfully:', validatedData);

            // For now, we'll simulate a successful submission
            // You would replace this with your actual processing logic

            return {
                success: true,
                message: 'Thank you for your message! We\'ll get back to you soon.'
            };

        } catch (error) {
            console.error('Error processing form:', error);

            return fail(500, {
                errors: {
                    _form: ['An error occurred while processing your request. Please try again.']
                },
                data: rawData
            });
        }
    }
};