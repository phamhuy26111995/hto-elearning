import ButtonTooltip from "@/components/project/common/ButtonTooltip";
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
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { Pencil, Trash2 } from "lucide-react";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";

export default function Users() {
  const navigate = useNavigate();
  const [dataList, setDataList] = useState<User[]>([]);
  const [totalRows, setTotalRows] = useState<number>(0);
  const [pageNumber, setPageNumber] = useState<number>(1);
  const [rowsPerPage, setRowsPerPage] = useState<number>(25);
  const [selectedUserId, setSelectedUserId] = useState<number | null>(null);
  const [selectedUsername, setSelectedUsername] = useState<string>("");
  const [isDialogOpen, setIsDialogOpen] = useState<boolean>(false);

  // const totalPage = Math.ceil(totalRows / rowsPerPage);

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
    fetchUsers(page);
  }

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

  async function fetchUsers(pageNumber : number) {
    const response = await userServices.getAllStudent({
      pageNumber: pageNumber,
      rowsPerPage: rowsPerPage,
    });
    const data = response.data as User[];

    setDataList(
      data.map((item, index) => ({
        ...item,
        numberNo: (pageNumber - 1) * rowsPerPage + 1 + index,
      }))
    );
    if (data.length > 0) {
      setTotalRows(data[0].totalRows);
    }
  }

  useEffect(() => {
    fetchUsers(pageNumber);
  }, []);

  return (
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
  );
}
