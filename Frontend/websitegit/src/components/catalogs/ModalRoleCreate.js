import React, { useState } from 'react';
import {
    Button,
    Divider,
    FormControl,
    FormHelperText,
    InputBase,
    makeStyles,
    withStyles
} from '@material-ui/core'
import AddIcon from '@material-ui/icons/Add';
import { useForm } from 'react-hook-form';
import { ErrorMessage } from '@hookform/error-message';
import { yupResolver } from '@hookform/resolvers';
import * as yup from 'yup';
import { useDispatch } from 'react-redux';
import { Creating } from '../../actions/persistenceActions';
import Grid from '@material-ui/core/Grid';
import FormLabel from '@material-ui/core/FormLabel';
import ListRole from '../helpers/ListRole';
import Add from '@material-ui/icons/Add';
import IconButton from '@material-ui/core/IconButton';
import ListModuleRolAndPermission from '../helpers/ListModuleRolAndPermission';

const ModalRoleCreate = () => {

    const BootstrapInput = withStyles((theme) => ({
        root: {
            'label + &': {
                marginTop: theme.spacing(1),
            },
        },
        input: {
            borderRadius: 4,
            backgroundColor: theme.palette.background.paper,
            border: '1px solid #ced4da',
            fontSize: 16,
            width: 220,
            textAlign: 'left',
            padding: '10px 0px 10px 12px',
            transition: theme.transitions.create(['border-color', 'box-shadow']),
            fontFamily: [
                '-apple-system',
                'BlinkMacSystemFont',
                '"Segoe UI"',
                'Roboto',
                '"Helvetica Neue"',
                'Arial',
                'sans-serif',
                '"Apple Color Emoji"',
                '"Segoe UI Emoji"',
                '"Segoe UI Symbol"',
            ].join(','),
            '&:focus': {
                borderRadius: 4,
                borderColor: '#80bdff',
                boxShadow: '0 0 0 0.2rem rgba(0,123,255,.25)',
            },
        },
    }))(InputBase);

    const useStyles = makeStyles((theme) => ({
        content: {
            padding: theme.spacing(2),
            width: 600,
            backgroundColor: theme.palette.background.paper,
            borderRadius: 5,
        },
        header: {
            display: 'flex',
            justifyContent: 'center',
            '& p': {
                fontSize: 24,
                marginBottom: theme.spacing(1),
            },
        },
        body: {
            '& #object': {
                marginTop: theme.spacing(2),
                marginBottom: theme.spacing(4),
            },
            '& div': {
                display: 'flex',
                justifyContent: 'center',
                '& #ButtonAdd': {
                    marginBottom: theme.spacing(1),
                    color: 'white',
                    backgroundColor: '#0367A6',
                    textTransform: 'none',
                },
            },

        },
        error: {
            '& span': {
                color: '#bf1650',
                '&::before': {
                    display: 'inline',
                    content: '"⚠ "',
                },
            },
        },
        margin: {
            margin: theme.spacing(2),
        },
        List: {
            margin: theme.spacing(1),
            width: 'auto',
        },
        subtittle3: {
            padding: '21px 0px 10px 5px',
            marginRight: 35,
        },
        subtittle2: {
            padding: '21px 10px 10px 5px',
            marginRight: 10,
        },
        subtittle: {
            margin: 14,
            marginLeft: 45,
        },
        icon: {
            marginRight: 20,
        },
    }));
    const classes = useStyles();

    const schema = yup.object().shape({
        Softskillname: yup.string().trim()
            .required(() => <span>Name is required</span>)
    });
    const { register, handleSubmit, errors } = useForm({
        resolver: yupResolver(schema),
    });

    const onSubmit = data => {
        dispatch(Creating("PonerWeasDeRolAqui/", data));
    }

    const [state, setstate] = useState(false)
    const dispatch = useDispatch();
    const handleShow = () => {
        setstate(!state);
    }

    return (
        <div className={classes.content}>
            <div className={classes.header}>
                <p>Create New Role</p>
            </div>
            <Divider />
            <form onSubmit={handleSubmit(onSubmit)} className={classes.body}>
                <div id="object">
                    <Grid container>
                        <Grid content direction="row" alignItems="flex-start">
                            <FormControl className={classes.margin}>
                                <FormLabel className={classes.subtittle} component="legend">Role Name</FormLabel>
                                <BootstrapInput
                                    name="RolName"
                                    inputRef={register}
                                    placeholder="Rol Name"
                                    className={classes.subtittle} />
                                <FormHelperText className={classes.error}>
                                    <ErrorMessage errors={errors} name="RolName" />
                                </FormHelperText>
                                <Grid container
                                    direction="row"
                                    justify="center"
                                    alignItems="flex-start">
                                    <Grid container justify="flex-start">
                                        <ListRole />
                                    </Grid>
                                </Grid>
                            </FormControl>
                            <Grid container
                                direction="column"
                                alignItems="flex-start">
                                <Grid item direction="column" className={classes.List}>
                                    <Grid item direction="row">
                                        <Grid item direction="column">
                                            <Grid item direction="row">
                                                <FormLabel component="legend" className={classes.subtittle2}>Module</FormLabel>
                                                <IconButton className={classes.icon} color="primary" aria-label="add to permissions" size="medium" onClick={handleShow}>
                                                    <Add />
                                                </IconButton>
                                            </Grid>
                                        </Grid>
                                        <Grid item direction="column">
                                            <FormLabel component="legend" className={classes.subtittle3}>Permissions</FormLabel>
                                        </Grid>
                                    </Grid>
                                    {
                                        (state)
                                            ? <ListModuleRolAndPermission />
                                            : null
                                    }
                                </Grid>
                            </Grid>
                        </Grid>
                    </Grid>
                </div>
                <div>
                    <Button
                        id="ButtonAdd"
                        variant="contained"
                        type="submit"
                        startIcon={<AddIcon />}>
                        Add role
                    </Button>
                </div>
            </form>
        </div>
    )
}
export default ModalRoleCreate