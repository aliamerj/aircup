import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { ThemeProvider } from "./components/theme/theme-provider.tsx";
import { AuthProvider } from "./contexts/auth/authContext.tsx";
import {
  createBrowserRouter,
  redirect,
  RouterProvider,
} from "react-router-dom";
import ErrorDisplay from "./components/errorDisplay/ErrorDisplay.tsx";
import { Register } from "./routes/Register.tsx";
import AuthWrapper from "./contexts/auth/AuthWrapper.tsx";
import { Toaster } from "./components/ui/toaster.tsx";
/// todo : remove AuthWrapper and use loader
// make api to save user auth
const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <AuthWrapper redirect="/register">
        <div>Hello world!, Dashboard</div>
      </AuthWrapper>
    ),
    errorElement: <ErrorDisplay message="404 Page not found" />,
  },
  {
    path: "/register",
    loader: async () => {
      const res = await fetch(import.meta.env.VITE_API_URL + "/auth/me", {
        method: "GET",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (res.ok) {
        return redirect("/");
      }
    },
    element: <Register />,
    errorElement: <ErrorDisplay message="Error. Please try again later." />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
      <AuthProvider>
        <Toaster />
        <RouterProvider router={router} />
      </AuthProvider>
    </ThemeProvider>
  </React.StrictMode>,
);
