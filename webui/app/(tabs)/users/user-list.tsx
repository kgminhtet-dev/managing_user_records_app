import { ScrollArea } from "@/components/ui/scroll-area";
import { Card } from "@/components/ui/card";
import Link from "next/link";
import { Separator } from "@radix-ui/react-separator";

interface User {
  id: string;
  name: string;
  email: string;
  password: string;
}

interface Props {
  users: User[];
}

export default async function UserList({ users }: Props) {
  users = [...users, ...users];
  return (
    <Card className="w-full h-max flex flex-col overflow-auto">
      <div className="grid grid-cols-4 gap-2  py-2 px-3 font-semibold">
        <div className="font-semibold">Name</div>
        <div className="font-semibold">Email</div>
        <div className="font-semibold">Password</div>
      </div>
      <Separator orientation="horizontal" />
      <ScrollArea className="h-full w-full rounded-md border p-2">
        {users.map((user, index) => (
          <Link
            key={index}
            // key={user.id}
            href={`/users/${user.id}`}
            className="text-sm grid grid-cols-4 gap-1 items-center rounded-md py-2 px-2 hover:bg-muted transition-colors"
          >
            <div>{user.name}</div>
            <div>{user.email}</div>
            <div>{user.password}</div>
          </Link>
        ))}
      </ScrollArea>
    </Card>
  );
}
