import { useEffect, useState } from 'react';
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from '@tanstack/react-table';
import { User } from '@/types/user';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import useStudentStore from '@/store/student';
import { Pencil, Trash2 } from 'lucide-react';
import { useNavigate } from 'react-router';
import ButtonTooltip from '@/components/project/common/ButtonTooltip';
import DialogRemove from '@/components/project/common/DialogRemove';

export default function Student() {
  const navigate = useNavigate();
  const { fetchStudents, students, deleteStudent } = useStudentStore();
  const [selectedUserId, setSelectedUserId] = useState<number | null>(null);
  const [selectedUsername, setSelectedUsername] = useState<string>('');
  const [isDialogOpen, setIsDialogOpen] = useState<boolean>(false);

  const onEdit = (id: number) => {
    if (!id || !Number(id)) return;
    navigate(`/student/${id}`);
  };

  const onDeleteConfirm = (id: number, username: string) => {
    setSelectedUserId(id);
    setSelectedUsername(username);
    setIsDialogOpen(true);
  };

  const onCloseDialog = () => {
    setIsDialogOpen(false);
    setSelectedUserId(null);
    setSelectedUsername('');
  };

  const onDelete = async (id: number) => {
    try {
      console.log('Deleting student with ID:', id);
      await deleteStudent(id);
      onCloseDialog();
    } catch (error) {
      console.error('Failed to delete student:', error);
    }
  };

  const columns: ColumnDef<User>[] = [
    {
      accessorKey: 'numberNo',
      header: 'No.',
    },
    {
      accessorKey: 'userId',
      header: 'ID',
    },
    {
      accessorKey: 'username',
      header: 'Name',
    },
    {
      accessorKey: 'email',
      header: 'Email',
    },
    {
      accessorKey: 'role',
      header: 'Role',
    },
    {
      accessorKey: 'Action',
      header: 'Action',
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
    data: students,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  useEffect(() => {
    fetchStudents();
  }, []);

  return (
    <div className="rounded-md border">
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
                data-state={row.getIsSelected() && 'selected'}
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
  );
}
