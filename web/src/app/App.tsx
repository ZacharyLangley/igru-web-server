import React, { useEffect, useMemo } from 'react';
import {RouterProvider} from 'react-router-dom';

import PublicRouter from './routers/public';
import PrivateRouter from './routers/private';
import useSession from 'src/store/useSession/useSession';

interface AppProps {}

const App: React.FC<AppProps> = () => {
  const {sessionValidated, validateSession} = useSession()
  const router = useMemo(() => sessionValidated ? PrivateRouter : PublicRouter, [sessionValidated]);

  useEffect(() => {
    validateSession();
  }, [validateSession])

  return <RouterProvider router={router} />;
};

export default App;
