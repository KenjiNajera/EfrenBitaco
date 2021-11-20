import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import Checkbox from "@material-ui/core/Checkbox";
import FormControl from '@material-ui/core/FormControl';

const useStyles = makeStyles((theme) => ({
  root: {
    width: "100%",
    maxWidth: 200,
    backgroundColor: theme.palette.background.paper,
    padding: 5,
  },
}));

export default function CheckboxList() {
  const classes = useStyles();
  const [checked, setChecked] = React.useState([0]);

  const handleToggle = (value) => () => {
    const currentIndex = checked.indexOf(value);
    const newChecked = [...checked];

    if (currentIndex === -1) {
      newChecked.push(value);
    } else {
      newChecked.splice(currentIndex, 1);
    }
    setChecked(newChecked);
  };

  return (
    <FormControl component="fieldset">
      <List className={classes.root}>
        {[0, 1, 2, 3].map((value) => {
          const labelId = `checkbox-list-label-${value}`;
          return (
            <ListItem
              key={value}
              role={undefined}
              dense
              button
              onClick={handleToggle(value)} >
              <ListItemIcon>
                <Checkbox
                  edge="start"
                  checked={checked.indexOf(value) !== -1}
                  tabIndex={-1}
                  disableRipple
                  inputProps={{ "aria-labelledby": labelId }} />
              </ListItemIcon>
              <ListItemText id={labelId} primary={`Line item ${value + 1}`} />
            </ListItem>
          );
        })}
      </List>
    </FormControl>
  );
}