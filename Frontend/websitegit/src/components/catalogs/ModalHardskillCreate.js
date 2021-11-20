import React, { useState } from 'react';
import { 
    Button, 
    Divider, 
    FormControl, 
    FormHelperText, 
    IconButton, 
    InputBase, 
    InputLabel, 
    makeStyles,
    Paper,
    withStyles
} from '@material-ui/core'
import AddIcon from '@material-ui/icons/Add';
import PhotoCamera from '@material-ui/icons/PhotoCamera';

import { useDispatch } from 'react-redux';
import { CreatingWithImage } from '../../actions/persistenceActions';

import { useForm } from 'react-hook-form';
import { ErrorMessage } from '@hookform/error-message';
import { yupResolver } from '@hookform/resolvers';
import * as yup from 'yup';
import imagen from '../../assets/images/bepensa.png';

const ModalHardskillCreate = ({ handleModalClose }) => {

    const dispatch = useDispatch();

    const BootstrapInput = withStyles((theme) => ({
        root: {
            'label + &': {
                marginTop: theme.spacing(3),
            },
        },
        input: {
            borderRadius: 4,
            // position: 'relative',
            backgroundColor: theme.palette.background.paper,
            border: '1px solid #ced4da',
            fontSize: 16,
            width:200,
            textAlign:'left',
            padding: '10px 26px 10px 12px',
            transition: theme.transitions.create(['border-color', 'box-shadow']),
            // Use the system font instead of the default Roboto font.
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
            width: 330,
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
                display: 'block',
                marginTop: theme.spacing(2),
                marginBottom: theme.spacing(4),
            },
            '& div': {
                display:'flex',
                justifyContent:'center',
                '& #ButtonAdd': {
                    marginBottom: theme.spacing(1),
                    color:'white',
                    backgroundColor: '#0367A6',
                    textTransform: 'none',
                },
            },
            
        },
        input: {
            display: 'none',
        },
        superposition: {
            position: 'absolute',
            bottom:-2,
            right: -2,
            '& #buttonPhoto': {
                color: '#000000',
            },
        },
        paper: {
            width: 85, 
            height:85,
            borderRadius: '100%',
            position: 'relative', 
            '& img':{ 
                width: '100%',
                height: '100%',
                margin: 'auto',
                borderRadius: '100%',
            },
        },
        error: {
            '& span':{
                color: '#bf1650',
                '&::before':{
                    display: 'inline',
                    content: '"⚠ "',
                },
            },
        },
        margin: {
            margin: theme.spacing(1),
        },
    }));
    const classes = useStyles();

    const schema = yup.object().shape({
        Hardskillname: yup.string().trim()
            .required( () => <span>Name is required</span> ),
    });
    const { register, handleSubmit, errors } = useForm({
        resolver: yupResolver(schema),
    });

    const onSubmit = data => {
        var formdata = new FormData();
        formdata.append("Image", data.Image[0]);
        formdata.append("Hardskillname", data.Hardskillname);
        dispatch( CreatingWithImage( "hardskill/", data, formdata ) );
    }
    
    const [file, setFile] = useState({
        file: imagen,
    });

    const handleChange = ( event ) => {
        event.preventDefault();
        setFile({
            file: URL.createObjectURL( event.target.files[0] ),
        })
    }

    return (
        <div className={ classes.content }>
            <div className={ classes.header }>
                <p>Harskill Form</p>
            </div>            
            <Divider />
            <form onSubmit={ handleSubmit( onSubmit ) } className={ classes.body }>
                <div id="object">
                    <div>
                        <Paper  className={ classes.paper } >
                            <img src={file.file} alt="Hardskill"/>
                            <input 
                                name="Image"
                                ref={ register }
                                required
                                accept="image/*" 
                                className={ classes.input } 
                                onChange={ handleChange }
                                id="icon-button-file" 
                                type="file" />
                            <label htmlFor="icon-button-file" className={ classes.superposition }>
                                <IconButton id="buttonPhoto" size="small" aria-label="upload picture" component="span">
                                    <PhotoCamera fontSize="small" />
                                </IconButton>
                            </label>
                        </Paper>
                    </div>
                    <div>
                        <FormControl className={ classes.margin }>
                            <InputLabel shrink htmlFor="Hardskillname">Hardskill Name</InputLabel>
                            <BootstrapInput 
                                name="Hardskillname"
                                inputRef={register}
                                placeholder="Hardskill Name" />
                            <FormHelperText className={ classes.error }>
                                <ErrorMessage errors={errors} name="Hardskillname" />
                            </FormHelperText>
                        </FormControl>
                    </div>
                </div>
                <div>
                    <Button 
                        id="ButtonAdd" 
                        variant="contained"
                        type="submit"
                        startIcon={ <AddIcon /> }>
                        Add
                    </Button>
                </div>
            </form>
        </div>
    )
}

export default ModalHardskillCreate
