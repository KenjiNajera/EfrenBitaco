import React from 'react';
import { makeStyles, withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import CardActionArea from '@material-ui/core/CardActionArea';
import Paper from '@material-ui/core/Paper';

const CardActionAreas = withStyles(theme => ({
    root: {
        marginTop:'2%',
        padding:'0',
        width:'100%'
    }
}))(CardActionArea);
const useStyles = makeStyles((theme) => ({
    layout: {
        width: '18%',
    },
    paper: {    
        [theme.breakpoints.up(600 + theme.spacing(3) * 2)]: {
            marginTop: theme.spacing(2.1), 
        },
    },
    settings: {
        padding: '4%',
        fontSize: '1.2vw',
        marginLeft: '12%',
        fontWeight: 700,
    },
    profile:{
        color: '#00539f',
        "&:before": {
            content: "'Profile'",
            background: '#EBEBEB',
            borderLeft:'6px solid',
            display: 'inline-block',
            minWidth:'28vh',
            maxWidth:'28vh',
            borderBottomLeftRadius: '2px',
            borderBottomRightRadius: '2px',
            borderTopLeftRadius: '2px',
            borderTopRightRadius: '2px',
            padding:'.45em ',
            backgroundColor:'#FFFFFF',
            fontfamily:'Helvetica,arial,sans-serif',
            fontSize: '2vh',
            fontWeight: 400,    
        }
    },
    password:{
        color: '#00539f',
        "&:before": {
            content: "'Password'",
            borderLeft:'6px solid',
            display: 'inline-block',
            borderBottomLeftRadius: '2px',
            borderBottomRightRadius: '2px',
            borderTopLeftRadius: '2px',
            borderTopRightRadius: '2px',
            padding:'.45em ',
            backgroundColor:'#FFFFFF',
            fontfamily:'Helvetica,arial,sans-serif',
            fontSize: '2vh',
            fontWeight: 400,    
        }
    }
}));

export default function MediaCard() {
    const classes = useStyles();

    return (
        <main className={classes.layout}>
        <Paper className={classes.paper}>

            <div className="Content-Settings">
                <Typography className={classes.settings}>
                    Profile settings
                </Typography>

                <hr></hr>
                <CardActionAreas>
                    <span className={classes.profile} ></span>
                </CardActionAreas>
                <CardActionAreas>
                    <span className={classes.password}></span>
                </CardActionAreas>
            </div>
       
            </Paper>
        </main>
    );
}

