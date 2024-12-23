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
import { login } from "@/lib/data";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

const formSchema = z.object({
  email: z.string().email("Invalid email address."),
  password: z.string().min(8, "Password must be at least 8 characters long."),
});

export default function SignInForm() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    await login(values);
  };

  return (
    <main className="w-screen h-screen flex justify-center items-center">
      <div className="w-1/4 h-auto flex flex-col justify-between items-center border rounded-lg p-3 shadow-md border-blue-300">
        <p className="font-bold text-3xl">Login</p>
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="w-full space-y-2"
          >
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input type={"email"} placeholder="" {...field} />
                  </FormControl>
                  {/*<FormDescription></FormDescription>*/}
                  <div className={"h-3"}>
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
                    <Input type={"password"} placeholder="" {...field} />
                  </FormControl>
                  {/*<FormDescription></FormDescription>*/}
                  <div className={"h-3"}>
                    <FormMessage />
                  </div>
                </FormItem>
              )}
            />
            <Button className={"w-full"} type="submit">
              Login
            </Button>
          </form>
        </Form>
      </div>
    </main>
  );
}
