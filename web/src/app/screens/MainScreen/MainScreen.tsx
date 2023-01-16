import React, {useState} from 'react';
import {Outlet, useNavigate} from 'react-router-dom';

import Layout from '../../../common/components/Screen/Screen';
import Header from '../../../common/components/Screen/Header/Header';
import Body from '../../../common/components/Screen/Body/Body';
import Sidebar, {
  SidebarOptions,
} from '../../../common/components/Screen/Sidebar/Sidebar';

import allIcon from '../../../common/assets/icons/nav/all_icon.png';
import tasksIcon from '../../../common/assets/icons/nav/tasks_icon.png';
import gardensIcon from '../../../common/assets/icons/nav/gardens_icon.png';
import plantsIcon from '../../../common/assets/icons/nav/plants_icon.png';
import strainsIcon from '../../../common/assets/icons/nav/strains_icon.png';
import recipesIcon from '../../../common/assets/icons/nav/recipes_icon.png';
import settingsIcon from '../../../common/assets/icons/nav/settings_icon.png';

import {RouteLabel, RoutePath} from '../../types/routes';

interface MainScreenProps {}

const sidebarOptions: SidebarOptions[] = [
  {
    label: RouteLabel.HOME,
    path: RoutePath.HOME,
    icon: allIcon,
  },
  {
    label: RouteLabel.TASKS,
    path: RoutePath.TASKS,
    icon: tasksIcon,
  },
  {
    label: RouteLabel.GARDENS,
    path: RoutePath.GARDENS,
    icon: gardensIcon,
  },
  {
    label: RouteLabel.PLANTS,
    path: RoutePath.PLANTS,
    icon: plantsIcon,
  },
  {
    label: RouteLabel.STRAINS,
    path: RoutePath.STRAINS,
    icon: strainsIcon,
  },
  {
    label: RouteLabel.RECIPES,
    path: RoutePath.RECIPES,
    icon: recipesIcon,
  },
  {
    label: RouteLabel.SETTINGS,
    path: RoutePath.SETTINGS,
    icon: settingsIcon,
  },
];

const MainScreen: React.FC<MainScreenProps> = () => {
  const [openSidebar, setOpenSidebar] = useState(false);
  const navigate = useNavigate();
  const toggleSidebar = () => {
    setOpenSidebar(!openSidebar);
  };

  const onOptionSelection = (path?: string) => {
    if (path) navigate(path);
    toggleSidebar();
  };

  return (
    <Layout>
      <Header onClick={toggleSidebar} />
      <Body>
        <Outlet />
      </Body>
      <Sidebar
        isOpen={openSidebar}
        onClose={toggleSidebar}
        onOptionSelection={onOptionSelection}
        options={sidebarOptions}
      />
    </Layout>
  );
};

export default MainScreen;
