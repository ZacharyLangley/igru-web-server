import {Action} from 'redux';
import {AlertTypes} from '../../common/types';
import {AlertActionTypes} from '../types/alert';

export interface Alert {
  id?: number;
  message?: string;
  alertType?: AlertTypes;
}

export interface DispatchMessageAlertAction extends Action {
  type: AlertActionTypes.DISPATCH_ALERT_MESSAGE;
  value: Alert;
}

export interface DispatchSuccessAlertAction extends Action {
  type: AlertActionTypes.DISPATCH_ALERT_SUCCESS;
  value: Alert;
}

export interface DispatchFailureAlertAction extends Action {
  type: AlertActionTypes.DISPATCH_ALERT_FAILURE;
  value: Alert;
}

export interface DispatchErrorAlertAction extends Action {
  type: AlertActionTypes.DISPATCH_ALERT_ERROR;
  value: Alert;
}

export interface DispatchCriticalAlertAction extends Action {
  type: AlertActionTypes.DISPATCH_ALERT_CRITICAL;
  value: Alert;
}

export interface DispatchNoncriticalAlertAction extends Action {
  type: AlertActionTypes.DISPATCH_ALERT_NONCRITICAL;
  value: Alert;
}

export interface DismissSelectedAlertAction extends Action {
  type: AlertActionTypes.DISMISS_SELECTED_ALERT;
  value: Alert;
}

export interface DismissAllAlertsAction extends Action {
  type: AlertActionTypes.DISMISS_ALL_ALERTS;
}
