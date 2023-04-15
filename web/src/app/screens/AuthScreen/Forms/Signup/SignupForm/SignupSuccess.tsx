import React from 'react';
import { useNavigate } from 'react-router-dom';

import language from 'src/common/language/index';
import { RoutePath } from 'src/app/types/routes';
import PrimaryButton from '../../../../../../common/components/Button/PrimaryButton/PrimaryButton';
import useUser from 'src/store/useUser/useUser';

interface SignupSuccessProps {}

const text = {
    messageLineTop: language("auth.sign_up.success.message_top"),
    messageLineBottom: language("auth.sign_up.success.message_bottom"),
    buttonTitle: language("auth.sign_up.success.button_title")
}

export const SignupSuccess: React.FC<SignupSuccessProps> = () => {
    const {resetSignUpStatus} = useUser();
    const navigate = useNavigate();

    const onClick = () => {
        resetSignUpStatus();
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
