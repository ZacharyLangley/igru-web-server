import {createSelector} from 'reselect';

const getAlertState = (state: any) => state.alert;

export const getAllAlertsSelector = createSelector(
  getAlertState,
  (alert) => alert.alerts
);
