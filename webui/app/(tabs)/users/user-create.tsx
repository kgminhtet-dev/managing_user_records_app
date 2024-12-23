"use client";

import React from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Button } from "@/components/ui/button";
import { Label } from "@radix-ui/react-label";
import { PlusIcon } from "lucide-react";
import { Input } from "@/components/ui/input";
import { DialogHeader, DialogFooter } from "@/components/ui/dialog";
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { ICreateUser } from "@/lib/types";
import { createUser } from "@/lib/data";

const createUserFormschema = z.object({
  name: z.string().min(1, "Name is required"),
  email: z.string().min(1, "Email is required").email("Invalid email address"),
  password: z.string().min(6, "Password must be at least 6 characters long"),
});

export default function CreateUserDialog() {
  const {
    reset,
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(createUserFormschema),
  });

  const onSubmit = async (values: any) => {
    const { ok, message } = await createUser(values as ICreateUser);
    if (ok) {
      reset();
      alert(message);
    } else {
      alert(message);
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="absolute right-2" variant="outline">
          Create user
          <PlusIcon className="ml-2" />
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Edit profile</DialogTitle>
          <DialogDescription>
            Create a new user. Click save when you&apos;re done.
          </DialogDescription>
        </DialogHeader>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="grid gap-4 py-4">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="name" className="text-right">
                Name
              </Label>
              <Input
                id="name"
                placeholder="Name"
                className="col-span-3"
                {...register("name")}
              />
              {errors.name && (
                <span className="col-span-4 text-red-500">
                  {errors.name.message}
                </span>
              )}
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="email" className="text-right">
                Email
              </Label>
              <Input
                type="email"
                id="email"
                placeholder="Email"
                className="col-span-3"
                {...register("email")}
              />
              {errors.email && (
                <span className="col-span-4 text-red-500">
                  {errors.email.message}
                </span>
              )}
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="password" className="text-right">
                Password
              </Label>
              <Input
                id="password"
                placeholder="Password"
                type="password"
                className="col-span-3"
                {...register("password")}
              />
              {errors.password && (
                <span className="col-span-4 text-red-500">
                  {errors.password.message}
                </span>
              )}
            </div>
          </div>
          <DialogFooter>
            <Button type="submit">Create</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
