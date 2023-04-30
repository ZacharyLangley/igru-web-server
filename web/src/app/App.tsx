import React, { useEffect, useMemo } from 'react';
import {RouterProvider} from 'react-router-dom';

import PublicRouter from './routers/public';
import PrivateRouter from './routers/private';
import useSession from 'src/store/useSession/useSession';
import useUser from 'src/store/useUser/useUser';

interface AppProps {}

const App: React.FC<AppProps> = () => {
  const {sessionValidated, validateSession} = useSession();
  const {setUser} = useUser();
  const router = useMemo(() => sessionValidated ? PrivateRouter : PublicRouter, [sessionValidated]);

  useEffect(() => {
    const validateUser = async () => {
      const user = await validateSession();
      setUser(user);
    }

    validateUser();
  }, [validateSession, setUser])

  return <RouterProvider router={router} />;
};

export default App;
