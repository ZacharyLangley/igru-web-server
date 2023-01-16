import React from 'react';

import './styles.scss';

export interface GlobalLayoutProps {
  testID?: string;
  children: JSX.Element | JSX.Element[];
}

const Screen: React.FC<GlobalLayoutProps> = ({testID, children}) => {
  return (
    <div id={testID} className={'screen-container'}>
      {children}
    </div>
  );
};

export default Screen;
