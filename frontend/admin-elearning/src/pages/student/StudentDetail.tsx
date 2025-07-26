import ErrorMessage from '@/components/project/common/ErrorMessage';
import { RequiredLabel } from '@/components/project/common/RequiredLabel';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { DEFAULT_TIMEOUT, ROLES } from '@/consts/const';
import studentServices from '@/services/student';
import { User } from '@/types/user';
import React, { useEffect } from 'react';
import { Controller, useForm } from 'react-hook-form';
import { useParams, useLocation, useNavigate } from 'react-router';
import { toast } from 'sonner';
import { Toaster } from '@/components/ui/sonner'

export default function StudentDetail() {
  const location = useLocation();
  const navigate = useNavigate();
  const { studentId: id } = useParams<{ studentId: string }>();
  const isAddMode = location.pathname.includes('/add');
  const isEditMode = location.pathname.includes('/edit');
  const listRoles = ROLES;

  const pageTitle = isAddMode
    ? 'Thêm mới học viên'
    : isEditMode
    ? 'Cập nhật học viên'
    : 'Chi tiết học viên';

  const {
    register,
    formState: { errors },
    control,
    handleSubmit,
    reset,
  } = useForm<User>({
    defaultValues: {
      role: 'STUDENT',
      username: '',
      password: '',
      email: '',
    },
  });

  useEffect(() => {
    if (!id || !Number(id)) {
      return;
    }

    getData();
  }, [id]);

  async function getData() {
    const response = await studentServices.getById(Number(id));

    const data = response.data;

    const studentDetail = data.user as User;

    reset({ ...studentDetail });
  }

  const onGoBack = () => {
    reset();
    navigate(`/student`);
  }

  const onSubmit = async (data: User) => {
    try {
      if (isAddMode) {
        await studentServices.createStudent(data);
        toast.success('Thêm mới học viên thành công');
      } else if (isEditMode) {
        data.userId = Number(id);
        await studentServices.updateStudent(data);
        toast.success('Cập nhật học viên thành công');
      }

      setTimeout(() => {
        navigate(`/student`);
      }, DEFAULT_TIMEOUT);
    } catch (error) {
      console.error('Error saving student:', error);
      toast.error('Có lỗi xảy ra khi lưu học viên');
    }
  }

  return (
    <React.Fragment>
      <div className="p-5">
        <h1 className="scroll-m-20 text-2xl font-extrabold tracking-tight text-balance mb-2.5">
          {pageTitle}
        </h1>
        <div className="grid grid-cols-2 gap-5">
          {/* USERNAME */}
          <div>
            <RequiredLabel htmlFor="username">Tên đăng nhập</RequiredLabel>
            <Input
              {...register('username', {
                required: 'Vui lòng nhập tên đăng nhập',
              })}
              aria-invalid={!!errors.username}
              placeholder="Nhập tên đăng nhập"
            />
            <ErrorMessage message={errors.username?.message} />
          </div>

          {/* PASSWORD */}
          <div>
            <RequiredLabel htmlFor="password">Mật khẩu</RequiredLabel>
            <Input
              type="password"
              {...register('password', {
                required: 'Vui lòng nhập mật khẩu',
              })}
              aria-invalid={!!errors.password}
              placeholder="Nhập mật khẩu"
            />
            <ErrorMessage message={errors.password?.message} />
          </div>

          {/* EMAIL */}
          <div>
            <RequiredLabel htmlFor="email">Email</RequiredLabel>
            <Input
              {...register('email', {
                required: 'Vui lòng nhập email',
                pattern: {
                  value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
                  message: 'Email không hợp lệ',
                },
              })}
              aria-invalid={!!errors.email}
              placeholder="Nhập email"
            />
            <ErrorMessage message={errors.email?.message} />
          </div>

          {/* ROLE */}
          <div>
            <RequiredLabel htmlFor="role">Quyền</RequiredLabel>
            <Controller
              name="role"
              control={control}
              rules={{ required: 'Vui lòng chọn quyền' }}
              render={({ field }) => (
                <Select value={field.value} onValueChange={field.onChange}>
                  <SelectTrigger
                    aria-invalid={!!errors.role}
                    className="w-full"
                  >
                    <SelectValue placeholder="Chọn quyền" />
                  </SelectTrigger>
                  <SelectContent>
                    {listRoles.map((role) => (
                      <SelectItem key={role.value} value={role.value}>
                        {role.name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              )}
            />
            <ErrorMessage message={errors.role?.message} />
          </div>

          <div className="col-span-2 flex justify-end gap-2">
            <Button variant="outline" className='hover:cursor-pointer' onClick={() => onGoBack()}>
              Hủy
            </Button>
            <Button className='hover:cursor-pointer' onClick={handleSubmit(onSubmit)}>Lưu</Button>
          </div>
        </div>
      </div>
      <Toaster position="top-right" richColors />
    </React.Fragment>
  );
}
