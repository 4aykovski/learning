import React from 'react';
import {BrowserRouter, NavLink, Route, Routes} from "react-router-dom";
import UserPage from "./components/UserPage";
import TodosPage from "./components/TodosPage";
import UserItemsPage from "./components/UserItemsPage";
import TodosItemsPage from "./components/TodosItemPage";

const App = () => {
    return (
        <BrowserRouter>
            <div>
                <NavLink to="/users">Users</NavLink>
                <NavLink to="/todos">Todos</NavLink>
            </div>
            <Routes>
                <Route path="/users" element={<UserPage/>}/>
                <Route path="/todos" element={<TodosPage/>}/>
                <Route path="/users/:id" element={<UserItemsPage/>}/>
                <Route path="/todos/:id" element={<TodosItemsPage/>}/>
            </Routes>
        </BrowserRouter>
    );
};

export default App;