import React, { useEffect, useState } from 'react'
import { Switch, Route, Redirect, useLocation } from 'react-router-dom'

import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Drawer from '@material-ui/core/Drawer';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import IconButton from '@material-ui/core/IconButton';
import Badge from '@material-ui/core/Badge';
import ExitToAppTwoToneIcon from '@material-ui/icons/ExitToAppTwoTone';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import NotificationsIcon from '@material-ui/icons/Notifications';
import { SidebarListIcons } from '../components/helpers/SidebarListIcons';
import { List } from '@material-ui/core';

import logo from '../assets/images/ic_git_logo_icon.svg';
import HomeScreen from '../components/HomeScreen';
import EmployeeScreen from '../components/EmployeeScreen'
import TeamScreen from '../components/TeamScreen';
import RequestScreen from '../components/RequestScreen';
import ClientScreen from '../components/ClientScreen';
import ModalCustom from '../components/helpers/ModalCustom';
import ModalLogout from '../components/helpers/ModalLogout';
import SoftskillCatalog from '../components/catalogs/SoftskillCatalog';
import TypeRequestCatalog from '../components/catalogs/TypeRequestCatalog';
import PermissionCatalog from '../components/catalogs/PermissionCatalog';
import ModuleCatalog from '../components/catalogs/ModuleCatalog';
import LanguageCatalog from '../components/catalogs/LanguageCatalog';
import HardskillCatalog from '../components/catalogs/HardskillCatalog';

import { useDispatch, useSelector } from 'react-redux';
import { Reset } from '../actions/persistenceActions';
import { ChangeStatus } from '../actions/meunuActions';
import CountryCatalog from '../components/catalogs/CountryCatalog';
import EmployeeDetailScreen from '../components/EmployeeDetailScreen';
import RoleCatalog from '../components/catalogs/RoleCatalog';
const HomeRouter = () => {
    const drawerWidth = 240;
    const useStyles = makeStyles((theme) => ({
        root: {
            display: 'flex',
        },
        toolbar: {
            paddingRight: 24, // keep right padding when drawer closed
        },
        toolbarIcon: {
            display: 'flex',
            backgroundColor: '#0367A6',
            alignItems: 'center',
            justifyContent: 'flex-end',
            padding: '0 8px',
            ...theme.mixins.toolbar,
        },
        appBar: {
            zIndex: theme.zIndex.drawer + 1,
            backgroundColor: '#0367A6',
            transition: theme.transitions.create(['width', 'margin'], {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.leavingScreen,
            }),
        },
        appBarShift: {
            marginLeft: drawerWidth,
            width: `calc(100% - ${drawerWidth}px)`,
            transition: theme.transitions.create(['width', 'margin'], {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.enteringScreen,
            }),
        },
        menuButton: {
            marginRight: 16,
        },
        menuButtonHidden: {
            display: 'none',
        },
        title: {
            flexGrow: 1,
            fontSize:28,
        },
        drawerPaper: {
            position: 'relative',
            color:'white',
            whiteSpace: 'nowrap',
            flexDirection: 'column',
            overflowX: 'hidden',
            height: '100vh',
            width: drawerWidth,
            backgroundColor: '#0367A6',
            transition: theme.transitions.create('width', {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.enteringScreen,
            }),
        },
        drawerPaperClose: {
            overflowX: 'hidden',
            width: theme.spacing(7),
            [theme.breakpoints.up('sm')]: {
                width: theme.spacing(9),
            },
            [theme.breakpoints.down('sm')]:{
                display: 'none',
            },
            transition: theme.transitions.create('width', {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.leavingScreen,
            }),
        },
        appBarSpacer: {
            height: 64,
        },
        content: {
            flexGrow: 1,
            height: '100vh',
            overflow: 'auto',
        },
        container: {
            width: '100%',
            height: `calc(100% - 96px)`,
            paddingTop: theme.spacing(4),
            marginBottom: theme.spacing(4),
            paddingLeft: theme.spacing(4),
            paddingRight:theme.spacing(4),
        },
        '@global': {
            '*::-webkit-scrollbar': {
                width: '0.5em',
            },
            '*::-webkit-scrollbar-track': {
                '-webkit-box-shadow': 'inset 0 0 6px rgb(5,151,242)',
                borderRadius:'5em',
            },
            '*::-webkit-scrollbar-thumb': {
                backgroundColor: 'rgb(255,255,255)',
                // outline: '1px solid slategrey',
                borderRadius:'5em',
            }
        },
        contentNested: {
            flex: '1 1 auto',
            maxHeight:'400px',
            overflowY: 'scroll',
            overflowX:'hidden',
        },
        nested: {
            paddingLeft: theme.spacing(4),
            overflow: 'hidden',
        },
        textLinks: {
            textDecoration: 'none',
            color: 'white',
        },  
        logo: {
            width: theme.spacing(4),
            maxWidth: theme.spacing(4),
            opacity: 0.5,
        },
    }));

    const dispatch = useDispatch();
    const open = useSelector( state => state.menu.open )

    const location = useLocation();

    useEffect(() => {
        if(location.pathname !== '/Developer/Detail/7'){
            dispatch( Reset() );
        }
        console.log(location.pathname)
    },[])
    
    const classes = useStyles();

    const [openModal, setOpenModal] = useState(false);

    const handleDrawerOpen = (event) => {
        event.preventDefault();
        dispatch(ChangeStatus(true))
    };
    const handleDrawerClose = (event) => {
        event.preventDefault();
        dispatch(ChangeStatus(false))
    };

    const handleModalOpen = () => {
        setOpenModal(true);
    };
    const handleModalClose = () => {
        setOpenModal(false);
    };
    
    return (
        <div className={classes.root}>
            <CssBaseline />
            <AppBar position="absolute" className={ clsx(classes.appBar, open && classes.appBarShift) }>
                <Toolbar className={ classes.toolbar }>
                    <IconButton
                        edge="start"
                        color="inherit"
                        aria-label="open drawer"
                        onClick={ handleDrawerOpen }
                        className={ clsx(classes.menuButton, open && classes.menuButtonHidden) }
                    >
                        <img className={classes.logo} alt="Portal Logo" src={logo} align="center" />
                        {/* <MenuIcon /> */}
                    </IconButton>
                    <Typography component="h1" variant="h6" color="inherit" noWrap className={ classes.title }>
                        Grupo GIT
                    </Typography>
                    <IconButton color="inherit">
                        <Badge badgeContent={ 4 } color="secondary">
                            <NotificationsIcon />
                        </Badge>
                    </IconButton>
                    <IconButton color="inherit">
                        <Badge>
                            <AccountCircleIcon/>
                        </Badge>
                    </IconButton>
                    <IconButton color="inherit" onClick={ handleModalOpen }>
                        <Badge>
                            <ExitToAppTwoToneIcon/>
                        </Badge>
                        
                    </IconButton>
                </Toolbar>
            </AppBar>
            <ModalCustom 
                            handleModalClose={ handleModalClose } 
                            openModal={ openModal }
                            component={ ModalLogout } />
            <Drawer
                variant="permanent"
                classes={{
                paper: clsx(classes.drawerPaper, !open && classes.drawerPaperClose),
                }}
                open={open} >
                <div className={classes.justifystart}>
                <div className={classes.toolbarIcon}>
                    <IconButton onClick={handleDrawerClose}>
                        <ChevronLeftIcon />
                    </IconButton>
                </div>
                <Divider />
                <List  >
                    <SidebarListIcons useStyles={useStyles} handleDrawerOpen={handleDrawerOpen} handleDrawerClose={handleDrawerClose} />
                </List>
                </div>
            </Drawer>
                <main className={classes.content}>
                    <div className={classes.appBarSpacer} />
                    <div className={classes.container} >
                        <Switch>
                            <Route 
                                exact path="/"
                                component={ HomeScreen } />
                            <Route 
                                path="/Developers"
                                component={ EmployeeScreen } />
                            <Route 
                                path="/Teams"
                                component={ TeamScreen } />
                            <Route 
                                path="/Clients"
                                component={ ClientScreen } />
                            <Route 
                                path="/Requests"
                                component={ RequestScreen } />

                            {/* Catalogos */}
                            <Route 
                                path="/Catalog/Softskill"
                                component={ SoftskillCatalog } />
                            <Route 
                                path="/Catalog/Typerequest"
                                component={ TypeRequestCatalog } />
                            <Route 
                                path="/Catalog/Permission"
                                component={ PermissionCatalog } />
                            <Route 
                                path="/Catalog/Module"
                                component={ ModuleCatalog } />
                            <Route 
                                path="/Catalog/Language"
                                component={ LanguageCatalog } />
                            <Route 
                                path="/Catalog/Hardskill"
                                component={ HardskillCatalog } />
                            <Route 
                                path="/Catalog/Country"
                                component={ CountryCatalog } />
                            <Route 
                                path="/Catalog/Roles"
                                component={ RoleCatalog } />

                            {/* Details */}
                            <Route 
                                path="/Developer/Detail/:userId"
                                component={ EmployeeDetailScreen }
                                />
                            <Redirect to="/" />
                        </Switch>
                    </div>
                </main>
    </div>
    )
}

export default HomeRouter
