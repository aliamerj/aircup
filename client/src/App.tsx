import "@/index.css";
import { LoginForm } from "./components/sign-up/LoginForm";
import { fetcher } from "./api-handler";
import useSWR from "swr";
import { ConfigRes } from "./api-handler/types/Config";
import { Loader } from "./components/loader/Loader";
import ErrorDisplay from "./components/errorDisplay/ErrorDisplay";
function App() {
  const apiUrl = import.meta.env.VITE_API_URL + "/config";
  const {
    data: configData,
    error,
    isLoading,
  } = useSWR(apiUrl, fetcher<ConfigRes>);
  if (error)
    return (
      <ErrorDisplay message="Failed to load data. Please try again later." />
    );

  if (isLoading)
    return (
      <div className="flex h-screen items-center justify-center">
        <Loader />
      </div>
    );
  if (!configData || configData.type === "NO_ACCOUNT")
    return (
      <div className="flex h-screen items-center justify-center">
        <LoginForm />
      </div>
    );

  return (
    <div>
      <h1>dashboard</h1>
    </div>
  );
}

export default App;
