import { useRouteError } from "react-router-dom";

export default function ErrorPage() {
    const error = useRouteError();
    console.error(error);

    return (
        <div className="bg-gray-100 flex items-center justify-center h-screen">

            <div id="error-page" className="bg-white p-8 rounded-lg shadow-md text-center max-w-md">
                <h1 className="text-2xl font-bold text-red-500 mb-2">Ошибка</h1>
                <p className="text-gray-500 mb-4">
                    <i>{error.statusText || error.message}</i>
                </p>
                <a href="/" className="text-gray-700 mb-4 text-blue-400">Вернуться на главную</a>
            </div>

        </div>
    );
}