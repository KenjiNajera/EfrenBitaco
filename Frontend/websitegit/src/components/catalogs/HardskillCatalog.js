import React, { useEffect, useState } from 'react'
import { 
    TableCell, 
    TableRow, 
    IconButton,
    TextField, 
    FormHelperText, 
    Paper,
} from '@material-ui/core';
import PhotoCamera from '@material-ui/icons/PhotoCamera';
import ClearIcon from '@material-ui/icons/Clear';
import DeleteIcon from '@material-ui/icons/Delete';
import DoneIcon from '@material-ui/icons/Done';
import EditIcon from '@material-ui/icons/Edit';


import { useDispatch, useSelector } from 'react-redux';

import TableCustom from './TableCustom';
import { ChargeData, Deleting, UpdatingWithImage } from '../../actions/persistenceActions';
import Toast from '../helpers/Toast';
import ModalCustom from '../helpers/ModalCustom';
import { useSearch } from '../../hooks/useSearch';

import { useForm } from 'react-hook-form';
import { ErrorMessage } from '@hookform/error-message';
import { yupResolver } from '@hookform/resolvers';
import * as yup from 'yup';
import ModalHardskillCreate from './ModalHardskillCreate';

const HardskillCatalog = () => {

    const [ 
        isSearch, 
        search, 
        isCreating, 
        whereUpdating, 
        handleChange, 
        reset, 
        handleSetValues,
        handleModalopen, 
        handleModalClose ] = useSearch();

    const dispatch = useDispatch();
    const state = useSelector( state => state.persistence.data );
    const { type, message, isOpen } = useSelector( state => state.toast );

    const [data, setData] = useState([]);
    const schema = yup.object().shape({
        Hardskillname: yup.string().trim()
            .required( () => <span>Name is required</span> ),
        id: yup.number(),
    });
    const { register, handleSubmit, errors } = useForm({
        resolver: yupResolver(schema),
    });
    
    const [file, setFile] = useState('');

    useEffect(() => {  
        dispatch( ChargeData("hardskill/hardskills/") );
    }, [dispatch]);
    
    useEffect(() => {
        if(search !== ''){
            setData(
                state.filter((data) =>
                    data.Hardskillname.toLowerCase().includes(search.toLowerCase())
                )
            );
        } else {
            setData( state );
        }
    }, [search, state]);

    const onSubmit = data => {
        var formdata = new FormData();
        formdata.append("Image", data.Image[0]);
        formdata.append("Hardskillname", data.Hardskillname);
        if(data.Image.length !== 0){
            formdata.append("image", data.Image[0].name);
            data.Image = data.Image[0].name;
        } else {
            formdata.append("image", file);
        }
        dispatch( UpdatingWithImage(`hardskill/${data.id}`, data, formdata ) );
        reset();
    }
    const headCells = [
        { id: 'id', numeric: false, disablePadding: true, label: 'ID' },
        { id: 'Image', numeric: false, disablePadding: false, label: 'Image' },
        { id: 'Hardskillname', numeric: false, disablePadding: false, label: 'Harskill' },
    ];
    const row = ( row, classes ) => (
        <TableRow
            hover tabIndex={-1}
            key={row.id}>
            <TableCell  component="th" scope="row" padding="none" className={ classes.row }>
                {row.id}
            </TableCell>
            {
                (whereUpdating === row.id)
                    ?   <TableCell>
                            <Paper  className={ classes.papertable } >
                                <img
                                    src={file}
                                    alt="imagen"
                                    />
                                <input 
                                    name="Image"
                                    ref={ register }
                                    accept="image/*" 
                                    className={ classes.input } 
                                    onChange={(event) => setFile(URL.createObjectURL(event.target.files[0])) }
                                    id="icon-button-file" 
                                    type="file" />
                                <label htmlFor="icon-button-file" className={ classes.superposition }>
                                    <IconButton id="buttonPhoto" size="small" aria-label="upload picture" component="span">
                                        <PhotoCamera fontSize="small" />
                                    </IconButton>
                                </label>
                            </Paper>
                        </TableCell>
                    :   <TableCell>
                            <Paper  className={ classes.papertable } >
                                <img
                                    src={row.Image}
                                    alt="imagen"
                                    />
                            </Paper>
                        </TableCell>
            }
            { 
                (whereUpdating === row.id)
                    ?   <TableCell align="left">  
                            <form onSubmit={ handleSubmit( onSubmit ) }>
                                <input type="hidden" name="id" ref={register} defaultValue={row.id} />
                                <TextField 
                                    inputRef={register}
                                    name="Hardskillname"
                                    defaultValue={row.Hardskillname} size="small" /> 
                                    <FormHelperText className={ classes.error }>
                                        <ErrorMessage errors={errors} name="Hardskillname" />
                                    </FormHelperText>
                            </form> 
                        </TableCell>   
                    :   <TableCell align="left">{row.Hardskillname}</TableCell>
            }
            {
                (whereUpdating === row.id)
                    ?   <TableCell align="right"> 
                            <IconButton aria-label="filter list" onClick={  handleSubmit( onSubmit ) } >
                                <DoneIcon style={{ color: '#28A745' }}/>
                            </IconButton>
                            <IconButton aria-label="filter list"  onClick={ handleUpdating(0) } >
                                <ClearIcon  style={{ color: '#DC3545' }}/>
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
        </TableRow>);

    const handleUpdating = (id) => ( event ) => {
        event.preventDefault();
        if(whereUpdating === 0){
            let datas = state.filter(data => data.id === id); 
            console.log(datas)
            setFile(datas[0].Image)
            handleSetValues({
                isCreating:  isCreating,
                whereUpdating: id
            });
        } else {
            
            handleSetValues({
                isCreating:  isCreating,
                whereUpdating: id
            });
        }
    }
    
    const handleDelete = ( id ) => ( event ) => {
        event.preventDefault();
        dispatch( Deleting(`hardskill/${id}`, id) );
    }
    return (
        <>
            <TableCustom 
                name={ "Hardskill" }
                data={ isSearch ? data : state }  
                headCells={ headCells} 
                handleModalopen={ handleModalopen }
                handleChange={ handleChange }
                rowCustom={row}
                />
            <Toast 
                isOpen={ isOpen }
                message={ message }
                type={ type }
                />
            <ModalCustom 
                handleModalClose={ handleModalClose }
                openModal={ isCreating }
                component={ ModalHardskillCreate }
                />
        </>
    )
}

export default HardskillCatalog
