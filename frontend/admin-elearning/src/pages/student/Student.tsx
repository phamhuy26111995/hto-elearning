// import useUserStore from '@/store/user';
import React, { useEffect } from 'react';
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
import { Button } from '@/components/ui/button';
import { Link } from 'react-router';
import { Pencil, Trash2 } from 'lucide-react';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog';

export default function Student() {
  const { fetchStudents, students } = useStudentStore();

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
        console.log('🚀 ~ row:', row);
        const infoStudent = row.original; // or row.original.id if your object uses 'id'
        return (
          <div>
            {/* BUTTON_EDIT */}
            <Button className="text-blue-500 hover:pointer">
              <Link to={`/student/${infoStudent.userId}`}>
                <Pencil />
              </Link>
            </Button>

            {/* BUTTON_DELETE */}
            <AlertDialog>
              <AlertDialogTrigger asChild>
                <Button className="ml-2 text-red-500 hover:pointer">
                  <Trash2 />
                </Button>
              </AlertDialogTrigger>
              <AlertDialogContent>
                <AlertDialogHeader>
                  <AlertDialogTitle>
                    Bạn có chắc muốn xóa học sinh{' '}
                    <span className="font-bold text-red-500">
                      {infoStudent.username}
                    </span>{' '}
                    này không?
                  </AlertDialogTitle>
                  <AlertDialogDescription></AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel>Không</AlertDialogCancel>
                  <AlertDialogAction
                    onClick={() => onDelete(infoStudent.userId)}
                  >
                    Có
                  </AlertDialogAction>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
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

  // Function to handle deletion of a student
  const onDelete = async (id: number) => {
    try {
      console.log('Deleting student with ID:', id);
    } catch (error) {
      console.error('Failed to delete student:', error);
    }
  };

  useEffect(() => {
    fetchStudents();
  }, []);

  return (
    <React.Fragment>
      <div className="rounded-md border">
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
                  data-state={row.getIsSelected() && 'selected'}
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
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </React.Fragment>
  );
}
