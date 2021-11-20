import React from 'react';
import { 
  CssBaseline,
  Paper,
  Stepper,
  Step,
  StepLabel,
  Button,
  Typography,
  makeStyles 
} from '@material-ui/core/';

import DataPersonal from './DataPersonal';
import EducationExperience from './EducationExperience';
import SkillsLanguages from './SkillsLanguages';

import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers';
import * as yup from 'yup';
import { useDispatch, useSelector } from 'react-redux';
import { FormChange } from '../../actions/persistenceActions';



const useStyles = makeStyles((theme) => ({
  layout: {
    width: 'auto',
    marginLeft: theme.spacing(2),
    marginRight: theme.spacing(2),
    [theme.breakpoints.up(600 + theme.spacing(2) * 2)]: {
      width: 800,
      marginLeft: 'auto',
      marginRight: 'auto',
    },
  },
  paper: {
    padding: theme.spacing(2),
    [theme.breakpoints.up(600 + theme.spacing(3) * 2)]: {
      padding: theme.spacing(4),
    },
  },
  stepper: {
    padding: theme.spacing(3, 0, 5),
  },
  buttons: {
    display: 'flex',
    justifyContent: 'flex-end',
  },
  button: {
    marginTop: theme.spacing(3),
    marginLeft: theme.spacing(1),
  },
}));

const steps = ['Datos', 'Educación & Experiencia', 'Habilidades & Lenguajes'];

export default function Checkout() {
  const classes = useStyles();
  const [activeStep, setActiveStep] = React.useState(0);

  const handleNext = () => {
    setActiveStep(activeStep + 1);
  };

  const handleBack = () => {
    setActiveStep(activeStep - 1);
  };

  
  const dispatch = useDispatch();
  const form = useSelector(state => state.persistence.form);
  
  const schema = yup.object().shape({
    user: yup.object({
      firstname: yup.string().trim().required(() => <span>First name is required</span>),
      lastname: yup.string().trim().required(() => <span>Last name is required</span>),
    }),
  });

  const { register, handleSubmit, errors } = useForm({
    resolver: yupResolver(schema),
    defaultValues: form,
  });

  const onSubmit = data => {
    // data.image = data.image[0];
    console.log(form);
    dispatch( FormChange(data) );
    handleNext();
  }

  const getStepContent = (step) => {
    switch (step) {
      case 0:
        return <DataPersonal register={ register } errors={ errors } />;
      case 1:
        return <EducationExperience />;
      case 2:
        return <SkillsLanguages />;
      default:
        throw new Error('Unknown step');
    }
  }

  return (
    <React.Fragment>
      {
        console.log(form)
      }
      <CssBaseline />
      <form onSubmit={ handleSubmit(onSubmit) } className={classes.layout}>
        <Paper className={classes.paper}>
          <Typography component="h1" variant="h4" align="center">
              Formulario de Registro de datos
          </Typography>
          <Stepper activeStep={activeStep} className={classes.stepper}>
            {steps.map((label) => (
              <Step key={label}>
                <StepLabel>{label}</StepLabel>
              </Step>
            ))}
          </Stepper>
          <React.Fragment>
            {activeStep === steps.length ? (
              <React.Fragment>
                <Typography variant="h5" gutterBottom>
                  Thank you for your order.
                </Typography>
                <Typography variant="subtitle1">
                  Your order number is #2001539. We have emailed your order confirmation, and will
                  send you an update when your order has shipped.
                </Typography>
              </React.Fragment>
            ) : (
              <React.Fragment>
                {getStepContent(activeStep)}
                
                <div className={classes.buttons}>
                  {activeStep !== 0 && (
                    <Button onClick={handleBack} className={classes.button}>
                      Back
                    </Button>
                  )}
                  <Button
                    variant="contained"
                    color="primary"
                    type="submit"
                    className={classes.button}
                  >
                    {activeStep === steps.length - 1 ? 'Place order' : 'Next'}
                  </Button>
                </div>
              </React.Fragment>
            )}
          </React.Fragment>
        </Paper>  
      </form>
    </React.Fragment>
  );
}