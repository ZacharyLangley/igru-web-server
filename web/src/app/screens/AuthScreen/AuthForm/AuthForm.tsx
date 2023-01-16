import React from 'react';

import './styles.scss';

interface AuthFormProps {
  testID?: string;
  title?: string;
  form?: JSX.Element | JSX.Element[];
}

const AuthForm: React.FC<AuthFormProps> = ({
  testID = 'auth-form',
  title,
  form,
}) => {
  return (
    <div className={'auth-form-container'}>
      <div className={'auth-form-title'}>{title}</div>
      <div className={'auth-form-container'}>{form}</div>
    </div>
  );
};

export default AuthForm;
