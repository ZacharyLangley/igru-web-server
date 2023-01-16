import React from 'react';
import {useSelector} from 'react-redux';
import {RouterProvider} from 'react-router-dom';
import {userIDSelector} from '../domain/selectors/user';

import PublicRouter from './routers/public';
import PrivateRouter from './routers/private';

interface AppProps {}

const App: React.FC<AppProps> = () => {
  const userID = useSelector(userIDSelector);

  const router = userID ? PrivateRouter : PublicRouter;
  return <RouterProvider router={router} />;
};

export default App;
