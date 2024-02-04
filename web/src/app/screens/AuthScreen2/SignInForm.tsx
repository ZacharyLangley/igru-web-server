import React from 'react';
import { PasswordInput, Stack, TextInput } from '@mantine/core';

import { useAuthFormContext } from '../../../common/contexts/authenticationContext';

const inputStyle = { width: '100%' }

export const SignInForm = React.memo(() => {

    const form = useAuthFormContext();

    return (
        <Stack gap="md" justify="space-between" align="flex-start" flex={1}>
            <TextInput
                withAsterisk
                radius={'md'}
                label={"Email"}
                style={inputStyle}
                {...form.getInputProps('email')}
            />
            <PasswordInput
                withAsterisk
                label={'Password'}
                radius={'md'}
                style={inputStyle}
                {...form.getInputProps('password')}
            />
        </Stack>
    )
});

export default SignInForm;