import React from 'react'
import dm_logo from '../images/logo.svg'

export default function Header() {
    return (
        <header className="bg-white border-gray-500 border-b h-10 flex justify-between items-center">

            {/* Логотип */}
            <div className="flex items-center pl-4">
                <img src={dm_logo} alt="Logo" className="w-6 h-6"></img>
                <h2 className="text-2xl font-bold text-center pl-2">МОЛОКОД</h2>
            </div>

            {/* Текст в центре */}
            <div className="grow items-center text-center">Управление маркировкой</div>

            {/* Кнопка выхода и отображение текущего пользователя */}
            <div className="items-center mr-4">
                <span className="text-gray-700 mr-4">Администратор</span>
                <button className="bg-white hover:bg-gray-200 px-3 rounded border-solid border-2 border-gray-500 ">
                    <span className="text-gray-700">Выйти</span>
                </button>
            </div>
        </header>
    )
}