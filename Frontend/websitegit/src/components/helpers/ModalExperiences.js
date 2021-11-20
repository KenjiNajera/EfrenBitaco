import React, { useState } from 'react'
import {
    withStyles, 
    InputBase,
    makeStyles,
    FormControl,
    InputLabel,
    Select,
    Divider,
    MenuItem,
    Button,
    TextField,} from '@material-ui/core';
import AddIcon from '@material-ui/icons/Add';

const ModalExperiences = ({ handleModalClose }) => {
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
            width:400,
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

            '& #project': {
                marginTop: theme.spacing(2),
            },
            '& #description': {
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
        margin: {
            margin: theme.spacing(1),
            width: 400,
        },
    }));
    const classes = useStyles();

    const CHARACTER_LIMIT = 200;
    const [year, setYear] = useState('');  
    const [values, setValues] = useState('');
    
    const handleChangeYear = ( e ) => {
        setYear(e.target.value);
    }

    const handleChangeDescription = (event) => {
        event.preventDefault();
        setValues(event.target.value);
    };

    return (
        <div className={ classes.content }>
            <div className={ classes.header }>
                <p>Experiences</p>
            </div>            
            <Divider />
            <div className={ classes.body }>
                <div id="project">
                    <FormControl className={classes.margin}>
                        <InputLabel shrink htmlFor="NameProjectInput">Name Project</InputLabel>
                        <BootstrapInput placeholder="Name Project" />
                    </FormControl>
                </div>
                <div>
                    <FormControl className={ classes.margin }>
                        <InputLabel shrink htmlFor="SelectYear">Year</InputLabel>
                        <Select
                            labelId="SelectYear"
                            id="SelectYear"
                            value={ year }
                            onChange={ handleChangeYear }
                            input={ <BootstrapInput /> }
                            >
                            <MenuItem value="">
                                <em>None</em>
                            </MenuItem>
                            <MenuItem value={10}>Ten</MenuItem>
                        </Select>
                    </FormControl>
                </div>
                <div id="description">
                    <FormControl className={classes.margin}>
                        <InputLabel shrink htmlFor="DescriptionInput">Description</InputLabel>
                        <TextField
                            fullWidth
                            multiline
                            rows={4}
                            inputProps={{
                                maxLength: CHARACTER_LIMIT
                            }}
                            value={values}
                            helperText={   `    ${values.length}/${CHARACTER_LIMIT}`}
                            onChange={handleChangeDescription}
                            margin="normal"
                            />
                    </FormControl>
                </div> 
                <div>
                    <Button id="ButtonAdd" 
                        variant="contained"
                        startIcon={ <AddIcon /> }>
                        Add
                    </Button>
                </div>
            </div>
        </div>
    )
}

export default ModalExperiences
