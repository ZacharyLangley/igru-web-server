import React from 'react';
import Logo from '../Logo/Logo';

import './styles.scss'

const Branding = () => {
    return (
        <div className={'branding'}>
            <Logo height={100} width={100} hideName />
            <div className={'title'}>{'IGRU'}</div>
        </div>
    )
};

export default Branding;