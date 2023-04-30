import React from 'react';

import './styles.scss';

interface DialogProps {
  testID?: string;
  header?: JSX.Element | JSX.Element[] | string;
  body?: JSX.Element | JSX.Element[] | string;
  footer?: JSX.Element | JSX.Element[] | string;
}

const Dialog: React.FC<DialogProps> = ({
  testID = 'dialog',
  header,
  body,
  footer,
}) => {
  return (
    <div id={testID} className={'dialog-container'}>
      {header && (
        <div id={`${testID}:header`} className={'dialog-header'}>
          {header}
        </div>
      )}
      {body && (
        <div id={`${testID}:body`} className={'dialog-body'}>
          {body}
        </div>
      )}
      {footer && (
        <div id={`${testID}:footer`} className={'dialog-footer'}>
          {footer}
        </div>
      )}
    </div>
  );
};

export default Dialog;
