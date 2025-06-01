import { Input } from "@/components/ui/input";
import { FormCourse } from "@/types/course";
import React from "react";
import { useFieldArray, useFormContext } from "react-hook-form";

interface QuizOptionProps {
  moduleIndex: number;
  questionIndex: number;
  quizIndex: number;
}

export default function QuizOption({
  questionIndex,
  quizIndex,
  moduleIndex,
}: QuizOptionProps) {
  const {
    register,
    control,
    formState: { errors },
  } = useFormContext<FormCourse>();

  const {
    fields: optionFields,
    append: appendOption,
    remove: removeOption,
  } = useFieldArray<
    FormCourse,
    `modules.${number}.quizzes.${number}.questions.${number}.options`
  >({
    control,
    name: `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options`,
  });

  return (
    <div>
      {optionFields.map((field, index) => (
        <div key={field.id}>
          <div className="grid w-full max-w-sm items-center gap-3">
            <Input
              {...register(
                `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options.${index}.optionContent`
              )}
            />
          </div>
        </div>
      ))}
    </div>
  );
}
