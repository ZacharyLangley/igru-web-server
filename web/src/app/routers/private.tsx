import React from 'react';
import {createBrowserRouter} from 'react-router-dom';

import MainScreen from '../screens/MainScreen/MainScreen';
import Error from '../screens/ErrorScreen/ErrorScreen';
import Home from '../screens/MainScreen/Home/Home';
import Tasks from '../screens/MainScreen/Tasks/Tasks';
import Gardens from '../screens/MainScreen/Gardens/Gardens';
import Plants from '../screens/MainScreen/Plants/Plants';
import Recipes from '../screens/MainScreen/Recipes/Recipes';
import Settings from '../screens/MainScreen/Settings/Settings';
import Strains from '../screens/MainScreen/Strains/Strains';
import {RoutePath} from '../types/routes';

const PrivateRouter = createBrowserRouter([
  {
    path: RoutePath.HOME,
    element: <MainScreen />,
    errorElement: <Error />,
    children: [
      {
        path: RoutePath.HOME,
        element: <Home />,
      },
      {
        path: RoutePath.TASKS,
        element: <Tasks />,
      },
      {
        path: RoutePath.GARDENS,
        element: <Gardens />,
      },
      {
        path: RoutePath.PLANTS,
        element: <Plants />,
      },
      {
        path: RoutePath.RECIPES,
        element: <Recipes />,
      },
      {
        path: RoutePath.STRAINS,
        element: <Strains />,
      },
      {
        path: RoutePath.SETTINGS,
        element: <Settings />,
      },
    ],
  },
]);

export default PrivateRouter;
