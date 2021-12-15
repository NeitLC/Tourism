import React from 'react';
import { Link } from 'react-router-dom';

const Header = () => (
    <nav className="navbar navbar-toggleable-md navbar-light bg-faded">
        <button className="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
        </button>
        <div className="container">
            <div className="collapse navbar-collapse" id="navbarNavDropdown">
                <ul className="navbar-nav">
                    <li className="nav-item active">
                        <Link className="nav-link" to="/homepage">Домой</Link>
                        <Link className="nav-link" to="/reservation">Заказы</Link>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
)

export default Header;
