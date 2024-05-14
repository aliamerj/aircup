import { AuthSchema } from "@/schema/user";
import { ReactNode, createContext, useState } from "react";
import { z } from "zod";

export type User = z.infer<typeof AuthSchema>;
export interface AuthContextType {
  login: (user: User) => void;
  logout: () => void;
  getSession: () => Promise<User | null>;
}

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [user, setUser] = useState<User | null>(null);

  const login = (userData: User) => {
    setUser(userData);
  };

  const logout = () => {
    setUser(null);
  };

  async function getSession(): Promise<User | null> {
    // Return the user if already available in context
    if (user) return user;
    const apiUrl = import.meta.env.VITE_API_URL + "/auth/me";
    try {
      const response = await fetch(apiUrl, {
        method: "GET",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        throw new Error("Failed to fetch user session");
      }

      const userData = await response.json();
      const data = AuthSchema.safeParse(userData.body);
      if (!data.success) {
        throw new Error("corrupted data");
      }
      setUser(data.data);
      return userData;
    } catch (__error: any) {
      return null;
    }
  }

  return (
    <AuthContext.Provider value={{ login, logout, getSession }}>
      {children}
    </AuthContext.Provider>
  );
};
