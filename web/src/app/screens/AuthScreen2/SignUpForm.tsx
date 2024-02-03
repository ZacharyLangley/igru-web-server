import React from 'react'

import { PasswordInput, Stack, TextInput } from '@mantine/core';

const inputStyle = { width: '100%' }

export const SignUpForm = React.memo(() => {
    return (
        <Stack gap="md" justify="space-between" align="flex-start" flex={1}>
            <TextInput
                withAsterisk
                radius={'md'}
                label={"Email"}
                style={inputStyle}
            />
            <TextInput
                withAsterisk
                radius={'md'}
                label={"Name"}
                description={'Custom Name for the Account'}
                style={inputStyle}
            />
            <PasswordInput
                withAsterisk
                label={'Password'}
                description={'Min. 8 Chars, Uppercase, Lowercase, Special Chars'}
                radius={'md'}
                style={inputStyle}
            />
            <PasswordInput
                withAsterisk
                label={'Confirm Password'}
                description={'Must match above Password'}
                radius={'md'}
                style={inputStyle}
            />
        </Stack>
    )
})