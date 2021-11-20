import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { ErrorMessage } from "@hookform/error-message";
import { yupResolver } from "@hookform/resolvers";
import * as yup from "yup";
import axios from 'axios';


import Toast from '../helpers/Toast';
const ForgottenPasword = () => {
    const emailSchema = yup.object().shape({
        email: yup.string()
            .email( () => <span>Email must be a valid email</span> )
            .required( () => <span>Emal is required</span> )
            .trim(),
    });

    const { register, handleSubmit, errors } = useForm({
        resolver: yupResolver(emailSchema),
    });

    const [list, setlist] = useState([])

    const onSubmit = data => {
        axios.post(`https://jsonplaceholder.typicode.com/users`, { data })
            .then( resp => {
                
            });
            
    }

    return (
        <div className="content-forgotten">
            <label className="title">¿Olvidaste su contraseña?</label>
            <label className="title">¡No hay problema!</label>
            <label className="subtitle">
                Sólo escribe el correo electrónico con el que te registraste y
                te enviaremos un enlace para reestablecer tu contraseña.
                </label>

            <form onSubmit={handleSubmit(onSubmit)}>
                <div className="content-input">
                    <label className="subtitlemail">Email:</label>
                    <input
                        className="emailinput"
                        name="email"
                        type="text"
                        placeholder="Email"
                        ref={register} />
                    <ErrorMessage errors={errors} name="email" as="p" />
                </div>
                <div className="">
                <button type="submit" className="buttonforgett">Solicitar Contraseña</button>
                </div>
            </form>

            <div className="note">
                <label >
                    Nota: Si no recibes nuestro e-mail revisa que hayas escrito tu correo correctamente o tal vez llegó a tu bandeja de correo no deseado,
                    si todo está bien y aún tienes problemas para ingresar escribe dascontacto@grupogit.com
                </label>
            </div>

            <Toast 
                    toastList={list}
                    position={'top-right'}
                    />
        </div>

    )
}

export default ForgottenPasword;