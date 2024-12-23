import { getUserById } from "@/lib/data";
import UserProfile from "./user-profile";
import Link from "next/link";
import { ArrowLeft } from "lucide-react";

interface User {
  id: string;
  name: string;
  email: string;
  password: string;
  created_at: Date;
  updated_at: Date;
}

export default async function UserPage({
  params,
}: {
  params: { user_id: string };
}) {
  const user_id = params.user_id;
  const resp = await getUserById(user_id);
  const user = resp.data as User;

  return (
    <main className="w-full h-full">
      <Link
        href={"/users"}
        className="absolute border rounded-md hover:bg-blue-50 hover:text-blue-500 m-2"
      >
        <ArrowLeft />
      </Link>
      <UserProfile user={user} />
    </main>
  );
}
