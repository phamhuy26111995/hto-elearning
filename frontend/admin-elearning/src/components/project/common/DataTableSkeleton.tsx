import { Skeleton } from "@/components/ui/skeleton";

export function DataTableSkeleton({
  rows = 10,
  columns = ["50px", "80px", "220px", "280px", "120px", "100px"],
  withToolbar = true,
}: {
  rows?: number;
  columns?: (string | number)[]; // widths per column
  withToolbar?: boolean;
}) {
  return (
    <div className="w-full">
      {withToolbar && (
        <div className="mb-3 flex items-center justify-between gap-3">
          <div className="flex items-center gap-2">
            <Skeleton className="h-9 w-48 rounded-xl" />
            <Skeleton className="h-9 w-28 rounded-xl" />
          </div>
          <div className="flex items-center gap-2">
            <Skeleton className="h-9 w-28 rounded-xl" />
            <Skeleton className="h-9 w-9 rounded-xl" />
            <Skeleton className="h-9 w-9 rounded-xl" />
          </div>
        </div>
      )}

      <div className="overflow-hidden rounded-2xl border border-border/40">
        {/* Header */}
        <div
          className="grid w-full bg-muted/40 px-4 py-2"
          style={{ gridTemplateColumns: columns.join(" ") }}
        >
          {columns.map((w, i) => (
            <Skeleton key={i} className="h-4 w-[80%]" />
          ))}
        </div>

        {/* Body */}
        <div className="divide-y divide-border/40">
          {Array.from({ length: rows }).map((_, r) => (
            <div
              key={r}
              className="grid items-center px-4 py-3"
              style={{ gridTemplateColumns: columns.join(" ") }}
            >
              {columns.map((_, c) => (
                <Skeleton key={c} className="h-4 w-[70%]" />
              ))}
            </div>
          ))}
        </div>
      </div>

      {/* Pagination row */}
      <div className="mt-4 flex items-center justify-end gap-2">
        <Skeleton className="h-9 w-9 rounded-xl" />
        <Skeleton className="h-9 w-9 rounded-xl" />
        <Skeleton className="h-9 w-16 rounded-xl" />
        <Skeleton className="h-9 w-9 rounded-xl" />
        <Skeleton className="h-9 w-9 rounded-xl" />
      </div>
    </div>
  );
}
