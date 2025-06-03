import React, { useEffect, useRef, useState } from 'react';
import { useParams } from 'react-router';
import useStudentStore from '@/store/student';
import { Input } from '@/components/ui/input';
import { useForm } from 'react-hook-form';
import { User } from '@/types/user';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { ROLES } from '@/consts/const';

export default function StudentDetail() {
  const { studentId: id } = useParams<{ studentId: string }>();
  const { fetchStudentDetail, studentDetail } = useStudentStore();
  const listRoles = ROLES;

  const {
    register,
    reset,
    formState: { errors },
    setValue, // ✅ để set role vào form nếu cần submit sau này
  } = useForm<User>();

  const [selectedRole, setSelectedRole] = useState<string>(''); // ✅ local state

  const hasFetchedRef = useRef(false);

  useEffect(() => {
    if (id && !hasFetchedRef.current) {
      hasFetchedRef.current = true;
      fetchStudentDetail(Number(id));
    }
  }, [id]);

  useEffect(() => {
    if (studentDetail) {
      reset(studentDetail);
      setSelectedRole(studentDetail.role); // ✅ gán giá trị role
      setValue('role', studentDetail.role); // ✅ set vào form nếu bạn dùng form submit
    }
  }, [studentDetail, reset, setValue]);

  return (
    <React.Fragment>
      <div className="p-5">
        <h1 className="scroll-m-20 text-2xl font-extrabold tracking-tight text-balance mb-2.5">
          Chi tiết học viên
        </h1>
        <div className="grid grid-cols-2 gap-5">
          {/* USERNAME */}
          <div>
            <label>Tên đăng nhập</label>
            <Input {...register('username')} placeholder='Nhập tên đăng nhập' />
            {errors.username && (
              <span className="text-red-500">Vui lòng nhập họ tên</span>
            )}
          </div>

          {/* PASSWORD */}
          <div>
            <label>Mật khẩu</label>
            <Input type="password" {...register('password')} placeholder='Nhập mật khẩu' />
            {errors.password && (
              <span className="text-red-500">Vui lòng nhập mật khẩu</span>
            )}
          </div>

          {/* EMAIL */}
          <div>
            <label>Email</label>
            <Input {...register('email')} placeholder='Nhập email' />
            {errors.email && (
              <span className="text-red-500">Vui lòng nhập email hợp lệ</span>
            )}
          </div>

          {/* ROLE */}
          <div>
            <label>Quyền</label>
            <Select
              value={selectedRole}
              onValueChange={(value) => {
                setSelectedRole(value);
                setValue('role', value); // ✅ cập nhật form
              }}
            >
              <SelectTrigger className="w-full">
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
          </div>
        </div>
      </div>
    </React.Fragment>
  );
}
