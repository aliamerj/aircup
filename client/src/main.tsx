import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { ThemeProvider } from "./components/theme/theme-provider.tsx";
import {
  createBrowserRouter,
  redirect,
  RouterProvider,
} from "react-router-dom";
import ErrorDisplay from "./components/errorDisplay/ErrorDisplay.tsx";
import { Register } from "./routes/Register.tsx";
import { Toaster } from "./components/ui/toaster.tsx";
import { Login } from "./routes/login.tsx";
import { getSession } from "./api-handler/auth-actions.ts";
import { Dashboard } from "./routes/Dashboard.tsx";
import { getSavedDisks } from "./api-handler/disk-actions.ts";

const router = createBrowserRouter([
  {
    path: "/",
    loader: async () => {
      const user = await getSession();
      if (!user) {
        throw redirect("/login");
      }
      const disks = await getSavedDisks();
      return disks;
    },
    element: <Dashboard />,
    errorElement: <ErrorDisplay message="404 Page not found" />,
  },
  {
    path: "/register",
    loader: async () => {
      const user = await getSession();
      if (user) {
        throw redirect("/");
      }
      return "register";
    },
    element: <Register />,
    errorElement: <ErrorDisplay message="Error. Please try again later." />,
  },
  {
    path: "/login",
    loader: async () => {
      const user = await getSession();
      if (user) {
        throw redirect("/");
      }
      return "login";
    },
    element: <Login />,
    errorElement: <ErrorDisplay message="Error. Please try again later." />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <Toaster />
      <RouterProvider router={router} />
    </ThemeProvider>
  </React.StrictMode>,
);
