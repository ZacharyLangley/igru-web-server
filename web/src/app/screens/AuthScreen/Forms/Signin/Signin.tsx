import React, {useState} from 'react';
import AuthForm from '../../AuthForm/AuthForm';
import SigninForm, {defaultSignInFormData} from './SigninForm/SigninForm';

interface SigninProps {
  testID?: string;
  title?: string;
  form?: JSX.Element | JSX.Element[];
}

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
      title={'SIGN INTO YOUR ACCOUNT'}
      form={<SigninForm formData={formData} onChange={updateFormData} />}
    />
  );
};

export default Signin;
