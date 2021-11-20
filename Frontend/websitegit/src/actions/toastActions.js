import { types } from '../types/types';

export const setSuccessToast = () => ({
    type: types.success
});

export const setErrorToast = () => ({
    type: types.error
});

export const setInfoToast = () => ({
    type: types.info
});

export const setWarningToast = () => ({
    type: types.warning
});

export const setclearToast = () => ({
    type: types.clear
})