import React from 'react';
import ReactDOM from 'react-dom/client';
import axios from 'axios';
import {Provider} from 'react-redux';
import {createStore, applyMiddleware} from 'redux';
import createSagaMiddleware from '@redux-saga/core';

import reducers from './domain/reducers';
import rootSaga from './domain/sagas';
import App from './app/App';
import reportWebVitals from './reportWebVitals';

import 'bootstrap/dist/css/bootstrap.min.css';

axios.defaults.withCredentials = true;
axios.defaults.baseURL = '';

const sagaMiddleware = createSagaMiddleware();
const store = createStore(reducers, applyMiddleware(sagaMiddleware));
sagaMiddleware.run(rootSaga);

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

root.render(
  <React.StrictMode>
    <Provider store={store}>
      <App />
    </Provider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
