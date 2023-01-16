import React, {useState} from 'react';
import AuthForm from '../../AuthForm/AuthForm';
import SignupForm, {defaultSignupFormData} from './SignupForm/SignupForm';

interface SignupProps {
  testID?: string;
  title?: string;
  form?: JSX.Element | JSX.Element[];
}

const Signup: React.FC<SignupProps> = ({
  testID = 'sign-up-form-container',
  title,
  form,
}) => {
  const [formData, setFormData] = useState(defaultSignupFormData);

  const updateFormData = (key: string, value: number | string) => {
    setFormData({
      ...formData,
      [key]: value,
    });
  };

  return (
    <AuthForm
      title={'CREATE YOUR ACCOUNT'}
      form={<SignupForm formData={formData} onChange={updateFormData} />}
    />
  );
};

export default Signup;
