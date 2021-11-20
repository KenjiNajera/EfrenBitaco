import React from 'react';
import { 
  Card,
  CardContent,
  LinearProgress,
  Typography,
  Grid,
  FormControlLabel,
  Checkbox,
} from '@material-ui/core';

import { makeStyles,withStyles  } from '@material-ui/core/styles';
import Add from '@material-ui/icons/Add';
import IconButton from '@material-ui/core/IconButton';


const TechnicalsLenguagesCardContent = withStyles(theme => ({
  root: {
    padding: 10
  }
}))(CardContent);
const SoftCardContent = withStyles(theme => ({
  root: {
    padding: 0,
    marginLeft:10,
  }
}))(CardContent);
const BorderLinearProgress = withStyles((theme) => ({
  root: {
    height: 4,
    borderRadius: 5,
    width:200,
    display: 'inline-block',
  },
  colorPrimary: {
    backgroundColor: theme.palette.grey[theme.palette.type === 'light' ? 200 : 700],
  },
  bar: {
    borderRadius: 5,
    backgroundColor: '#1a90ff',
  },
}))(LinearProgress);

const useStyles = makeStyles({
  root: {
    width: 350,
    height: 200,
    overflowY: 'scroll',
    padding: 12,
  },
  title: {
    display: 'inline-block',
    margin: '0 6px',
    fontSize:10,
  },
  titlecard: {
    margin: '0 6px',
    fontSize:20,
    display: 'inline-block',
  },
  cardp:{
    width: 310,
    height:40,
    margin: '0 10px',
    marginTop:10,
    paddingBottom:-100,
  },
  cardsoft:{
    width: 310,
    height:40,
    margin: '0 10px',
    marginTop:10,
  },
  icon:{
    marginLeft:180,
  },
});

const SkillsLanguages= ({hard,lang,sof}) => {
  const classes = useStyles();
  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        Habilidades & Lenguajes
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} sm={6}>
          <Card className={classes.root}>
            <Typography className={classes.titlecard}>
            Technicals
            </Typography>
            <IconButton className={classes.icon} color="primary" aria-label="add to Experience" size="small">
              <Add/>
            </IconButton>
              <Card className={classes.cardp}>
                <TechnicalsLenguagesCardContent>
                  <Typography className={classes.title} color="textSecondary" gutterBottom>
                  C#
                  </Typography>
                  <BorderLinearProgress variant="determinate" value={80} />
                  <Typography className={classes.title} color="textSecondary" >
                  10%
                  </Typography>
                </TechnicalsLenguagesCardContent>
              </Card>
          </Card>
        </Grid>
        <Grid item xs={12} sm={6} >
          <Card className={classes.root}>
            <Typography className={classes.titlecard}>
            Language
            </Typography>
            <IconButton className={classes.icon} color="primary" aria-label="add to Experience" size="small">
              <Add/>
            </IconButton>
              <Card className={classes.cardp}>
                <TechnicalsLenguagesCardContent>
                <Typography className={classes.title} color="textSecondary" gutterBottom>
                  C#
                  </Typography>
                  <BorderLinearProgress variant="determinate" value={80} />
                  <Typography className={classes.title} color="textSecondary" >
                  10%
                  </Typography>
                </TechnicalsLenguagesCardContent>
              </Card>
          </Card>
            </Grid>
            <Grid item xs={12} sm={6} >
          <Card className={classes.root}>
            <Typography className={classes.titlecard}>
              Softskill
            </Typography>
            <IconButton className={classes.icon} color="primary" aria-label="add to Experience" size="small">
              <Add/>
            </IconButton>
              <Card className={classes.cardsoft}>
                <SoftCardContent>
                <FormControlLabel control={<Checkbox name="checkedC" />} label="Uncontrolled" />
                </SoftCardContent>
              </Card>
          </Card>
            </Grid>
      </Grid>
    </React.Fragment>
    
  );
  
}
export default SkillsLanguages;
