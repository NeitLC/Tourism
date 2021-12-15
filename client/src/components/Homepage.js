import React from 'react';
import FlightSearchContainer from '../containers/FlightSearchContainer';

const Homepage = () => (
    <div className="row">
        <div className="col-md-12">
            <div className="search-tabs search-tabs-bg mt-5">
                <h1>Найдите Ваш Полет Мечты</h1>
                <div className="tabbable">
                    <div className="tab-content">
                        <div className="tab-pane active" id="flights" role="tabpanel">
                            <FlightSearchContainer position="homepage" />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
)

export default Homepage;
