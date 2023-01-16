import {
  DispatchMessageAlertAction,
  DispatchSuccessAlertAction,
  DispatchFailureAlertAction,
  DispatchErrorAlertAction,
  DispatchCriticalAlertAction,
  DispatchNoncriticalAlertAction,
  DismissSelectedAlertAction,
  DismissAllAlertsAction,
} from '../interfaces/alert';
import {AlertActionTypes} from '../types/alert';

export const dispatchMessageAlertAction = (
  message: string
): DispatchMessageAlertAction => ({
  type: AlertActionTypes.DISPATCH_ALERT_MESSAGE,
  value: {
    message,
  },
});

export const dispatchSuccessAlertAction = (
  message: string
): DispatchSuccessAlertAction => ({
  type: AlertActionTypes.DISPATCH_ALERT_SUCCESS,
  value: {
    message,
  },
});

export const dispatchFailureAlertAction = (
  message: string
): DispatchFailureAlertAction => ({
  type: AlertActionTypes.DISPATCH_ALERT_FAILURE,
  value: {
    message,
  },
});

export const dispatchErrorAlertAction = (
  message: string
): DispatchErrorAlertAction => ({
  type: AlertActionTypes.DISPATCH_ALERT_ERROR,
  value: {
    message,
  },
});

export const dispatchCriticalAlertAction = (
  message: string
): DispatchCriticalAlertAction => ({
  type: AlertActionTypes.DISPATCH_ALERT_CRITICAL,
  value: {
    message,
  },
});

export const dispatchNoncriticalAlertAction = (
  message: string
): DispatchNoncriticalAlertAction => ({
  type: AlertActionTypes.DISPATCH_ALERT_NONCRITICAL,
  value: {
    message,
  },
});

export const dismissSelectedAlertAction = (
  id: number
): DismissSelectedAlertAction => ({
  type: AlertActionTypes.DISMISS_SELECTED_ALERT,
  value: {
    id,
  },
});

export const dismissAllAlertsAction = (): DismissAllAlertsAction => ({
  type: AlertActionTypes.DISMISS_ALL_ALERTS,
});
