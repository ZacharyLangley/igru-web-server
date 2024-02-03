import React from 'react'
import { Outlet } from 'react-router-dom';
import { Card, Image, Title, Group, Divider, Button, Anchor, Stack } from '@mantine/core';

import logo from '../../../common/assets/branding/IGRU_White_logo.png';
import {AuthFormProvider, useAuthForm} from '../../../common/contexts/authenticationContext';

export const AuthDialog = React.memo(() => {
    const form = useAuthForm({
        initialValues: {
          email: '',
          userName: '',
          password: '',
          confirmPassword: ''
        },
    })

    return (
        <Card shadow="xl" padding="xl" radius="md" withBorder>
            <AuthDialogHeader />
            <Divider my={'md'}/>
            <AuthFormProvider form={form}>
                <Outlet />
            </AuthFormProvider>
            <Divider my={'md'}/>
            <AuthDialogFooter/>
        </Card>
    )
});

const AuthDialogHeader = React.memo(() => {
    return (
        <Stack mih={50} gap="xs" justify="center" align="center">
            <Group>
                <Image src={logo} height={150} width={150} alt={'branding'}/>
            </Group>
            <Title order={1} style={{color: '#494850'}}>{'IGRU'}</Title>
            <Title order={2} style={{color: '#469d4b'}}>{'Sign into your account'}</Title>
        </Stack>
    )
});

const AuthDialogFooter = React.memo(() => {
    return (
        <Group justify="space-between">
            <Anchor size={'sm'} underline={'never'} onClick={() => console.log('CLICK')}>{'Need an Account?'}</Anchor>
            <Button radius={'md'}>{'Login'}</Button>
        </Group>
    )
})

export default AuthDialog