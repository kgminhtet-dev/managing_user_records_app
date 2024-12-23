"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { IUpdateUser } from "./types";
import { revalidatePath } from "next/cache";

const uri = process.env.URI;

async function setCookie(key: string, value: string) {
  const cookie = await cookies();
  cookie.set(key, value, {
    httpOnly: true,
    secure: true,
    sameSite: "none",
    maxAge: 60 * 60 * 24,
  });
}

async function getCookie(key: string) {
  const cookie = await cookies();
  return cookie.get(key);
}

async function login(values: { email: string; password: string }) {
  try {
    const resp = await fetch(`${uri}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(values),
    });

    if (!resp.ok) {
      throw new Error(`HTTP error! status: ${resp.status}`);
    }

    const { user_id, _token } = await resp.json();
    if (user_id && _token) {
      await setCookie("user_id", user_id);
      await setCookie("_token", _token);
    }
    redirect("/users");
  } catch (err) {
    console.error("Failed to login:", err);
    throw err;
  }
}

async function getUsers(page: string) {
  const token = await getCookie("_token");
  try {
    const resp = await fetch(`${uri}/api/v1/users?page=${page}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token?.value}`,
      },
    });

    if (!resp.ok) {
      throw new Error(`HTTP error! status: ${resp.status}`);
    }
    const json = await resp.json();
    return json;
  } catch (err) {
    console.error("Failed to fetch users:", err);
    throw err;
  }
}

async function getUserById(id: string) {
  const token = await getCookie("_token");
  try {
    const resp = await fetch(`${uri}/api/v1/users/${id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token?.value}`,
      },
    });

    if (!resp.ok) {
      throw new Error(`HTTP error! status: ${resp.status}`);
    }
    const json = await resp.json();
    return json;
  } catch (err) {
    console.error("Failed to fetch users:", err);
    throw err;
  }
}

async function updateUser(userData: IUpdateUser) {
  const token = await getCookie("_token");
  try {
    const resp = await fetch(`${uri}/api/v1/users/${userData.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token?.value}`,
      },
      body: JSON.stringify({
        id: userData.id,
        name: userData.name,
        email: userData.email,
      }),
    });

    if (!resp.ok) {
      console.error(await resp.json());
      return;
    }

    revalidatePath(`/users/${userData.id}`);
  } catch (err) {
    console.error("Failed to login:", err);
    throw err;
  }
}

async function deleteUser(id: string) {
  const token = await getCookie("_token");
  try {
    const resp = await fetch(`${uri}/api/v1/users/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token?.value}`,
      },
    });

    if (!resp.ok) {
      throw new Error(`HTTP error! status: ${resp.status}`);
    }

    revalidatePath("/users");
  } catch (err) {
    console.error("Failed to login:", err);
    throw err;
  }
}

export { getUsers, getUserById, updateUser, deleteUser, login };
