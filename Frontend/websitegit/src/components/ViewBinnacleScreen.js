import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useForm } from 'react-hook-form';
import { ErrorMessage } from '@hookform/error-message';
import { yupResolver } from '@hookform/resolvers';
import * as yup from 'yup';
import { 
    TableCell, 
    TableRow, 
    IconButton,
    TextField, 
    FormHelperText, 
} from '@material-ui/core';
import moment from 'moment';
import ClearIcon from '@material-ui/icons/Clear';
import DeleteIcon from '@material-ui/icons/Delete';
import DoneIcon from '@material-ui/icons/Done';
import EditIcon from '@material-ui/icons/Edit';

import { useSearch } from '../hooks/useSearch';
import TableCustom from './catalogs/TableCustom';
import { ChargeData } from '../actions/persistenceActions';

const ViewBinnacleScreen = () => {

    const [ 
        isSearch, 
        search, 
        isCreating, 
        whereUpdating, 
        handleChange, 
        reset, 
        handleSetValues,
        handleModalopen, 
        handleModalClose 
    ] = useSearch();

    const schema = yup.object().shape({
        Hardskillname: yup.string().trim()
            .required( () => <span>Name is required</span> ),
        id: yup.number(),
    });
    const { register, handleSubmit, errors } = useForm({
        resolver: yupResolver(schema),
    });

    const dispatch = useDispatch();
    const state = useSelector( state => state.persistence.data );

    useEffect(() => {
        dispatch( ChargeData( "binnacle/getbinnacle/5" ) );
    }, [])

    const onSubmit = data => {
        // reset();
    }
    const handleUpdating = (id) => ( event ) => {
        event.preventDefault();
        // if(whereUpdating === 0){
        //     let datas = state.filter(data => data.id == id); 
        //     console.log(datas)
        //     setFile(datas[0].Image)
        //     handleSetValues({
        //         isCreating:  isCreating,
        //         whereUpdating: id
        //     });
        // } else {
            
        //     handleSetValues({
        //         isCreating:  isCreating,
        //         whereUpdating: id
        //     });
        // }
    }
    
    const handleDelete = ( id ) => ( event ) => {
        event.preventDefault();
        // dispatch( Deleting(`hardskill/${id}`, id) );
    }

    const headCells = [
        { id: 'Dia', numeric: false, disablePadding: false, label: 'Dia' },
        { id: 'Fecha', numeric: false, disablePadding: false, label: 'Fecha' },
        { id: 'Hr. Entrada', numeric: false, disablePadding: false, label: 'Hr. Entrada' },
        { id: 'Hr. Salida', numeric: false, disablePadding: false, label: 'Hr. Salida' },
        { id: 'Hr. Dia', numeric: false, disablePadding: false, label: 'Hr. Dia' },
        { id: 'Hr.Comida', numeric: false, disablePadding: false, label: 'Hr.Comida' },
        { id: 'Total Horas/Día', numeric: false, disablePadding: false, label: 'Total Horas/Día' },
        { id: 'Actividades del dia', numeric: false, disablePadding: false, label: 'Actividades del dia' },
    ];

    const row = ( row, classes ) => {
        
        return (
            <TableRow
                hover tabIndex={-1}
                key={row.id}>
                <TableCell  component="th" scope="row" padding="none" className={ classes.row }>
                    {row.daydata.Dayname}
                </TableCell>
                <TableCell align="left">{ moment(row.daydata.Daydate).format('DD/MM/YYYY') }</TableCell>
                <TableCell align="left">{ moment(row.checktime.Checktime).format('hh:mm:ss') }</TableCell>
                <TableCell align="left">{ moment(row.departuretime.Departuretime).format('hh:mm:ss') }</TableCell>
                <TableCell align="left">{ moment(row.daydata.Dailytime).format('hh:mm:ss') }</TableCell>
                <TableCell align="left">{ moment(row.mealtime.Mealtime).format('hh:mm:ss') }</TableCell>
                <TableCell align="left">{ moment(row.daydata.Totalhoursday).format('hh:mm:ss') }</TableCell>
                <TableCell align="left">{ row.activity.Summary }</TableCell>
                {
                    (whereUpdating === row.id)
                        ?   <TableCell align="right"> 
                                <IconButton aria-label="filter list" onClick={  handleSubmit( onSubmit ) } >
                                    <DoneIcon style={{ color: '#28A745' }}/>
                                </IconButton>
                                <IconButton aria-label="filter list"  onClick={ handleUpdating(0) } >
                                    <ClearIcon style={{ color: '#DC3545' }}/>
                                </IconButton>
                            </TableCell>   
                        :   <TableCell align="right"> 
                                <IconButton aria-label="filter list" onClick={ handleUpdating( row.id ) } >
                                    <EditIcon style={{ color: '#FFC107' }}/>
                                </IconButton>
                                <IconButton aria-label="filter list" onClick={ handleDelete( row.id ) } >
                                    <DeleteIcon  style={{ color: '#DC3545' }}/>
                                </IconButton>
                            </TableCell>
                } 
            </TableRow>
        );
    }

    return (
        <>
            <TableCustom 
                name={ "" }
                data={ Array.isArray(state.days) ? state.days : [] }  
                headCells={ headCells} 
                handleModalopen={ null }
                handleChange={ null }
                isOrderHeadTable={false}
                rowCustom={row}
            />
        </>
    )
}

export default ViewBinnacleScreen;
