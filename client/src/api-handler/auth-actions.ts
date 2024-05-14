import { AuthSchema } from "@/schema/user";
import { redirect } from "react-router-dom";
import { z } from "zod";
export type User = z.infer<typeof AuthSchema>;
export async function getSession(): Promise<User | null> {
  // Check localStorage for user data
  const cachedUser = localStorage.getItem("user");
  if (!cachedUser) return null;
  // If user data is found, parse and return it
  const validate = AuthSchema.safeParse(JSON.parse(cachedUser));
  if (!validate.success) return null;

  // Make a request to verify the session token
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
      // If the token is invalid or expired, clear localStorage and return null
      localStorage.removeItem("user");
      return null;
    }

    // If the token is valid, refresh the localStorage data
    const userData = await response.json();
    const data = AuthSchema.safeParse(userData.body);
    if (!data.success) {
      throw new Error("Corrupted data");
    }

    localStorage.setItem("user", JSON.stringify(data.data));
    return data.data;
  } catch (__error) {
    localStorage.removeItem("user");
    return null;
  }
}

export async function logout(): Promise<boolean> {
  // Check localStorage for user data
  const cachedUser = localStorage.getItem("user");
  if (!cachedUser) return true;
  // If user data is found, parse and return it
  const validate = AuthSchema.safeParse(JSON.parse(cachedUser));
  if (!validate.success) return true;
  // Make a request to Logout
  const apiUrl = import.meta.env.VITE_API_URL + "/auth/logout";
  try {
    const response = await fetch(apiUrl, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (response.ok) {
      localStorage.removeItem("user");
      return true;
    }
    return false;
  } catch (__err) {
    return false;
  }
}

// If no user data is found in localStorage, return null

export function saveUserToLocalStorage(user: User) {
  localStorage.setItem("user", JSON.stringify(user));
}

export function clearSessionCache() {
  localStorage.removeItem("user");
}
