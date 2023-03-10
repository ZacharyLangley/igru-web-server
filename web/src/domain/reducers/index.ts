import {combineReducers} from 'redux';

import alert from './alert';
import user from './user';
import session from './session';

export default combineReducers({
  alert,
  user,
  session,
});
