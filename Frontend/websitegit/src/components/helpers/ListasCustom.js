import React, { useEffect, useState } from 'react';
import { 
    GridList, 
    GridListTile, 
    isWidthUp, 
    makeStyles, 
    withWidth } from '@material-ui/core';
import ToolBarCustom from './ToolBarCustom';
import { useSelector } from 'react-redux';
import { useSearch } from '../../hooks/useSearch';
import ModalCustom from './ModalCustom';
import Filterbar from './Filterbar';

const ListasCustom = ({
    ChangeStatusButtons,
    clickModal,
    component: Component,
    data,
    filterbar,
    handleSetData,
    modalcomponent: Modal,
    ...props
}) => {

    const useStyles = makeStyles(theme => ({
        secundario: {
            display: 'flex',
            margin: '0px',
            height: '100',
            width: '100',
        },
        cardsContainer: {
            overflow: 'hidden',
            marginLeft: theme.spacing(2),
            padding: theme.spacing(2),
        }, 
        root: {
            width: '100%',
            padding: theme.spacing(2),
        }
    }));
    const classes = useStyles();
    
    const state = useSelector(state => state.persistence.data)

    const [ isDelete, setIsDelete ] = useState(false); 
    const [ 
        isSearch, 
        search, 
        isCreating, 
        handleChange,
        handleModalopen, 
        handleModalClose ] = useSearch();

    useEffect(() => {
            handleSetData( state, search )
    }, [search, handleSetData, state]);

    const handleDeleting = (event) => {
        event.preventDefault();
        setIsDelete(!isDelete)
    }

    let cols = 5;
    if(isWidthUp('xs', props.width)){
        cols = 1;
    }
    if(isWidthUp('sm', props.width)){
        cols = 2;
    }
    if(isWidthUp('md', props.width)){
        cols = 3;
    }
    if(isWidthUp('lg', props.width)){
        cols = 4;
    }
    if(isWidthUp('xl', props.width)){
        cols = 5;
    }

    return (
        <>
            <div className={ classes.secundario}>
                {
                    (filterbar)
                        ? <Filterbar />
                        : null
                }
                <div className={ classes.root }>
                    <ToolBarCustom handleDeleting={ handleDeleting } handleModalopen={ handleModalopen } handleChange={ handleChange } ChangeStatusButtons={ ChangeStatusButtons }/>
                    <GridList cellHeight={ 'auto' } cols={cols} className={ classes.cardsContainer }>
                        {
                            (isSearch) 
                                ?   data.map((item, index) => (
                                        <GridListTile key={ item.id } cols={1}>
                                            <Component   
                                                handleModalopen={ handleModalopen } 
                                                item={ item } 
                                                isDelete={ isDelete } />
                                                {
                                                    (clickModal)
                                                        ? <ModalCustom 
                                                            handleModalClose={ handleModalClose }
                                                            openModal={ isCreating }
                                                            component={() => (<Modal item={item}/>)} />
                                                        : null
                                                }
                                        </GridListTile>
                                    )) 
                                :   state.map((item, index) => (
                                        <GridListTile key={ item.id } cols={1}>
                                            <Component   
                                                handleModalopen={ handleModalopen } 
                                                item={ item } 
                                                isDelete={ isDelete } />
                                                {
                                                    (clickModal)
                                                        ? <ModalCustom 
                                                            handleModalClose={ handleModalClose }
                                                            openModal={ isCreating }
                                                            component={() => (<Modal item={item}/>)} />
                                                        : null
                                                }
                                        </GridListTile>
                                )) 
                        }
                    </GridList>
                    {
                        (clickModal)
                            ? null
                            : <ModalCustom 
                                handleModalClose={ handleModalClose }
                                openModal={ isCreating }
                                component={ Modal } />
                    }
                </div>
            </div>
        </>
    )
}

export default withWidth()(ListasCustom);
