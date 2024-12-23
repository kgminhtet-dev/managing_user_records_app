"use client";

import { Button } from "@/components/ui/button";
import { deleteUser } from "@/lib/data";
import { Trash2 } from "lucide-react";

interface Props {
  id: string;
}

export default function UserDelete({ id }: Props) {
  return (
    <Button
      variant={"destructive"}
      className="w-8 h-8 rounded-md"
      onClick={async () => {
        await deleteUser(id);
      }}
      size={"icon"}
    >
      <Trash2 />
    </Button>
  );
}
