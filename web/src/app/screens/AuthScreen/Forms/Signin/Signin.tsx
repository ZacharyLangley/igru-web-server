import React, {useState} from 'react';

import AuthForm from '../../AuthForm/AuthForm';
import SigninForm, {defaultSignInFormData} from './SigninForm/SigninForm';
import language from '../../../../../common/language/index';
import AuthDialog from '../../AuthDialog/AuthDialog';
import Branding from '../../../../../common/components/Branding/Branding';
import AuthFooter from '../../AuthFooter/AuthFooter';
import { RoutePath } from '../../../../types/routes';
import useSession from 'src/store/useSession/useSession';
import useUser from 'src/store/useUser/useUser';

interface SigninProps {
  testID?: string;
}

const headerTitle = language("auth.sign_in.header");

const isValid = (formData: any) => (
  formData.email &&
  formData.email.length > 6 &&
  formData.password &&
  formData.password.length > 8
)

const Signin: React.FC<SigninProps> = ({
  testID = 'sign-in-form-container',
  ...props
}) => {
  const {signIn} = useSession();
  const {setUser} = useUser();
  const [formData, setFormData] = useState(defaultSignInFormData);

  const updateFormData = (key: string, value: number | string) => {
    setFormData({
      ...formData,
      [key]: value,
    });
  };

  const handleAuthentication = async () => {
    if (isValid(formData)) {
      const user = await signIn(formData.email, formData.password);
      console.log(user);
      setUser(user);
    }
  }

  return (
    <AuthDialog
      header={<Branding />}
      body={
        <AuthForm
          title={headerTitle}
          form={<SigninForm formData={formData} onChange={updateFormData} />}
        />
      }
      footer={
        <AuthFooter
          title={language("auth.sign_in.title")}
          linkTitle={language("auth.sign_in.link_title")}
          buttonTitle={language("auth.sign_in.button_title")}
          linkUrl={RoutePath.SIGN_UP}
          onButtonClick={handleAuthentication}
          disableButton={!isValid(formData)}
        />
      }
    />
  );
};

export default Signin;
