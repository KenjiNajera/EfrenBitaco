import React from 'react';
import { makeStyles, withStyles, ThemeProvider, createMuiTheme} from '@material-ui/core/styles';
import { 
  FormControl,
  CardContent,
  InputAdornment,
  CardActions,
  Typography,
  TextField,
  Button,
  Grid,
  Box, } from '@material-ui/core';
import Speaker_notes from '@material-ui/icons/Comment';
import Event from '@material-ui/icons/Event';
import Flag from '@material-ui/icons/Flag';

import { useForm } from 'react-hook-form';
import { ErrorMessage } from '@hookform/error-message';
import { yupResolver } from '@hookform/resolvers';
import * as yup from 'yup';

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
        borderColor: '#0597F2',
      },
    '&.Mui-focused fieldset': {
      borderColor: '#0367A6',
    },
  },
},
}))(TextField);

const theme = createMuiTheme({
  palette: {
    default: {
      main: '#213568',
      contrastText: '#FFFFFF',
    },
    primary: {
      main: '#4caf50',
      contrastText: '#FFFFFF',
    },
    secondary:{
      main: '#d32f2f',
      contrastText: '#FFFFFF'
    },
  },
});

const textfiel = createMuiTheme({
  palette: {
    primary: {
      main: '#0367A6',
      contrastText: '#0367A6',
    },
    secondary:{
      main: '#858585',
      contrastText: '#0367A6',
    },
    default: {
      main: '#0367A6',
    },
  },
});

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'block',
    minWidth: 422,
    maxWidth: 500,
    maxHeight: '100%',
    height: 550,
    width: 418,
    margin: 10,
  },
  bullet: {
    display: 'inline-block',
    margin: '0 2px',
    transform: 'scale(0.8)',
  },
  title: {
    fontSize: 20,
    color: '#0367A6',
  },
  pos: {
    marginBottom: 8,
    marginLeft: 8,
  },
  formControl: {
    display: 'flex',
    color: '#0367A6',
    marginBottom: 8,
    marginTop: 8,    
  },
  grid: {
    background: '#E6E6E6',
  },
  btndown: {
    margin: '10px',
  },
}));

const ModalApprovedRequest = ({item}) => {
  const classes = useStyles();
  
  // const schema = yup.object().shape({
  //   details: yup.string().required()
  // });
  const { handleSubmit, errors, register, setValue } = useForm({
    // resolver: yupResolver(schema),
    defaultValues: {
      requestid: item.id,
      details: "",
      status: "",
    },
  });

  const onSubmit = data => {
    console.log(data)
  }

  return (
    <form className={classes.root}>
    <CardContent>
      <Box display="flex" justification="center" borderBottom={1}>
        <Typography className={classes.title} color="textSecondary" gutterBottom>
          Solicitud #{item.id} : Confirmar/Rechazar
      </Typography>
      </Box>
      <FormControl className={classes.formControl} disabled>
        <CssTextField
          className={classes.margin}
          placeholder="Vacaciones"
          variant="outlined"
          value={item.Typerequestname}
          // disabled
          InputProps={{
            style: { color: '#0367A6' },
            startAdornment: (
              <InputAdornment position="start">
                <Flag />
              </InputAdornment>
            ),
          }}
        />
      </FormControl>
      <FormControl className={classes.formControl} disabled>
        <CssTextField
          className={classes.margin}
          placeholder="Vacaciones"
          variant="outlined"
          value={item.CreateAt}
          // disabled
          InputProps={{
            style: { color: '#0367A6' },
            startAdornment: (
              <InputAdornment position="start">
                <Event />
              </InputAdornment>
            ),
          }}
        />
      </FormControl>
      <FormControl className={classes.formControl}>
        <ThemeProvider theme={textfiel}>
          <CssTextField
            multiline
            name="details"
            ref={register}
            rows={6}
            placeholder="Indique las razones de su decisión"
            variant="outlined"
            InputProps={{
              style: { color: '#0367A6' },
              startAdornment: (
                <InputAdornment position="start">
                  <Speaker_notes />
                </InputAdornment>
              ),
            }}
          />
        </ThemeProvider>
      </FormControl>
      <Box display="flex" justification="center" borderTop={5} color="#0367A6">
        <Grid
          container
          direction="column"
          justify="space-between"
          className={classes.grid} >
          <Box margin={1}>
            <Typography className={classes.pos} >
              Importante:
            </Typography>
            <Typography variant="body2" className={classes.pos}>
              Es requisito indicar un comentario si rechaza la solicitud.
            </Typography>
            <Typography variant="body2" className={classes.pos}>
              Este comentario se comparte con el usuario solicitante.
            </Typography>
          </Box>
        </Grid>
      </Box>
    </CardContent>
    <CardActions>
      <Grid
        display="flex"
        container
        direction="row"
        justify="space-around"
        className={classes.btndown}>
        <ThemeProvider theme={theme}>
          <Button 
            onClick={ () => { handleSubmit(onSubmit); setValue("status", false) } }
            variant="contained"
            color="secondary" >
            Rechazar
          </Button>
        </ThemeProvider>
        <ThemeProvider theme={theme}>
          <Button onClick={ () => { handleSubmit(onSubmit); setValue("status", true) }  } variant="contained" color="primary">
            Aprobar
          </Button>
        </ThemeProvider>
      </Grid>
    </CardActions>
  </form>
  )
}

export default ModalApprovedRequest