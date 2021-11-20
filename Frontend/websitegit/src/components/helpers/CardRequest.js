import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import { CardActionArea } from '@material-ui/core';
import Box from '@material-ui/core/Box';

const useStyles = makeStyles((theme) => ({
    cardContainer: {
        position: 'relative',
        marginLeft: theme.spacing(3),
        marginRight: theme.spacing(3),
        marginBottom: theme.spacing(3),
        marginTop: theme.spacing(3),
    },
    media: {
        height: 200,
        padding: theme.spacing(1),

    },
    name: {
        fontSize: '.9vw',
        marginTop: 20,
        marginLeft: 4,
    },
    aboutimage: {
        display: 'flex',
        marginLeft: '10px',
        marginTop: 6,
        width: '10vh',
        height: '10vh',
        borderRadius: '100%',
        objectFit: 'contain',
    },
    card: {
        display: 'flex',
    },
    description: {
        fontSize: '1vw',
    },
    typerequest: {
        fontSize: '.9vw',
        marginTop: 8,
        marginLeft: 4,
    },
}));

const CardRequest = ({ item, handleModalopen }) => {
    const classes = useStyles();
    return (
        <div className={classes.cardContainer} onClick={ handleModalopen }>
            <CardActionArea>
                <Card>
                    <div className={classes.card}>
                        <CardMedia className={classes.aboutimage} component="img" image={item.Photouser} />
                        <div>
                            <Typography className={classes.name}>
                                {item.Firstname + " " + item.Lastname}
                            </Typography>
                            <Typography className={classes.typerequest}>
                                Solicitud por:  { item.Typerequestname }
                            </Typography>
                        </div>
                    </div>
                    <div>
                        <Box display="" justification="center" borderBottom={1}></Box>
                        <CardContent >
                            <Typography className={classes.description}>
                                Descripcion:
                            </Typography>
                            <Typography className={classes.description}>
                                {item.Description}
                            </Typography>
                        </CardContent>
                    </div>
                </Card>
            </CardActionArea>
        </div>
    );
}

export default CardRequest