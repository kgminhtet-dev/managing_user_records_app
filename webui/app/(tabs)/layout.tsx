import Link from "next/link";

export default function AppLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <main className={"w-screen h-screen grid grid-flow-row grid-rows-12"}>
      <div className="row-span-1 flex justify-start items-center p-1 bg-primary text-white">
        <p>Admin Panel</p>
      </div>
      <div className="row-span-11 grid grid-flow-col grid-cols-12">
        <nav className="text-sm col-span-2 border-r flex flex-col justify-start p-1">
          <Link
            className="p-1 pr-3 pl-3 border bg-gray-50 rounded-md border-white hover:rounded-3xl hover:bg-blue-50 hover:text-blue-600"
            href={"/users"}
          >
            Users
          </Link>
          <Link
            className="p-1 pr-3 pl-3 border bg-gray-50 rounded-md border-white hover:rounded-3xl hover:bg-blue-50 hover:text-blue-600"
            href={"/logs"}
          >
            Logs
          </Link>
        </nav>
        <div className="col-span-10 p-1 overflow-auto">{children}</div>
      </div>
    </main>
  );
}
