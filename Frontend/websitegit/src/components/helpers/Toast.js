import React from 'react'
import { Snackbar } from '@material-ui/core';
import MuiAlert from '@material-ui/lab/Alert';
import { useDispatch } from 'react-redux';
import { setclearToast } from '../../actions/toastActions';

function Alert(props) {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const Toast = ({
    isOpen,
    message,
    type
}) => {
    const dispatch = useDispatch();
    const handleClose = (event) => {
        // event.preventDefault();
        dispatch( setclearToast() )
    }

    return (
        <Snackbar open={isOpen} autoHideDuration={4000} onClose={handleClose}>
            <Alert onClose={handleClose} severity={type}>
                { message }
            </Alert>
        </Snackbar>
    )
}

export default Toast
