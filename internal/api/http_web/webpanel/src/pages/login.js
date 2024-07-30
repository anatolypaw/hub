import React from "react";

// Форма для авторизации

export default function Login() {
    return (
        <>
            <div className="flex items-center justify-center h-screen font-mono">
                <div className="bg-white border-2 border-gray-300 rounded-lg p-4 w-full max-w-sm">

                    <h2 className="text-2xl font-bold mb-2 text-center">МОЛОКОД</h2>
                    <h2 className="text-1xl mb-2 text-center">Авторизация</h2>

                    <form>
                        <div className="mb-4">
                            <label className="block text-sm font-bold mb-2" htmlFor="username">Логин</label>
                            <input
                                className="appearance-none border-2 border-gray-300 rounded w-full py-2 px-3  leading-tight focus:outline-none focus:shadow-outline"
                                id="username" type="text" placeholder="Введите логин"></input>
                        </div>

                        <div className="mb-6">
                            <label className="block text-sm font-bold mb-2" htmlFor="password">Пароль</label>
                            <input
                                className="appearance-none border-2 border-gray-300 rounded w-full py-2 px-3  mb-2 leading-tight focus:outline-none focus:shadow-outline"
                                id="password" type="password" placeholder="Введите пароль"></input>
                        </div>

                        <button
                            className="bg-white hover:bg-gray-200 px-3 rounded border-solid border-2 border-gray-500 "
                            type="button">
                            Войти
                        </button>
                    </form>
                </div>
            </div>
        </>
    )
}