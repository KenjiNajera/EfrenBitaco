import React from 'react';
import { makeStyles } from "@material-ui/core/styles";
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormControl from '@material-ui/core/FormControl';

const useStyles = makeStyles((theme) => ({
  root: {
    width: "100%",
    maxWidth: 200,
    backgroundColor: theme.palette.background.paper,
    padding: 5,
  },
  subtittle: {
    margin: theme.spacing(1),
  },
}));

export default function RadioButtonsGroup() {
  const classes = useStyles();

  const [value, setValue] = React.useState('Catalogos');

  const handleChange = (event) => {
    setValue(event.target.value);
  };

  return (
    <FormControl component="fieldset">
      <RadioGroup aria-label="gender" name="gender1" value={value} onChange={handleChange} className={classes.root}>
        <FormControlLabel value="Catalogs" control={<Radio />} label="Catalogs" />
        <FormControlLabel value="Users" control={<Radio />} label="Users" />
        <FormControlLabel value="Clients" control={<Radio />} label="Clients" />
        <FormControlLabel value="Teams" control={<Radio />} label="Teams" />
        <FormControlLabel value="Projects" control={<Radio />} label="Projects" />
      </RadioGroup>
    </FormControl>
  );
}
