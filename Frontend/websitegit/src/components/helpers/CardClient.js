import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import { CardActionArea, IconButton } from '@material-ui/core';
import buttondelete from '../../assets/icons/menos.svg';
import { useDispatch } from 'react-redux';
import { Deleting } from '../../actions/persistenceActions';

const useStyles = makeStyles((theme) =>({
    cardContainer:{
        position: 'relative',
        marginLeft:theme.spacing(3),
        marginRight: theme.spacing(3),
        marginBottom: theme.spacing(3),
        marginTop: theme.spacing(3),
    },
    buttonDelete:{
        position: 'absolute',
        top: -22,
        right:-22,
        [theme.breakpoints.up(2800)]:{
            top: -25,
            right:-25,
        },
    },
    imgDelete: {
        width: theme.spacing(3),
        [theme.breakpoints.up(2800)]:{
            width: theme.spacing(5),
        },
    },
    media: {
        height: 200,
        padding: theme.spacing(1),
        objectFit: 'contain',
    },
    name:{
        fontSize:'1vw',
    },
}));
const CardClient = ({ item, isDelete }) => {
    const classes = useStyles();

    const dispatch = useDispatch();

    const handleDelete = ( id ) => (e) => {
        e.preventDefault();
        dispatch( Deleting( `client/delete/${id}`, id ) )
    }
    return (
        <div className={classes.cardContainer}>
            <CardActionArea>
            <Card>
                {
                    (isDelete)
                        ?   <IconButton size="medium" className={ classes.buttonDelete } onClick={ handleDelete(item.id) }>
                                <img src={ buttondelete }  className={ classes.imgDelete } alt="boton delete" />
                            </IconButton> 
                        :   null  
                }
                
                <CardMedia
                    component="img"
                    className={classes.media}
                    image={ item.Image }
                    title={ item.Clientname }
                />
                <CardContent>
                    <Typography  className={classes.name}>
                    { item.Clientname }
                    </Typography>
                </CardContent>
            </Card>  
            </CardActionArea>          
        </div>
    );
}

export default CardClient