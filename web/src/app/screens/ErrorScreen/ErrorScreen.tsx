import React, { useCallback } from 'react';
import language from '../../../common/language/index';

import './styles.scss';

import { AppShell, Center, Card, Image, Title, Group, Divider, Button, Stack, Text, MantineStyleProp } from '@mantine/core';
import logo from '../../../common/assets/branding/IGRU_White_logo.png';
import { useNavigate } from 'react-router-dom';
import { RoutePath } from 'src/app/types/routes';

const errorFooterNavigateUnauthLabel = language("error.label.footer.navigate.unauth")

const errorTitleLabel = language("error.label.title");
const errorMessageLabel= language("error.label.message");
const messageStyle: MantineStyleProp = { textAlign: 'center', maxWidth: 400 };

export interface ErrorScreenProps {
  isPublic?: boolean;
}

const ErrorScreen = React.memo((props: ErrorScreenProps) => {
  const navigate = useNavigate();

  const handleButtonClick = useCallback(() => navigate(RoutePath.HOME), []);

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
                    <Text size={'xl'} style={messageStyle}>{errorTitleLabel}</Text>
                    <Text style={messageStyle}>{errorMessageLabel}</Text>
                  </Stack>
                  <Divider my={'md'}/>
                  <Button color={'#469d4b'} radius={'md'} fullWidth onClick={handleButtonClick}>{'Go Home'}</Button>
                  </Card>
                </Center>
            </Group>
        </AppShell.Main>
    </AppShell>
  )
});

const ErrorHeader = React.memo(() => {
  return (
      <Stack mih={50} gap="xs" justify="center" align="center">
          <Group>
              <Image src={logo} height={150} width={150} alt={'branding'}/>
          </Group>
          <Title order={1} style={{color: '#494850'}}>{'IGRU'}</Title>
          <Title order={2} style={{color: '#469d4b'}}>{'Woah there, lil\' buddy!'}</Title>
      </Stack>
  )
});
export default ErrorScreen;
