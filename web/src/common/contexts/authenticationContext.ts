import { createFormContext } from '@mantine/form';

interface AuthFormValues {
    email: string;
    userName: string;
    password: string;
    confirmPassword: string;
}

export const [AuthFormProvider, useAuthFormContext, useAuthForm] = createFormContext<AuthFormValues>();
