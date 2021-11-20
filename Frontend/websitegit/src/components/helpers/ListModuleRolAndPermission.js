import React from 'react';
import { makeStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';
import ListModuleRol from './ListModule';
import ListModulePermissions from '../helpers/ListModulePermissions';
import Button from '@material-ui/core/Button';
import AddIcon from '@material-ui/icons/Add';

const useStyles = makeStyles((theme) => ({
    root: {
        width: "100%",
        maxWidth: 360,
        padding: 20,
        borderRadius: 5,
        padding: theme.spacing(1),
    },
    subtittle: {
        margin: theme.spacing(1),
    },
    subtittle2: {
        padding: '0px 5px 10px 45px',
    },
    icon: {
        marginLeft: 5,
    },
}));

const ListModuleRolAndPermission = () => {
    const classes = useStyles();

    return (
        <div className={classes.root}>
            <from>
                <Grid Container >
                    <Grid item direction="row">
                        <Grid item direction="column">
                            <ListModuleRol />
                        </Grid>
                        <Grid item direction="column">
                            <ListModulePermissions />
                        </Grid>
                    </Grid>
                </Grid>
                <div>
                    <Grid item direction="column">
                        <Button
                            id="ButtonAdd"
                            variant="contained"
                            type="submit"
                            startIcon={<AddIcon />}
                            className={classes.btn2}>
                            Select
                         </Button>
                    </Grid>
                </div>
            </from>
        </div>
    )
}
export default ListModuleRolAndPermission