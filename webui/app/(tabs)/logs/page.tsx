import { ScrollArea } from "@/components/ui/scroll-area";
import { getLogs } from "@/lib/data";
import { ILog } from "@/lib/types";
import Log from "./log";

export default async function LogPage() {
  const logs = (await getLogs("1")) as ILog[];
  return (
    <main className="w-full h-full flex flex-col gap-1">
      <ScrollArea>
        {logs.map((log) => (
          <Log key={log.id} log={log} />
        ))}
      </ScrollArea>
    </main>
  );
}
