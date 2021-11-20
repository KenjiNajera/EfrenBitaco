import React from 'react';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import { makeStyles, withStyles} from '@material-ui/core/styles';
import gitlogo from '../../assets/images/ic_git_logo_icon.svg';
import gittext from '../../assets/images/ic_git_logo_text.svg';

import { createMuiTheme } from '@material-ui/core/styles';
import blue from '@material-ui/core/colors/blue';
import { ThemeProvider } from '@material-ui/styles';


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
    secondary:{
      main: '#858585',
    },
  },
});

const useStyles = makeStyles((theme) => ({
  root:{
    height: '100%',
    backgroundColor: '#858585',
    display: 'block',
    borderTopRightRadius: 10,
    borderBottomRightRadius: 10,
    color: 'white',
  },
  paper: {
    margin: theme.spacing(8,4),
    flexDirection: 'column',
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
    height:'55px',
    maxHeight: '100%',
  },
  image: {
    backgroundRepeat: 'no-repeat',
    backgroundColor:
      theme.palette.type === 'light' ? theme.palette.grey[50] : theme.palette.grey[900],
    backgroundSize: 'cover',
    backgroundPosition: 'center',
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
  links:{
    color: '#ffffff',
  },
}));

export default function SignIn() {
  const classes = useStyles();

  return (
    <Grid container component="main" className={classes.root}>
      <CssBaseline />
      <Grid item>
        <div className={classes.paper}>
          <form className={classes.form} noValidate>
            <Grid container direction="row" alignItems="center" ClassName={classes.image} >
              <img className={classes.img2} alt='texto' src={gitlogo} />
              <img className={classes.img3} alt='logo' src={gittext} />
            </Grid>
            <ThemeProvider theme={theme}>
              <CssTextField
                variant="filled"
                margin="normal"
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                autoFocus
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
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
                className={classes.textf}
                inputProps={{ style: { color: '#0367A6' } }}
              />
            </ThemeProvider>
            <ThemeProvider theme={theme}>
              <FormControlLabel
                control={<Checkbox value="remember" color="primary" />}
                label="Recuérdame"
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
              Sign In
            </Button>
            </ThemeProvider>
            <Grid
              container
              justify="center">
              <Grid item>
                <Link href="/auth/recovery+password" className={classes.links}>
                  ¿Se te olvidó tu contraseña?
                </Link>
              </Grid>
            </Grid>
          </form>
        </div>
      </Grid>
    </Grid>
  );
}
