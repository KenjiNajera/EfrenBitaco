import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles, withStyles } from '@material-ui/core/styles';

const CssTextField = withStyles(theme => ({
    root: {
        marginTop:40,
        '& label.Mui-focused': {
            color: 'white',
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

const useStyles = makeStyles((theme) => ({
    root: {
        height: '69.9vh',
        borderBottomRightRadius: 10,
        borderTopRightRadius: 10,
        backgroundColor: '#858585',
        color: '#ffffff',
    },
    paper: {

        margin: theme.spacing(8, 4),
        display: 'flex',
        flexDirection: 'column',
    },
    title: {
        textAlign: 'center'
    },
    form: {
        width: '100%',
        marginBottom: theme.spacing(5),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
        backgroundColor: '#0367A6',
        marginTop:20,
        width:'70%',
        marginLeft:'16%',
    },
}));

export default function SignInSide() {
    const classes = useStyles();

    return (
        <Grid container component="main" className={classes.root}>
            <div className={classes.paper}>
                <Typography variant="h4" gutterBottom className={classes.title}> 
                    ¿Olvidaste su contraseña?
                </Typography>
                <Typography variant="h6" gutterBottom >
                    ¡No hay problema!
                </Typography>
                <Typography variant="body1" gutterBottom>
                    Sólo escribe el correo electrónico con el que te registraste y te enviaremos un enlace para reestablecer tu contraseña.
                </Typography>
                <form className={classes.form} noValidate>
                    <CssTextField
                        variant="outlined"
                        margin="normal"
                        color="#ffffff"
                        fullWidth
                        id="email"
                        label="Email Address"
                        name="email"
                        autoComplete="email"
                        autoFocus
                        inputProps={{ style: { color: 'white' } }}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={classes.submit} >
                        Change Password
                    </Button>
                </form>
                <Typography variant="body2">
                    Nota: Si no recibes nuestro e-mail revisa que hayas escrito tu correo correctamente o
                    tal vez llegó a tu bandeja de correo no deseado, si todo está bien y
                    aún tienes problemas para ingresar escribe a contacto@grupogit.com
                </Typography>
            </div>

        </Grid>
    );
}
