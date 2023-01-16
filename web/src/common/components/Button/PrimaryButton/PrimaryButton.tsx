import React from 'react';
import {Button} from 'reactstrap';

import './styles.scss';

interface PrimaryButtonProps {
  testID?: string;
  title?: string;
  onClick?: () => void;
}

const PrimaryButton: React.FC<PrimaryButtonProps> = ({
  testID = 'primary-button',
  title,
  onClick,
}) => {
  return (
    <Button id={testID} color={'primary'} size={'sm'} block onClick={onClick}>
      {title}
    </Button>
  );
};

export default PrimaryButton;
