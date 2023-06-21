import React from 'react';
import {Button as RSButton, ButtonProps as RSButtonProps} from 'reactstrap';

import './styles.scss';

interface ButtonProps extends RSButtonProps {
  testID?: string;
  title?: string;
  onClick?: () => void;
  color?: string;
  disable?: boolean;
}

const Button: React.FC<ButtonProps> = ({
  testID = 'primary-button',
  title,
  color = 'primary',
  onClick,
  disable,
  ...props
}) => {
  return (
    <RSButton id={testID} color={color} size={'sm'} block onClick={onClick} disabled={disable} {...props}>
      {title}
    </RSButton>
  );
};

export default Button;
