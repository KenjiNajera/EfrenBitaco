import React from 'react'
import CardEmployeeDetail from './helpers/CardEmployeeDetail'
import { useSelector } from 'react-redux';
import { useHistory, useParams } from 'react-router-dom';


const EmployeeDetailScreen = () => {

    const { userId } = useParams();

    const history = useHistory();

    const state = useSelector(state => state.persistence.data);

    if(state.find(data => userId) === undefined){
        history.push(`/Developers`)
    }

    return (
        <>
            {
                (state.find(data => userId) === undefined)
                ? history.push(`/Developers`)
                : <CardEmployeeDetail item={state.find(data => userId)}/>
            }
        </>
    )
}
export default EmployeeDetailScreen