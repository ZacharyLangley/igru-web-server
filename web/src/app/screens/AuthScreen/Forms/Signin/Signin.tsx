import React, {useState} from 'react';
import AuthForm from '../../AuthForm/AuthForm';
import SigninForm, {defaultSignInFormData} from './SigninForm/SigninForm';
import language from '../../../../../common/language/index';

interface SigninProps {
  testID?: string;
  title?: string;
  form?: JSX.Element | JSX.Element[];
}

const headerTitle = language("auth.sign_in.header");

const Signin: React.FC<SigninProps> = ({
  testID = 'sign-in-form-container',
  title,
  form,
}) => {
  const [formData, setFormData] = useState(defaultSignInFormData);

  const updateFormData = (key: string, value: number | string) => {
    setFormData({
      ...formData,
      [key]: value,
    });
  };

  return (
    <AuthForm
      title={headerTitle}
      form={<SigninForm formData={formData} onChange={updateFormData} />}
    />
  );
};

export default Signin;
