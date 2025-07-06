import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { FormCourse } from "@/types/course";
import { CirclePlus } from "lucide-react";
import React from "react";
import {
  FieldArray,
  FieldValues,
  useFieldArray,
  useFormContext,
  UseFormRegister,
} from "react-hook-form";

interface LessonProps {
  moduleIndex: number;
}

export default function Lesson({ moduleIndex }: LessonProps) {
  const {
    register,
    control,
    formState: { errors },
  } = useFormContext<FormCourse>();

  const {
    fields: lessonFields,
    append: appendLesson,
    remove: removeLesson,
  } = useFieldArray<FormCourse, `modules.${number}.lessons`>({
    control,
    name: `modules.${moduleIndex}.lessons`,
  });

  return (
    <div>
      <div className="flex flex-wrap gap-5">
        {lessonFields.length === 0 && (
          <Card className="min-w-80 min-h-80 flex items-center justify-center ">
            <CardContent>
              <>
                <span
                  className="cursor-pointer"
                  onClick={() =>
                    appendLesson({
                      title: "",
                      content: "",
                      moduleId: moduleIndex,
                      createdAt: "",
                      updatedAt: "",
                      lessonId: 0,
                      videoUrl: "",
                      orderIndex: 0,
                    })
                  }
                >
                  <CirclePlus size={50} />
                </span>
              </>
            </CardContent>
          </Card>
        )}
        {lessonFields.map((field, index) => (
          <>
            <LessonItem
              register={register}
              moduleIndex={moduleIndex}
              index={index}
              key={field.id}
            />
            {index === lessonFields.length - 1 && (
              <Card className="min-w-80 min-h-80 flex items-center justify-center ">
                <CardContent>
                  <>
                    <span
                      className="cursor-pointer"
                      onClick={() =>
                        appendLesson({
                          title: "",
                          content: "",
                          moduleId: moduleIndex,
                          createdAt: "",
                          updatedAt: "",
                          lessonId: 0,
                          videoUrl: "",
                          orderIndex: 0,
                        })
                      }
                    >
                      <CirclePlus size={50} />
                    </span>
                  </>
                </CardContent>
              </Card>
            )}
          </>
        ))}
      </div>

      <Button>
        <span onClick={() => removeLesson(moduleIndex)}>Remove Lesson</span>
      </Button>
    </div>
  );
}

function LessonItem({
  moduleIndex,
  index,
  register,
}: {
  moduleIndex: number;
  index: number;
  register: UseFormRegister<FormCourse>;
}) {
  return (
    <>
      <Card className="min-w-80">
        <CardHeader>
          <CardTitle>Card Title</CardTitle>
          <CardDescription>Card Description</CardDescription>
          <CardAction>Card Action</CardAction>
        </CardHeader>
        <CardContent>
          <div className="grid w-full max-w-sm items-center gap-3">
            <Label>Tiêu đề bài học :</Label>
            <Input
              {...register(`modules.${moduleIndex}.lessons.${index}.title`)}
            />
          </div>
          <div className="grid w-full max-w-sm items-center gap-3">
            <Label>Mô tả bài học:</Label>
            <Input
              {...register(`modules.${moduleIndex}.lessons.${index}.content`)}
            />
          </div>
          <div className="grid w-full max-w-sm items-center gap-3">
            <Label>Video Url : </Label>
            <Input
              {...register(`modules.${moduleIndex}.lessons.${index}.videoUrl`)}
            />
          </div>
        </CardContent>
        <CardFooter>
          <p>Card Footer</p>
        </CardFooter>
      </Card>
    </>
  );
}
