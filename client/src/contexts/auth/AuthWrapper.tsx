import { ReactNode, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "./authHooks";
import { User } from "./authContext";
import { Loader } from "@/components/loader/Loader";

const AuthWrapper = ({
  children,
  redirect,
}: {
  children: ReactNode;
  redirect: string;
}) => {
  const [user, setUser] = useState<User>();
  const { getSession } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    const checkAuth = async () => {
      const fetchedUser = await getSession();
      if (!fetchedUser) {
        return navigate(redirect);
      }
      setUser(fetchedUser);
    };

    checkAuth();
  }, [navigate, getSession, redirect]);

  if (!user) {
    return <Loader />;
  }

  return <>{children}</>;
};

export default AuthWrapper;
