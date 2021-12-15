import { connect } from 'react-redux';
import { searchDepartureFlights, searchReturnFlights } from '../actions/';
import SearchResult from '../components/FlightSearch/SearchResult';

const ttt = (obj) => {
    const flightDetails = [];

    for (let temp in obj) {
        let o = obj[temp];
        const flightDetail = [];
        for (let t in o) {
            flightDetail[t] = o[t]
        }
        flightDetails.push(flightDetail);
    }

    return flightDetails;
}

const mapStateToProps = state => {
    const search = state.search;
    const ret    = {
        departureResult: ttt(search.response.departure.entities.responses.undefined),
        request: search.request,
        isLoading: search.isLoading,
        isRoundTrip: search.request.isRoundTrip,
        isDeparture: search.request.requestType === "departure"
    }

    return ret;
};

const mapDispatchToProps = {
    searchDepartureFlights,
    searchReturnFlights
}

export default connect(mapStateToProps, mapDispatchToProps)(SearchResult);
