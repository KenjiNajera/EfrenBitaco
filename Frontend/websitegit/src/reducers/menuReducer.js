import { types } from "../types/types"

const initialState = {
    open:false
}

export const menuReducer = ( state = initialState, action ) => {

    switch ( action.type ) {
        case types.changestatus:
            return {
                open:action.payload.open,
            }
        default:
            return state;
    }
}