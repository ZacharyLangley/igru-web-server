import React from 'react';
import {Link} from 'react-router-dom';
import PrimaryButton from '../../../../common/components/Button/PrimaryButton/PrimaryButton';

import './styles.scss';

export interface AuthFooterProps {
  testID?: string;
  title?: string;
  linkTitle?: string;
  linkUrl: string;
  buttonTitle?: string;
  onButtonClick?: () => void;
}

const AuthFooter: React.FC<AuthFooterProps> = ({
  testID = 'auth-footer',
  title,
  linkTitle,
  linkUrl,
  buttonTitle,
  onButtonClick,
}) => {
  return (
    <div className={'footer-container'}>
      <div className={'footer-link-container'}>
        <span className={'footer-title'}>{title}</span>
        <Link to={linkUrl} className={'link'}>
          {linkTitle}
        </Link>
      </div>
      <div className={'footer-button-container'}>
        <PrimaryButton title={buttonTitle} onClick={onButtonClick} />
      </div>
    </div>
  );
};

export default AuthFooter;
