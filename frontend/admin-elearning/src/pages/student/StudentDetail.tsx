import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { ROLES } from "@/consts/const";
import studentServices from "@/services/student";
import { User } from "@/types/user";
import React, { useEffect } from "react";
import { Controller, useForm } from "react-hook-form";
import { useParams } from "react-router";

export default function StudentDetail() {
  const { studentId: id } = useParams<{ studentId: string }>();
  const listRoles = ROLES;

  const {
    register,
    formState: { errors },
    control,
    handleSubmit,
    reset, 
  } = useForm<User>({
    defaultValues: {
      role: "STUDENT",
      username: "",
      password: "",
      email: "",
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

    reset({...studentDetail});
  }

  function onSubmit(data: User) {
    console.log(data);
  }

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
            <Input {...register("username")} placeholder="Nhập tên đăng nhập" />
            {errors.username && (
              <span className="text-red-500">Vui lòng nhập họ tên</span>
            )}
          </div>

          {/* PASSWORD */}
          <div>
            <label>Mật khẩu</label>
            <Input
              type="password"
              {...register("password")}
              placeholder="Nhập mật khẩu"
            />
            {errors.password && (
              <span className="text-red-500">Vui lòng nhập mật khẩu</span>
            )}
          </div>

          {/* EMAIL */}
          <div>
            <label>Email</label>
            <Input {...register("email")} placeholder="Nhập email" />
            {errors.email && (
              <span className="text-red-500">Vui lòng nhập email hợp lệ</span>
            )}
          </div>

          {/* ROLE */}
          <div>
            <label>Quyền</label>
            <Controller
              name="role"
              control={control}
              render={({ field }) => (
                <Select value={field.value} onValueChange={field.onChange}>
                  <SelectTrigger>
                    <SelectValue placeholder="Select a role" />
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
          </div>

          <div>
            <Button onClick={handleSubmit(onSubmit)}>Submit</Button>
          </div>
        </div>
      </div>
    </React.Fragment>
  );
}
