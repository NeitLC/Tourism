import React from 'react';
import AutoSuggest from '../AutoSuggest/';
import CalendarPicker from '../CalendarPicker/';
import { airPortSelector, autoSuggestRender } from '../../utils/';

const SidebarSearchForm = (props) =>
    <form className="booking-item-dates-change">
        <div className="row">
            <div className="col-md-12 trip-tabs mt-2 mb-2 sidebar-tabs">
                <ul className="nav nav-tabs" role="tablist">
                    <li className="nav-item">
                        <a className={toggleClass(props.isRoundTrip)} data-toggle="tab" role="tab" onClick={() => props.onSetField('isRoundTrip', true)}>В оба направления</a>
                    </li>
                    <li className="nav-item">
                        <a className={toggleClass(!props.isRoundTrip)} data-toggle="tab" role="tab" onClick={() => props.onSetField('isRoundTrip', false)}>В одном направлении</a>
                    </li>
                </ul>
            </div>
        </div>
        <div className="form-group form-group-sm form-group-icon-left">
            <i className="fa fa-map-marker input-icon"></i>
            <label>Вылет из</label>
            <AutoSuggest
                inputValue={props.departureAirport}
                items={props.departureAirports}
                getItem={airPortSelector}
                onChange={(event, value) => props.onAirportChange(event, value, true)}
                onSelect={(event, value) => props.onAirportSelect(event, value, true)}
                renderItem={autoSuggestRender}
            />
        </div>
        <div className="form-group form-group-sm form-group-icon-left">
            <i className="fa fa-map-marker input-icon"></i>
            <label>Прилет в</label>
            <AutoSuggest
                inputValue={props.arrivalAirport}
                items={props.arrivalAirports}
                getItem={airPortSelector}
                onChange={props.onAirportChange}
                onSelect={props.onAirportSelect}
                renderItem={autoSuggestRender}
            />
        </div>
        <div className="form-group form-group-sm form-group-icon-left">
            <i className="fa fa-calendar input-icon input-icon-highlight"></i>
            <label>Дата вылета</label>
            <CalendarPicker onChange={(field, value) => props.onSetField('departureDate', value)} value={props.departureDate} />
        </div>
        <div className={"form-group form-group-sm form-group-icon-left" + (props.isRoundTrip ? " show" : " hide")}>
            <i className="fa fa-calendar input-icon input-icon-highlight"></i>
            <label>Дата прилета</label>
            <CalendarPicker onChange={(field, value) => props.onSetField('returnDate', value)} value={props.returnDate} startDate={props.departureDate} />
        </div>
        {props.errorMessage ?
            <div className="input-group mb-3 has-danger">
                <div className="form-control-feedback">{props.errorMessage}</div>
            </div>
            : null
        }
        <button className="btn btn-primary" type="button" onClick={props.onSearchClick} disabled={props.isLoading}>Изменить</button>
    </form>

const toggleClass = (value) => "nav-link" + (value ? " active" : "");

export default SidebarSearchForm;
