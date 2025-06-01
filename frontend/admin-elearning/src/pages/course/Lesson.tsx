import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { FormCourse } from "@/types/course";
import React from "react";
import { useFieldArray, useFormContext } from "react-hook-form";

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
      {lessonFields.map((field, index) => (
        <div key={field.id}>
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
        </div>
      ))}
      <Button>
        <span
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
          Add Lesson
        </span>
      </Button>

      <Button>
        <span onClick={() => removeLesson(moduleIndex)}>Remove Lesson</span>
      </Button>
    </div>
  );
}
