import React, { useState} from 'react';
import { useForm } from 'react-hook-form';
import {Link} from "react-router-dom";
import { ErrorMessage } from '@hookform/error-message';
import { yupResolver } from '@hookform/resolvers';
import * as yup from 'yup';

import Fondo from '../../assets/images/ic_git_logo_icon.svg';
import Fondo2 from '../../assets/images/ic_git_logo_text.svg';

const LoginScreen = () => {
    
    const [view, setView] = useState(true);

    const LoginSchema =  yup.object().shape({
        email: yup.string()
                .email( () => <span>Email must be a valid email</span> )
                .required( () => <span>Email is required</span> ),
        password: yup.string().required( () => <span>Password is required</span> )
    });

    const { register, handleSubmit, errors } = useForm({
        resolver: yupResolver(LoginSchema),
    });

    const onSubmit = data => {
        
    }
    
    const handleChange = () => {
        if(view)
        {
            document.getElementById('eyepassword').type = 'text';
            document.getElementById('password').style.color="#0597F2";
            setView(false);
        }
        if(!view){
            document.getElementById('eyepassword').type = 'password';
            document.getElementById('password').style.color="#0367A6";
            setView(true);
        }
    }

    return (
        <>
            <div className="loginScreen">
                <div id="Logo">
                    <img src={Fondo} alt="Hola" id="logo-icon"/>
                    <img src={Fondo2} alt="Holas" id="logo-text"/>
                </div>
                <form onSubmit={ handleSubmit(onSubmit) }>
                    <div className="content-input">
                        <div>
                            <i id="email" className="tiny material-icons">account_circle</i>
                            <input 
                                type="text" 
                                name="email"
                                placeholder="Email"  
                                className="letter"
                                ref={register}
                                />
                        </div>
                        <ErrorMessage errors={errors} name="email" />
                    </div>
                    <div className="content-input">
                        <div>
                            <i id="password" onClick={handleChange} className="tiny material-icons">remove_red_eye</i>
                            <input 
                                id="eyepassword"
                                type="password" 
                                name="password"
                                placeholder="Contraseña"  
                                className="validate"
                                ref={register}
                                />
                        </div>
                        <ErrorMessage errors={errors} name="password" />
                    </div>
                    <div className="content-checkbox">
                        <label>
                            <input type="checkbox" className="filled-in" />
                            <span>Remember me</span>
                        </label>
                    </div>
                    <div>
                        <button type="submit">Sing In</button>
                    </div>
                </form>
                <Link to="/auth/recovery+password" id="link">You have forgotten your password?</Link>            
            </div>
        </>
    )
}

export default LoginScreen;