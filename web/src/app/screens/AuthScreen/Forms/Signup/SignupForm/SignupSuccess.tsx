import React from 'react';
import { useDispatch } from 'react-redux';
import { useNavigate } from 'react-router-dom';

import language from 'src/common/language/index';
import { RoutePath } from 'src/app/types/routes';
import PrimaryButton from '../../../../../../common/components/Button/PrimaryButton/PrimaryButton';
import { signUpResetAction } from '../../../../../../domain/actions/user';

interface SignupSuccessProps {}

const text = {
    messageLineTop: language("auth.sign_up.error.message_top"),
    messageLineBottom: language("auth.sign_up.error.message_bottom"),
    buttonTitle: language("auth.sign_up.error.button_title")
}

export const SignupSuccess: React.FC<SignupSuccessProps> = () => {
    const navigate = useNavigate();
    const dispatch = useDispatch();

    const onClick = () => {
        dispatch(signUpResetAction())
        navigate(RoutePath.HOME)
    }

    return (
        <div style={{display: 'flex', flex: 1, flexDirection: 'column'}}>
            <div style={{textAlign: 'center', paddingTop: 35, paddingBottom: 60}}>
                <p>{text.messageLineTop}</p>
                <p>{text.messageLineBottom}</p>
            </div>
            <PrimaryButton title={text.buttonTitle} onClick={onClick} />
        </div>
    )
};
