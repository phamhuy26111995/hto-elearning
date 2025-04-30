import React from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

type FormData = {
  username: string;
  password: string;
};

const LoginForm: React.FC = () => {
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>();

  const onSubmit = (data: FormData) => {
    const { username, password } = data;
    if (username === "admin" && password === "123456") {
      navigate("/home");
    } else {
      alert("Sai tài khoản hoặc mật khẩu");
    }
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="bg-white p-6 rounded shadow-md w-96 space-y-4"
    >
      <h2 className="text-2xl font-semibold text-center">Đăng nhập</h2>

      <div>
        <Input
          placeholder="Tên đăng nhập"
          {...register("username", { required: "Vui lòng nhập tên đăng nhập" })}
        />
        {errors.username && (
          <p className="text-red-500 text-sm mt-1">{errors.username.message}</p>
        )}
      </div>

      <div>
        <Input
          type="password"
          placeholder="Mật khẩu"
          {...register("password", { required: "Vui lòng nhập mật khẩu" })}
        />
        {errors.password && (
          <p className="text-red-500 text-sm mt-1">{errors.password.message}</p>
        )}
      </div>

      <Button type="submit" className="w-full">
        Đăng nhập
      </Button>
    </form>
  );
};

export default LoginForm;
