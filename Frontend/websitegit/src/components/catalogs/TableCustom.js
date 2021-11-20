import React, { useState } from 'react'
import { 
    Paper, 
    makeStyles, 
    TablePagination, 
    TableContainer, 
    Table, 
    TableBody,
} from '@material-ui/core';

import ToolbarTable from './ToolbarTable';
import HeadTable from './HeadTable';

function descendingComparator(a, b, orderBy) {
    if (b[orderBy] < a[orderBy]) {
        return -1;
    }
    if (b[orderBy] > a[orderBy]) {
        return 1;
    }
    return 0;
}

function getComparator(order, orderBy) {
    return order === 'desc'
        ? (a, b) => descendingComparator(a, b, orderBy)
        : (a, b) => -descendingComparator(a, b, orderBy);
}

function stableSort(array, comparator) {    
    const stabilizedThis = array.map((el, index) => [el, index]);
    stabilizedThis.sort((a, b) => {
        const order = comparator(a[0], b[0]);
        if (order !== 0) return order;
        return a[1] - b[1];
    });
    return stabilizedThis.map((el) => el[0]);
}

const TableCustom = ({ 
    data, 
    headCells, 
    name,
    rowCustom,
    isOrderHeadTable,
    handleModalopen,
    handleChange,
}) => {
    const useStyles = makeStyles((theme) => ({
        root: {
            margin: 'auto',
            width: '95%',
        },
        paper: {
            width: '100%',
            marginBottom: theme.spacing(2),
        },
        table: {
            width: '100%',
            paddingLeft: theme.spacing(5),
            paddingRight: theme.spacing(5),
        },
        visuallyHidden: {
            border: 0,
            clip: 'rect(0 0 0 0)',
            height: 1,
            margin: -1,
            overflow: 'hidden',
            padding: 0,
            position: 'absolute',
            top: 20,
            width: 1,
        },
        input: {
            display: 'none',
        },
        superposition: {
            position: 'absolute',
            bottom:-7,
            right: -7,
            '& #buttonPhoto': {
                color: '#000000',
            },
        },
        papertable: {
            width: 40, 
            height:40,
            borderRadius: '100%',
            position: 'relative', 
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            '& img':{ 
                width: '100%',
                height: '100%',
                borderRadius: '100%'
            },
        },
        error: {
            '& span':{
                color: '#bf1650',
                '&::before':{
                    display: 'inline',
                    content: '"⚠ "',
                },
            },
        },
        row: {
            paddingLeft: theme.spacing(2),
            paddingRight: theme.spacing(5),
        },  
    }));
    const classes = useStyles();
    
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(5);
    const [order, setOrder] = useState("asc");
    const [orderBy, setOrderBy] = useState("ID");

    const handleChangePage = (event, newPage) => {
        setPage(newPage);
    };
    
    const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
    };

    const handleRequestSort = (event, property) => {
        const isAsc = orderBy === property && order === 'asc';
        setOrder(isAsc ? 'desc' : 'asc');
        setOrderBy(property);
    };
    
    return (
        <div className={ classes.root }>
            <Paper className={ classes.paper }>
                <ToolbarTable name={ name } handleModalopen={ handleModalopen } handleChange={ handleChange } />
                <TableContainer 
                    className={classes.table}>
                    <Table 
                        aria-labelledby="tableTitle"
                        size="medium"
                        aria-label="enhanced table" >
                        <HeadTable 
                            classes={classes}
                            order={order}
                            orderBy={orderBy}
                            isOrder={isOrderHeadTable}
                            onRequestSort={handleRequestSort}
                            headCells={headCells}
                            rowCount={data.length} />
                        <TableBody>
                            {
                                stableSort( Array.isArray(data) ? data : [] , getComparator(order, orderBy))
                                    .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                                    .map((row, index) => (rowCustom(row, classes))
                                )
                            }
                        </TableBody>
                    </Table>
                </TableContainer>
                <TablePagination 
                    component="div"
                    rowsPerPageOptions={[5,10,25]}
                    count={data.length}
                    rowsPerPage={rowsPerPage}
                    page={page}
                    onChangePage={ handleChangePage }
                    onChangeRowsPerPage={ handleChangeRowsPerPage } />
            </Paper>
        </div>
    );
}

export default TableCustom
