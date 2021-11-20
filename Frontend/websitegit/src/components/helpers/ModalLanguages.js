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
    Button, } from '@material-ui/core';
import AddIcon from '@material-ui/icons/Add';

const ModalLanguages = () => {
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
            '& #language': {
                marginTop: theme.spacing(2),
            },
            '& #domain': {
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

    const [language, setLanguage] = useState('');
    const [domain, setDomain] = useState('');
    
    const handleChangeLanguage = ( e ) => {
        setLanguage(e.target.value);
    }
    const handleChangeDomain = ( e ) => {
        setDomain(e.target.value);
    }
    return (
        <div className={ classes.content }>
            <div className={ classes.header }>
                <p>Languages</p>
            </div>            
            <Divider />
            <div className={ classes.body }>
                <div id="language">
                    <FormControl className={ classes.margin }>
                        <InputLabel shrink htmlFor="SelectLanguage">Language</InputLabel>
                        <Select
                            labelId="SelectLanguage"
                            id="SelectLanguage"
                            value={ language }
                            onChange={ handleChangeLanguage }
                            input={ <BootstrapInput /> }
                            >
                            <MenuItem value="">
                                <em>None</em>
                            </MenuItem>
                            <MenuItem value={10}>Ten</MenuItem>
                        </Select>
                    </FormControl>
                </div>
                <div id="domain"> 
                    <FormControl className={ classes.margin }>
                        <InputLabel shrink htmlFor="SelectDomain">Domain</InputLabel>
                        <Select
                            labelId="SelectDomain"
                            id="SelectDomain"
                            value={ domain }
                            onChange={ handleChangeDomain }
                            input={ <BootstrapInput /> }
                            >
                            <MenuItem value="">
                                <em>None</em>
                            </MenuItem>
                            <MenuItem value={10}>10%</MenuItem>
                            <MenuItem value={20}>20%</MenuItem>
                            <MenuItem value={30}>30%</MenuItem>
                            <MenuItem value={40}>40%</MenuItem>
                            <MenuItem value={50}>50%</MenuItem>
                            <MenuItem value={60}>60%</MenuItem>
                            <MenuItem value={70}>70%</MenuItem>
                            <MenuItem value={80}>80%</MenuItem>
                            <MenuItem value={90}>90%</MenuItem>
                            <MenuItem value={100}>100%</MenuItem>
                        </Select>
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

export default ModalLanguages
