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
  UseFieldArrayRemove,
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
              removeLesson={removeLesson}
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
    </div>
  );
}

function LessonItem({
  moduleIndex,
  index,
  register,
  removeLesson,
}: {
  moduleIndex: number;
  index: number;
  register: UseFormRegister<FormCourse>;
  removeLesson: UseFieldArrayRemove;
}) {
  return (
    <>
      <Card className="min-w-80">
        <CardHeader>
          <CardTitle>Card Title</CardTitle>
          <CardDescription>Card Description</CardDescription>
          <CardAction>
            <Button>
              <span onClick={() => removeLesson(index)}>
                Remove Lesson
              </span>
            </Button>
          </CardAction>
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
      </Card>
    </>
  );
}
