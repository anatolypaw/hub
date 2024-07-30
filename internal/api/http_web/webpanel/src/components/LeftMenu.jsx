import React from "react";
import { Link } from "gatsby";

export default function LeftMenu() {
    return (
        <nav className="bg-white w-48 border-r border-gray-500">
            <ul>
                <li>
                    <Link to="c1" className="hover:bg-gray-200 block py-2 px-4">
                        Главная
                    </Link>
                </li>
                <li>
                    <Link to="t2" className="hover:bg-gray-200 block py-2 px-4">
                        О нас
                    </Link>
                </li>
                <li >
                    <Link to="#" className="hover:bg-gray-200 block py-2 px-4">
                        Блог
                    </Link>
                </li>
                <li >
                    <Link to="login" className="hover:bg-gray-200 block py-2 px-4">
                        Логин
                    </Link>
                </li>
            </ul>
        </nav>
    )
}