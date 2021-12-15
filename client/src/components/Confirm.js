import React from "react";
import PropTypes from 'prop-types'

class Confirm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            items: [],
            reservation: []
        };
    }

    static contextTypes = {
        router: PropTypes.object
    }

    componentDidMount() {
        let book_id = window.sessionStorage.getItem("book-id");
        fetch(`http://localhost:8000/book/${book_id}`);
        let detail_id = window.sessionStorage.getItem("detail-id");
        fetch(`http://localhost:8000/detail/${detail_id}`)
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        items: result
                    });
                },
                (error) => {
                    this.setState({
                        isLoaded: true,
                        error
                    });
                }
            );
        fetch(`http://localhost:8000/retrieve/${book_id}`)
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
    }

    render() {
        const { error, isLoaded, items, reservation } = this.state;
        const res = obj_to_arr(reservation);
        if (error) {
            return <div>Ошибка: {error.message}</div>;
        } else if (!isLoaded) {
            return <div>Загрузка...</div>;
        } else  {
            return (
                <div className="row">
                    <div className="col-md-8">
                        <div className="search-tabs search-tabs-bg mt-5">
                            <h1>Ваш заказ</h1>
                            <div className="tabbable">
                                <div className="tab-content">
                                    <div className="tab-pane active" id="flights" role="tabpanel">
                                        <div className="row">
                                            <div className="col-md-12">
                                                {res.map((detail) =>
                                                    <div className="row">
                                                        <div className="col-md-6">
                                                            Локатор: {detail.locator}
                                                        </div>
                                                        <div className="col-md-6">
                                                            Пассажир: {detail.firstname} {detail.lastname}
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
                    <div className="col-md-4 mb-3">
                        <div className="search-tabs search-tabs-bg mt-5">
                            <h1>Ваш полет</h1>
                            <div className="tabbable">
                                <div className="tab-content">
                                    <div className="tab-pane active" id="flights" role="tabpanel">
                                        <div className="row">
                                            <div className="col-md-12">
                                                <ul className="booking-list">
                                                    {items.map((detail) =>
                                                        <li key={detail.id}>
                                                            <div className="booking-item-container">
                                                                <div className="booking-item">
                                                                    {detail.outbound.map(outbound =>
                                                                        <div className="row">
                                                                            <div className="col-md-12">
                                                                                <div className="booking-item-airline-logo">
                                                                                    {outbound.segments[0].airline} {outbound.segments[0].number}
                                                                                </div>
                                                                            </div>
                                                                            <div className="col-md-12">
                                                                                {outbound.segments.map(segment =>
                                                                                    <div className="booking-item-flight-details" key={segment.id}>
                                                                                        <div className="booking-item-departure">
                                                                                            <i className="fa fa-plane"></i>
                                                                                            <h5>{segment.dep_tim}</h5>
                                                                                            <p className="booking-item-date">{segment.dep_dat}</p>
                                                                                            <p className="booking-item-destination">{segment.dep_arp}</p>
                                                                                        </div>
                                                                                        <div className="booking-item-arrival">
                                                                                            <i className="fa fa-plane fa-flip-vertical"></i>
                                                                                            <h5>{segment.arr_tim}</h5>
                                                                                            <p className="booking-item-date">{segment.arr_dat}</p>
                                                                                            <p className="booking-item-destination">{segment.arr_arp}</p>
                                                                                        </div>
                                                                                    </div>
                                                                                )}
                                                                            </div>
                                                                            <div className="col-md-12">
                                                                                <h5>{outbound.duration} min</h5>
                                                                                <p>{outbound.segments.length > 1 ? outbound.segments.length + ' пересадка' : 'без пересадок'}</p>
                                                                            </div>
                                                                            <div className="col-md-12">
                                                                                <p className="booking-item-flight-class">{outbound.segments[0].cabin} class</p>
                                                                            </div>
                                                                        </div>
                                                                    )}
                                                                    {detail.inbound !== null ?
                                                                        detail.inbound.map(inbound =>
                                                                            <div className="row">
                                                                                <div className="col-md-12">
                                                                                    <div className="booking-item-airline-logo">
                                                                                        <p>{inbound.segments[0].airline}</p>
                                                                                        <p>{inbound.segments[0].number}</p>
                                                                                    </div>
                                                                                </div>
                                                                                <div className="col-md-12">
                                                                                    {inbound.segments.map(segment =>
                                                                                        <div className="booking-item-flight-details" key={segment.id}>
                                                                                            <div className="booking-item-departure">
                                                                                                <i className="fa fa-plane"></i>
                                                                                                <h5>{segment.dep_tim}</h5>
                                                                                                <p className="booking-item-date">{segment.dep_dat}</p>
                                                                                                <p className="booking-item-destination">{segment.dep_arp}</p>
                                                                                            </div>
                                                                                            <div className="booking-item-arrival">
                                                                                                <i className="fa fa-plane fa-flip-vertical"></i>
                                                                                                <h5>{segment.arr_tim}</h5>
                                                                                                <p className="booking-item-date">{segment.arr_dat}</p>
                                                                                                <p className="booking-item-destination">{segment.arr_arp}</p>
                                                                                            </div>
                                                                                        </div>
                                                                                    )}
                                                                                </div>
                                                                                <div className="col-md-12">
                                                                                    <h5>{inbound.duration} min</h5>
                                                                                    <p>{inbound.segments.length > 1 ? --inbound.segments.length + ' пересадка' : 'без пересадок'}</p>
                                                                                </div>
                                                                                <div className="col-md-12">
                                                                                    <p className="booking-item-flight-class">{inbound.segments[0].cabin} class</p>
                                                                                </div>
                                                                            </div>
                                                                        ) : null}
                                                                    <div className="row">
                                                                        <div className="col-md-12">
                                                                            <span className="booking-item-price">{detail.price} {detail.currency}</span>
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                        </li>
                                                    )}
                                                </ul>
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
        if (temp === 'reservation') {
            let o = obj[temp];
            const res = [];
            for (let t in o) {
                res[t] = o[t]
            }
            reservation.push(res);
        }
    }

    return reservation;
}

export default Confirm;
