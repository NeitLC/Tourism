import { normalize } from 'normalizr';
import moment from 'moment';
import responseSchema from '../utils/schema';
import { SEARCH_REQUEST, DEPARTURE_SEARCH_RESULT, RETURN_SEARCH_RESULT } from './actionTypes';
import configuration from '../configuration';

const apiUrl = configuration.apiUrl;

export const searchAirports = async (searchString, isDeparture) => {
//    let response = await fetch(`http://localhost:8000/airports/${searchString}`);
    let response = await fetch(`${apiUrl}/airports?searchString=${searchString}`);
    return await response.json();
}

export const searchDepartureFlights = (request) => async (dispatch, getState) => {
    const departureRequest = {
        ...request,
        requestType: "departure"
    }
    return dispatch(searchFlights(departureRequest));
}

export const searchReturnFlights = (request) => async (dispatch, getState) => {
    const returnRequest = {
        ...request,
        requestType: "return"
    }
    return dispatch(searchFlights(returnRequest));
}

export const searchFlights = (request) => async (dispatch, getState) => {
    dispatch({
        type: SEARCH_REQUEST,
        request: request,
        receivedAt: Date.now()
    })

    const apiRequest  = generateApiRequest(request);
    const apiResponse = await fetch(`${apiUrl}/search`,
        {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(apiRequest)
        }
    );
    const responseData   = await apiResponse.json();
    const normalizedData = normalize(responseData, responseSchema);

    return dispatch({
        type: request.requestType === "departure" ? DEPARTURE_SEARCH_RESULT : RETURN_SEARCH_RESULT,
        response: normalizedData,
        receivedAt: Date.now()
    });
}

const generateApiRequest = (request) => {
    let origin      = request.departureAirport.code;
    let destination = request.arrivalAirport.code;
    let date        = request.departureDate;
    let passengers  = 1;
    let date2;
    if (request.isRoundTrip) {
        date2 = request.returnDate;
    }

    const apiRequest = {
        from: origin,
        to: destination,
        date: moment(date).format('YYYY-MM-DD'),
        date2: request.isRoundTrip ? moment(date2).format('YYYY-MM-DD') : null,
        passengers: passengers,
        solutions: 20
    }

    return apiRequest;
}
