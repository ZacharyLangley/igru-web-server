import { createFormContext } from '@mantine/form';

export interface AuthFormValues {
    email: string;
    userName: string;
    password: string;
    confirmPassword: string;
}

export const [AuthFormProvider, useAuthFormContext, useAuthForm] = createFormContext<AuthFormValues>();
