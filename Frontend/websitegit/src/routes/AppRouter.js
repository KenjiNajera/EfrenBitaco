import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Switch } from 'react-router-dom';
// import { useDispatch } from 'react-redux';
import PrivateRoutes from './PrivateRoutes';
import PublicRoutes from './PublicRoutes';
import AuthRouter from './AuthRouter';
import HomeRouter from './HomeRouter';



const AppRouter = () => {
    
    // const dispatch = useDispatch();
    
    const [checking, setChecking] = useState( true );
    const [isLoggedIn, setIsLoggedIn] = useState( false );
    const sd = 6;
    // Logica para logueo
    useEffect(() => {
        if( sd === 6 ) {
            //dispatch( login( 1,'sd' ) );
            setIsLoggedIn( true );
        } else {
            setIsLoggedIn( false );
        }

        setChecking( false );
    }, [ setChecking, setIsLoggedIn])

    if( checking ) {
        return(
            // componente de carga de pantalla
            <h1>Espere...</h1>
        )
    } 
    
    return (
        <Router>
            <div>
                <Switch>
                    {/* Rutas Publicas */}
                    <PublicRoutes 
                        path="/auth"
                        component={ AuthRouter }
                        isLoggedIn={ isLoggedIn }/>
                    {/* Rutas Privadas */}
                    <PrivateRoutes 
                        path="/"
                        component={ HomeRouter }
                        isLoggedIn={ isLoggedIn }/>

                </Switch>
            </div>
        </Router>
    )
}

export default AppRouter 
