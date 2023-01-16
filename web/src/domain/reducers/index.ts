import {combineReducers} from 'redux';

import alert from './alert';
import user from './user';

export default combineReducers({
  alert,
  user,
});
