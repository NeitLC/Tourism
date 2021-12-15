import React from 'react';
import { Switch, Route, Redirect } from 'react-router-dom';
import Header from './Header';
import Footer from './Footer';
import Homepage from './Homepage';
import Detail from './Detail';
import Payment from './Payment';
import Confirm from './Confirm';
import Reservation from './Reservation';
import Ticket from './Ticket';
import Checkin from './Checkin';
import SearchResultContainer from '../containers/SearchResultContainer';

const App = () => (
    <div>
        <Header />
        <div className="container main">
            <Switch>
                <Route path="/homepage" component={Homepage} />
                <Route path="/search" component={SearchResultContainer} />
                <Route path="/detail" component={Detail} />
                <Route path="/payment" component={Payment} />
                <Route path="/confirm" component={Confirm} />
                <Route path="/reservation" component={Reservation} />
                <Route path="/ticket" component={Ticket} />
                <Route path="/checkin" component={Checkin} />
                <Redirect from="/" to="/homepage" />
            </Switch>
        </div>
        <Footer />
    </div>
)

export default App;
