import React from 'react';

import './styles.scss';

interface AuthDialogProps {
  testID?: string;
  header?: JSX.Element | JSX.Element[] | string;
  body?: JSX.Element | JSX.Element[] | string;
  footer?: JSX.Element | JSX.Element[] | string;
}

const AuthDialog: React.FC<AuthDialogProps> = ({
  testID = 'auth-dialog',
  header,
  body,
  footer,
}) => {
  return (
    <div id={testID} className={'auth-dialog-container'}>
      {header && (
        <div id={`${testID}:header`} className={'auth-dialog-header'}>
          {header}
        </div>
      )}
      {body && (
        <div id={`${testID}:body`} className={'auth-dialog-body'}>
          {body}
        </div>
      )}
      {footer && (
        <div id={`${testID}:footer`} className={'auth-dialog-footer'}>
          {footer}
        </div>
      )}
    </div>
  );
};

export default AuthDialog;
