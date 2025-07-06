import { Card } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import Lesson from "./Lesson";
import Quiz from "./Quiz";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useFormContext } from "react-hook-form";

interface ModuleProps {
  id: number;
}

export default function Module({ id }: ModuleProps) {
  const {
    register,
    formState: { errors },
  } = useFormContext();

  return (
    <>
      <div>Modules {id  + 1}</div>
      <div>
        <div className="grid w-full max-w-sm items-center gap-3">
          <Label>Tiêu đề :</Label>
          <Input {...register(`modules.${id}.title`)} />
        </div>

        <div className="grid w-full max-w-sm items-center gap-3">
          <Label>Mô tả :</Label>
          <Input {...register(`modules.${id}.description`)} />
        </div>
      </div>
      <Tabs defaultValue="lessons">
        <TabsList>
          <TabsTrigger value="lessons">Lessons</TabsTrigger>
          <TabsTrigger value="quizzes">Quizzes</TabsTrigger>
        </TabsList>
        <TabsContent value="lessons">
          <Lesson moduleIndex={id} />  
        </TabsContent>
        <TabsContent value="quizzes">
          <Quiz moduleIndex={id} />
        </TabsContent>
      </Tabs>
    </>
  );
}
