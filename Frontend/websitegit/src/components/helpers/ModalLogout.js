import React from 'react';

import { makeStyles, Button } from '@material-ui/core';
import LockIcon from '@material-ui/icons/Lock';

const ModalLogout = ({ handleModalClose }) => {
    const useStyles = makeStyles((theme) => ({
        title:{
            display: 'flex',
            marginLeft: theme.spacing(2),
            marginTop: theme.spacing(1),
            '& p': {
                opacity:0.8,
                marginRight: theme.spacing(1),    
            },
        },
        body: {
            display: 'flex',
            justifyContent: 'center',
            margin: theme.spacing(2),
        },
        buttonsContent: {
            justifyContent: 'right',
            display: 'flex',
            backgroundColor: '#0367A6',
            padding: theme.spacing(1),
            '& Button': {
                textTransform: 'none',
                padding: theme.spacing(0.3, 1, 0.3, 1),
            },
            '& #cancel': {
                color: '#434E65',
                marginRight: theme.spacing(2),
                backgroundColor: '#EBEBEB',
            },
            '& #yes': {
                color:'white',
                backgroundColor: '#DC3545',
            },
        },
    }));
    const classes = useStyles();

    const handleLogOut = () => {
        // cierra sesion
    }

    return (
        <>
            <div className={ classes.title } >
                <p id="transition-modal-title" >Log Out <LockIcon /></p> 
            </div>
            <div className={ classes.body } > 
                <p>Are you sure you want to log out?</p>
            </div>
            <div className={ classes.buttonsContent } >
                <Button variant="contained" id="cancel" onClick={ handleModalClose } >Cancel</Button>
                <Button variant="contained" id="yes" onClick={ handleLogOut } >Yes, I'm sure</Button>
            </div>
        </>
    )
}

export default ModalLogout

