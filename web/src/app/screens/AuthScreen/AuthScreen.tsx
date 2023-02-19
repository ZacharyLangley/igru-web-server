import React from 'react';
import {Outlet} from 'react-router-dom';
import './styles.scss';

interface AuthScreenProps {
  testID?: string;
}

const AuthScreen: React.FC<AuthScreenProps> = ({testID = 'auth-screen'}) => {
  return (
    <div id={'auth-screen'} className={'auth-screen-container'}>
      <div className={'auth-screen-content'}>
        <Outlet />
      </div>
    </div>
  );
};

export default AuthScreen;
