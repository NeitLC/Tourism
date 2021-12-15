import React from "react";
import PropTypes from 'prop-types'
import configuration from '../configuration';

const apiUrl = configuration.apiUrl;

class Detail extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            items: []
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    static contextTypes = {
        router: PropTypes.object
    }

    handleChange(event) {
        const target = event.target;
        const value  = target.value;
        const name   = target.name;
        this.setState({
            [name]: value
        });
    }

    handleSubmit(event) {
        window.sessionStorage.setItem("payment-id", this.state.items[0].id);
        store(this.state, this.state.items[0].id);
        this.context.router.history.push('/payment');
    }

    componentDidMount() {
        let id = window.sessionStorage.getItem("detail-id");
        fetch(`http://localhost:8000/detail/${id}`)
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
    }

    render() {
        const { error, isLoaded, items } = this.state;
        if (error) {
            return <div>Ошибка: {error.message}</div>;
        } else if (!isLoaded) {
            return <div>Загрузка...</div>;
        } else  {
            return (
                <div className="row">
                    <div className="col-md-8">
                        <div className="search-tabs search-tabs-bg mt-5">
                            <h1>Введите данные</h1>
                            <div className="tabbable">
                                <div className="tab-content">
                                    <div className="tab-pane active" id="flights" role="tabpanel">
                                        <div className="row">
                                            <div className="col-md-12">
                                                <form onSubmit={this.handleSubmit}>
                                                    <div className="row">
                                                        <div className="col-md-6">
                                                            <div className="form-group form-group-sm">
                                                                <label>Имя</label>
                                                                <input type="text" name="firstname" onChange={this.handleChange} />
                                                            </div>
                                                        </div>
                                                        <div className="col-md-6">
                                                            <div className="form-group form-group-sm">
                                                                <label>Фамилия</label>
                                                                <input type="text" name="lastname" onChange={this.handleChange} />
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div className="row">
                                                        <div className="col-md-6">
                                                            <div className="form-group form-group-sm">
                                                                <label>Дата рождения</label>
                                                                <input type="text" name="birthdate" onChange={this.handleChange} />
                                                            </div>
                                                        </div>
                                                        <div className="col-md-6">
                                                            <div className="form-group form-group-sm">
                                                                <label>Номер паспорта</label>
                                                                <input type="text" name="passport" onChange={this.handleChange} />
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div className="row">
                                                        <div className="col-md-6">
                                                            <div className="form-group form-group-sm">
                                                                <label>Почта</label>
                                                                <input type="text" name="email" onChange={this.handleChange} />
                                                            </div>
                                                        </div>
                                                        <div className="col-md-6">
                                                            <div className="form-group form-group-sm">
                                                                <label>Телефон</label>
                                                                <input type="text" name="phone" onChange={this.handleChange} />
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div className="row">
                                                        <div className="col-md-9">
                                                            <Button />
                                                        </div>
                                                    </div>
                                                </form>
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

const Button = () => (
    <input className="btn btn-primary btn-lg" type="submit" value="Отправить" />
)

export const generateRequest = (request, id) => {
    return {
        id_routes: id,
        firstname: request.firstname,
        lastname: request.lastname,
        birthdate: request.birthdate,
        passport: request.passport,
        email: request.email,
        phone: request.phone
    }
}

export const store = (request, id) => {
    const apiRequest = generateRequest(request, id);
    const apiOptions = {
        method: "POST",
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(apiRequest)
    };
    fetch(`${apiUrl}/store`, apiOptions)
        .then(res => res.json())
        .then((result) => {
            window.sessionStorage.setItem("book-id", result);
        });
}

export default Detail;
