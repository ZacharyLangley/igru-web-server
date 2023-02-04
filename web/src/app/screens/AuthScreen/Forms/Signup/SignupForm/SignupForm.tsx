import React, {ChangeEvent} from 'react';
import {Form, FormGroup, Label, Input} from 'reactstrap';
import language from '../../../../../../common/language';

export interface SignInFormData {
  email?: string;
  password?: string;
  confirmPassword?: string;
}

interface SignupFormProps {
  formData: SignInFormData;
  onChange: (name: string, value: string) => void;
}

export const defaultSignupFormData: SignInFormData = {
  email: '',
  password: '',
  confirmPassword: '',
};

const emailLabel = language("input.label.email");
const passwordLabel = language("input.label.password");
const confirmPasswordLabel = language("input.label.confirm_password");

const SignupForm: React.FC<SignupFormProps> = ({formData, onChange}) => {
  const onFieldChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (e === undefined) return;
    e.preventDefault();
    onChange(e.target.name, e.target.value);
  };

  return (
    <Form>
      <FormGroup floating>
        <Input
          id={'email'}
          name={'email'}
          type={'email'}
          value={formData.email}
          placeholder={''}
          onChange={onFieldChange}
          valid={formData.email && formData.email.length > 0 ? formData.email.length > 8 : undefined}
          invalid={formData.email && formData.email.length > 0 ? formData.email.length < 8 : undefined}
        />
        <Label for={'email'}>{emailLabel}</Label>
      </FormGroup>
      <FormGroup floating>
        <Input
          id={'password'}
          name={'password'}
          type={'password'}
          value={formData.password}
          placeholder={''}
          onChange={onFieldChange}
          valid={formData.password && formData.password.length > 0 && formData.confirmPassword && formData.confirmPassword.length > 0 ? formData.password === formData.confirmPassword : undefined}
          invalid={formData.password  && formData.password.length > 0 && formData.confirmPassword && formData.confirmPassword.length > 0 ? formData.password !== formData.confirmPassword : undefined}
        />
        <Label for={'password'}>{passwordLabel}</Label>
      </FormGroup>
      <FormGroup floating>
        <Input
          id={'confirmPassword'}
          name={'confirmPassword'}
          type={'password'}
          value={formData.confirmPassword}
          placeholder={''}
          onChange={onFieldChange}
          valid={formData.password && formData.password.length > 0 && formData.confirmPassword && formData.confirmPassword.length > 0 ? formData.password === formData.confirmPassword : undefined}
          invalid={formData.password  && formData.password.length > 0 && formData.confirmPassword && formData.confirmPassword.length > 0 ? formData.password !== formData.confirmPassword : undefined}
        />
        <Label for={'confirmPassword'}>{confirmPasswordLabel}</Label>
      </FormGroup>
    </Form>
  );
};

export default SignupForm;
