import React from 'react';

import { PasswordInput, Stack, TextInput } from '@mantine/core';

const inputStyle = { width: '100%' }

export const SignInForm = React.memo(() => {
    return (
        <Stack gap="md" justify="space-between" align="flex-start" flex={1}>
            <TextInput
                withAsterisk
                radius={'md'}
                label={"Email"}
                style={inputStyle}
            />
            <PasswordInput
                withAsterisk
                label={'Password'}
                radius={'md'}
                style={inputStyle}
            />
        </Stack>
    )
});

export default SignInForm;