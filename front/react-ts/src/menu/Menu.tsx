import React, {Dispatch, FC, SetStateAction} from 'react';
import classes from './Menu.module.css';
import {IItem} from "../types/types";



interface MenuProps {
    header: string;
    items: IItem[];
    active: boolean;
    setActive: Dispatch<SetStateAction<boolean>>;
}

const Menu: FC<MenuProps> = ({header, items, active, setActive}) => {
    return (
        <div className={active ? classes.menu + ' ' + classes.active : classes.menu} onClick={() => setActive(false)}>
            <div className={classes.blur}></div>
            <div className={classes.menuContent} onClick={e => e.stopPropagation()}>
                <div className={classes.menuHeader}>
                    {header}
                    <ul>
                        {items.map(item =>
                        <li>
                            <a href={item.href}>{item.value}</a>
                        </li>)}
                    </ul>
                </div>
            </div>
        </div>
    );
};

export default Menu;