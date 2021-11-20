import React from 'react'
import { makeStyles, fade, TextField } from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import InputAdornment from '@material-ui/core/InputAdornment';

const Search = ({ handleChange }) => {

    const useStyles = makeStyles((theme) => ({
        root: {
            flexGrow: 1,
        },
        search: {
            position: 'relative',
            borderRadius: theme.shape.borderRadius,
            backgroundColor: fade(theme.palette.common.white, 0.15),
            '&:hover': {
                backgroundColor: fade(theme.palette.common.white, 0.25),
            },
            marginLeft: 0,
            width: '100%',
            [theme.breakpoints.up('sm')]: {
                marginLeft: theme.spacing(1),
                width: 'auto',
            },
        },
        searchIcon: {
            color: 'white',
            padding: theme.spacing(0, 2),
            height: '100%',
            position: 'absolute',
            pointerEvents: 'none',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
        },
        inputRoot: {
            color: 'white',
        },
        inputInput: {
            padding: theme.spacing(1, 1, 1, 0),
            paddingLeft: `calc(1em + ${theme.spacing(4)}px)`,
            transition: theme.transitions.create('width'),
            width: '100%',
            [theme.breakpoints.up('sm')]: {
                width: '12ch',
                '&:focus': {
                width: '20ch',
                },
            },
        },
    }));
    const classes = useStyles();

    return (
        <div className={classes.search}>
            <TextField
                variant="outlined"
                placeholder="Search…"
                classes={{
                    root: classes.inputRoot,
                    input: classes.inputInput,
                }}
                name="search"
                onChange={ handleChange }
                InputProps={{
                    'aria-label': 'search',
                    startAdornment: (
                      <InputAdornment position="start">
                        <SearchIcon />
                      </InputAdornment>
                    ),
                  }}
                />

        </div>
    )
}

export default Search
