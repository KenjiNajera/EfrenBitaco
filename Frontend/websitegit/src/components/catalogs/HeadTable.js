import React from 'react'
import { 
    TableHead, 
    TableRow, 
    TableCell, 
    TableSortLabel,
} from '@material-ui/core'


const HeadTable = (props) => {
    const { classes, order, orderBy, onRequestSort, headCells, isOrder } = props;
    const createSortHandler = (property) => (event) => {
        onRequestSort(event, property);
    };
    return (
        <TableHead>
            <TableRow>
                {headCells.map((headCell) => (
                    <TableCell
                        key={headCell.id}
                        className={classes.row}
                        align={headCell.numeric ? 'right' : 'left'}
                        padding={headCell.disablePadding ? 'none' : 'default'}
                        sortDirection={orderBy === headCell.id ? order : false}
                    >
                        {
                            !isOrder
                            ? headCell.label
                            : <TableSortLabel
                            active={orderBy === headCell.id}
                            direction={orderBy === headCell.id ? order : 'asc'}
                            onClick={createSortHandler(headCell.id)}
                            >
                                {headCell.label}
                                {orderBy === headCell.id ? (
                                    <span className={classes.visuallyHidden}>
                                    {order === 'desc' ? 'sorted descending' : 'sorted ascending'}
                                    </span>
                                ) : null}
                            </TableSortLabel>
                        }
                    </TableCell>
                ))}
                <TableCell
                    className={classes.row}
                    align="right"
                    padding="default"
                    >
                        Actions
                </TableCell>
            </TableRow>
        </TableHead>
    )
}

export default HeadTable
