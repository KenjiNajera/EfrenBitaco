import React from 'react';
import PropTypes from 'prop-types';

import { Modal, makeStyles, IconButton } from '@material-ui/core';
import Backdrop from '@material-ui/core/Backdrop';
import Fade from '@material-ui/core/Fade';
import CloseIcon from '@material-ui/icons/Close';

const ModalCustom = ({ 
    handleModalClose, 
    openModal, 
    component: Component,
}) => {

    const useStyles = makeStyles((theme) => ({
        modal: {
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            fontSize: 20,
            color:'#434E65',
            // fontFamily: ['Segoe UI',],
        },
        paper: {
            position: 'relative',
            backgroundColor: theme.palette.background.paper,
            boxShadow: theme.shadows[5],
        },
        icon: {
            position: 'absolute',
            right: 5,
            top: 5,
        }
    }))
    const classes = useStyles();

    return (
        <>
            <Modal   
                aria-labelledby="transition-modal-title"
                aria-describedby="transition-modal-description"
                className={ classes.modal }
                open={ openModal }
                onClose={ handleModalClose } 
                closeAfterTransition
                BackdropComponent={ Backdrop }
                BackdropProps={{
                    timeout: 500,
                }} >
                    <Fade in={ openModal }  >
                        <div className={classes.paper}>
                            <IconButton onClick={ handleModalClose } size="small" className={ classes.icon }>
                                <CloseIcon fontSize="small"/>
                            </IconButton>
                            <Component 
                                handleModalClose={ handleModalClose } />
                        </div>
                    </Fade>
            </Modal>
        </>
    )
}

ModalCustom.propType = {
    handleModalClose: PropTypes.func.isRequired,
    openModal: PropTypes.bool.isRequired,
    component: PropTypes.func.isRequired
}

export default ModalCustom
