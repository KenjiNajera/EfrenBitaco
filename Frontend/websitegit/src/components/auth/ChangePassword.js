import React from 'react';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import { makeStyles, withStyles } from '@material-ui/core/styles';
import gitlogo from '../../assets/images/ic_git_logo_icon.svg';
import gittext from '../../assets/images/ic_git_logo_text.svg';
import { createMuiTheme } from '@material-ui/core/styles';
import { ThemeProvider } from '@material-ui/styles';
import { Typography } from '@material-ui/core';


const CssTextField = withStyles(theme => ({
    root: {
        '& label.Mui-focused': {
            color: '#0367A6',
        },
        '& .MuiInput-underline:after': {
            borderBottomColor: '#0597F2',
        },
        '& .MuiOutlinedInput-root': {
            '& fieldset': {
                borderColor: '#0367A6',
            },
            '&:hover fieldset': {
                borderColor: '#0367A6',
            },
            '&.Mui-focused fieldset': {
                borderColor: '#0597F2',
            }
        }
    },
    palette: {
        primary: {
            main: '#0367A6',
        },
    },
}))(TextField);

const theme = createMuiTheme({
    palette: {
        primary: {
            main: '#0367A6',
        },
        secondary: {
            main: '#E6E6E6',
        },
    },
});

const useStyles = makeStyles((theme) => ({
    root: {
        backgroundColor: '#858585',
        display: 'flex',
        borderRadius: 10,
        justifyContent: 'center',
        alignItems: 'center',
    },
    paper: {
        margin: theme.spacing(8, 4),
        alignItems: 'center',
    },
    form: {
        marginTop: theme.spacing(5),
    },
    textf: {
        borderRadius: '5px',
        backgroundColor: theme.palette.common.white,
    },
    submit: {
        margin: theme.spacing(4, 0, 4),
        backgroundColor: 'white',
        color: '#0367A6',
        height: '55px',
        maxHeight: '100%',
    },
    img2: {
        display: 'block',
        maxWidth: '35%',
        maxHeight: '100%',
    },
    img3: {
        display: 'block',
        maxWidth: '65%',
        maxHeight: '100%',
    },
}));

export default function SignIn() {
    const classes = useStyles();

    return (
        <Grid container justify="center" component="main" className={classes.root}>
            <CssBaseline />
            <Grid item>
                <div className={classes.paper}>
                    <form className={classes.form} noValidate>
                        <Grid container direction="row" alignItems="center" justify="center" ClassName={classes.image} >
                            <img className={classes.img2} alt='texto' src={gitlogo} />
                            <img className={classes.img3} alt='logo' src={gittext} />
                        </Grid>
                        <ThemeProvider theme={theme}>
                            <ThemeProvider theme={theme}>
                                <Grid container justify="center">
                                    <Typography>
                                        <h2 >Change your password</h2>
                                    </Typography>
                                </Grid>
                            </ThemeProvider>
                            <CssTextField
                                variant="filled"
                                margin="normal"
                                fullWidth
                                name="password"
                                label="Current Password"
                                placeholder="Current Password"
                                type="password"
                                id="password"
                                autoComplete="current-password"
                                className={classes.textf}
                                inputProps={{ style: { color: '#0367A6' } }}
                            />
                        </ThemeProvider>
                        <ThemeProvider theme={theme}>
                            <CssTextField
                                variant="filled"
                                margin="normal"
                                fullWidth
                                name="password"
                                label="New Password"
                                placeholder="New Password"
                                type="password"
                                id="password"
                                autoComplete="current-password"
                                className={classes.textf}
                                inputProps={{ style: { color: '#0367A6' } }}
                            />
                        </ThemeProvider>
                        <ThemeProvider theme={theme}>
                            <CssTextField
                                variant="filled"
                                margin="normal"
                                fullWidth
                                name="password"
                                label="Confirm New Password"
                                placeholder="Confirm New Password"
                                type="password"
                                id="password"
                                autoComplete="current-password"
                                className={classes.textf}
                                inputProps={{ style: { color: '#0367A6' } }}
                            />
                        </ThemeProvider>
                        <ThemeProvider theme={theme}>
                            <Button
                                type="submit"
                                fullWidth
                                variant="contained"
                                color="secondary"
                                className={classes.submit}
                            >
                                Change my password
                            </Button>
                        </ThemeProvider>
                    </form>
                </div>
            </Grid>
        </Grid>
    );
}
