import React from "react";
import PropTypes from 'prop-types'
import { withRouter } from 'react-router-dom'

class Checkin extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            reservation: []
        };
    }

    static contextTypes = {
        router: PropTypes.object
    }

    componentDidMount() {
	let id = window.sessionStorage.getItem("checkin-id");
        fetch(`http://localhost:8000/checkin/${id}`)
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        reservation: result
                    });
                },
                (error) => {
                    this.setState({
                        isLoaded: true,
                        error
                    });
                }
            );
        this.context.router.history.push('/reservation');
    }

    render() {
        const { error, isLoaded, reservation } = this.state;
        const res = obj_to_arr(reservation);
        if (error) {
            return <div>Ошибка: {error.message}</div>;
        } else if (!isLoaded) {
            return <div>Загрузка...</div>;
        } else  {
            return (
                <div className="row">
                    <div className="col-md-12">
                        <div className="search-tabs search-tabs-bg mt-5">
                            <h1>Заказы</h1>
                            <div className="tabbable">
                                <div className="tab-content">
                                    <div className="tab-pane active" id="flights" role="tabpanel">
                                        <div className="row">
                                            <div className="col-md-12">
                                                <div className="row">
                                                    <div className="col-md-1">
                                                        Локатор
                                                    </div>
                                                    <div className="col-md-2">
                                                        Пассажир
                                                    </div>
                                                    <div className="col-md-1">
                                                        Откуда
                                                    </div>
                                                    <div className="col-md-1">
                                                        Куда
                                                    </div>
                                                    <div className="col-md-1">
                                                        Дата
                                                    </div>
                                                    <div className="col-md-2">
                                                        Цена
                                                    </div>
                                                    <div className="col-md-4">
                                                        &nbsp;
                                                    </div>
                                                </div>
                                                <div className="row">
                                                    <div className="col-md-12">
                                                        <hr/>
                                                    </div>
                                                </div>
                                            </div>
                                            <div className="col-md-12">
                                                {res.map((detail) => 
                                                    <div className="row">
                                                        <div className="col-md-1">
                                                            {detail['reservation']['locator']}
                                                        </div>
                                                        <div className="col-md-2">
                                                            {detail['reservation']['firstname']} {detail['reservation']['lastname']}
                                                        </div>
                                                        <div className="col-md-1">
                                                            {detail['routes'][0]['outbound'][0]['from']}<br/>
                                                            {detail['routes'][0]['inbound'] !== null ?
                                                                detail['routes'][0]['inbound'][0]['from'] :
                                                                null
                                                            }
                                                        </div>
                                                        <div className="col-md-1">
                                                            {detail['routes'][0]['outbound'][0]['to']}<br/>
                                                            {detail['routes'][0]['inbound'] !== null ?
                                                                detail['routes'][0]['inbound'][0]['to'] :
                                                                null
                                                            }
                                                        </div>
                                                        <div className="col-md-1">
                                                            {detail['routes'][0]['outbound'][0]['date']}<br/>
                                                            {detail['routes'][0]['inbound'] !== null ?
                                                                detail['routes'][0]['inbound'][0]['date'] :
                                                                null
                                                            }
                                                        </div>
                                                        <div className="col-md-2">
                                                            {detail['routes'][0]['price']} {detail['routes'][0]['currency']}
                                                        </div>
                                                        <div className="col-md-2">
                                                            <ButtonTicket id={detail.id} />
                                                        </div>
                                                        <div className="col-md-2">
                                                            <ButtonCheckin id={detail.id} />
                                                        </div>
                                                        <div className="col-md-12">
                                                            <hr/>
                                                        </div>
                                                    </div>
                                                )}
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            );
        }
    }
}

const obj_to_arr = (obj) => {
    const reservation = [];

    for (let temp in obj) {
        let o = obj[temp];
        const res = [];
        for (let t in o) {
            res[t] = o[t]
        }
        reservation.push(res);
    }

    return reservation;
}

const ButtonTicket = withRouter(({ history, id }) => (
    <button className="btn btn-primary btn-lg" type='button' onClick={() => ticket(history, id)}>Тикет</button>
))

const ButtonCheckin = withRouter(({ history, id }) => (
    <button className="btn btn-primary btn-lg" type='button' onClick={() => checkin(history, id)}>Чек-ин</button>
))

export const ticket = async (history, id) => {
    history.push('/ticket');
    window.sessionStorage.setItem("ticket-id", id);
}

export const checkin = async (history, id) => {
    history.push('/checkin');
    window.sessionStorage.setItem("checkin-id", id);
}

export default Checkin;
