import Axios from "axios";
import { types } from '../types/types';
import { setSuccessToast, setErrorToast } from "./toastActions";

const BaseUrl = "https://resource-grupogit.herokuapp.com/";

export const Reset = () => ({
    type: types.reset
});

export const FormChange = (data) => ({
    type: types.formchange,
    payload: {
        data
    }
});

// Metodos de carga
export const DataChange = (data) => ({
    type: types.datachange,
    payload: {
        data
    }
});
export const ChargeData = ( url ) => {
    return async( dispatch ) => {
        await Axios.get(BaseUrl+url).then( resp => {
            dispatch( DataChange( resp.data ) );
        }).catch(error => {
        });
    }
}

// Metodos Create
export const CreatingWithImage = ( url, data, formdata ) => {
    return async ( dispatch ) => {
        await Axios.post(BaseUrl + url, formdata).then( resp => {
            if( resp.data.id !== 0 ) {
                data.id = resp.data.id;
                data.Image = resp.data.Image
                dispatch( Create( data ) );
                dispatch( setSuccessToast() );
            } else {
                dispatch( setErrorToast() );
            }
        }).catch( (error) => {
            dispatch( setErrorToast() );
        });
    }
}
export const Creating = ( url, data ) => {
    return async ( dispatch ) => {
        await Axios.post(BaseUrl + url, data).then( resp => {
            if( resp.data !== 0 ) {
                data.id = resp.data;
                console.log(resp.data)
                console.log("==========")
                console.log(data)
                dispatch( Create( data ) );
                dispatch( setSuccessToast() );
            } else {
                dispatch( setErrorToast() );
            }
        }).catch( (error) => {
            dispatch( setErrorToast() );
        });
    }
}
export const Create = ( data ) => ({
    type: types.createData,
    payload: data,
});

// Metodos Update
export const UpdatingWithImage = ( url, data, formdata ) => {
    return async ( dispatch ) => {
        await Axios.put(BaseUrl + url, formdata).then( resp => {
            if( resp.data.id !== 0 ) {
                data.Image = resp.data.Image
                dispatch( Update( data ) );
                dispatch( setSuccessToast() );
            } else {
                dispatch( setSuccessToast() );
            }
        }).catch( (error) => {
            dispatch( setErrorToast() );
        });
    }
}
export const Updating = ( url, data) => {
    return async ( dispatch ) => {
        await Axios.put(BaseUrl+url, data).then( resp => {
            if ( resp.data ){
                dispatch( Update( data ) );
                dispatch( setSuccessToast() );
            } else {
                dispatch( setErrorToast() );
            }
        }).catch((error) => {
            dispatch( setErrorToast() );
        });
    }
}
export const Update = ( data ) => ({
    type: types.updateDataById,
    payload: {
        id: data.id,
        data: data,
    }
});

// Metodos Delete
export const Deleting = ( url, id ) => {
    return async (dispatch)  => {
        await Axios.delete(BaseUrl+url).then( resp => {
            if (resp.data){
                dispatch( Delete(id) );
                dispatch( setSuccessToast() );
            } else {
                dispatch( setErrorToast() );
            }
        }).catch((error) => {
            dispatch( setErrorToast() );
        });
    }
}
export const Delete = ( id ) => ({
    type:  types.deleteDataById,
    payload: id,
});




