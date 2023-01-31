import React, {useState} from 'react';
import AuthForm from '../../AuthForm/AuthForm';
import SignupForm, {defaultSignupFormData} from './SignupForm/SignupForm';
import language from '../../../../../common/language/index';

interface SignupProps {
  testID?: string;
  title?: string;
  form?: JSX.Element | JSX.Element[];
}

const headerTitle = language("auth.sign_up.header")

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
      title={headerTitle}
      form={<SignupForm formData={formData} onChange={updateFormData} />}
    />
  );
};

export default Signup;
