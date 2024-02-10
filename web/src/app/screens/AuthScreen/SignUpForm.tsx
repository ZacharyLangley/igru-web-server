import React from 'react'
import { MantineStyleProp, PasswordInput, Stack, Text, TextInput } from '@mantine/core';

import { useAuthFormContext } from '../../../common/contexts/authenticationContext';
import language from '../../../common/language/index';

const lang = language();
const inputStyle = { width: '100%' };

export const SignUpForm = React.memo(() => {
    const form = useAuthFormContext();
    return (
        <Stack gap="md" justify="space-between" align="flex-start" flex={1}>
            <TextInput
                withAsterisk
                radius={'md'}
                label={lang.inputs.email.label}
                style={inputStyle}
                {...form.getInputProps('email')} />
            <TextInput
                withAsterisk
                radius={'md'}
                label={lang.inputs.name.label}
                description={lang.inputs.name.description}
                style={inputStyle}
                {...form.getInputProps('userName')} />
            <PasswordInput
                withAsterisk
                label={lang.inputs.password.label}
                description={lang.inputs.password.description}
                radius={'md'}
                style={inputStyle}
                {...form.getInputProps('password')} />
            <PasswordInput
                withAsterisk
                label={lang.inputs.confirmPassword.label}
                description={lang.inputs.confirmPassword.description}
                radius={'md'}
                style={inputStyle}
                {...form.getInputProps('confirmPassword')} />
        </Stack>
    )
})

const messageStyle: MantineStyleProp = { textAlign: 'center', maxWidth: 400 };

export const SignUpFailure = React.memo(() => {
    return (
        <Stack>
            <Text size={'xl'} style={messageStyle}>{lang.auth.signUpError.title}</Text>
            <Text style={messageStyle}>{lang.auth.signUpError.message}</Text>
        </Stack>
    )
});

export const SignUpSuccess = React.memo(() => {
    return (
        <Stack>
            <Text size={'xl'} style={messageStyle}>{lang.auth.signUpSuccess.title}</Text>
            <Text style={messageStyle}>{lang.auth.signUpSuccess.message}</Text>
        </Stack>
    )
});