import { types } from '../types/types';

export const toastReducer = ( state = {}, action ) => {

    switch (action.type) {
        case types.success:
            return {
                type: 'success',
                message: 'This is a success toast component',
                isOpen: true,
            }

        case types.error:
            return {
                type: 'error',
                message: 'This is a error toast component',
                isOpen: true,
            }
        case types.info:
            return {
                type: 'info',
                message: 'This is an info toast component',
                isOpen: true,
            }
        case types.warning:
            return {
                type: 'warning',
                message: 'This is a warning toast component',
                isOpen: true,
            }
        
        case types.clear:
            return {
                ...state,
                isOpen: false,
            }
            
        default:
            return {}
    }
}
