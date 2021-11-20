import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import CardActionArea from '@material-ui/core/CardActionArea';

const useStyles = makeStyles((theme) => ({
    status: {
        display: "flex",
        height: 'auto',
    },
    cardContainer: {
        position: 'relative',
        marginLeft: theme.spacing(3),
        marginRight: theme.spacing(0),
        marginBottom: theme.spacing(0),
        marginTop: theme.spacing(3),
    },
    card: {
        display: "flex",
        flexDirection: "column"
    },
    active: {
        width: 30,
        backgroundColor: '#28A745',
    },
    desactive: {
        width: 30,
        backgroundColor: 'red',
    },
    controls: {
        display: "flex",
    },
    hards: {
        marginTop: 6,
        width: '6vh',
        height: '6vh',
        borderRadius: '100%',
        marginLeft: 10,
        objectFit: 'contain',
    },
    nameproject: {
        fontSize: '1.2vw',
        paddingLeft: 8,

        marginTop: 20,
    },
    subtitle: {
        fontSize: '2vh',
    },
    aboutimage: {
        marginLeft: '10px',
        marginTop: 6,
        width: '8vh',
        height: '8vh',
        borderRadius: '100%',
        objectFit: 'contain',
    },
    paddinContent:{
       padding:9,
       paddingTop:2,
    }
}));

const CardTeam = ({ item ,StatusProject}) => {
    const classes = useStyles();

    return (
        <div className={classes.cardContainer}>

            <Card className={classes.status}>
                <CardActionArea>
                    <div className={classes.card}>
                        <div className={classes.controls}>
                            <CardMedia className={classes.aboutimage} component="img" image='http://localhost:8080/helper/images/203-9322409668.jpg' />
                            <Typography className={classes.nameproject}>
                                {item.Projectname}
                            </Typography>
                        </div>
                    </div>
                    <CardContent className={classes.paddinContent}>
                        <Typography className={classes.subtitle}>
                            Technologies
                        </Typography>
                        <div className={classes.controls}>
                            {
                                item.Hardskills.map((item, index) => (
                                    <CardMedia className={classes.hards} component="img" image={item.Image} key={item.id}/>
                                  
                                ))
                            }
                        </div>
                        <Typography className={classes.subtitle} >
                            Teams:
                            </Typography>
                        <div className={classes.controls}>
                            {
                                item.Resources.map((item, index) => (
                                    <CardMedia className={classes.hards} component="img" image={item.Photouser} key={item.Userid}/>
                                    
                                ))
                            }
                        </div>
                    </CardContent>
                </CardActionArea>
                {
                    (item.Status)
                        ? <CardMedia
                            className={classes.active}       
                          />
                          
                        : <CardMedia
                            className={classes.desactive}    
                          />
                }
            </Card>
        </div>
    );
}
export default CardTeam