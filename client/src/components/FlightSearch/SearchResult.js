import React from 'react';
import FlightSearchContainer from '../../containers/FlightSearchContainer';
import ResultList from './ResultList';

const SearchResult = (props) =>
    <div className="row">
        <div className="col-md-12">
            <div className="list-page mt-5">
                <h1>Результаты поиска</h1>
                <div className="row">
                    <div className="col-md-9">
                    {props.isLoading ?
                        <div className="text-center mt-5">
                            <img src="/img/loading.gif" alt="Loading" />
                        </div>
                        :
                        <div className="result-tabs">
                            <div className="tab-content" id="myTabContent">
                                <div role="tabpanel" id="departure" aria-labelledby="home-tab" aria-expanded="true">
                                    <ResultList result={props.departureResult} />
                                </div>
                            </div>
                        </div>
                    }
                    </div>
                    <div className="col-md-3 mb-3">
                        <FlightSearchContainer position="sidebar" />
                    </div>
                </div>
            </div>
        </div>
    </div>

export default SearchResult;
