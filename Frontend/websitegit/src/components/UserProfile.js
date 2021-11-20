import React from 'react'
import Toolbar from '@material-ui/core/Toolbar';
import { makeStyles, withStyles, fade } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import { Typography } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import ProfileSettings from './helpers/ProfileSettings'
import CardProfile from './helpers/CardProfile'


const useStyles = makeStyles((theme) => ({
    root: {
        position: 'fixed',
        height: '100vh',
        width: '100vw',
        overflow: 'scroll',
      
    },
    secundario:{
        display: 'flex',
        margin: '0px',
        height: '100%',
        width: '100%',
    }
}));


export default function UserProfile() {
    const classes = useStyles();
    return (
        <>
            <Grid className={classes.secundario}>
            <ProfileSettings />
            <main>
            <CardProfile/>
            </main>
            </Grid>
        
        </>
    )
}
