import React from 'react'
import Toolbar from '@material-ui/core/Toolbar';
import { makeStyles, withStyles, fade } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import { Typography } from '@material-ui/core';
import DataPersonal from './DataPersonal';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    color: {
        backgroundColor: '#FFFFFF'
    },

    profile: {
        fontSize: '2.5vh',
        fontWeight: 600,
        flexGrow: 1,
    },
    title: {
        fontSize: '2vh',
    },
    range: {
        fontSize: '2.5vh',
        fontWeight: 600,
        marginLeft: '1vh'
    },
    pad: {
        padding: '4%',
        marginBottom: '2vh',
    },
    submit: {
        justifyContent: 'center',
        marginLeft: '40%',
        marginTop: '4%',
        backgroundColor: '#0367A6',
        height: '6vh',
        width: '25vh',
    },
    layout: {
        width: 'auto',
        marginLeft: '1vh',
        marginTop: '2.3vh',

    },
    paper: {
        marginTop: theme.spacing(3),
        marginBottom: theme.spacing(3),
        padding: theme.spacing(2),
        [theme.breakpoints.up(600 + theme.spacing(3) * 2)]: {
            marginTop: theme.spacing(1),
            marginBottom: theme.spacing(6),
            padding: theme.spacing(3),
        },
    },
}));


export default function Cardprofile() {
    const classes = useStyles();
    return (
        <main className={classes.layout}>
            <Paper className={classes.paper}>

                <Toolbar className={classes.color}>
                    <Typography className={classes.profile}>
                        Profile
            </Typography>
                    <Typography className={classes.title}>
                        Title:
            </Typography>
                    <Typography className={classes.range}>
                        Internship
            </Typography>
                </Toolbar>
                <hr></hr>
                <Grid className={classes.pad}>
                    <DataPersonal />
                </Grid>
                <hr></hr>
                <Button
                    type="submit"
                    fullWidth
                    variant="contained"
                    color="primary"
                    className={classes.submit}
                >
                    Upload
            </Button>

            </Paper>
        </main>
    )
}

