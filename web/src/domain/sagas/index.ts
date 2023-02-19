import {all} from 'redux-saga/effects';

import userSagas from './user';
import sessionsSagas from './sessions';

export default function* rootSaga() {
  yield all([
    ...userSagas,
    ...sessionsSagas
  ]);
}
