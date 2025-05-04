import { apiService } from "@/api/apiService";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { cn } from "@/lib/utils";
import useUserStore from "@/store/user";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { toast } from "sonner";

type FormData = {
  username: string;
  password: string;
};

export function LoginForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  // 1. Khởi tạo React Hook Form
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<FormData>();

  let navigate = useNavigate();
  const { setCurrentUserLogin } = useUserStore();

  async function onSubmit(data: FormData) {
    
    try {
      const requestBody = {
        username: data.username,
        password: data.password,
      };
      const response = await apiService.post("/login", requestBody);

      const { token, userInfo }: any = response.data;

      if (!token) {
        return;
      }

      localStorage.setItem("token", token);

      setCurrentUserLogin(userInfo);

      navigate("/courses");

    } catch (error) {
      toast.error("Username hoặc password không chính xác");
    }
  }

  const handleKeyDown: React.KeyboardEventHandler<HTMLFormElement> = (e) => {
    if (e.key === "Enter") {
      e.preventDefault();
      handleSubmit(onSubmit)();
    }
  };

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card className="overflow-hidden">
        <CardContent className="grid p-0 md:grid-cols-2">
          {/* 3. Đăng ký handleSubmit */}
          <form
            className="p-6 md:p-8"
            onSubmit={handleSubmit(onSubmit)}
            noValidate
            onKeyDown={handleKeyDown}
          >
            <div className="flex flex-col gap-6">
              <div className="flex flex-col items-center text-center">
                <h1 className="text-2xl font-bold">Welcome back</h1>
                <p className="text-balance text-muted-foreground">
                  Login to your Acme Inc account
                </p>
              </div>

              {/* Username Field */}
              <div className="grid gap-2">
                <Label htmlFor="username">Username</Label>
                <Input
                  id="username"
                  type="text"
                  placeholder="Vui lòng nhập username của bạn"
                  // 4. Đăng ký field với tên 'username' và rule required
                  {...register("username", {
                    required: "Username is required",
                  })}
                />
                {errors.username && (
                  <p className="text-sm text-red-500">
                    {errors.username.message}
                  </p>
                )}
              </div>

              {/* Password Field */}
              <div className="grid gap-2">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                  <a
                    href="#"
                    className="ml-auto text-sm underline-offset-2 hover:underline"
                  >
                    Forgot your password?
                  </a>
                </div>
                <Input
                  id="password"
                  type="password"
                  // 5. Đăng ký field với tên 'password' và rule required
                  {...register("password", {
                    required: "Password is required",
                    minLength: {
                      value: 6,
                      message: "Password must be at least 6 characters",
                    },
                  })}
                />
                {errors.password && (
                  <p className="text-sm text-red-500">
                    {errors.password.message}
                  </p>
                )}
              </div>

              {/* Submit Button */}
              <Button type="submit" className="w-full" disabled={isSubmitting}>
                {isSubmitting ? "Logging in..." : "Login"}
              </Button>

              {/* Social / Divider… */}
            </div>
          </form>

          {/* Side image */}
          <div className="relative hidden bg-muted md:block mr-5 rounded-3xl">
            <img
              src="/public/login.jpg"
              alt="Image"
              loading="lazy"
              className="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale rounded-3xl"
            />
          </div>
        </CardContent>
      </Card>

      <div className="text-balance text-center text-xs text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-primary">
        By clicking continue, you agree to our <a href="#">Terms of Service</a>{" "}
        and <a href="#">Privacy Policy</a>.
      </div>
    </div>
  );
}
