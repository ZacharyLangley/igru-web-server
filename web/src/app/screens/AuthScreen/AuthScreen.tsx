import React from 'react';
import {Location, Outlet, useLocation, useNavigate} from 'react-router-dom';
import Logo from '../../../common/components/Logo/Logo';
import {RoutePath} from '../../types/routes';

import AuthDialog from './AuthDialog/AuthDialog';
import AuthFooter, {AuthFooterProps} from './AuthFooter/AuthFooter';
import './styles.scss';

interface AuthScreenProps {
  testID?: string;
}

const signInFooter: AuthFooterProps = {
  title: 'Need an account?',
  linkUrl: RoutePath.SIGN_UP,
  linkTitle: 'Click Here',
  buttonTitle: 'LOGIN',
};

const signUpFooter: AuthFooterProps = {
  title: 'Already have an account?',
  linkUrl: RoutePath.HOME,
  linkTitle: 'Click Here',
  buttonTitle: 'SIGN UP',
};

const determineFooter = (location: Location) => {
  const {pathname} = location;
  if (pathname === RoutePath.HOME) return signInFooter;
  if (pathname === RoutePath.SIGN_UP) return signUpFooter;
  return signInFooter;
};

const AuthScreen: React.FC<AuthScreenProps> = ({testID = 'auth-screen'}) => {
  const navigate = useNavigate();
  const location = useLocation();
  const branding = (
    <div className={'auth-screen-branding'}>
      <Logo height={100} width={100} hideName />
      <div className={'auth-screen-title'}>{'IGRU'}</div>
    </div>
  );

  const onClick = () => {
    const {pathname} = location;
    if (pathname === RoutePath.HOME) {
      // TODO: ZL | Dispatch Signin Action to Execute Signin Saga
      navigate(RoutePath.HOME);
    }
    if (pathname === RoutePath.SIGN_UP) {
      // TODO: ZL | Dispatch Signup Action to Execute Signup Saga
      navigate(RoutePath.HOME);
    }
  };

  const footerConfig = determineFooter(location);
  return (
    <div id={'auth-screen'} className={'auth-screen-container'}>
      <div className={'auth-screen-content'}>
        <AuthDialog
          header={branding}
          body={<Outlet />}
          footer={<AuthFooter {...footerConfig} onButtonClick={onClick} />}
        />
      </div>
    </div>
  );
};

export default AuthScreen;
