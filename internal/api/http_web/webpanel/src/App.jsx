import { useState } from 'react'
import './App.css'
import Header from './components/Header'
import LeftMenu from './components/LeftMenu'

export default function App() {
    return (
        <>
            <Header />
            <div className="flex h-[calc(100vh-3rem)]">
                <LeftMenu />

                {/* Здесь размещается основной контент страницы */}
                <div id="content" className="flex-1 bg-white shadow-md p-4 ml-2 overflow-auto">
                    <p>Основной контент страницы...</p>
                    <p>Основной контент страницы...</p>
                    <p>Основной контент страницы...</p>
                    <p>Основной контент страницы...</p>
                    <p>Основной контент страницы...</p>
                    <p>Основной контент страницы...</p>
                    <p>Основной контент страницы...</p>
                </div>
            </div >
        </>
    )
}
