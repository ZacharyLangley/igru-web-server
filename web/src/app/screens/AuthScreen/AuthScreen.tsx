import React from 'react'
import { AppShell, Center, Group } from '@mantine/core';
import AuthDialog from './AuthDialog';

export const AuthScreen = React.memo(() => {
    return (
        <AppShell padding="md">
            <AppShell.Main h={'100%'} flex={1}>
                <Group
                    mih={0}
                    h={'95vh'}
                    w={'95vw'}
                    flex={1}
                    gap="sm"
                    justify="center"
                    align="center"
                    wrap="wrap">
                    <Center w={900}>
                        <AuthDialog />
                    </Center>
                </Group>
            </AppShell.Main>
        </AppShell>
    )
})

export default AuthScreen;