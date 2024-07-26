import React from 'react'
import ReactDOM from 'react-dom/client'
import {
    createBrowserRouter,
    RouterProvider,
} from "react-router-dom";
import ErrorPage from './pages/error-page.jsx';
import Login from './pages/login.jsx';
import Content1 from './pages/content1.jsx';
import Content2 from './pages/content2.jsx';
import Content3 from './pages/content3.jsx';


import App from './App.jsx'



const router = createBrowserRouter([
    {
        path: "/",
        element: <App />,
        errorElement: <ErrorPage />,
    },
    {
        path: "/login",
        element: <Login />,
    },
    {
        path: "/с1",
        element: <Content1 />,
    },
    {
        path: "/с2",
        element: <Content2 />,
    },
    {
        path: "/с3",
        element: <Content3 />,
    },

]);


ReactDOM.createRoot(document.getElementById('root')).render(
    <React.StrictMode>
        <RouterProvider router={router} />
    </React.StrictMode>,
)
