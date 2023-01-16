import React from 'react';
import {createBrowserRouter} from 'react-router-dom';

import Error from '../screens/ErrorScreen/ErrorScreen';
import AuthScreen from '../screens/AuthScreen/AuthScreen';
import Signin from '../screens/AuthScreen/Forms/Signin/Signin';
import Signup from '../screens/AuthScreen/Forms/Signup/Signup';
import {RoutePath} from '../types/routes';

const PrivateRouter = createBrowserRouter([
  {
    path: RoutePath.HOME,
    element: <AuthScreen />,
    errorElement: <Error />,
    children: [
      {
        path: RoutePath.HOME,
        element: <Signin />,
      },
      {
        path: RoutePath.SIGN_UP,
        element: <Signup />,
      },
    ],
  },
]);

export default PrivateRouter;
