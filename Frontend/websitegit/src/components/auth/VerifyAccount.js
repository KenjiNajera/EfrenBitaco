import { 
    Button, 
    Grid, 
    makeStyles, 
    TextField, 
    Typography, 
    withStyles 
} from '@material-ui/core';
import React from 'react';
import { createMuiTheme } from '@material-ui/core/styles';
import { ThemeProvider } from '@material-ui/styles';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers';
import * as yup from 'yup';
import { ErrorMessage } from '@hookform/error-message';

const VerifyAccount = () => {
    const CssTextField = withStyles(theme => ({
        root: {
            marginTop:30,
            '& label.Mui-focused': {
                color: '#0597F2',
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
        }
    }))(TextField);
    
    const useStyles = makeStyles(theme => ({
        root:{
            color: 'white',
            padding: theme.spacing(3),
            justifyContent: 'center',
            alignItems: 'center',
            
        },
        title: {
            textAlign: 'center'
        },
        error: {
            '& span':{
                color: '#bf1650',
                '&::before':{
                    display: 'inline',
                    content: '"⚠ "',
                },
            },
        },
        submit: {
            margin: theme.spacing(3, 0, 2),
            backgroundColor: '#0367A6',
            marginTop:30,
            width:'70%',
            marginLeft:'16%',
        },
    }));

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

    const classes = useStyles();

    const schema = yup.object().shape({
        password: yup.string().trim()
            .required(() => <span>Password is required</span>),
        confirmPassword: yup.string().trim()
            // .oneOf([yup.ref('password'), null], 'Passwords must match')
            .required(() => <span>Confirm Password is required</span>),
    });
    const { register, handleSubmit, errors } = useForm({
        resolver: yupResolver(schema),
    });
    
    const onSubmit = data => {
        console.log(data)
    }

    return (
        <div>
            <Grid 
                container 
                spacing={5} 
                className={ classes.root }
                direction="column"
                justify="flex-start"
                alignItems="center">
                <Grid item xs={12} className={ classes.title }>
                    <ThemeProvider theme={theme}>
                    <Typography variant="h3" gutterBottom color="primary">
                        Verify Account
                    </Typography>
                    </ThemeProvider>
                </Grid>
                <Grid item xs={12}>
                <ThemeProvider theme={theme}>
                    <Typography variant="subtitle1" gutterBottom color="primary">
                        Hello, you have verified your account, now set your password to finish the registration process.
                    </Typography>
                    </ThemeProvider>
                </Grid>
                <Grid item xs={12}>
                    <form onSubmit={ e => e.preventDefault() }>
                        <CssTextField
                            variant="outlined"
                            margin="normal"
                            fullWidth
                            type="password"
                            label="Password"
                            name="password"
                            inputRef={ register }
                            className={ classes.error }
                            helperText={ <ErrorMessage errors={errors} name="password" /> }
                            inputProps={{ style: { color: 'white' } }} />
                        <CssTextField
                            variant="outlined"
                            margin="normal"
                            type="password"
                            fullWidth
                            label="Confirm Password"
                            name="confirmPassword"
                            inputRef={ register }
                            className={ classes.error }
                            helperText={ <ErrorMessage errors={errors} name="confirmPassword" /> }
                            inputProps={{ style: { color: 'white' } }} />
                        <ThemeProvider theme={theme}>
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            color="primary"
                            className={classes.submit}
                            onClick={ handleSubmit( onSubmit ) } >
                            Verify Account
                        </Button>
                        </ThemeProvider>
                    </form>
                </Grid>
            </Grid>
        </div>
    )
}

export default VerifyAccount
