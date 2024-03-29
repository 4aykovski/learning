import React, {useContext} from 'react';
import {Link} from "react-router-dom";
import {AuthContext} from "../../../context";
import MyButton from "../button/MyButton";

const Navbar = () => {
    const {isAuth, setIsAuth} = useContext(AuthContext);

    const logOut = () => {
        setIsAuth(false);
        localStorage.removeItem('auth')
    }

    return (
        <div>
            <div className='navbar'>
                <MyButton onClick={logOut}>Logout</MyButton>
                <div className="navbar__links">
                    <Link to="/about">О сайте</Link>
                    <Link to="/posts">Посты</Link>
                </div>
            </div>
        </div>
    );
};

export default Navbar;