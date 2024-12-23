import UserList from "@/app/(tabs)/users/user-list";
import { getUsers } from "@/lib/data";
import CreateUserDialog from "./user-create";

interface User {
  id: string;
  name: string;
  email: string;
  password: string;
}

export default async function UserPage() {
  const resp = await getUsers("1");
  const users = resp.data as User[];

  return (
    <div className="w-full h-full flex flex-col gap-1">
      <CreateUserDialog />
      <UserList users={users} />
    </div>
  );
}
