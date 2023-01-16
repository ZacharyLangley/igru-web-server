import React from 'react';
import {Alert as RSAlert, AlertProps as RSAlertProps} from 'reactstrap';

import {AlertTypes, AlertColors} from '../../types';

export interface AlertProps extends RSAlertProps {
  alertType: AlertTypes;
  onDismiss?: () => void;
}

const Alert: React.FC<AlertProps> = ({
  alertType,
  onDismiss,
  children,
  ...props
}) => {
  const determineAlertColor = (alertType: AlertTypes) => {
    if (alertType === AlertTypes.MESSAGE) return AlertColors.INFO;
    if (alertType === AlertTypes.FAILURE || alertType === AlertTypes.CRITICAL)
      return AlertColors.DANGER;
    if (alertType === AlertTypes.ERROR) return AlertColors.DARK;
    if (alertType === AlertTypes.NONCRITICAL) return AlertColors.WARNING;
    if (alertType === AlertTypes.SUCCESS) return AlertColors.SUCCESS;
    return AlertColors.DARK;
  };

  const color = determineAlertColor(alertType);
  return (
    <RSAlert color={color} toggle={onDismiss} {...props}>
      {children}
    </RSAlert>
  );
};

export default Alert;
