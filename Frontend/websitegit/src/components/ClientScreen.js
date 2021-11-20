import React, { useEffect, useState } from 'react';
import ListasCustom from './helpers/ListasCustom';
import { useDispatch } from 'react-redux';
import { ChargeData } from '../actions/persistenceActions';
import CardClient from './helpers/CardClient';
import ModalClientCreate from './helpers/ModalClientCreate';

const ClientScreen = () => {
    const dispatch = useDispatch();
    const [data, setData] = useState([]);

    const handleSetData = ( state, search ) => {
        if(search !== ''){
            setData(
                state.filter((data) =>
                    data.Clientname.toLowerCase().includes(search.toLowerCase())
                )
            );
        } else {
            setData( state );
        }
    }

    useEffect(() => {
        dispatch( ChargeData("client/clients/") )
    }, [dispatch]);
    
    return (
        <div>
            <ListasCustom 
                ChangeStatusButtons={true}
                component={ CardClient } 
                modalcomponent={ ModalClientCreate } 
                filterbar={ false }
                data={ data }
                handleSetData={ handleSetData } />
        </div>
    )
}

export default ClientScreen;