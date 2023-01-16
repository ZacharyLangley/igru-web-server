import React from 'react';

import logo from '../../assets/branding/IGRU_White_logo.png';
import './styles.scss';

export interface LogoProps {
  testID?: string;
  onClick?: () => void;
  height?: number;
  width?: number;
  hideName?: boolean;
}

const Logo: React.FC<LogoProps> = ({
  testID = 'igru-branding-logo',
  onClick,
  height = 40,
  width = 40,
  hideName = false,
}) => (
  <div
    id={`${testID}:container`}
    className={'logo-container'}
    onClick={onClick}
  >
    <div id={`${testID}:img:container`} className={'logo-img-container'}>
      <img
        id={`${testID}:img:asset`}
        className={'logo-img'}
        src={logo}
        alt={'branding'}
        height={height}
        width={width}
      />
    </div>
    {!hideName && (
      <div id={`${testID}:text`} className={'logo-text'}>
        {'IGRU'}
      </div>
    )}
  </div>
);

export default Logo;
