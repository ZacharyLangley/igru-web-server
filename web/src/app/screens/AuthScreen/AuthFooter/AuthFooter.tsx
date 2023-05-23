import React from 'react';
import {Link} from 'react-router-dom';
import Button from '../../../../common/components/Button/Button';

import './styles.scss';

export interface AuthFooterProps {
  testID?: string;
  title?: string;
  linkTitle?: string;
  linkUrl: string;
  buttonTitle?: string;
  onButtonClick?: () => void;
  disableButton?: boolean;
}

const AuthFooter: React.FC<AuthFooterProps> = ({
  testID = 'auth-footer',
  title,
  linkTitle,
  linkUrl,
  buttonTitle,
  onButtonClick,
  disableButton,
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
        <Button title={buttonTitle} onClick={onButtonClick} disable={disableButton} />
      </div>
    </div>
  );
};

export default AuthFooter;
