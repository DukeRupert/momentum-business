import { fail } from '@sveltejs/kit';
import type { Actions } from './$types';
import { schema } from './schema';
import { z } from 'zod/v4'
import { POSTMARK_TO, POSTMARK_FROM, POSTMARK_API_KEY } from '$env/static/private'
import { sendContactFormEmail, sendThankYouEmail, type ContactFormData } from './email';

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
            return fail(400, {
                errors: e,
                data: rawData
            });
        }

        // If validation passes, you can process the form data here
        const validatedData = result.data;

        try {
            // Send notification email to client
            const clientEmailResult = await sendContactFormEmail(validatedData, POSTMARK_API_KEY, POSTMARK_TO, POSTMARK_FROM);

            if (!clientEmailResult.success) {
                throw new Error(`Failed to send client notification: ${clientEmailResult.error}`);
            }

            // Send thank you email to form submitter
            const thankYouResult = await sendThankYouEmail({
                formData: validatedData,
                postmarkToken: POSTMARK_API_KEY,
                fromEmail: POSTMARK_FROM
            });

            if (!thankYouResult.success) {
                console.warn('Thank you email failed, but client notification succeeded:', thankYouResult.error);
                // Don't fail the entire process if thank you email fails
            }

            return {
                success: true,
                clientEmailSent: clientEmailResult.success,
                thankYouEmailSent: thankYouResult.success,
                data: {
                    clientEmail: clientEmailResult.data,
                    thankYouEmail: thankYouResult.data
                }
            };

        } catch (error) {
            console.error('Contact form submission failed:', error);
            return {
                success: false,
                error: (error as Error).message,
                clientEmailSent: false,
                thankYouEmailSent: false
            };
        }
    }
};