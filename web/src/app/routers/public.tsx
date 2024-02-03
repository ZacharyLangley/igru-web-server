import React from 'react';
import {createBrowserRouter} from 'react-router-dom';

import Error from '../screens/ErrorScreen/ErrorScreen';
import Signup from '../screens/AuthScreen/Forms/Signup/Signup';
import {RoutePath} from '../types/routes';

import AuthScreen2 from '../screens/AuthScreen2/AuthScreen2';
import SignInForm from '../screens/AuthScreen2/SignInForm';
import { SignUpForm } from '../screens/AuthScreen2/SignUpForm';

const PrivateRouter = createBrowserRouter([
  {
    path: RoutePath.HOME,
    element: <AuthScreen2 />,
    errorElement: <Error isPublic={true}/>,
    children: [
      {
        path: RoutePath.HOME,
        element: <SignInForm />,

      },
      {
        path: RoutePath.SIGN_UP,
        element: <SignUpForm />,
      },
    ],
  },
]);

export default PrivateRouter;
