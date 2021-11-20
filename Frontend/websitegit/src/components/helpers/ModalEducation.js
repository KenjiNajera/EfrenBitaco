import React, { useState } from 'react';
import { 
    makeStyles, 
    FormControl, 
    InputLabel,  
    withStyles, 
    InputBase,
    Select,
    MenuItem,
    Divider,
    Button, } from '@material-ui/core';
import AddIcon from '@material-ui/icons/Add';

const ModalEducation = ({ handleModalClose }) => {
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
            '& #institution': {
                marginTop: theme.spacing(2),
            },
            '& #year': {
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
        },
    }));
    const classes = useStyles();
    const [institution, setInstitution] = useState('');
    const [year, setYear] = useState('');

    const handleChange = ( e ) => {
        setInstitution(e.target.value);
    }
    const handleChangeYear = ( e ) => {
        setYear(e.target.value);
    }

    return (
        <div className={ classes.content }>
            <div className={ classes.header }>
                <p>Education</p>
            </div>            
            <Divider />
            <div className={ classes.body }>
                <div id="institution">
                    <FormControl className={ classes.margin }>
                        <InputLabel shrink htmlFor="SelectInstitution">Institution</InputLabel>
                        <Select
                            labelId="SelectInstitution"
                            id="SelectInstitution"
                            value={ institution }
                            onChange={ handleChange }
                            input={ <BootstrapInput /> }
                            >
                            <MenuItem value="">
                                <em>None</em>
                            </MenuItem>
                            <MenuItem value={10}>Ten</MenuItem>
                        </Select>
                    </FormControl>
                </div>
                <div>
                    <FormControl className={classes.margin}>
                        <InputLabel shrink htmlFor="AcademicTitleInput">Academic Title</InputLabel>
                        <BootstrapInput placeholder="Academic Title" />
                    </FormControl>
                </div>
                <div id="year"> 
                    <FormControl className={ classes.margin }>
                        <InputLabel shrink htmlFor="SelectYear">year</InputLabel>
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
                <div>
                    <Button id="ButtonAdd" 
                        variant="contained"
                        // color="primary"
                        startIcon={ <AddIcon /> }>
                        Add
                    </Button>
                </div>
            </div>
        </div>
    )
}


export default ModalEducation;
