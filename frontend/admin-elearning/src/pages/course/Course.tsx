import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { ROUTES } from "@/consts/const";
import courseServices from "@/services/course";
import { Course as CourseType } from "@/types/course";
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  Row,
  // RowData,
  useReactTable,
} from "@tanstack/react-table";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";

const columns: ColumnDef<CourseType>[] = [
  {
    accessorKey: "numberNo",
    header: "No.",
  },
  {
    accessorKey: "courseId",
    header: "ID",
  },
  {
    accessorKey: "title",
    header: "Title",
  },
  {
    accessorKey: "description",
    header: "Description",
  },
  {
    accessorKey: "Action",
    header: "Action",
    cell: ({ row }) => <ActionCell row={row} />,
  },
];

export default function Course() {
  const [courses, setCourses] = useState<CourseType[]>([]);
  const table = useReactTable({
    data: courses,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  // useEffect(() => {
  //   fetchCourse();
  // }, []);

  async function fetchCourse() {
    const response = await courseServices.getAll();
    const data = response.data as CourseType[];

    setCourses(
      data.map((item: CourseType, index: number) => ({
        ...item,
        numberNo: index + 1,
      }))
    );
  }

  return (
    <div className="p-5">
      <Table>
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => {
                return (
                  <TableHead key={header.id}>
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                  </TableHead>
                );
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map((row) => (
              <TableRow
                key={row.id}
                data-state={row.getIsSelected() && "selected"}
              >
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}

interface ActionCellProps {
  row: Row<CourseType>;
}

function ActionCell({ row }: ActionCellProps) {
  const navigate = useNavigate();

  function onClickAction() {
    navigate(`${ROUTES.COURSE}/edit/${row.original.courseId}`);
  }

  return <button onClick={onClickAction}>
    View
  </button>
}
