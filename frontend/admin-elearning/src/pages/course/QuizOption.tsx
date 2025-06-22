import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
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
            <span>Quiz Option</span>
            <Input
              {...register(
                `modules.${moduleIndex}.quizzes.${quizIndex}.questions.${questionIndex}.options.${index}.optionContent`
              )}
            />

            <Select>
              <SelectTrigger className="w-[180px]">
                <SelectValue placeholder="Theme" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="light">Light</SelectItem>
                <SelectItem value="dark">Dark</SelectItem>
                <SelectItem value="system">System</SelectItem>
              </SelectContent>
            </Select>
          </div>
         
        </div>
      ))}
       <Button
            onClick={() =>
              appendOption({
                optionContent: "",
                isCorrect: false,
                createdAt: "",
                updatedAt: "",
                optionId: 0,
                orderIndex: 0,
                questionId: 0,
              })
            }
          >Add option</Button>
    </div>
  );
}
