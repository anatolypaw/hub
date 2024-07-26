import dm_logo from '/logo.svg'

export default function Header() {
    const now = new Date()
    return (
        <header className="bg-white shadow-md py-1 mb-2">
            <div className="container mx-auto flex justify-between items-center">
                {/* Логотип */}
                <div className="flex items-center">
                    <img src={dm_logo} alt="Logo" className="w-7 h-7"></img>
                </div>
                <span className="flex items-center align-left">Управление маркировкой</span>
                {/* Кнопка выхода и отображение текущего пользователя */}
                <div className="flex items-center">
                    <span className="text-gray-700 mr-4">Администратор</span>
                    <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-4 rounded">
                        Выйти
                    </button>
                </div>
            </div>
        </header>
    )
}