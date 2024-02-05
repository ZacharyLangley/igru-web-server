import React from 'react'
import { Center, MantineStyleProp, PasswordInput, Stack, Text, TextInput, Title } from '@mantine/core';

import { useAuthFormContext } from '../../../common/contexts/authenticationContext';

const inputStyle = { width: '100%' }

export const SignUpForm = React.memo(() => {
    const form = useAuthFormContext();
    return (
        <Stack gap="md" justify="space-between" align="flex-start" flex={1}>
            <TextInput
                withAsterisk
                radius={'md'}
                label={"Email"}
                style={inputStyle}
                {...form.getInputProps('email')} />
            <TextInput
                withAsterisk
                radius={'md'}
                label={"Name"}
                description={'Custom Name for your Account'}
                style={inputStyle}
                {...form.getInputProps('userName')} />
            <PasswordInput
                withAsterisk
                label={'Password'}
                description={'Min. 8 Chars, Uppercase, Lowercase, Special Chars'}
                radius={'md'}
                style={inputStyle}
                {...form.getInputProps('password')} />
            <PasswordInput
                withAsterisk
                label={'Confirm Password'}
                description={'Must match above Password'}
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
            <Text size={'xl'} style={messageStyle}>{'You registration to IGRU failed unexpectedly.'}</Text>
            <Text style={messageStyle}>{'Feel free to click the link below to go try again'}</Text>
        </Stack>
    )
});

export const SignUpSuccess = React.memo(() => {
    return (
        <Stack>
            <Text size={'xl'} style={messageStyle}>{'Thank you for signing up for IGRU.'}</Text>
            <Text style={messageStyle}>{'Feel free to click the link below to go back to the login page and authenticate with your account.'}</Text>
        </Stack>
    )
});