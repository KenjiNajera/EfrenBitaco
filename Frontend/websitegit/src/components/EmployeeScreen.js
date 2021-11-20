import React, { useEffect, useState } from 'react';
import ListasCustom from './helpers/ListasCustom';
import CardEmployee from './helpers/CardEmployee';
import { useDispatch } from 'react-redux';
import { ChargeData } from '../actions/persistenceActions';

const EmployeeScreen = () => {
    const dispatch = useDispatch(); 
    const [data, setData] = useState([]);

    const handleSetData = ( state, search ) => {
        if(search !== ''){
            setData(
                state.filter((data) =>
                    data.Resourcedata.Firstname.toLowerCase().includes(search.toLowerCase())
                )
            );
        } else {
            setData( state );
        }
    }
    useEffect(() => {
        dispatch( ChargeData("user/getusers/") )
    }, [dispatch])
    return (
        <div>
            <ListasCustom 
                ChangeStatusButtons={true}
                component={ CardEmployee } 
                filterbar={ true }
                data={ data }
                handleSetData={ handleSetData } />
        </div>
    )
}
export default EmployeeScreen;