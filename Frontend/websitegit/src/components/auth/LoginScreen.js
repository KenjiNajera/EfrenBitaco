import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import { makeStyles, withStyles } from '@material-ui/core/styles';
import logo from '../../assets/images/ic_git_logo_icon.svg';
import logo2 from '../../assets/images/ic_git_logo_text.svg';

const CssTextField = withStyles(theme => ({
  root: {
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
    alignItems: 'center',
  },
  logo: {
    
    width:62,
   marginLeft:50,
  },
  logos: {
    
    width:106,
   },
   lg: {
    marginTop:20,
    marginLeft:100,
   },
  form: {
    width: '100%', // Fix IE 11 issue.
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
    backgroundColor: '#0367A6',
  },
  links:{
    color: '#ffffff',
    marginLeft:'28%',
  }
}));

export default function SignInSide() {
  const classes = useStyles();

  return (
    <Grid container component="main" className={classes.root}>
      <Grid container spacing={3}>
      <Grid item xs={12} sm={6} className={classes.lg}>
        <img className={classes.logo}  alt="Portal Logo" src={logo} align="center" />
        <img className={classes.logos}  alt="log" src={logo2} align="center" />
      </Grid>
      </Grid>
        <div className={classes.paper}>
      
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
              inputProps={{ style: { color: 'white'}}}
            />
            <CssTextField
              variant="outlined"
              margin="normal"
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
              inputProps={{ style: { color: 'white'}}}
            />
            <FormControlLabel
              control={<Checkbox value="remember" color="primary" />}
              label="Remember me"
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              className={classes.submit}
            >
              Sign In
            </Button>
            <Grid container>
              <Grid item xs>
                <Link href="/auth/recovery+password" className={classes.links}>
                ¿Se te olvidó tu contraseña?
                </Link>
              </Grid>
            </Grid>
            
          </form>
        </div>
      
    </Grid>
  );
}