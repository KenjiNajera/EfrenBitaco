import React from 'react';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import { Card, CardContent } from '@material-ui/core';
import { red } from '@material-ui/core/colors';
import { makeStyles } from '@material-ui/core/styles';
import Add from '@material-ui/icons/Add';
import IconButton from '@material-ui/core/IconButton';


// npm add react-native-web
const useStyles = makeStyles({
  root: {
    width: 350,
    backgroundColor: red,
    height: 310,
    overflowY: 'scroll',
  },
  title: {
    display: 'inline-block',
    margin: '0 6px',
    fontSize:20,
  },
  titlecard: {
    margin: '0 6px',
    fontSize:20,
    display: 'inline-block',
  },
  school: {
    margin: '0 6px',
    fontSize:13,
  },
  cardp:{
    width: 320,
    height:78,
    margin: '0 10px',
    marginTop:10,
  },
  icon:{
    marginLeft:180,
  },
});

const EducationExperience = ({cert,educ,expe}) => {
  const classes = useStyles();
  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        Educación & Experiencia
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} sm={6}>
          <Card className={classes.root}>
            <Typography className={classes.titlecard}>
              Educación
            </Typography>
            <IconButton className={classes.icon} color="primary" aria-label="add to Experience" size="small">
              <Add/>
            </IconButton>
              <Card className={classes.cardp}>
                <CardContent>
                  <Typography className={classes.title} color="textSecondary" gutterBottom>
                    Universidad UTM
                  </Typography>
                  <Typography className={classes.title} color="textSecondary" >
                    2022
                  </Typography>
                  <Typography className={classes.school} color="textSecondary" >
                  Ingenieria en tecnologías de la informacion
                  </Typography>
                </CardContent>
              </Card>
              <Card className={classes.cardp}>
                <CardContent>
                  <Typography className={classes.title} color="textSecondary" gutterBottom>
                    Universidad UTM
                  </Typography>
                  <Typography className={classes.title} color="textSecondary" >
                    2022
                  </Typography>
                  <Typography className={classes.school} color="textSecondary" >
                  Ingenieria en tecnologías de la informacion
                  </Typography>
                </CardContent>
              </Card> <Card className={classes.cardp}>
                <CardContent>
                  <Typography className={classes.title} color="textSecondary" gutterBottom>
                    Universidad UTM
                  </Typography>
                  <Typography className={classes.title} color="textSecondary" >
                    2022
                  </Typography>
                  <Typography className={classes.school} color="textSecondary" >
                  Ingenieria en tecnologías de la informacion
                  </Typography>
                </CardContent>
              </Card> 
          </Card>
        </Grid>
        <Grid item xs={12} sm={6}>
          <Card className={classes.root}>
            <Typography className={classes.titlecard}>
              Experience
            </Typography>
            <IconButton className={classes.icon} color="primary" aria-label="add to Experience" size="small">
              <Add/>
            </IconButton>
              <Card className={classes.cardp}>
                <CardContent>
                  <Typography className={classes.title} color="textSecondary" gutterBottom>
                    Universidad UTM
                  </Typography>
                  <Typography className={classes.title} color="textSecondary" >
                    2022
                  </Typography>
                  <Typography className={classes.school} color="textSecondary" >
                  Ingenieria en tecnologías de la informacion
                  </Typography>
                </CardContent>
              </Card>
              <Card className={classes.cardp}>
                <CardContent>
                  <Typography className={classes.title} color="textSecondary" gutterBottom>
                    Universidad UTM
                  </Typography>
                  <Typography className={classes.title} color="textSecondary" >
                    2022
                  </Typography>
                  <Typography className={classes.school} color="textSecondary" >
                  Ingenieria en tecnologías de la informacion
                  </Typography>
                </CardContent>
              </Card> <Card className={classes.cardp}>
                <CardContent>
                  <Typography className={classes.title} color="textSecondary" gutterBottom>
                    Universidad UTM
                  </Typography>
                  <Typography className={classes.title} color="textSecondary" >
                    2022
                  </Typography>
                  <Typography className={classes.school} color="textSecondary" >
                  Ingenieria en tecnologías de la informacion
                  </Typography>
                </CardContent>
              </Card> 
          </Card>
        </Grid>
      </Grid>
    </React.Fragment>
  );
}
export default EducationExperience;