"use client";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { updateUser } from "@/lib/data";
import { IUpdateUser } from "@/lib/types";

const updateFormSchema = z.object({
  id: z.string(),
  name: z.string().min(1, "Name is required."),
  email: z.string().email("Invalid email address."),
  password: z.string().min(8, "Password must be at least 8 characters long."),
  created_at: z.string(),
  updated_at: z.string(),
});

interface User {
  id: string;
  name: string;
  email: string;
  password: string;
  created_at: Date;
  updated_at: Date;
}

interface Props {
  user: User;
}

export default function UserProfile({ user }: Props) {
  const form = useForm<z.infer<typeof updateFormSchema>>({
    resolver: zodResolver(updateFormSchema),
    defaultValues: {
      id: user.id,
      name: user.name,
      email: user.email,
      password: user.password,
      created_at: new Date(user.created_at).toLocaleString(),
      updated_at: new Date(user.updated_at).toLocaleString(),
    },
  });

  const onSubmit = async (values: z.infer<typeof updateFormSchema>) => {
    await updateUser(values as IUpdateUser);
  };

  return (
    <div className="w-full h-full flex flex-col justify-between border rounded-lg p-3 shadow-md border-blue-300">
      <p className="font-bold text-2xl text-center">User Information</p>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="w-1/2">
          <FormField
            control={form.control}
            name="id"
            render={({ field }) => (
              <FormItem>
                <FormLabel>ID</FormLabel>
                <FormControl>
                  <Input type={"text"} {...field} value={user.id} disabled />
                </FormControl>
                <div>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Name</FormLabel>
                <FormControl>
                  <Input type={"text"} {...field} />
                </FormControl>
                <div>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input type={"email"} {...field} />
                </FormControl>
                <div>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input type={"password"} disabled {...field} />
                </FormControl>
                <div>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />
          <div className="flex flex-row justify-start gap-1">
            <FormField
              control={form.control}
              name="created_at"
              render={({ field }) => (
                <FormItem className="w-full">
                  <FormLabel>Created At</FormLabel>
                  <FormControl>
                    <Input type={"text"} disabled {...field} />
                  </FormControl>
                  <div>
                    <FormMessage />
                  </div>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="updated_at"
              render={({ field }) => (
                <FormItem className="w-full">
                  <FormLabel>Updated At</FormLabel>
                  <FormControl>
                    <Input type={"text"} disabled {...field} />
                  </FormControl>
                  <div>
                    <FormMessage />
                  </div>
                </FormItem>
              )}
            />
          </div>
          <Button type="submit">Save</Button>
        </form>
      </Form>
    </div>
  );
}
