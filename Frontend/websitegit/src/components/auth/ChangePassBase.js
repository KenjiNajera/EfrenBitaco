import React from 'react'
import { useParams } from 'react-router-dom'
import { useForm } from 'react-hook-form';
import { ErrorMessage } from '@hookform/error-message';
import  axios from 'axios';

const ChangePassword = () => {
    const { userId: userid} = useParams();

    const { register, handleSubmit, errors, setError, watch } = useForm();

    const onSubmit = data => {
        axios.post(`https://jsonplaceholder.typicode.com/users`, { 
            hashuserid: userid,
            newpassword: data.password
        })
            .then(resp => {
                console.log(resp)
                console.log(resp.data)
            })
    }

    return (
        <div>
            <form onSubmit={ handleSubmit(onSubmit) }>
                <div>
                    <label>Password</label>
                    <input
                        name="password"
                        type="password"
                        ref={ register({ required: true })}
                        />
                    <ErrorMessage errors={errors} name="password" as="p" message="Password is required"/>
                </div>
                <div>
                    <label>Confirm password</label>
                    <input
                        name="confirpass"
                        type="password"
                        ref={ register({ required: true, 
                            validate: (value) => value === watch('password') }) }
                        />
                    <ErrorMessage errors={errors} name="confirpass" as="span" message="Passwords do not match"/>
                </div>
                <button type="submit" className="buttonforgett">Change password</button>
            </form>
        </div>
    )
}

export default ChangePassword