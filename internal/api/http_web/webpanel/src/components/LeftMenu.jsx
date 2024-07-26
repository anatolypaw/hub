export default function LeftMenu() {
    return (
        <nav className="bg-white shadow-md w-52 p-2">
            <ul>
                <li className="mb-2">
                    <a href="/c1" className="text-gray-700 hover:bg-blue-100 block py-2 px-4 rounded">
                        Главная
                    </a>
                </li>
                <li className="mb-2">
                    <a href="/c2" className="text-gray-700 hover:bg-blue-100 block py-2 px-4 rounded">
                        О нас
                    </a>
                </li>
                <li className="mb-2">
                    <a href="/c3" className="text-gray-700 hover:bg-blue-100 block py-2 px-4 rounded">
                        Услуги
                    </a>
                </li>
                <li className="mb-2">
                    <a href="/login" className="text-gray-700 hover:bg-blue-100 block py-2 px-4 rounded">
                        Контакты
                    </a>
                </li>
                <li className="mb-2">
                    <a href="#" className="text-gray-700 hover:bg-blue-100 block py-2 px-4 rounded">
                        Блог
                    </a>
                </li>
            </ul>
        </nav>
    )
}