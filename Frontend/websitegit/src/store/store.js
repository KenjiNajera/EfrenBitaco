import thunk from 'redux-thunk';
import { createStore, combineReducers, applyMiddleware, compose } from 'redux';

import { authReducer } from '../reducers/authReducer';
import { persistenceReducer } from '../reducers/persistenceReducer';
import { toastReducer } from '../reducers/toastReducer';
import { menuReducer } from '../reducers/menuReducer';

const composeEnhancers = (typeof window !== 'undefined' && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) || compose;

const reducers = combineReducers({
    auth: authReducer,
    persistence: persistenceReducer,
    toast: toastReducer,
    menu: menuReducer,
})


const store = createStore( 
    reducers,
    composeEnhancers(
        applyMiddleware( thunk )
    )
    );

export default store;