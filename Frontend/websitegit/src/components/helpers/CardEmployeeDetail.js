import React from 'react';
import Card from '@material-ui/core/Card';
import EducationExperience from './EducationExperience';
import SkillsLanguages from './SkillsLanguages';
import { AppBar, Button, Grid, Tabs, Tab, Box, Typography, makeStyles } from '@material-ui/core';
import DataPersonal from './DataPersonal';

const useStyles = makeStyles((theme) => ({
    cardContainer: {
        width: 'auto',
        position: 'relative',
        marginLeft: theme.spacing(3),
        marginRight: theme.spacing(3),
        [theme.breakpoints.up(600 + theme.spacing(2) * 2)]: {
            width: 800,
            marginLeft: 'auto',
            marginRight: 'auto',
        },
        marginBottom: theme.spacing(6),
        marginTop: theme.spacing(3),
    },
    media: {
        height: 200,
        padding: theme.spacing(1),

    },
    name: {
        fontSize: '.9vw',
        marginTop: 20,
        marginLeft: 4,
    },
    aboutimage: {
        display: 'flex',
        marginLeft: '10px',
        marginTop: 6,
        width: '10vh',
        height: '10vh',
        borderRadius: '100%',
        objectFit: 'contain',
    },
    card: {
        display: 'flex',
    },
    submit: {
        marginRight: theme.spacing(3),
        marginBottom: theme.spacing(3),
        marginTop: theme.spacing(3),
        justifyContent: 'center',
        marginLeft: '40%',
        backgroundColor: '#0367A6',
        height: '6vh',
        width: '25vh',
        fontSize:'1vw'
    },
    pad: {
        marginLeft: theme.spacing(3),
        marginRight: theme.spacing(3),
        marginBottom: theme.spacing(3),
        marginTop: theme.spacing(3),
    },
    color: {
        color: 'black',
        backgroundColor: '#FFFFFF'
    },
}));

function a11yProps(index) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
    };
}
function TabPanel(props) {
    const { children, value, index, ...other } = props;
    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`simple-tabpanel-${index}`}
            aria-labelledby={`simple-tab-${index}`}
            {...other}
        >
            {value === index && (
                <Box p={1}>
                    <Typography>{children}</Typography>
                </Box>
            )}
        </div>
    );
}

const CardEmployeeDetail = ({ item }) => {
    const classes = useStyles();
    
    const [value, setValue] = React.useState(0);
    const handleChange = (event, newValue) => {
        setValue(newValue);
    };

    console.log(item)
    return (
        <div className={classes.cardContainer}>
                <Card>
                    <AppBar className={classes.color} position="static">
                        <Tabs value={value} onChange={handleChange} aria-label="simple tabs example">
                            <Tab label="Datos" {...a11yProps(0)} />
                            <Tab label="Educación & Experiencia" {...a11yProps(1)} />
                            <Tab label="Habilidades & Lenguajes" {...a11yProps(2)} />
                        </Tabs>
                    </AppBar>
                    <Box display="" justification="center" borderBottom={1}></Box>
                    <Grid className={classes.pad}>
                        <TabPanel value={value} index={0}>
                            <DataPersonal datas={item}/> 
                        </TabPanel>
                        <TabPanel value={value} index={1} >
                            <EducationExperience cert={item.Resourcedata.Certificates} educ={item.Resourcedata.Educations} expe={item.Resourcedata.Experiences}/> 
                        </TabPanel>
                        <TabPanel value={value} index={2}>
                            <SkillsLanguages hard={item.Resourcedata.Hardskills} lang={item.Resourcedata.Languages} sof={item.Resourcedata.Softskills}/>
                        </TabPanel>
                    </Grid>
                    <Box display="" justification="center" borderBottom={1}></Box>
                    <Button
                        type="submit"
                        
                        variant="contained"
                        color="primary"
                        className={classes.submit}>
                        Upload
                    </Button>
                </Card>
        </div>
    );
}

export default CardEmployeeDetail;