import ButtonTooltip from "@/components/project/common/ButtonTooltip";
import { DataTableSkeleton } from "@/components/project/common/DataTableSkeleton";
import DialogRemove from "@/components/project/common/DialogRemove";
import Paging from "@/components/project/common/Paging";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { ROUTES } from "@/consts/const";
import userServices from "@/services/user";
import { User } from "@/types/user";
import { useQuery } from "@tanstack/react-query";
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { Pencil, Trash2 } from "lucide-react";
import { useEffect, useMemo, useState } from "react";
import { useNavigate } from "react-router";

export default function Users() {
  const navigate = useNavigate();
  const [pageNumber, setPageNumber] = useState<number>(1);
  const [rowsPerPage, setRowsPerPage] = useState<number>(10);
  const [selectedUserId, setSelectedUserId] = useState<number | null>(null);
  const [selectedUsername, setSelectedUsername] = useState<string>("");
  const [isDialogOpen, setIsDialogOpen] = useState<boolean>(false);

  const { data, isPending } = useQuery<User[]>({
    queryKey: ["users", { pageNumber, rowsPerPage }],
    queryFn: () => fetchUsers(pageNumber, rowsPerPage),
    staleTime: 1000 * 60,
    gcTime: 1000 * 60 * 3,
  });

  const onEdit = (id: number) => {
    if (!id || !Number(id)) return;
    navigate(`${ROUTES.USER}/edit/${id}`);
  };

  const onDeleteConfirm = (id: number, username: string) => {
    setSelectedUserId(id);
    setSelectedUsername(username);
    setIsDialogOpen(true);
  };

  const onCloseDialog = () => {
    setIsDialogOpen(false);
    setSelectedUserId(null);
    setSelectedUsername("");
  };

  function onChangePage(page: number) {
    setPageNumber(page);
  }

  const dataList = useMemo(
    () =>
      (data ?? []).map((item, idx) => ({
        ...item,
        numberNo: (pageNumber - 1) * rowsPerPage + idx + 1,
      })),
    [data, pageNumber, rowsPerPage]
  );

  const onDelete = async (id: number) => {
    try {
      onCloseDialog();
    } catch (error) {
      console.error("Failed to delete student:", error);
    }
  };

  const columns: ColumnDef<User>[] = [
    {
      accessorKey: "numberNo",
      header: "No.",
    },
    {
      accessorKey: "userId",
      header: "ID",
    },
    {
      accessorKey: "username",
      header: "Name",
    },
    {
      accessorKey: "email",
      header: "Email",
    },
    {
      accessorKey: "role",
      header: "Role",
    },
    {
      accessorKey: "Action",
      header: "Action",
      cell: ({ row }) => {
        const infoStudent = row.original;
        return (
          <div className="flex items-center gap-2">
            {/* EDIT */}
            <ButtonTooltip
              tooltip="Sửa"
              icon={<Pencil size={16} className="text-blue-500" />}
              onClick={() => onEdit(infoStudent.userId)}
            />

            {/* DELETE */}
            <ButtonTooltip
              tooltip="Xóa"
              icon={<Trash2 size={16} className="text-red-500" />}
              onClick={() =>
                onDeleteConfirm(infoStudent.userId, infoStudent.username)
              }
            />
          </div>
        );
      },
    },
  ];

  const table = useReactTable({
    data: dataList,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  async function fetchUsers(pageNumber: number, rowsPerPage: number) {
    const response = await userServices.getAllStudent({
      pageNumber: pageNumber,
      rowsPerPage: rowsPerPage,
    });
    const data = response.data as User[];

    const result = data.map((item, index) => ({
      ...item,
      numberNo: (pageNumber - 1) * rowsPerPage + 1 + index,
    }));

    return result;
  }

  const totalRows = useMemo(() => data?.[0]?.totalRows ?? 0, [data]);

  return (
    <>
      {isPending ? (
        <DataTableSkeleton />
      ) : (
        <>
          <div className="rounded-md border p-4">
            <Table>
              <TableHeader>
                {table.getHeaderGroups().map((headerGroup) => (
                  <TableRow key={headerGroup.id}>
                    {headerGroup.headers.map((header) => (
                      <TableHead key={header.id}>
                        {header.isPlaceholder
                          ? null
                          : flexRender(
                              header.column.columnDef.header,
                              header.getContext()
                            )}
                      </TableHead>
                    ))}
                  </TableRow>
                ))}
              </TableHeader>
              <TableBody>
                {table.getRowModel().rows.length ? (
                  table.getRowModel().rows.map((row) => (
                    <TableRow
                      key={row.id}
                      data-state={row.getIsSelected() && "selected"}
                    >
                      {row.getVisibleCells().map((cell) => (
                        <TableCell key={cell.id}>
                          {flexRender(
                            cell.column.columnDef.cell,
                            cell.getContext()
                          )}
                        </TableCell>
                      ))}
                    </TableRow>
                  ))
                ) : (
                  <TableRow>
                    <TableCell
                      colSpan={columns.length}
                      className="h-24 text-center"
                    >
                      Không có kết quả.
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>

            {/* DIALOG_DELETE */}
            <DialogRemove
              open={isDialogOpen}
              onOpenChange={setIsDialogOpen}
              username={selectedUsername}
              onConfirm={() => {
                if (selectedUserId !== null) {
                  onDelete(selectedUserId);
                }
              }}
              title="Bạn có chắc muốn xóa học sinh này không?"
            />
          </div>
          <div className="mt-4">
            <Paging
              totalRows={totalRows}
              rowsPerPage={rowsPerPage}
              currentPage={pageNumber}
              onChange={onChangePage} // TODO;
            />
          </div>
        </>
      )}
    </>
  );
}
