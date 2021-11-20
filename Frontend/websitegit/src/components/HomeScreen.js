import React from 'react';

import { makeStyles } from '@material-ui/core';
import logo from '../assets/images/ic_git_logo_icon.svg';
import ViewBinnacleScreen from './ViewBinnacleScreen';

const HomeScreen = () => {
    const styles = makeStyles((theme) => ({
        divContent: {
            width: '100%',
            height: '100%',
            display:'flex',
            justifyContent: 'center',
            alignItems: 'center',
            margin: '0 auto',
        },
        logo: {
            padding: "1rem",
            width: 500,
            maxWidth: 500,
            opacity: 0.5,
            transition: theme.transitions.create('width', {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.enteringScreen,
            }),
            [theme.breakpoints.down("sm")]: {
                width: 280,
                transition: theme.transitions.create('width', {
                    easing: theme.transitions.easing.sharp,
                    duration: theme.transitions.duration.leavingScreen,
                }),
            }
        },
    }));
    const classes = styles();
    
    return (
        // Para pruebas comentar estas lineas
        
        <ViewBinnacleScreen />

        // <div className={classes.divContent}>
        //     <img className={classes.logo} alt="Portal Logo" src={logo} align="center" />
        // </div>

    )
}

export default HomeScreen
