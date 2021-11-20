import React, { useState } from 'react';
import logo from '../../assets/images/hero.jpg'
import { useForm } from 'react-hook-form';
import { useDispatch, useSelector } from 'react-redux';
import { FormChange } from '../../actions/persistenceActions';
import { useHistory } from 'react-router-dom';

const WorkerData = () => {

    const dispatch = useDispatch()
    
    const { push } = useHistory();
    const datacharge = useSelector( state => state.persistence )

    const [imagen, setImagen] = useState({
        file: null,
        imgURL: logo
    })
    const {register, handleSubmit, errors } = useForm({
        defaultValues: datacharge.form        
    })

    const onSubmit = data => {

        dispatch( FormChange( data ) );

        push("/step2");
    }

    const handleImageChange = (e) => {
        e.preventDefault();
        let file = e.target.files[0];
        let reader = new FileReader();
        reader.onloadend = function ( e )  {
            setImagen({
                file: file,
                imgURL: reader.result
            })
        }.bind(this)
        reader.readAsDataURL(file)
    }
    return (
        <div className="Container-data">
            <label className="titulos">Data</label>
            <form>
                <div>
                    <img src={imagen.imgURL}></img>
                    <input 
                        className="emailinputwork" 
                        name="firstname" 
                        type="text"
                        ref={register} 
                        placeholder="First Name" />

                    <input 
                        className="emailinputwork" 
                        name="lastname" 
                        type="text" 
                        ref={register}
                        placeholder="Last Name" />

                    <div className="drag-drop">
                        <input 
                            type="file" 
                            multiple="multiple" 
                            name="photo"
                            onChange={handleImageChange} 
                            ref={register}
                            />
                        <span className="fa-stack ">
                            <i className="fa fa-camera  bottom pulsating"></i>
                        </span>
                    </div>
                </div >
                <div className="centrado">
                    <input 
                        className="address" 
                        name="email" 
                        type="text" 
                        ref={register} 
                        placeholder="Address" />
                </div>
                <div className="centradomails">
                    <input 
                        className="emailinputwork" 
                        name="PersonalEmail" 
                        type="text" 
                        ref={register}
                        placeholder="Personal mail" />

                    <input 
                        className="emailinputwork" 
                        name="BussinessMail" 
                        ref={register}
                        type="text" 
                        placeholder="Business mail" />
                </div>
                <div className="centradomails">
                    <input 
                        className="emailinputwork" 
                        name="Cellphone" 
                        type="text"
                        ref={register} 
                        placeholder="Cellphone" />
                    <input 
                        className="emailinputwork" 
                        name="Datebirth" 
                        type="date" 
                        ref={register}
                        />
                </div>
                <div className="centradomails">
                    <input 
                        className="emailinputwork" 
                        name="City" 
                        type="text"
                        ref={register} 
                        placeholder="City" />
                    <input 
                        className="emailinputwork" 
                        name="State" 
                        type="text" 
                        ref={register}
                        placeholder="State/Province/Region" />
                </div>
                <div className="centradomails">
                    <input 
                        className="emailinputwork" 
                        name="Country" 
                        type="text" 
                        ref={register}
                        placeholder="Country" />
                </div>
            </form>
            <div className="centradomails">
                <button onClick={handleSubmit(onSubmit)} className="buttonNext">Next</button>
            </div>
        </div>
    );
}

export default WorkerData;