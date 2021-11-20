
import { types } from '../types/types';

const initialState = {
    form: {},
    data: [],
}

export const persistenceReducer = ( state = initialState , action) => {
    switch (action.type) {
        case types.reset:
            return {
                ...state,
                form: {},
                data: [],
            }

        case types.formchange:
            return {
                ...state,
                form: action.payload.data,
            }

        case types.datachange:
            return {
                ...state,
                data: action.payload.data,
            }
            
        case types.deleteDataById:
            return {
                ...state,
                data: state.data.filter( data => data.id !== action.payload )
            }
        
        case types.createData:
            return {
                ...state,
                data: [ action.payload, ...state.data ]
            }

        case types.updateDataById:
            return {
                ...state,
                data: state.data.map(
                    data => data.id === action.payload.id 
                        ?   action.payload.data
                        :   data
                )
            }

        default:
            return state 
    }
}
