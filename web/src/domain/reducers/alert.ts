import {AlertTypes} from '../../common/types';
import {
  Alert,
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

export interface AlertState {
  alerts: Alert[];
}

const INITIAL_STATE: AlertState = {
  alerts: [],
};

const addAlertItem = (
  state: AlertState,
  item: Alert,
  alertType: AlertTypes
) => ({
  ...state,
  alerts: [{...item, alertType}, ...state.alerts],
});

export default function alerts(
  state: AlertState = INITIAL_STATE,
  action:
    | DispatchMessageAlertAction
    | DispatchSuccessAlertAction
    | DispatchFailureAlertAction
    | DispatchErrorAlertAction
    | DispatchCriticalAlertAction
    | DispatchNoncriticalAlertAction
    | DismissSelectedAlertAction
    | DismissAllAlertsAction
) {
  switch (action.type) {
    case AlertActionTypes.DISPATCH_ALERT_MESSAGE:
      return addAlertItem(state, action.value, AlertTypes.MESSAGE);
    case AlertActionTypes.DISPATCH_ALERT_SUCCESS:
      return addAlertItem(state, action.value, AlertTypes.SUCCESS);
    case AlertActionTypes.DISPATCH_ALERT_FAILURE:
      return addAlertItem(state, action.value, AlertTypes.FAILURE);
    case AlertActionTypes.DISPATCH_ALERT_ERROR:
      return addAlertItem(state, action.value, AlertTypes.ERROR);
    case AlertActionTypes.DISPATCH_ALERT_CRITICAL:
      return addAlertItem(state, action.value, AlertTypes.CRITICAL);
    case AlertActionTypes.DISPATCH_ALERT_NONCRITICAL:
      return addAlertItem(state, action.value, AlertTypes.MESSAGE);
    case AlertActionTypes.DISMISS_SELECTED_ALERT:
      if (action?.value?.id === undefined) return state;
      return {
        ...state,
        alerts: state.alerts.filter(
          (alert, index) => index !== action.value.id
        ),
      };
    case AlertActionTypes.DISMISS_ALL_ALERTS:
      return {
        ...state,
        alerts: [],
      };
    default:
      return state;
  }
}
