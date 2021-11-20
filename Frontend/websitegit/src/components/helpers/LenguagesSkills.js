import React from 'react';
import { useHistory } from 'react-router-dom';

const LenguagesSkills = () => {
    const { push } = useHistory()
    const handlerBack = () => {
        push("/step2")
    }
    return (
        <div className="Container-datas">
            <label className="titulos">Skills & Languages</label>
            <div className="card-Technicals">
                <label className="Subtitulos">Technicals</label>
                <button className="buttonadd"><i className="fa fa-plus"></i></button>
                <div className="Technicals">
                    <div className="CardGenLeng">
                        <label className="TitleLenguage">c#</label>
                        <progress className="barprogress" max="100" value="40"></progress>
                        <span className="barprogress">40</span>
                    </div>
                </div>
            </div>
            <div className="card-content">
                <label className="Subtitulos">Lenguage</label>
                <button className="buttonadd"><i className="fa fa-plus"></i></button>
                <div className="Language">
                    <div className="CardGenLeng">
                        <label className="TitleLenguage">Spanish</label>
                        <progress className="barprogress" max="100" value="20"></progress>
                        <span className="barprogress">20</span>
                    </div>
                </div>
            </div>
            <div className="card-Softskill">
                <label className="Subtitulos">Softskill</label>
                <button className="buttonadd"><i className="fa fa-plus"></i></button>
                <div className="Softskill">
                    <div className="CardGenLeng">
                        <label className="TitleLenguage">Spanish</label>
                        <progress className="barprogress" max="100" value="20"></progress>
                        <span className="barprogress">20</span>
                    </div>
                </div>
            </div>
            <button onClick={handlerBack} className="buttonBack"><a>Back</a></button>
            <button className="buttonFinish"><a>Finish</a></button>
        </div>
    );
}

export default LenguagesSkills; 