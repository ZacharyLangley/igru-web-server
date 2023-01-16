import React, {ChangeEvent} from 'react';
import {Form, FormGroup, Label, Input} from 'reactstrap';

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
  email: undefined,
  password: undefined,
  confirmPassword: undefined,
};

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
          name={'name'}
          type={'email'}
          value={formData.email}
          placeholder={''}
          onChange={onFieldChange}
        />
        <Label for={'email'}>{'Email'}</Label>
      </FormGroup>
      <FormGroup floating>
        <Input
          id={'password'}
          name={'password'}
          type={'password'}
          value={formData.password}
          placeholder={''}
          onChange={onFieldChange}
        />
        <Label for={'password'}>{'Password'}</Label>
      </FormGroup>
      <FormGroup floating>
        <Input
          id={'confirm-password'}
          name={'confirm-password'}
          type={'password'}
          value={formData.password}
          placeholder={''}
          onChange={onFieldChange}
        />
        <Label for={'confirm-password'}>{'Confirm Password'}</Label>
      </FormGroup>
    </Form>
  );
};

export default SignupForm;
