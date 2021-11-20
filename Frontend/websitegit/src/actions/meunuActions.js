import { types } from '../types/types';

export const ChangeStatus = (open) => ({
    type: types.changestatus,
    payload: {
        open
    }
});