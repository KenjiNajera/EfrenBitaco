import React from 'react'
import clsx from 'clsx';
import { Toolbar, Typography, Tooltip, IconButton, makeStyles, lighten } from '@material-ui/core'
import AddIcon from '@material-ui/icons/Add';
import Search from '../helpers/Search';

const ToolbarTable = (props) => {
    const useToolbarStyles = makeStyles((theme) => ({
        root: {
            borderRadius: `4px 4px 0 0`,
            backgroundColor: '#0367A6',
            paddingLeft: theme.spacing(3),
            paddingRight: theme.spacing(1),
        },
        highlight:
            theme.palette.type === 'light'
            ? {
                color: theme.palette.secondary.main,
                backgroundColor: lighten(theme.palette.secondary.light, 0.85),
                }
            : {
                color: theme.palette.text.primary,
                backgroundColor: theme.palette.secondary.dark,
                },
        title: {
            flex: '1 1 100%',
            color: 'white',
        },
    }));
    const classes = useToolbarStyles();
    const { numSelected, name, handleModalopen, handleChange } = props;

    return (
        <Toolbar
            className={clsx(classes.root, {
                [classes.highlight]: numSelected > 0,
            })}
            >
                <Typography className={classes.title} variant="h6" id="tableTitle" component="div">
                    { name }
                </Typography>
                <Search handleChange={ handleChange } />
                <Tooltip title="Add">
                    <IconButton  aria-label="filter list" onClick={ handleModalopen } >
                        <AddIcon fontSize="small" style={ {color:'white'} } />
                    </IconButton>
                </Tooltip>
        </Toolbar>
    )
}

export default ToolbarTable
