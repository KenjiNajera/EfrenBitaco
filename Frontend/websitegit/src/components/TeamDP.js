import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import ButtonBase from '@material-ui/core/ButtonBase';
import gitlogo from '../assets/images/ic_git_logo_icon.svg';
import CardEmployee from './helpers/CardEmployee';
import ListasCustom from './helpers/ListasCustom';
import GridList from '@material-ui/core/GridList';
import GridListTile from '@material-ui/core/GridListTile';
import GridListTileBar from '@material-ui/core/GridListTileBar';
import IconButton from '@material-ui/core/IconButton';
import StarBorderIcon from '@material-ui/icons/StarBorder';
import CardClient from '../components/helpers/CardClient';

import logob from '.././assets/images/bepensa.png';
import buttondelete from '../assets/icons/menos.svg';
import buttonadd from '../assets/icons/mas.svg';

import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import Slide from '@material-ui/core/Slide';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    width: 'auto',
  },
  paper: {
    padding: theme.spacing(2),
    margin: 'auto',
    maxWidth: 'auto',
  },
  image: {
    width: 128,
    height: 128,
  },
  imagelg: {
    width: 60,
    height: 60,
  },
  containerIconButton: {
    width: 32,
    height: 32,
    margin: 10,
  },
  iconbutton: {
    width: 32,
    height: 32,
  },
  img: {
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
  },
  img2: {
    display: 'block',
    maxWidth: '35%',
    maxHeight: '100%',
  },
  btn: {
    margin: 15,
    backgroundColor: '#0367A6',
  },
  gridList: {
    flexWrap: 'nowrap',
    transform: 'translateZ(0)',
    maxWidth: '91%',
  },
  icons: {
    margin: 11,
    maxWidth: '20%'
  },
}));

const Transition = React.forwardRef(function Transition(props, ref){
  return <Slide direction ="up" ref={ref} {...props} />;
});

export default function ModalTeamDP() {
  const [open, setOpen] = React.useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Paper className={classes.paper}>
        <Grid container  >
          <Box display="flex" justification="center" borderBottom={1}>
            <Grid item>
              <ButtonBase className={classes.image} disabled>
                <img className={classes.img} alt="complex" src={gitlogo} />
              </ButtonBase>
            </Grid>
            <Grid item xs={12} sm container>
              <Grid item xs container direction="column" spacing={0}>
                <Grid item
                  container
                  direction="column"
                  xs
                  display="block" >
                  <Typography gutterBottom variant="h3">
                    Resource Manger
                </Typography>
                  <Grid container direction="row" xs spacing={4} >
                    <Grid item xs={3} >
                      <Typography variant="h4" gutterBottom>
                        Description:
                      </Typography>
                    </Grid>
                    <Grid item xs={9} className={classes.prueba}>
                      <Typography variant="h5" color="textSecondary">
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit,
                        sed do eiusmod tempor incididunt ut labore et
                        dolore magna aliqua. Ut enim ad minim veniam,
                     </Typography>
                    </Grid>
                    <Grid item xs={3} >
                      <Typography variant="h4" gutterBottom>
                        Technologies:
                </Typography>
                    </Grid>
                    <Grid container item xs={9}>
                      <Box>
                        <ButtonBase className={classes.imagelg} disabled>
                          <img className={classes.img} alt="complex" src={logob} />
                        </ButtonBase>
                      </Box>
                    </Grid>
                  </Grid>
                </Grid>
                <Grid item
                  container
                  direction="row"
                  justify="flex-end"
                  alignItems="flex-end">
                  <Button className={classes.btn}
                    onClick={handleClickOpen}>
                    Conclude Project
                 </Button>
                </Grid>
              </Grid>
              <Grid item>
                <IconButton className={classes.containerIconButton} onClick={handleClose}>
                  <img className={classes.iconbutton} alt="complex" src={buttondelete} />
                </IconButton>
              </Grid>
            </Grid>
          </Box>
          <Grid
            item
            container
            direction="row"
            xs={12}
            display="block"
          >
            <GridList className={classes.gridList} cols={2.5} xs={11}>
              <CardEmployee/>
              <CardEmployee/>
              <CardEmployee/>
            </GridList>
            <Grid item container xs={1} ClassName={classes.icons}>
              <Dialog
                open={open}
                TransitionComponent={Transition}
                keepMounted
                onClose={handleClose}
                aria-labelledby="alert-dialog-slide-title"
                aria-describedby="alert-dialog-slide-description"
              >
                <DialogTitle id="alert-dialog-slide-title">{"Are you sure about this?"}</DialogTitle>
                <DialogContent>
                  <DialogContentText id="alert-dialog-slide-description">
                    Example
                 </DialogContentText>
                </DialogContent>
                <DialogActions>
                  <Button onClick={handleClose} color="primary">
                    Disagree
                 </Button>
                  <Button onClick={handleClose} color="primary">
                    Agree
                  </Button>
                </DialogActions>
              </Dialog>
              <IconButton className={classes.containerIconButton}>
                <img className={classes.iconbutton} alt="complex" src={buttonadd} />
              </IconButton>
              <IconButton className={classes.containerIconButton}>
                <img className={classes.iconbutton} alt="complex" src={buttondelete} />
              </IconButton>
            </Grid>
          </Grid>
        </Grid>
      </Paper>
    </div>
  );
}
