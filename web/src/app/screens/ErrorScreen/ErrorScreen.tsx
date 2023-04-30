import React from 'react';
import Screen from '../../../common/components/Screen/Screen';
import Dialog from '../../../common/components/Dialog/Dialog';
import Branding from '../../../common/components/Branding/Branding';
import PrimaryButton from '../../../common/components/Button/PrimaryButton/PrimaryButton';
import language from '../../../common/language/index';

import './styles.scss';
import useSession from 'src/store/useSession/useSession';
import { useNavigate } from 'react-router-dom';
import { RoutePath } from 'src/app/types/routes';

export interface ErrorScreenProps {
  isPublic?: boolean;
}

export interface ErrorContentProps {
  title: string;
  message: string;
}

export interface ErrorFooterProps {
  isPublic?: boolean;
}

const ErrorContent: React.FC<ErrorContentProps> = (props) => (
  <div className={'error-text-container'}>
    <div className={'error-title'}>{props.title}</div>
    <div className={'error-message'}>{props.message}</div>
  </div>
)
const errorFooterNavigateUnauthLabel = language("error.label.footer.navigate.unauth")
const errorFooterNavigateAuthLabel = language("error.label.footer.navigate.auth")
const errorFooterSignOutLabel = language("error.label.footer.signout")

const ErrorFooter: React.FC<ErrorFooterProps> = (props) => {
  const navigate = useNavigate();
  const {signOut} = useSession();

  const navigateToHome = () => navigate(RoutePath.HOME)

  return (
    <div className={'error-footer-container'}>
        <div className='error-footer-button'>
          <PrimaryButton title={props.isPublic ? errorFooterNavigateUnauthLabel : errorFooterNavigateAuthLabel} onClick={navigateToHome}/>
        </div>
        {
          !props.isPublic &&
          <div className='error-footer-button'><PrimaryButton title={errorFooterSignOutLabel} onClick={signOut}/></div>
        }
    </div>
  )
}

const errorTitleLabel = language("error.label.title");
const errorMessageLabel= language("error.label.message");

const ErrorScreen: React.FC<ErrorScreenProps> = (props) => {
  return (
    <Screen>
      <div id={'error-screen'} className={'error-screen-container'}>
        <div className={'error-screen-content'}>
          <Dialog
            header={<Branding />}
            body={<ErrorContent title={errorTitleLabel} message={errorMessageLabel}/>}
            footer={<ErrorFooter isPublic={props.isPublic}/>}
          />
        </div>
      </div>
    </Screen>
  )
};

export default ErrorScreen;
