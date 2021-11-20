import React, { useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';
import { ChargeData } from '../actions/persistenceActions';
import CardTeam from './helpers/CardTeam';
import ListasCustom from './helpers/ListasCustom';

const TeamScreen = () => {
    const dispatch = useDispatch(); 
    const [data, setData] = useState([]);

    const handleSetData = ( state, search ) => {
        if(search !== ''){
            setData(
                state.filter((data) =>
                    data.Projectname.toLowerCase().includes(search.toLowerCase())
                )
            );
        } else {
            setData( state );
        }
    }
   
    useEffect(() => {
        dispatch( ChargeData("project/getprojects/") )
    }, [dispatch]);

    return (
        <div>
            <ListasCustom 
                ChangeStatusButtons={false}
                component={ CardTeam } 
                filterbar={ true }
                data={ data }
                handleSetData={ handleSetData } />
        </div>
    )
}
export default TeamScreen
