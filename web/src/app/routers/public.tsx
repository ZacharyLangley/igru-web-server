import React from 'react';
import {createBrowserRouter} from 'react-router-dom';

import Error from '../screens/ErrorScreen/ErrorScreen';
import {RoutePath} from '../types/routes';

import AuthScreen from '../screens/AuthScreen/AuthScreen';
import SignInForm from '../screens/AuthScreen/SignInForm';
import { SignUpForm, SignUpSuccess, SignUpFailure } from '../screens/AuthScreen/SignUpForm';

const PrivateRouter = createBrowserRouter([
  {
    path: RoutePath.HOME,
    element: <AuthScreen />,
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
      {
        path: RoutePath.SIGN_UP_SUCCESS,
        element: <SignUpSuccess />,
      },
      {
        path: RoutePath.SIGN_UP_FAILURE,
        element: <SignUpFailure />,
      },
    ],
  },
]);

export default PrivateRouter;
