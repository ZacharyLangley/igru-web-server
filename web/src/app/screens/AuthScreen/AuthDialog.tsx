import React, { useCallback, useMemo } from 'react'
import { Outlet, useNavigate, useLocation } from 'react-router-dom';
import { Card, Image, Title, Group, Divider, Button, Anchor, Stack } from '@mantine/core';

import logo from '../../../common/assets/branding/IGRU_White_logo.png';
import {AuthFormProvider, useAuthForm, AuthFormValues, useAuthFormContext} from '../../../common/contexts/authenticationContext';
import useSession from 'src/store/useSession/useSession';
import useUser, { Status } from 'src/store/useUser/useUser';
import { RoutePath } from '../../types/routes';
import language from '../../../common/language/index';

const isLoginValid = (formData: AuthFormValues) => (
    (formData?.email.length >= 6 ) &&
    (formData?.password?.length >= 8)
)

const isSignUpValid = (formData: AuthFormValues) => (
    formData?.email?.length >= 6 &&
    formData?.userName?.length > 3 &&
    formData?.password.length >= 8 &&
    formData?.password === formData.confirmPassword &&
    formData?.userName.length >= 4
)

export const AuthDialog = React.memo(() => {
    const location = useLocation();
    const navigate = useNavigate();

    const {signUp} = useUser();
    const {signIn} = useSession();
    const {setUser} = useUser();

    const form = useAuthForm({
        initialValues: {
          email: '',
          userName: '',
          password: '',
          confirmPassword: ''
        },
    })

    const handleButtonClick = useCallback(async (values: AuthFormValues) => {
        if (location.pathname === RoutePath.HOME && isLoginValid(values)) {
            const user = await signIn(values.email, values.password);
            setUser(user);
            return;
        }
        if (location.pathname === RoutePath.SIGN_UP) {
            if (isSignUpValid(values)) {
                const signUpStatus = await signUp(values.email, values.password);
                if (signUpStatus === Status.SUCCESS) navigate(RoutePath.SIGN_UP_SUCCESS);
                else if (signUpStatus === Status.FAILURE) navigate(RoutePath.SIGN_UP_FAILURE);
                return;
            }
        }
        else if (location?.pathname === RoutePath.SIGN_UP_SUCCESS) navigate(RoutePath.HOME);
        else if (location?.pathname === RoutePath.SIGN_UP_FAILURE) navigate(RoutePath.SIGN_UP);
        else return;
    }, [location?.pathname, signIn, signUp, setUser, navigate])

    return (
        <Card shadow="xl" padding="xl" radius="md" withBorder>
            <AuthDialogHeader />
            <Divider my={'md'}/>
            <AuthFormProvider form={form}>
                <form onSubmit={form.onSubmit(handleButtonClick)}>
                    <Outlet />
                    <Divider my={'md'}/>
                    <AuthDialogFooter/>
                </form>
            </AuthFormProvider>
        </Card>
    )
});

const brandingTitle = language('branding.name')
const authDialogLoginTitle = language('auth.sign_in.header');
const authDialogSignUpTitle = language('auth.sign_up.header');

const AuthDialogHeader = React.memo(() => {
    const location = useLocation();
    const title = location?.pathname === RoutePath.HOME ? authDialogLoginTitle : authDialogSignUpTitle;

    return (
        <Stack mih={50} gap="xs" justify="center" align="center">
            <Group>
                <Image src={logo} height={150} width={150} alt={'branding'}/>
            </Group>
            <Title order={1} style={{color: '#494850'}}>{brandingTitle}</Title>
            <Title order={2} style={{color: '#469d4b'}}>{title}</Title>
        </Stack>
    )
});

const loginButtonTitle = language('auth.sign_in.button_title');

const signUpButtonTitle = language('auth.sign_up.button_title');
const signUpButtonSuccessTitle = language('auth.sign_up.success.button_title');
const signUpButtonFailureTitle = language('auth.sign_up.error.button_title');

const AuthDialogFooter = React.memo(() => {
    const navigate = useNavigate();
    const location = useLocation();
    const form = useAuthFormContext();
    const buttonTitle = useMemo(() => {
        if (location?.pathname === RoutePath.HOME) return loginButtonTitle;
        else if (location?.pathname === RoutePath.SIGN_UP) return signUpButtonTitle;
        else if (location?.pathname === RoutePath.SIGN_UP_SUCCESS) return signUpButtonSuccessTitle;
        else if (location?.pathname === RoutePath.SIGN_UP_FAILURE) return signUpButtonFailureTitle;
    }, [location?.pathname]);

    const linkTitle = useMemo(() => {
        if (location?.pathname === RoutePath.HOME) return 'Need an Account?';
        else if (location?.pathname === RoutePath.SIGN_UP) return 'Already have an Account?';
        else return;
    }, [location?.pathname]);

    const isFullwidth = useMemo(() => (location?.pathname !== RoutePath.HOME && location?.pathname !== RoutePath.SIGN_UP), [location?.pathname])

    const handleNavigation = useCallback(() => {
        if (location?.pathname === RoutePath.SIGN_UP) navigate(RoutePath.HOME);
        else if (location?.pathname === RoutePath.HOME) navigate(RoutePath.SIGN_UP)
        else navigate(RoutePath.HOME)
    }, [location?.pathname, navigate]);

    const isDisabled = useMemo(() => {
        if (location.pathname === RoutePath.HOME) return !(isLoginValid(form.values));
        else if (location.pathname === RoutePath.SIGN_UP) return !(isSignUpValid(form.values));
        else return false;
    }, [location?.pathname, form.values])

    return (
        <Group justify="space-between">
            { linkTitle && <Anchor size={'sm'} underline={'never'} onClick={handleNavigation}>{linkTitle}</Anchor> }
            <Button color={'#469d4b'} radius={'md'} type={'submit'} fullWidth={isFullwidth} disabled={isDisabled}>{buttonTitle}</Button>
        </Group>
    )
})

export default AuthDialog