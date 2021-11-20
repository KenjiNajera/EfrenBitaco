import React from 'react';
import { Switch, Route, Redirect } from 'react-router-dom';

// imports de componentes pantallas
import LoginScreen from '../components/auth/LoginScreen';
import ForgottenPasword from '../components/auth/ForgottenPasword';
import Fondo from '../assets/images/hero.jpg'
import ChangePassword from '../components/auth/ChangePassword';
import VerifyAccount from '../components/auth/VerifyAccount';

//Pantalla Login y password?
const AuthRouter = () => {
    return (
        <div className="containerbackground foundColor">
            <div className="login-container">
                <div className="imageLogin" >
                    <img src={Fondo} alt="Hola"/>
                </div>
                <div className="content-components-login">
                    <Switch>
                        <Route 
                            exact path="/auth/login"
                            component={ LoginScreen }/>
                        <Route 
                            path="/auth/recovery+password"
                            component={ ForgottenPasword }/>
                        <Route
                            path="/auth/change+password/:userId"
                            component={ ChangePassword } />
                        <Route
                            path="/auth/verify+account/:userId"
                            component={ VerifyAccount } />
                            
                        <Redirect to="/auth/login" />
                    </Switch>
                </div>
            </div>
        </div>
    )
}

export default AuthRouter






