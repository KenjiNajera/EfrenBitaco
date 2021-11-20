import React, { useEffect, useState } from 'react';
import { 
    Checkbox,
    FormControlLabel,
    Grid,
    makeStyles, 
    Paper, 
    Typography
} from '@material-ui/core';
// import Axios from 'axios';


const useStyles = makeStyles((theme) => ({
    container:{
        width: 230,
    },
    textFilter: {
        textAlign:'center',
    },
    card: {
        marginBottom: theme.spacing(3),
    },
    textCard: {
        backgroundColor: '#EBEBEB',
        textAlign: 'center',
        color: '#434E65',
        borderTopRightRadius: 4,
        borderTopLeftRadius: 4,
        marginTop: 10,
    },
    containerOptions:{
        // backgroundColor: 'green',
        height: 200,
        justifyContent:'center',
        paddingLeft:theme.spacing(2),
        overflow: 'scroll',
        overflowX: 'hidden',
    },
    paper: {
        padding: theme.spacing(2),
        marginRight: theme.spacing(1),
        [theme.breakpoints.down('sm')]:{
            display: 'none',
        },
    },
}));

const Filterbar = () => {
    const classes = useStyles();

    // const [state, setstate] = useState([])
    // useEffect(() => {
        
    // }, [])


    const array = [1]
    const arrays = [1]
    return (
        <div className={ classes.container }>
            <Paper className={ classes.paper }>
                <Grid item xs className={ classes.textFilter } >
                    <Typography variant="h4" gutterBottom>
                        Filter
                    </Typography>
                </Grid>
                {
                    arrays.map(card =>(
                        <div className={ classes.card}>
                            <Grid className={ classes.textCard }>
                                <Typography variant="subtitle1" >
                                    HardSkills
                                </Typography>
                            </Grid>
                            <Grid className={ classes.containerOptions }>
                                {
                                    array.map( key => (
                                        <FormControlLabel control={<Checkbox name="checkedC" />} label="Uncontrolled" />
                                    ))
                                }
                            </Grid>
                        </div>
                    ))
                }
            </Paper>
        </div>
    )
}

export default Filterbar;

