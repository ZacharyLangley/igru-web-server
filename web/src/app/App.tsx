import React, { useEffect, useMemo } from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {RouterProvider} from 'react-router-dom';
import {isSessionValidatedSelector} from '../domain/selectors/sessions';

import PublicRouter from './routers/public';
import PrivateRouter from './routers/private';
import { dispatchValidateSessionAction } from 'src/domain/actions/sessions';

interface AppProps {}

const App: React.FC<AppProps> = () => {
  const dispatch = useDispatch();
  // TODO: Need Loading Screen while session auth is confirming
  const isSessionValidated = useSelector(isSessionValidatedSelector);
  const router = useMemo(() => isSessionValidated ? PrivateRouter : PublicRouter, [isSessionValidated]);

  useEffect(() => {
    dispatch(dispatchValidateSessionAction({}));
  }, [dispatch])

  return <RouterProvider router={router} />;
};

export default App;
