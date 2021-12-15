import React from 'react';
import { withRouter } from 'react-router-dom'

const ResultList = ({ result }) =>
    <ul className="booking-list">
        {result.map(detail =>
            <li key={detail.id}>
                <div className="booking-item-container">
                    <div className="booking-item">
                        {detail.outbound.map(outbound =>
                            <div className="row">
                                <div className="col-md-2">
                                    <div className="booking-item-airline-logo">
                                        <p>{outbound.segments[0].airline}</p>
                                        <p>{outbound.segments[0].number}</p>
                                    </div>
                                </div>
                                <div className="col-md-5">
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
                                <div className="col-md-2">
                                    <h5>{outbound.duration} min</h5>
                                    <p>{outbound.segments.length > 1 ? outbound.segments.length + ' пересадка' : 'без пересадок'}</p>
                                </div>
                                <div className="col-md-3">
                                    <span className="booking-item-price">{outbound.price} {outbound.currency}</span>
                                    <p className="booking-item-flight-class">{outbound.segments[0].cabin} class</p>
                                </div>
                            </div>
                        )}
                        {detail.inbound !== null ?
                            detail.inbound.map(inbound =>
                                <div className="row">
                                    <div className="col-md-2">
                                        <div className="booking-item-airline-logo">
                                            <p>{inbound.segments[0].airline}</p>
                                            <p>{inbound.segments[0].number}</p>
                                        </div>
                                    </div>
                                    <div className="col-md-5">
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
                                    <div className="col-md-2">
                                        <h5>{inbound.duration} min</h5>
                                        <p>{inbound.segments.length > 1 ? --inbound.segments.length + ' пересадка' : 'без пересадок'}</p>
                                    </div>
                                    <div className="col-md-3">
                                        <span className="booking-item-price">{inbound.price} {inbound.currency}</span>
                                        <p className="booking-item-flight-class">{inbound.segments[0].cabin} class</p>
                                    </div>
                                </div>
                            ) : null}
                        <div className="row">
                            <div className="col-md-9">
                                <Button id={detail.id} />
                            </div>
                            <div className="col-md-3">
                                <span className="booking-item-price">{detail.price} {detail.currency}</span>
                            </div>
                        </div>
                    </div>
                </div>
            </li>
        )}
    </ul>

const Button = withRouter(({ history, id }) => (
    <button className="btn btn-primary btn-lg" type='button' onClick={() => detailFlights(history, id)}>Выбрать</button>
))

export const detailFlights = async (history, searchString) => {
    history.push('/detail');
    window.sessionStorage.setItem("detail-id", searchString);
}

export default ResultList;
