import React, { useEffect, useState } from 'react'
import CardRequest from './helpers/CardRequest'
import ListasCustom from './helpers/ListasCustom'
import { useDispatch } from 'react-redux';
import { ChargeData } from '../actions/persistenceActions';
import ModalApprovedRequest from './helpers/ModalApprovedRequest';

const RequestScreen = () => {
    const dispatch = useDispatch(); 
    const [data, setData] = useState([]);

    const handleSetData = ( state, search ) => {
        if(search !== ''){
            setData(
                state.filter((data) =>
                    data.Firstname.toLowerCase().includes(search.toLowerCase())
                )
            );
        } else {
            setData( state );
        }
    }
    useEffect(() => {
        dispatch( ChargeData("request/getrequests/") )
    }, [dispatch])
    return (
        <div>
            <ListasCustom 
                clickModal={true}
                modalcomponent={ ModalApprovedRequest }
                ChangeStatusButtons={false}
                component={ CardRequest } 
                filterbar={ false }
                data={ data }
                handleSetData={ handleSetData } 
                />
        </div>
    )
}

export default RequestScreen
