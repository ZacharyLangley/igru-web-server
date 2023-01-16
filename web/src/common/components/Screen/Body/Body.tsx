import React from 'react';

import './styles.scss';

export interface BodyProps {
  testID?: string;
  children?: JSX.Element | JSX.Element[] | string;
}

const Body: React.FC<BodyProps> = ({
  testID = 'global-layout-body',
  children,
}) => {
  return (
    <div id={testID} className={'body-container'}>
      {children}
    </div>
  );
};

export default Body;
