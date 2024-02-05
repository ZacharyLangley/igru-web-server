import React, { useCallback } from 'react';

import { AppShell, Center, Card, Image, Title, Group, Divider, Button, Stack, Text, MantineStyleProp } from '@mantine/core';
import logo from '../../../common/assets/branding/IGRU_White_logo.png';
import { useNavigate } from 'react-router-dom';
import { RoutePath } from 'src/app/types/routes';

import language from '../../../common/language/index';

const lang = language();
const messageStyle: MantineStyleProp = { textAlign: 'center', maxWidth: 400 };

export interface ErrorScreenProps {
  isPublic?: boolean;
}

const ErrorScreen = React.memo((props: ErrorScreenProps) => {
  const navigate = useNavigate();

  const handleButtonClick = useCallback(() => navigate(RoutePath.HOME), []);
  const buttonTitle = props?.isPublic ? lang.error.unauthButtonTitle : lang.error.authButtonTitle;

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
                  <Card shadow="xl" padding="xl" radius="md" withBorder>
                  <ErrorHeader />
                  <Divider my={'md'}/>
                  <Stack>
                    <Text size={'xl'} style={messageStyle}>{lang.error.subtitle}</Text>
                    <Text style={messageStyle}>{lang.error.message}</Text>
                  </Stack>
                  <Divider my={'md'}/>
                  <Button color={'#469d4b'} radius={'md'} fullWidth onClick={handleButtonClick}>{buttonTitle}</Button>
                  </Card>
                </Center>
            </Group>
        </AppShell.Main>
    </AppShell>
  )
});

const nameStyle = {color: '#494850'};
const titleStyle = {color: '#469d4b'};

const ErrorHeader = React.memo(() => {
  return (
      <Stack mih={50} gap="xs" justify="center" align="center">
          <Group>
              <Image src={logo} height={150} width={150} alt={'branding'}/>
          </Group>
          <Title order={1} style={nameStyle}>{lang.branding.name}</Title>
          <Title order={2} style={titleStyle}>{lang.error.title}</Title>
      </Stack>
  )
});
export default ErrorScreen;
