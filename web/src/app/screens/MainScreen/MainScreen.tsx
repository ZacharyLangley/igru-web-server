import React, {useCallback} from 'react';
import {Outlet, useNavigate} from 'react-router-dom';

import { AppShell, Burger, Group, Image, Stack, Title, UnstyledButton, Avatar } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import logo from '../../../common/assets/branding/IGRU_White_logo_mini.png';

import allIcon from '../../../common/assets/icons/nav/all_icon.png';
import tasksIcon from '../../../common/assets/icons/nav/tasks_icon.png';
import gardensIcon from '../../../common/assets/icons/nav/gardens_icon.png';
import plantsIcon from '../../../common/assets/icons/nav/plants_icon.png';
import strainsIcon from '../../../common/assets/icons/nav/strains_icon.png';
import recipesIcon from '../../../common/assets/icons/nav/recipes_icon.png';
import settingsIcon from '../../../common/assets/icons/nav/settings_icon.png';

import {RoutePath} from '../../types/routes';
import language from '../../../common/language/index';
import GroupSelect from './components/GroupSelect/GroupSelect';

interface MainScreenProps {}

export interface SidebarOptions {
  label: string;
  path: string;
  icon: string;
}

const lang = language();
const sidebarOptions: SidebarOptions[] = [
  {
    label: lang.sidebar.options.home,
    path: RoutePath.HOME,
    icon: allIcon,
  },
  {
    label: lang.sidebar.options.tasks,
    path: RoutePath.TASKS,
    icon: tasksIcon,
  },
  {
    label: lang.sidebar.options.gardens,
    path: RoutePath.GARDENS,
    icon: gardensIcon,
  },
  {
    label: lang.sidebar.options.plants,
    path: RoutePath.PLANTS,
    icon: plantsIcon,
  },
  {
    label: lang.sidebar.options.strains,
    path: RoutePath.STRAINS,
    icon: strainsIcon,
  },
  {
    label: lang.sidebar.options.recipes,
    path: RoutePath.RECIPES,
    icon: recipesIcon,
  },
  {
    label: lang.sidebar.options.settings,
    path: RoutePath.SETTINGS,
    icon: settingsIcon,
  },
];

const headerConfig = { height: 60 };

const MainScreen: React.FC<MainScreenProps> = () => {
  const [mobileOpened, { toggle: toggleMobile }] = useDisclosure(false);
  const [desktopOpened, { toggle: toggleDesktop }] = useDisclosure(false);

  return (
    <AppShell
      header={headerConfig}
      navbar={{
        width: 200,
        breakpoint: 'sm',
        collapsed: { mobile: !mobileOpened, desktop: !desktopOpened },
      }}
      padding="md"
    >
      <AppShell.Header>
        <Group h="100%" px="md">
          <Group flex={1} justify={'flex-start'}>
            <Burger opened={mobileOpened} onClick={toggleMobile} hiddenFrom="sm" size="sm" />
            <Burger opened={desktopOpened} onClick={toggleDesktop} visibleFrom="sm" size="sm" />
            <HeaderBranding />
          </Group>
          <Group justify='flex-end' flex={2}>
            <GroupSelect />
            <Avatar radius="xl" />
          </Group>
        </Group>
      </AppShell.Header>
      <AppShell.Navbar p="md">
        {sidebarOptions.map((item, index) => (
          <NavbarItem key={index} index={index} {...item}/>
        ))}
      </AppShell.Navbar>
      <AppShell.Main>
          <Outlet />
      </AppShell.Main>
    </AppShell>
  );
};

interface NavBarItemProps {
  key: number;
  index: number;
  label: string;
  path: string;
  icon: string; 
}

const NavbarItem = React.memo((props: NavBarItemProps) => {
  const navigate = useNavigate();
  const handleNavigate = useCallback(() => {
    navigate(props.path);
  }, [navigate, props?.path])
  return (
    <UnstyledButton onClick={handleNavigate}>
      <Stack gap={'xl'} justify='center' style={{ paddingTop: 8, paddingBottom: 8}}>
        <Group gap={'md'} align='center' style={{ paddingBottom: 8 }}>
          <Image src={props.icon} height={35} width={100} alt={'branding'} />
          <Title order={5} style={titleStyle}>{props.label}</Title>
        </Group>
      </Stack>
    </UnstyledButton>
  )
})

const titleStyle = {color: '#494850'};

const HeaderBranding = React.memo(() => {
  const navigate = useNavigate();
  const handleNavigate = useCallback(() => {
    navigate(RoutePath.HOME);
  }, [navigate])

  return (
    <UnstyledButton onClick={handleNavigate}>
      <Group gap={'sm'}>
        <Image src={logo} height={35} width={100} alt={'branding'}/>
        <Title order={2} style={titleStyle}>{lang.branding.name}</Title>
      </Group>
    </UnstyledButton>
  )
});

export default MainScreen;
