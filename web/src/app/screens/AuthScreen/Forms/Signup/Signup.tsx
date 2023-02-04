import React, {useState} from 'react';
import { useDispatch } from 'react-redux';

import AuthForm from '../../AuthForm/AuthForm';
import SignupForm, {defaultSignupFormData} from './SignupForm/SignupForm';
import language from '../../../../../common/language/index';
import AuthDialog from '../../AuthDialog/AuthDialog';
import Branding from '../../../../../common/components/Branding/Branding';
import AuthFooter from '../../AuthFooter/AuthFooter';
import { RoutePath } from '../../../../types/routes';
import { dispatchSignUpAction } from '../../../../../domain/actions/user';

interface SignupProps {
  testID?: string;
}

const headerTitle = language("auth.sign_up.header")

const isValid = (formData: any) => (
  formData.email &&
  formData.email.length > 6 &&
  formData.password &&
  formData.password.length > 8 &&
  formData.password === formData.confirmPassword
)

const Signup: React.FC<SignupProps> = ({
  testID = 'sign-up-form-container',
  ...props
}) => {
  const dispatch = useDispatch();
  const [formData, setFormData] = useState(defaultSignupFormData);

  const updateFormData = (key: string, value: number | string) => {
    setFormData({
      ...formData,
      [key]: value,
    });
  };

  const onClick = () => {
    if (isValid(formData)) {
      dispatch(dispatchSignUpAction({email: formData.email, password: formData.password}));
    }
  }

  return (
    <AuthDialog
      header={<Branding />}
      body={
        <AuthForm
          title={headerTitle}
          form={<SignupForm formData={formData} onChange={updateFormData} />}
        />
      }
      footer={
        <AuthFooter
          title={language("auth.sign_up.title")}
          linkTitle={language("auth.sign_up.link_title")}
          buttonTitle={language("auth.sign_up.button_title")}
          linkUrl={RoutePath.HOME}
          onButtonClick={onClick}
          disableButton={!isValid(formData)}
        />
      }
    />
  )
};

export default Signup;
