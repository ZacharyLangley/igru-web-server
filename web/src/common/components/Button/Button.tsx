import React from 'react';
import {Button as RSButton} from 'reactstrap';

import './styles.scss';

interface PrimaryButtonProps {
  testID?: string;
  title?: string;
  onClick?: () => void;
  color?: string;
  disable?: boolean;
}

const Button: React.FC<PrimaryButtonProps> = ({
  testID = 'primary-button',
  title,
  color = 'primary',
  onClick,
  disable,
}) => {
  return (
    <RSButton id={testID} color={color} size={'sm'} block onClick={onClick} disabled={disable}>
      {title}
    </RSButton>
  );
};

export default Button;
