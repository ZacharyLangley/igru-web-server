import React from 'react';
import {Button} from 'reactstrap';

import './styles.scss';

interface PrimaryButtonProps {
  testID?: string;
  title?: string;
  onClick?: () => void;
  disable?: boolean;
}

const PrimaryButton: React.FC<PrimaryButtonProps> = ({
  testID = 'primary-button',
  title,
  onClick,
  disable,
}) => {
  return (
    <Button id={testID} color={'primary'} size={'sm'} block onClick={onClick} disabled={disable}>
      {title}
    </Button>
  );
};

export default PrimaryButton;
