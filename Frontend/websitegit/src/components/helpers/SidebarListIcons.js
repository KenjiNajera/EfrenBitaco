import React from 'react'

import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import PersonIcon from '@material-ui/icons/Person';
import GroupWorkIcon from '@material-ui/icons/GroupWork';
import BookIcon from '@material-ui/icons/Book';
import GroupIcon from '@material-ui/icons/Group';
import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import Collapse from '@material-ui/core/Collapse';
import List from '@material-ui/core/List';
import { Link } from 'react-router-dom';

export const SidebarListIcons = ({useStyles, handleDrawerOpen, handleDrawerClose }) => {
    const classes = useStyles();

    const [selectedIndex, setSelectedIndex] = React.useState(null);
    
    const [open, setOpen] = React.useState(false);
    
    const handleClick = (event) => {
        open ? handleDrawerClose(event) : handleDrawerOpen(event);
        setOpen(!open);
    };
    const handleListItemClick = (event, index) => {
        if(index === 4) handleClick(event);
        
        setSelectedIndex(index);
    };

    return (
        <div>
            <Link to="/Developers" className={classes.textLinks} >
                <ListItem button selected={selectedIndex === 0} onClick={(event) => handleListItemClick(event, 0)} >
                    <ListItemIcon>
                        <PersonIcon style={ {color: 'white'} } />
                    </ListItemIcon>
                    
                    <ListItemText primary="Developers" />
                </ListItem>
            </Link>
            
            <Link to="/Teams" className={classes.textLinks} >
                <ListItem button selected={selectedIndex === 1} onClick={(event) => handleListItemClick(event, 1)} >
                    <ListItemIcon>
                        <GroupWorkIcon style={ {color: 'white'} } />
                    </ListItemIcon>
                    <ListItemText primary="Teams" />
                </ListItem>
            </Link>

            <Link to="/Clients" className={classes.textLinks} >
                <ListItem button  selected={selectedIndex === 2} onClick={(event) => handleListItemClick(event, 2)} >
                    <ListItemIcon>
                        <GroupIcon style={ {color: 'white'} } />
                    </ListItemIcon>
                    <ListItemText primary="Clients" /> 
                </ListItem>
            </Link>

            <Link to="/Requests" className={classes.textLinks} >
                <ListItem button selected={selectedIndex === 3} onClick={(event) => handleListItemClick(event, 3)} >
                    <ListItemIcon>
                        <BookIcon style={ {color: 'white'} } />
                    </ListItemIcon>
                    <ListItemText primary="Requests" />
                </ListItem>
            </Link>

            <ListItem button  selected={selectedIndex === 4} onClick={(event) => handleListItemClick(event, 4)} >
                <ListItemIcon >
                    <LibraryBooksIcon style={ {color: 'white'} } />
                </ListItemIcon>
                <ListItemText primary="Catalogs" />
            </ListItem>
            <Collapse className={classes.contentNested} in={open} timeout="auto" unmountOnExit>
                <List component="div" disablePadding>
                    <Link to="/Catalog/Softskill" className={classes.textLinks}>
                        <ListItem button selected={selectedIndex === 5} onClick={(event) => handleListItemClick(event, 5)} className={classes.nested}>
                            <ListItemIcon>
                                <LibraryBooksIcon style={ {color: 'white'} } />
                            </ListItemIcon>
                            <ListItemText primary="Softskills" />
                        </ListItem>
                    </Link>

                    <Link to="/Catalog/Typerequest" className={classes.textLinks}>
                        <ListItem button selected={selectedIndex === 6} onClick={(event) => handleListItemClick(event, 6)} className={classes.nested}>
                            <ListItemIcon>
                                <LibraryBooksIcon style={ {color: 'white'} } />
                            </ListItemIcon>
                            <ListItemText primary="Type Requests" />
                        </ListItem>
                    </Link>

                    <Link to="/Catalog/Permission" className={classes.textLinks}>
                        <ListItem button selected={selectedIndex === 7} onClick={(event) => handleListItemClick(event, 7)} className={classes.nested}>
                            <ListItemIcon>
                                <LibraryBooksIcon style={ {color: 'white'} } />
                            </ListItemIcon>
                            <ListItemText primary="Permissions" />
                        </ListItem>
                    </Link>

                    <Link to="/Catalog/Module" className={classes.textLinks}>
                        <ListItem button selected={selectedIndex === 8} onClick={(event) => handleListItemClick(event, 8)} className={classes.nested}>
                            <ListItemIcon>
                                <LibraryBooksIcon style={ {color: 'white'} } />
                            </ListItemIcon>
                            <ListItemText primary="Modules" />
                        </ListItem>
                    </Link>

                    <Link to="/Catalog/Language" className={classes.textLinks}>
                        <ListItem button selected={selectedIndex === 9} onClick={(event) => handleListItemClick(event, 9)} className={classes.nested}>
                            <ListItemIcon>
                                <LibraryBooksIcon style={ {color: 'white'} } />
                            </ListItemIcon>
                            <ListItemText primary="Languages" />
                        </ListItem>
                    </Link>

                    <Link to="/Catalog/Hardskill" className={classes.textLinks}>
                        <ListItem button selected={selectedIndex === 11} onClick={(event) => handleListItemClick(event, 11)} className={classes.nested}>
                            <ListItemIcon>
                                <LibraryBooksIcon style={ {color: 'white'} } />
                            </ListItemIcon>
                            <ListItemText primary="Hardskills" />
                        </ListItem>
                    </Link>

                    <Link to="/Catalog/Country" className={classes.textLinks}>
                        <ListItem button selected={selectedIndex === 12} onClick={(event) => handleListItemClick(event, 12)} className={classes.nested}>
                            <ListItemIcon>
                                <LibraryBooksIcon style={ {color: 'white'} } />
                            </ListItemIcon>
                            <ListItemText primary="Countries" />
                        </ListItem>
                    </Link>
                    
                    <Link to="/Catalog/Roles" className={classes.textLinks}>
                        <ListItem button selected={selectedIndex === 13} onClick={(event) => handleListItemClick(event, 13)} className={classes.nested}>
                            <ListItemIcon>
                                <LibraryBooksIcon style={ {color: 'white'} } />
                            </ListItemIcon>
                            <ListItemText primary="Roles" />
                        </ListItem>
                    </Link>
                </List>
            </Collapse>
        </div>
    )
}
