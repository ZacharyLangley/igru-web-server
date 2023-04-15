import React, {useCallback, useMemo, useState} from 'react';

import AuthForm from '../../AuthForm/AuthForm';
import SignupForm, {defaultSignupFormData} from './SignupForm/SignupForm';
import language from '../../../../../common/language/index';
import AuthDialog from '../../AuthDialog/AuthDialog';
import Branding from '../../../../../common/components/Branding/Branding';
import AuthFooter from '../../AuthFooter/AuthFooter';
import { RoutePath } from '../../../../types/routes';
import { SignupSuccess } from './SignupForm/SignupSuccess';
import { SignupFailure } from './SignupForm/SignupFailure';
import useUser, {Status} from 'src/store/useUser/useUser';

interface SignupProps {
  testID?: string;
}

const text = {
  headerTitle: language("auth.sign_up.header"),
  headerTitleSuccess: language("auth.sign_up.header.success"),
  headerTitleFailure: language("auth.sign_up.header.failure"),
  authFooterTitle: language("auth.sign_up.title"),
  authFooterLinkTitle: language("auth.sign_up.link_title"),
  authFooterButtonTitle: language("auth.sign_up.button_title")
}

const isValid = (formData: any) => (
  formData.email &&
  formData.email.length > 6 &&
  formData.password &&
  formData.password.length > 8 &&
  formData.password === formData.confirmPassword
)

const Signup: React.FC<SignupProps> = () => {
  const {signUpStatus, signUp} = useUser();
  const [formData, setFormData] = useState(defaultSignupFormData);
  
  const updateFormData = useCallback((key: string, value: number | string) => {
    setFormData({
      ...formData,
      [key]: value,
    });
  }, [formData, setFormData]);

  const onClick = () => {
    if (isValid(formData)) {
      signUp(formData.email, formData.password);
    }
  };

  const formTitle = useMemo(() => {
    switch(signUpStatus) {
      case Status.FAILURE: return text.headerTitleFailure;
      case Status.SUCCESS: return text.headerTitleSuccess;
      case Status.IDLE:
      default: return text.headerTitle;
    }
  }, [signUpStatus]);

  const formBody = useMemo(() => {
    switch(signUpStatus) {
      case Status.FAILURE: return <SignupFailure />
      case Status.SUCCESS: return <SignupSuccess />
      case Status.IDLE:
      default: return <SignupForm formData={formData} onChange={updateFormData} />;
    }
  }, [signUpStatus, formData, updateFormData]);

  const formFooter = (signUpStatus === Status.IDLE)
    ? (
      <AuthFooter
        title={text.authFooterTitle}
        linkTitle={text.authFooterLinkTitle}
        buttonTitle={text.authFooterButtonTitle}
        linkUrl={RoutePath.HOME}
        onButtonClick={onClick}
        disableButton={!isValid(formData)}
      />
    ) : undefined;

  return (
    <AuthDialog
      header={<Branding />}
      body={<AuthForm title={formTitle} form={formBody} />}
      footer={formFooter}
    />
  );
};

export default Signup;
