"use client";

import { Button } from "@/components/ui/button";
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";
import { ILog } from "@/lib/types";
import { ChevronUp, ChevronDown } from "lucide-react";
import { useState } from "react";

interface Props {
  log: ILog;
}

export default function Log({ log }: Props) {
  const [isExpanded, setIsExpanded] = useState(false);

  const toggleExpand = () => setIsExpanded(!isExpanded);

  return (
    <Card className="mb-4">
      <CardHeader>
        <CardTitle className="text-sm font-medium">
          <span className="text-blue-600">
            {new Date(log.timestamp).toLocaleString()}
          </span>{" "}
          <span className="text-green-600 ml-2">{log.event}</span>
        </CardTitle>
      </CardHeader>
      <CardContent>
        <p className="text-sm text-gray-600">User ID: {log.user_id}</p>
        <p className="text-sm text-gray-600">Log ID: {log.id}</p>
        <div className="mt-2">
          <Button
            variant="outline"
            size="sm"
            onClick={toggleExpand}
            className="text-xs"
          >
            {isExpanded ? (
              <>
                <ChevronUp className="mr-2 h-4 w-4" />
                Hide Details
              </>
            ) : (
              <>
                <ChevronDown className="mr-2 h-4 w-4" />
                Show Details
              </>
            )}
          </Button>
        </div>
        {isExpanded && (
          <pre className="mt-2 p-2 bg-gray-100 rounded text-xs overflow-x-auto">
            {JSON.stringify(log.data, null, 2)}
          </pre>
        )}
      </CardContent>
    </Card>
  );
}
